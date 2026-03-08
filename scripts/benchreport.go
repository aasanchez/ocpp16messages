package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	benchPackage = "./analysis_benchmak"
)

var benchLine = regexp.MustCompile(
	`^(Benchmark\S+)-\d+\s+\d+\s+([0-9.]+)\s+ns/op\s+([0-9.]+)\s+B/op\s+([0-9.]+)\s+allocs/op`,
)

type metric struct {
	NsOp     float64
	BytesOp  float64
	AllocsOp float64
}

type svgSeries struct {
	Name   string
	Color  string
	Values []float64
}

func main() {
	imgDir := flag.String("img-dir", "docs/img", "output directory for svg files")
	reportPath := flag.String("report", "docs/benchmark.md", "output markdown report")
	runBench := flag.Bool("run-bench", true, "run benchmarks before generating")
	inputPath := flag.String("in", "docs/img/benchmark_raw.txt", "benchmark input file")
	flag.Parse()

	if err := os.MkdirAll(*imgDir, 0o755); err != nil {
		exitErr(fmt.Errorf("create image dir: %w", err))
	}

	if *runBench {
		if err := runBenchmarkCommand(*inputPath); err != nil {
			exitErr(err)
		}
	}

	metrics, err := parseBenchMetrics(*inputPath)
	if err != nil {
		exitErr(err)
	}

	if len(metrics) == 0 {
		exitErr(errors.New("no benchmarks parsed"))
	}

	if err := generateCharts(*imgDir, metrics); err != nil {
		exitErr(err)
	}

	if err := writeReport(*reportPath, metrics); err != nil {
		exitErr(err)
	}

	fmt.Printf("Generated report: %s\n", *reportPath)
	fmt.Printf("Generated images in: %s\n", *imgDir)
}

func runBenchmarkCommand(outputPath string) error {
	command := exec.Command(
		"go",
		"test",
		"-tags=bench",
		"-run",
		"^$",
		"-bench",
		".",
		"-benchmem",
		benchPackage,
	)

	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output

	if err := command.Run(); err != nil {
		return fmt.Errorf("run benchmarks: %w\n%s", err, output.String())
	}

	if err := os.WriteFile(outputPath, output.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write benchmark output: %w", err)
	}

	fmt.Printf("Benchmark output saved: %s\n", outputPath)

	return nil
}

func parseBenchMetrics(path string) (map[string]metric, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open benchmark input: %w", err)
	}
	defer func() { _ = file.Close() }()

	results := make(map[string]metric)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := benchLine.FindStringSubmatch(line)
		if len(match) != 5 {
			continue
		}

		nsOp, err := strconv.ParseFloat(match[2], 64)
		if err != nil {
			continue
		}
		bytesOp, err := strconv.ParseFloat(match[3], 64)
		if err != nil {
			continue
		}
		allocsOp, err := strconv.ParseFloat(match[4], 64)
		if err != nil {
			continue
		}

		results[match[1]] = metric{NsOp: nsOp, BytesOp: bytesOp, AllocsOp: allocsOp}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan benchmark input: %w", err)
	}

	return results, nil
}

func generateCharts(imgDir string, metrics map[string]metric) error {
	sizes := []int{1, 25, 100, 250, 500, 1000}
	variants := []string{"Custom", "PrimitiveDirect", "PrimitiveValidated"}
	colors := []string{"#d1495b", "#2a9d8f", "#1f77b4"}

	sllNS := make([]svgSeries, 0, len(variants))
	sllBytes := make([]svgSeries, 0, len(variants))
	sllAllocs := make([]svgSeries, 0, len(variants))
	for index, variant := range variants {
		sllNS = append(sllNS, svgSeries{
			Name:   variant,
			Color:  colors[index],
			Values: scalingValues(metrics, "SendLocalListReq", variant, sizes, "ns"),
		})
		sllBytes = append(sllBytes, svgSeries{
			Name:   variant,
			Color:  colors[index],
			Values: scalingValues(metrics, "SendLocalListReq", variant, sizes, "bytes"),
		})
		sllAllocs = append(sllAllocs, svgSeries{
			Name:   variant,
			Color:  colors[index],
			Values: scalingValues(metrics, "SendLocalListReq", variant, sizes, "allocs"),
		})
	}

	gconfNS := make([]svgSeries, 0, len(variants))
	for index, variant := range variants {
		gconfNS = append(gconfNS, svgSeries{
			Name:   variant,
			Color:  colors[index],
			Values: scalingValues(metrics, "GetConfigurationReq", variant, sizes, "ns"),
		})
	}

	if err := writeLineChart(
		filepath.Join(imgDir, "sendlocallist_ns.svg"),
		"SendLocalListReq Scaling (ns/op)",
		"ns/op",
		sizes,
		sllNS,
	); err != nil {
		return err
	}

	if err := writeLineChart(
		filepath.Join(imgDir, "sendlocallist_bytes.svg"),
		"SendLocalListReq Scaling (B/op)",
		"B/op",
		sizes,
		sllBytes,
	); err != nil {
		return err
	}

	if err := writeLineChart(
		filepath.Join(imgDir, "sendlocallist_allocs.svg"),
		"SendLocalListReq Scaling (allocs/op)",
		"allocs/op",
		sizes,
		sllAllocs,
	); err != nil {
		return err
	}

	if err := writeLineChart(
		filepath.Join(imgDir, "getconfiguration_ns.svg"),
		"GetConfigurationReq Scaling (ns/op)",
		"ns/op",
		sizes,
		gconfNS,
	); err != nil {
		return err
	}

	coreCategories := []string{"DateTime", "ParentIdTag", "StartTransactionReq"}
	coreVariants := []string{"Custom", "PrimitiveDirect", "PrimitiveValidated"}
	coreValues := make([]svgSeries, 0, len(coreVariants))
	for index, variant := range coreVariants {
		dateTimeBenchmark := "BenchmarkDateTime_CustomType"
		if variant == "PrimitiveDirect" {
			dateTimeBenchmark = "BenchmarkDateTime_PrimitiveDirect"
		}
		if variant == "PrimitiveValidated" {
			dateTimeBenchmark = "BenchmarkDateTime_PrimitiveValidated"
		}

		parentBenchmark := "BenchmarkParentIdTag_CustomChain"
		if variant == "PrimitiveDirect" {
			parentBenchmark = "BenchmarkParentIdTag_PrimitiveDirect"
		}
		if variant == "PrimitiveValidated" {
			parentBenchmark = "BenchmarkParentIdTag_PrimitiveValidated"
		}

		coreValues = append(coreValues, svgSeries{
			Name:  variant,
			Color: colors[index],
			Values: []float64{
				lookupMetric(metrics, dateTimeBenchmark).NsOp,
				lookupMetric(metrics, parentBenchmark).NsOp,
				lookupMetric(metrics, "BenchmarkStartTransactionReq_"+variant).NsOp,
			},
		})
	}

	if err := writeGroupedBarChart(
		filepath.Join(imgDir, "core_constructors_ns.svg"),
		"Core Constructors and Message Path (ns/op)",
		coreCategories,
		coreValues,
		"ns/op",
	); err != nil {
		return err
	}

	ratioCategories := []string{
		"DateTime",
		"ParentIdTag",
		"StartTransactionReq",
		"SendLocalListReq_1000",
		"GetConfigurationReq_1000",
	}
	ratioValues := []float64{
		ratio(
			lookupMetric(metrics, "BenchmarkDateTime_CustomType").NsOp,
			lookupMetric(metrics, "BenchmarkDateTime_PrimitiveValidated").NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkParentIdTag_CustomChain").NsOp,
			lookupMetric(metrics, "BenchmarkParentIdTag_PrimitiveValidated").NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkStartTransactionReq_Custom").NsOp,
			lookupMetric(metrics, "BenchmarkStartTransactionReq_PrimitiveValidated").NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkSendLocalListReq_Custom_1000").NsOp,
			lookupMetric(metrics, "BenchmarkSendLocalListReq_PrimitiveValidated_1000").NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkGetConfigurationReq_Custom_1000").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkGetConfigurationReq_PrimitiveValidated_1000",
			).NsOp,
		),
	}

	if err := writeSingleBarChart(
		filepath.Join(imgDir, "custom_vs_validated_ratio.svg"),
		"Custom / PrimitiveValidated Ratio (ns/op)",
		ratioCategories,
		ratioValues,
		"ratio",
	); err != nil {
		return err
	}

	return nil
}

func scalingValues(
	metrics map[string]metric,
	family string,
	variant string,
	sizes []int,
	field string,
) []float64 {
	values := make([]float64, 0, len(sizes))
	for _, size := range sizes {
		name := fmt.Sprintf("Benchmark%s_%s_%d", family, variant, size)
		value := lookupMetric(metrics, name)
		switch field {
		case "bytes":
			values = append(values, value.BytesOp)
		case "allocs":
			values = append(values, value.AllocsOp)
		default:
			values = append(values, value.NsOp)
		}
	}

	return values
}

func lookupMetric(metrics map[string]metric, name string) metric {
	if value, ok := metrics[name]; ok {
		return value
	}

	return metric{}
}

func ratio(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func writeLineChart(
	path string,
	title string,
	yLabel string,
	xValues []int,
	series []svgSeries,
) error {
	const (
		width  = 980.0
		height = 560.0
		left   = 95.0
		right  = 50.0
		top    = 75.0
		bottom = 90.0
	)

	plotWidth := width - left - right
	plotHeight := height - top - bottom
	xMin := float64(xValues[0])
	xMax := float64(xValues[len(xValues)-1])
	yMax := maxSeriesValue(series) * 1.1
	if yMax <= 0 {
		yMax = 1
	}

	xToPx := func(x float64) float64 {
		return left + ((x-xMin)/(xMax-xMin))*plotWidth
	}
	yToPx := func(y float64) float64 {
		return top + (1-(y/yMax))*plotHeight
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create chart %s: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	writer := bufio.NewWriter(file)
	defer func() { _ = writer.Flush() }()

	writeSVGHeader(writer, width, height)
	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"490\" y=\"34\" font-size=\"21\" text-anchor=\"middle\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
		title,
	)

	drawAxes(writer, left, top, plotWidth, plotHeight, xValues, xToPx, yToPx, yMax, yLabel)

	for _, s := range series {
		points := make([]string, 0, len(xValues))
		for index, xValue := range xValues {
			x := xToPx(float64(xValue))
			y := yToPx(s.Values[index])
			points = append(points, fmt.Sprintf("%.2f,%.2f", x, y))
		}

		_, _ = fmt.Fprintf(
			writer,
			"  <polyline points=\"%s\" fill=\"none\" stroke=\"%s\" stroke-width=\"3\"/>\n",
			strings.Join(points, " "),
			s.Color,
		)

		for index, xValue := range xValues {
			x := xToPx(float64(xValue))
			y := yToPx(s.Values[index])
			_, _ = fmt.Fprintf(
				writer,
				"  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4.2\" fill=\"%s\"/>\n",
				x,
				y,
				s.Color,
			)
		}
	}

	drawLegend(writer, series, 640, 86)
	_, _ = fmt.Fprintln(writer, "</svg>")

	return nil
}

func writeGroupedBarChart(
	path string,
	title string,
	categories []string,
	series []svgSeries,
	yLabel string,
) error {
	const (
		width  = 980.0
		height = 560.0
		left   = 95.0
		right  = 50.0
		top    = 75.0
		bottom = 110.0
	)

	plotWidth := width - left - right
	plotHeight := height - top - bottom
	yMax := maxSeriesValue(series) * 1.1
	if yMax <= 0 {
		yMax = 1
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create chart %s: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	writer := bufio.NewWriter(file)
	defer func() { _ = writer.Flush() }()

	writeSVGHeader(writer, width, height)
	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"490\" y=\"34\" font-size=\"21\" text-anchor=\"middle\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
		title,
	)

	drawBarAxes(writer, left, top, plotWidth, plotHeight, categories, yMax, yLabel)

	groupWidth := plotWidth / float64(len(categories))
	barWidth := (groupWidth * 0.7) / float64(len(series))
	for categoryIndex := range categories {
		groupStart := left + float64(categoryIndex)*groupWidth + groupWidth*0.15
		for seriesIndex, s := range series {
			value := s.Values[categoryIndex]
			barHeight := (value / yMax) * plotHeight
			x := groupStart + float64(seriesIndex)*barWidth
			y := top + plotHeight - barHeight
			_, _ = fmt.Fprintf(
				writer,
				"  <rect x=\"%.2f\" y=\"%.2f\" width=\"%.2f\" height=\"%.2f\" fill=\"%s\"/>\n",
				x,
				y,
				barWidth,
				barHeight,
				s.Color,
			)
		}
	}

	drawLegend(writer, series, 640, 86)
	_, _ = fmt.Fprintln(writer, "</svg>")

	return nil
}

func writeSingleBarChart(
	path string,
	title string,
	categories []string,
	values []float64,
	yLabel string,
) error {
	series := []svgSeries{{Name: "ratio", Color: "#6c5ce7", Values: values}}
	return writeGroupedBarChart(path, title, categories, series, yLabel)
}

func writeSVGHeader(writer *bufio.Writer, width, height float64) {
	_, _ = fmt.Fprintln(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	_, _ = fmt.Fprintf(
		writer,
		"<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%.0f\" height=\"%.0f\" viewBox=\"0 0 %.0f %.0f\">\n",
		width,
		height,
		width,
		height,
	)
	_, _ = fmt.Fprintln(writer, "  <rect x=\"0\" y=\"0\" width=\"100%\" height=\"100%\" fill=\"white\"/>")
}

func drawAxes(
	writer *bufio.Writer,
	left,
	top,
	plotWidth,
	plotHeight float64,
	xValues []int,
	xToPx func(float64) float64,
	yToPx func(float64) float64,
	yMax float64,
	yLabel string,
) {
	_, _ = fmt.Fprintf(
		writer,
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n",
		left,
		top+plotHeight,
		left+plotWidth,
		top+plotHeight,
	)
	_, _ = fmt.Fprintf(
		writer,
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n",
		left,
		top,
		left,
		top+plotHeight,
	)

	for _, tick := range xValues {
		x := xToPx(float64(tick))
		_, _ = fmt.Fprintf(
			writer,
			"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#eee\"/>\n",
			x,
			top,
			x,
			top+plotHeight,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"middle\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%d</text>\n",
			x,
			top+plotHeight+24,
			tick,
		)
	}

	for tick := 0; tick <= 6; tick++ {
		value := (float64(tick) / 6.0) * yMax
		y := yToPx(value)
		_, _ = fmt.Fprintf(
			writer,
			"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#eee\"/>\n",
			left,
			y,
			left+plotWidth,
			y,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"end\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%.1f</text>\n",
			left-8,
			y+4,
			value,
		)
	}

	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"13\" text-anchor=\"middle\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">Input size</text>\n",
		left+plotWidth/2,
		top+plotHeight+52,
	)

	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"28\" y=\"%.2f\" font-size=\"13\" text-anchor=\"middle\" fill=\"#333\" transform=\"rotate(-90 28 %.2f)\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
		top+plotHeight/2,
		top+plotHeight/2,
		yLabel,
	)
}

func drawBarAxes(
	writer *bufio.Writer,
	left,
	top,
	plotWidth,
	plotHeight float64,
	categories []string,
	yMax float64,
	yLabel string,
) {
	_, _ = fmt.Fprintf(
		writer,
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n",
		left,
		top+plotHeight,
		left+plotWidth,
		top+plotHeight,
	)
	_, _ = fmt.Fprintf(
		writer,
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n",
		left,
		top,
		left,
		top+plotHeight,
	)

	groupWidth := plotWidth / float64(len(categories))
	for index, category := range categories {
		x := left + float64(index)*groupWidth + groupWidth/2
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"middle\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
			x,
			top+plotHeight+24,
			category,
		)
	}

	for tick := 0; tick <= 6; tick++ {
		value := (float64(tick) / 6.0) * yMax
		y := top + (1-(value/yMax))*plotHeight
		_, _ = fmt.Fprintf(
			writer,
			"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#eee\"/>\n",
			left,
			y,
			left+plotWidth,
			y,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"end\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%.2f</text>\n",
			left-8,
			y+4,
			value,
		)
	}

	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"26\" y=\"%.2f\" font-size=\"13\" text-anchor=\"middle\" fill=\"#333\" transform=\"rotate(-90 26 %.2f)\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
		top+plotHeight/2,
		top+plotHeight/2,
		yLabel,
	)
}

func drawLegend(
	writer *bufio.Writer,
	series []svgSeries,
	startX float64,
	startY float64,
) {
	y := startY
	for _, s := range series {
		_, _ = fmt.Fprintf(
			writer,
			"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"%s\" stroke-width=\"3\"/>\n",
			startX,
			y,
			startX+30,
			y,
			s.Color,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4\" fill=\"%s\"/>\n",
			startX+15,
			y,
			s.Color,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n",
			startX+40,
			y+4,
			s.Name,
		)
		y += 20
	}
}

func maxSeriesValue(series []svgSeries) float64 {
	max := 0.0
	for _, s := range series {
		for _, value := range s.Values {
			max = math.Max(max, value)
		}
	}
	if max == 0 {
		return 1
	}
	return max
}

func writeReport(path string, metrics map[string]metric) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create report dir: %w", err)
	}

	sllCustom1000 := lookupMetric(metrics, "BenchmarkSendLocalListReq_Custom_1000")
	sllValidated1000 := lookupMetric(
		metrics,
		"BenchmarkSendLocalListReq_PrimitiveValidated_1000",
	)
	getCustom1000 := lookupMetric(metrics, "BenchmarkGetConfigurationReq_Custom_1000")
	getValidated1000 := lookupMetric(
		metrics,
		"BenchmarkGetConfigurationReq_PrimitiveValidated_1000",
	)

	lines := []string{
		"# Benchmark Report: Custom Types vs Primitives",
		"",
		"This report is generated from benchmarks in `analysis_benchmak` and charts",
		"in `docs/img/`.",
		"",
		"## How To Reproduce",
		"",
		"```sh",
		"go run ./scripts/benchreport.go",
		"```",
		"",
		"## Scope",
		"",
		"- Core constructors: `DateTime`, `ParentIdTag`, `StartTransactionReq`",
		"- Scaling path #1: `SendLocalListReq` (1 to 1000 entries)",
		"- Scaling path #2: `GetConfigurationReq` (1 to 1000 keys)",
		"- Metrics: `ns/op`, `B/op`, `allocs/op`",
		"",
		"## Charts",
		"",
		"### 1) SendLocalListReq Scaling (ns/op)",
		"",
		"![SendLocalListReq ns/op](img/sendlocallist_ns.svg)",
		"",
		"### 2) SendLocalListReq Scaling (B/op)",
		"",
		"![SendLocalListReq B/op](img/sendlocallist_bytes.svg)",
		"",
		"### 3) SendLocalListReq Scaling (allocs/op)",
		"",
		"![SendLocalListReq allocs/op](img/sendlocallist_allocs.svg)",
		"",
		"### 4) GetConfigurationReq Scaling (ns/op)",
		"",
		"![GetConfigurationReq ns/op](img/getconfiguration_ns.svg)",
		"",
		"### 5) Core Constructors and Message Path (ns/op)",
		"",
		"![Core constructors ns/op](img/core_constructors_ns.svg)",
		"",
		"### 6) Custom / PrimitiveValidated Ratio (ns/op)",
		"",
		"![Custom vs PrimitiveValidated ratio](img/custom_vs_validated_ratio.svg)",
		"",
		"## Key Numbers",
		"",
		"| Case | Custom ns/op | PrimitiveValidated ns/op | Ratio |",
		"| ---- | -----------: | -----------------------: | ----: |",
		fmt.Sprintf(
			"| SendLocalListReq_1000 | %.2f | %.2f | %.2fx |",
			sllCustom1000.NsOp,
			sllValidated1000.NsOp,
			ratio(sllCustom1000.NsOp, sllValidated1000.NsOp),
		),
		fmt.Sprintf(
			"| GetConfigurationReq_1000 | %.2f | %.2f | %.2fx |",
			getCustom1000.NsOp,
			getValidated1000.NsOp,
			ratio(getCustom1000.NsOp, getValidated1000.NsOp),
		),
		fmt.Sprintf(
			"| StartTransactionReq | %.2f | %.2f | %.2fx |",
			lookupMetric(metrics, "BenchmarkStartTransactionReq_Custom").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkStartTransactionReq_PrimitiveValidated",
			).NsOp,
			ratio(
				lookupMetric(metrics, "BenchmarkStartTransactionReq_Custom").NsOp,
				lookupMetric(
					metrics,
					"BenchmarkStartTransactionReq_PrimitiveValidated",
				).NsOp,
			),
		),
		"",
		"## Analysis",
		"",
		"1. `PrimitiveDirect` is the fastest baseline because it skips validation.",
		"2. Against a fair baseline (`PrimitiveValidated`), custom types add",
		"   bounded overhead but keep all OCPP validation centralized.",
		"3. As input size grows, custom and validated primitive lines trend",
		"   similarly (same O(n) shape), which means scaling behavior is",
		"   predictable.",
		"4. Allocation charts show where object wrapping adds memory cost; this",
		"   is measurable but small compared to typical network/JSON costs in",
		"   end-to-end OCPP flows.",
		"",
		"## Conclusion",
		"",
		"Using first-class datatypes is a speed vs safety tradeoff. If you want",
		"maximum raw microbenchmark speed, direct primitives win. If you want",
		"stronger correctness guarantees, clearer APIs, and less repeated",
		"validation logic, custom datatypes provide a practical and predictable",
		"cost profile.",
	}

	content := strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write report: %w", err)
	}

	return nil
}

func exitErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
