// Command benchreport generates the benchmark markdown report and SVG charts.
package main

import (
	"bufio"
	"bytes"
	"context"
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

	fileModeDir  = 0o750
	fileModeRead = 0o600

	zeroInt         = 0
	oneInt          = 1
	twoInt          = 2
	regexSubmatches = 5
	floatBitSize    = 64
	yAxisTicks      = 6

	zeroFloat          = 0.0
	oneFloat           = 1.0
	twoFloat           = 2.0
	yAxisPaddingFactor = 1.1
	barWidthFactor     = 0.7
	groupOffsetFactor  = 0.15
	defaultChartMax    = oneFloat
	legendItemHeight   = 20.0
	gridOffsetX        = 8.0
	gridOffsetY        = 4.0
	legendLineWidth    = 30.0
	legendCircleOffset = 15.0
	legendTextOffset   = 40.0
	xTickLabelOffset   = 24.0
	xAxisTitleOffset   = 52.0
	legendStartX       = 640.0
	legendStartY       = 86.0
	titleCenterX       = 490.0
	titleY             = 34.0
	layoutWidth        = 980.0
	layoutHeight       = 560.0
	layoutLeft         = 95.0
	layoutRight        = 50.0
	layoutTop          = 75.0
	lineBottom         = 90.0
	barBottom          = 110.0
	svgRatioColor      = "#6c5ce7"

	sizeTwentyFive   = 25
	sizeOneHundred   = 100
	sizeTwoHundred50 = 250
	sizeFiveHundred  = 500
	sizeOneThousand  = 1000
	reportLineCap    = 64
)

const (
	metricNS     = "ns"
	metricBytes  = "bytes"
	metricAllocs = "allocs"
	yLabelNS     = "ns/op"
	yLabelBytes  = "B/op"
	yLabelAllocs = "allocs/op"
	yLabelRatio  = "ratio"
)

const (
	benchFamilySendLocalList = "SendLocalListReq"
	benchFamilyConfigReq     = "GetConfigurationReq"
	benchStartTransaction    = "StartTransactionReq"

	variantCustom             = "Custom"
	variantPrimitiveDirect    = "PrimitiveDirect"
	variantPrimitiveValidated = "PrimitiveValidated"
)

const (
	svgFontFamily = "Helvetica,Arial,sans-serif"
	svgAxisLine   = "" +
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" " +
		"stroke=\"#222\" stroke-width=\"1.5\"/>\n"
	svgGridLine = "" +
		"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" " +
		"stroke=\"#eee\"/>\n"
	svgTitle = "" +
		"  <text x=\"%.0f\" y=\"%.0f\" font-size=\"21\" " +
		"text-anchor=\"middle\" font-family=\"" + svgFontFamily + "\">%s" +
		"</text>\n"
	svgRoot = "" +
		"<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%.0f\" " +
		"height=\"%.0f\" viewBox=\"0 0 %.0f %.0f\">\n"
)

const (
	reportTitle = "# Benchmark Report: Custom Types vs Primitives"
	reportIntro = "" +
		"This report is generated from benchmarks in `analysis_benchmak` " +
		"and charts in `docs/img/`."
	reportAnalysisA = "" +
		"1. `PrimitiveDirect` is the fastest baseline because it skips " +
		"validation."
	reportAnalysisB = "" +
		"2. Against a fair baseline (`PrimitiveValidated`), custom types " +
		"add bounded overhead but keep all OCPP validation centralized."
	reportAnalysisC = "" +
		"3. As input size grows, custom and validated primitive lines " +
		"trend similarly (same O(n) shape), which means scaling behavior " +
		"is predictable."
	reportAnalysisD = "" +
		"4. Allocation charts show where object wrapping adds memory cost; " +
		"this is measurable but small compared to typical network/JSON " +
		"costs in end-to-end OCPP flows."
	reportConclusionA = "" +
		"Using first-class datatypes is a speed vs safety tradeoff. If " +
		"you want maximum raw microbenchmark speed, direct primitives win."
	reportConclusionB = "" +
		"If you want stronger correctness guarantees, clearer APIs, and " +
		"less repeated validation logic, custom datatypes provide a " +
		"practical and predictable cost profile."
	reportBlankLine = ""
)

var (
	errNoBenchmarksParsed = errors.New("no benchmarks parsed")

	benchLine = regexp.MustCompile(
		`^(Benchmark\S+)-\d+\s+\d+\s+([0-9.]+)\s+ns/op\s+` +
			`([0-9.]+)\s+B/op\s+([0-9.]+)\s+allocs/op`,
	)
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

type chartLayout struct {
	Width      float64
	Height     float64
	Left       float64
	Right      float64
	Top        float64
	Bottom     float64
	PlotWidth  float64
	PlotHeight float64
}

type options struct {
	ImgDir     string
	ReportPath string
	RunBench   bool
	InputPath  string
}

type scalingChartConfig struct {
	FileName string
	Title    string
	Family   string
	Field    string
	YLabel   string
}

func main() {
	opts := parseFlags()

	flag.Parse()

	err := run(opts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(oneInt)
	}
}

func parseFlags() options {
	imgDir := flag.String(
		"img-dir",
		"docs/img",
		"output directory for svg files",
	)
	reportPath := flag.String(
		"report",
		"docs/benchmark.md",
		"output markdown report",
	)
	runBench := flag.Bool("run-bench", true, "run benchmarks before generating")
	inputPath := flag.String(
		"in",
		"docs/img/benchmark_raw.txt",
		"benchmark input file",
	)

	return options{
		ImgDir:     *imgDir,
		ReportPath: *reportPath,
		RunBench:   *runBench,
		InputPath:  *inputPath,
	}
}

func run(opts options) error {
	err := ensureDir(opts.ImgDir)
	if err != nil {
		return fmt.Errorf("create image dir: %w", err)
	}

	err = maybeRunBenchmarks(opts)
	if err != nil {
		return err
	}

	metrics, err := loadMetrics(opts.InputPath)
	if err != nil {
		return err
	}

	err = generateCharts(opts.ImgDir, metrics)
	if err != nil {
		return err
	}

	err = writeReport(opts.ReportPath, metrics)
	if err != nil {
		return err
	}

	writeSuccessMessage(opts)

	return nil
}

func maybeRunBenchmarks(opts options) error {
	if !opts.RunBench {
		return nil
	}

	return runBenchmarkCommand(opts.InputPath)
}

func loadMetrics(inputPath string) (map[string]metric, error) {
	metrics, err := parseBenchMetrics(inputPath)
	if err != nil {
		return nil, err
	}

	if len(metrics) == zeroInt {
		return nil, errNoBenchmarksParsed
	}

	return metrics, nil
}

func writeSuccessMessage(opts options) {
	_, _ = fmt.Fprintf(os.Stdout, "Generated report: %s\n", opts.ReportPath)
	_, _ = fmt.Fprintf(os.Stdout, "Generated images in: %s\n", opts.ImgDir)
}

func runBenchmarkCommand(outputPath string) error {
	ctx := context.Background()
	command := exec.CommandContext(
		ctx,
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

	err := command.Run()
	if err != nil {
		return fmt.Errorf("run benchmarks: %w\n%s", err, output.String())
	}

	err = writeFile(outputPath, output.Bytes())
	if err != nil {
		return fmt.Errorf("write benchmark output: %w", err)
	}

	_, _ = fmt.Fprintf(os.Stdout, "Benchmark output saved: %s\n", outputPath)

	return nil
}

func parseBenchMetrics(path string) (map[string]metric, error) {
	file, err := openFileRead(path)
	if err != nil {
		return nil, fmt.Errorf("open benchmark input: %w", err)
	}

	defer closeFile(file)

	results := make(map[string]metric)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parseMetricLine(results, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("scan benchmark input: %w", err)
	}

	return results, nil
}

func parseMetricLine(results map[string]metric, line string) {
	match := benchLine.FindStringSubmatch(line)
	if len(match) != regexSubmatches {
		return
	}

	parsedMetric, parsed := parseMetricMatch(match)
	if !parsed {
		return
	}

	results[match[oneInt]] = parsedMetric
}

func parseMetricMatch(match []string) (metric, bool) {
	nsOp, parsed := parseFloatValue(match[twoInt])
	if !parsed {
		return zeroMetric(), false
	}

	bytesOp, parsed := parseFloatValue(match[3])
	if !parsed {
		return zeroMetric(), false
	}

	allocsOp, parsed := parseFloatValue(match[4])
	if !parsed {
		return zeroMetric(), false
	}

	return metric{
		NsOp:     nsOp,
		BytesOp:  bytesOp,
		AllocsOp: allocsOp,
	}, true
}

func parseFloatValue(raw string) (float64, bool) {
	value, err := strconv.ParseFloat(raw, floatBitSize)
	if err != nil {
		return zeroFloat, false
	}

	return value, true
}

func generateCharts(imgDir string, metrics map[string]metric) error {
	err := writeScalingCharts(imgDir, metrics)
	if err != nil {
		return err
	}

	err = writeCoreChart(imgDir, metrics)
	if err != nil {
		return err
	}

	err = writeRatioChart(imgDir, metrics)
	if err != nil {
		return err
	}

	return nil
}

func writeScalingCharts(imgDir string, metrics map[string]metric) error {
	for _, config := range scalingChartConfigs() {
		err := writeScalingChart(imgDir, metrics, config)
		if err != nil {
			return err
		}
	}

	return nil
}

func scalingChartConfigs() []scalingChartConfig {
	return []scalingChartConfig{
		{
			FileName: "sendlocallist_ns.svg",
			Title:    "SendLocalListReq Scaling (ns/op)",
			Family:   benchFamilySendLocalList,
			Field:    metricNS,
			YLabel:   yLabelNS,
		},
		{
			FileName: "sendlocallist_bytes.svg",
			Title:    "SendLocalListReq Scaling (B/op)",
			Family:   benchFamilySendLocalList,
			Field:    metricBytes,
			YLabel:   yLabelBytes,
		},
		{
			FileName: "sendlocallist_allocs.svg",
			Title:    "SendLocalListReq Scaling (allocs/op)",
			Family:   benchFamilySendLocalList,
			Field:    metricAllocs,
			YLabel:   yLabelAllocs,
		},
		{
			FileName: "getconfiguration_ns.svg",
			Title:    "GetConfigurationReq Scaling (ns/op)",
			Family:   benchFamilyConfigReq,
			Field:    metricNS,
			YLabel:   yLabelNS,
		},
	}
}

func writeScalingChart(
	imgDir string,
	metrics map[string]metric,
	config scalingChartConfig,
) error {
	sizes := []int{
		oneInt,
		sizeTwentyFive,
		sizeOneHundred,
		sizeTwoHundred50,
		sizeFiveHundred,
		sizeOneThousand,
	}
	series := buildScalingSeries(
		metrics,
		config.Family,
		config.Field,
		sizes,
		benchmarkVariants(),
		benchmarkColors(),
	)

	return writeLineChart(
		filepath.Join(imgDir, config.FileName),
		config.Title,
		config.YLabel,
		sizes,
		series,
	)
}

func buildScalingSeries(
	metrics map[string]metric,
	family string,
	field string,
	sizes []int,
	variants []string,
	colors []string,
) []svgSeries {
	series := make([]svgSeries, zeroInt, len(variants))

	for index, variant := range variants {
		series = append(series, svgSeries{
			Name:   variant,
			Color:  colors[index],
			Values: scalingValues(metrics, family, variant, sizes, field),
		})
	}

	return series
}

func writeCoreChart(imgDir string, metrics map[string]metric) error {
	categories := []string{
		"DateTime",
		"ParentIdTag",
		benchStartTransaction,
	}
	variants := benchmarkVariants()
	colors := benchmarkColors()
	series := make([]svgSeries, zeroInt, len(variants))

	for index, variant := range variants {
		series = append(series, svgSeries{
			Name:  variant,
			Color: colors[index],
			Values: []float64{
				lookupMetric(metrics, dateTimeBenchmarkName(variant)).NsOp,
				lookupMetric(metrics, parentTagBenchmarkName(variant)).NsOp,
				lookupMetric(
					metrics,
					"BenchmarkStartTransactionReq_"+variant,
				).NsOp,
			},
		})
	}

	return writeGroupedBarChart(
		filepath.Join(imgDir, "core_constructors_ns.svg"),
		"Core Constructors and Message Path (ns/op)",
		categories,
		series,
		yLabelNS,
	)
}

func writeRatioChart(imgDir string, metrics map[string]metric) error {
	categories := []string{
		"DateTime",
		"ParentIdTag",
		benchStartTransaction,
		"SendLocalListReq_1000",
		"GetConfigurationReq_1000",
	}
	values := []float64{
		ratio(
			lookupMetric(metrics, dateTimeBenchmarkName(variantCustom)).NsOp,
			lookupMetric(
				metrics,
				dateTimeBenchmarkName(variantPrimitiveValidated),
			).NsOp,
		),
		ratio(
			lookupMetric(metrics, parentTagBenchmarkName(variantCustom)).NsOp,
			lookupMetric(
				metrics,
				parentTagBenchmarkName(variantPrimitiveValidated),
			).NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkStartTransactionReq_Custom").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkStartTransactionReq_PrimitiveValidated",
			).NsOp,
		),
		ratio(
			lookupMetric(metrics, "BenchmarkSendLocalListReq_Custom_1000").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkSendLocalListReq_PrimitiveValidated_1000",
			).NsOp,
		),
		ratio(
			lookupMetric(
				metrics,
				"BenchmarkGetConfigurationReq_Custom_1000",
			).NsOp,
			lookupMetric(
				metrics,
				"BenchmarkGetConfigurationReq_PrimitiveValidated_1000",
			).NsOp,
		),
	}

	return writeSingleBarChart(
		filepath.Join(imgDir, "custom_vs_validated_ratio.svg"),
		"Custom / PrimitiveValidated Ratio (ns/op)",
		categories,
		values,
		yLabelRatio,
	)
}

func scalingValues(
	metrics map[string]metric,
	family string,
	variant string,
	sizes []int,
	field string,
) []float64 {
	values := make([]float64, zeroInt, len(sizes))

	for _, size := range sizes {
		name := fmt.Sprintf("Benchmark%s_%s_%d", family, variant, size)
		metricValue := lookupMetric(metrics, name)
		values = append(values, metricFieldValue(metricValue, field))
	}

	return values
}

func metricFieldValue(value metric, field string) float64 {
	switch field {
	case metricBytes:
		return value.BytesOp
	case metricAllocs:
		return value.AllocsOp
	default:
		return value.NsOp
	}
}

func lookupMetric(metrics map[string]metric, name string) metric {
	value, ok := metrics[name]
	if ok {
		return value
	}

	return zeroMetric()
}

func zeroMetric() metric {
	return metric{
		NsOp:     zeroFloat,
		BytesOp:  zeroFloat,
		AllocsOp: zeroFloat,
	}
}

func ratio(numerator, denominator float64) float64 {
	if denominator == zeroFloat {
		return zeroFloat
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
	layout := newLineChartLayout()
	yMax := chartYMax(series)
	xMin := float64(xValues[zeroInt])
	xMax := float64(xValues[len(xValues)-oneInt])

	xToPixel := func(value float64) float64 {
		return layout.Left + ((value-xMin)/(xMax-xMin))*layout.PlotWidth
	}
	yToPixel := func(value float64) float64 {
		return layout.Top + (oneFloat-(value/yMax))*layout.PlotHeight
	}

	writer, closeWriter, err := newSVGWriter(path)
	if err != nil {
		return fmt.Errorf("create chart %s: %w", path, err)
	}

	defer closeWriter()

	writeSVGHeader(writer, layout.Width, layout.Height)
	writeChartTitle(writer, title)
	drawLineAxes(writer, layout, xValues, xToPixel, yToPixel, yMax, yLabel)
	drawLineSeries(writer, xValues, xToPixel, yToPixel, series)
	drawLegend(writer, series, legendStartX, legendStartY)
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
	layout := newBarChartLayout()
	yMax := chartYMax(series)

	writer, closeWriter, err := newSVGWriter(path)
	if err != nil {
		return fmt.Errorf("create chart %s: %w", path, err)
	}

	defer closeWriter()

	writeSVGHeader(writer, layout.Width, layout.Height)
	writeChartTitle(writer, title)
	drawBarAxes(writer, layout, categories, yMax, yLabel)
	drawBarSeries(writer, layout, categories, series, yMax)
	drawLegend(writer, series, legendStartX, legendStartY)
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
	series := []svgSeries{
		{
			Name:   yLabelRatio,
			Color:  svgRatioColor,
			Values: values,
		},
	}

	return writeGroupedBarChart(path, title, categories, series, yLabel)
}

func newLineChartLayout() chartLayout {
	return buildLayout(
		layoutWidth,
		layoutHeight,
		layoutLeft,
		layoutRight,
		layoutTop,
		lineBottom,
	)
}

func newBarChartLayout() chartLayout {
	return buildLayout(
		layoutWidth,
		layoutHeight,
		layoutLeft,
		layoutRight,
		layoutTop,
		barBottom,
	)
}

func buildLayout(
	width float64,
	height float64,
	left float64,
	right float64,
	top float64,
	bottom float64,
) chartLayout {
	return chartLayout{
		Width:      width,
		Height:     height,
		Left:       left,
		Right:      right,
		Top:        top,
		Bottom:     bottom,
		PlotWidth:  width - left - right,
		PlotHeight: height - top - bottom,
	}
}

func newSVGWriter(path string) (*bufio.Writer, func(), error) {
	file, err := createFileWrite(path)
	if err != nil {
		return nil, nil, err
	}

	writer := bufio.NewWriter(file)
	closeWriter := func() {
		_ = writer.Flush()

		closeFile(file)
	}

	return writer, closeWriter, nil
}

func writeSVGHeader(writer *bufio.Writer, width, height float64) {
	_, _ = fmt.Fprintln(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	_, _ = fmt.Fprintf(writer, svgRoot, width, height, width, height)
	_, _ = fmt.Fprintln(
		writer,
		"  <rect x=\"0\" y=\"0\" width=\"100%\" height=\"100%\" "+
			"fill=\"white\"/>",
	)
}

func writeChartTitle(writer *bufio.Writer, title string) {
	_, _ = fmt.Fprintf(writer, svgTitle, titleCenterX, titleY, title)
}

func drawLineAxes(
	writer *bufio.Writer,
	layout chartLayout,
	xValues []int,
	xToPixel func(float64) float64,
	yToPixel func(float64) float64,
	yMax float64,
	yLabel string,
) {
	drawAxesSkeleton(writer, layout)
	drawLineXAxis(writer, layout, xValues, xToPixel)
	drawYAxis(writer, layout, yToPixel, yMax, yLabel, "%.1f")
	writeXAxisTitle(writer, layout, "Input size")
}

func drawBarAxes(
	writer *bufio.Writer,
	layout chartLayout,
	categories []string,
	yMax float64,
	yLabel string,
) {
	drawAxesSkeleton(writer, layout)
	drawBarXAxis(writer, layout, categories)
	drawYAxis(
		writer,
		layout,
		func(value float64) float64 {
			return layout.Top + (oneFloat-(value/yMax))*layout.PlotHeight
		},
		yMax,
		yLabel,
		"%.2f",
	)
}

func drawAxesSkeleton(writer *bufio.Writer, layout chartLayout) {
	_, _ = fmt.Fprintf(
		writer,
		svgAxisLine,
		layout.Left,
		layout.Top+layout.PlotHeight,
		layout.Left+layout.PlotWidth,
		layout.Top+layout.PlotHeight,
	)
	_, _ = fmt.Fprintf(
		writer,
		svgAxisLine,
		layout.Left,
		layout.Top,
		layout.Left,
		layout.Top+layout.PlotHeight,
	)
}

func drawLineXAxis(
	writer *bufio.Writer,
	layout chartLayout,
	xValues []int,
	xToPixel func(float64) float64,
) {
	for _, tick := range xValues {
		xValue := xToPixel(float64(tick))
		_, _ = fmt.Fprintf(
			writer,
			svgGridLine,
			xValue,
			layout.Top,
			xValue,
			layout.Top+layout.PlotHeight,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" "+
				"text-anchor=\"middle\" fill=\"#333\" "+
				"font-family=\"%s\">%d</text>\n",
			xValue,
			layout.Top+layout.PlotHeight+xTickLabelOffset,
			svgFontFamily,
			tick,
		)
	}
}

func drawBarXAxis(
	writer *bufio.Writer,
	layout chartLayout,
	categories []string,
) {
	groupWidth := layout.PlotWidth / float64(len(categories))

	for index, category := range categories {
		xValue := layout.Left +
			float64(index)*groupWidth +
			groupWidth/twoFloat
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" "+
				"text-anchor=\"middle\" fill=\"#333\" "+
				"font-family=\"%s\">%s</text>\n",
			xValue,
			layout.Top+layout.PlotHeight+xTickLabelOffset,
			svgFontFamily,
			category,
		)
	}
}

func drawYAxis(
	writer *bufio.Writer,
	layout chartLayout,
	yToPixel func(float64) float64,
	yMax float64,
	yLabel string,
	format string,
) {
	for tick := zeroInt; tick <= yAxisTicks; tick++ {
		value := (float64(tick) / float64(yAxisTicks)) * yMax
		yValue := yToPixel(value)

		_, _ = fmt.Fprintf(
			writer,
			svgGridLine,
			layout.Left,
			yValue,
			layout.Left+layout.PlotWidth,
			yValue,
		)

		writeYAxisTick(
			writer,
			layout.Left-gridOffsetX,
			yValue+gridOffsetY,
			format,
			value,
		)
	}

	writeYAxisTitle(writer, layout, yLabel)
}

func writeYAxisTick(
	writer *bufio.Writer,
	xValue float64,
	yValue float64,
	format string,
	value float64,
) {
	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" "+
			"text-anchor=\"end\" fill=\"#333\" font-family=\"%s\">"+
			format+"</text>\n",
		xValue,
		yValue,
		svgFontFamily,
		value,
	)
}

func writeXAxisTitle(
	writer *bufio.Writer,
	layout chartLayout,
	title string,
) {
	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"13\" "+
			"text-anchor=\"middle\" fill=\"#333\" font-family=\"%s\">%s"+
			"</text>\n",
		layout.Left+layout.PlotWidth/twoFloat,
		layout.Top+layout.PlotHeight+xAxisTitleOffset,
		svgFontFamily,
		title,
	)
}

func writeYAxisTitle(
	writer *bufio.Writer,
	layout chartLayout,
	yLabel string,
) {
	centerY := layout.Top + layout.PlotHeight/twoFloat
	_, _ = fmt.Fprintf(
		writer,
		"  <text x=\"28\" y=\"%.2f\" font-size=\"13\" "+
			"text-anchor=\"middle\" fill=\"#333\" "+
			"transform=\"rotate(-90 28 %.2f)\" "+
			"font-family=\"%s\">%s</text>\n",
		centerY,
		centerY,
		svgFontFamily,
		yLabel,
	)
}

func drawLineSeries(
	writer *bufio.Writer,
	xValues []int,
	xToPixel func(float64) float64,
	yToPixel func(float64) float64,
	series []svgSeries,
) {
	for _, currentSeries := range series {
		points := buildPolylinePoints(
			currentSeries,
			xValues,
			xToPixel,
			yToPixel,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <polyline points=\"%s\" fill=\"none\" stroke=\"%s\" "+
				"stroke-width=\"3\"/>\n",
			points,
			currentSeries.Color,
		)

		drawLineMarkers(writer, currentSeries, xValues, xToPixel, yToPixel)
	}
}

func buildPolylinePoints(
	currentSeries svgSeries,
	xValues []int,
	xToPixel func(float64) float64,
	yToPixel func(float64) float64,
) string {
	points := make([]string, zeroInt, len(xValues))

	for index, xValue := range xValues {
		points = append(
			points,
			fmt.Sprintf(
				"%.2f,%.2f",
				xToPixel(float64(xValue)),
				yToPixel(currentSeries.Values[index]),
			),
		)
	}

	return strings.Join(points, " ")
}

func drawLineMarkers(
	writer *bufio.Writer,
	currentSeries svgSeries,
	xValues []int,
	xToPixel func(float64) float64,
	yToPixel func(float64) float64,
) {
	for index, xValue := range xValues {
		_, _ = fmt.Fprintf(
			writer,
			"  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4.2\" fill=\"%s\"/>\n",
			xToPixel(float64(xValue)),
			yToPixel(currentSeries.Values[index]),
			currentSeries.Color,
		)
	}
}

func drawBarSeries(
	writer *bufio.Writer,
	layout chartLayout,
	categories []string,
	series []svgSeries,
	yMax float64,
) {
	groupWidth := layout.PlotWidth / float64(len(categories))
	barWidth := (groupWidth * barWidthFactor) / float64(len(series))

	for categoryIndex := range categories {
		groupStart := layout.Left +
			float64(categoryIndex)*groupWidth +
			groupWidth*groupOffsetFactor
		drawBarGroup(
			writer,
			layout,
			groupStart,
			barWidth,
			categoryIndex,
			series,
			yMax,
		)
	}
}

func drawBarGroup(
	writer *bufio.Writer,
	layout chartLayout,
	groupStart float64,
	barWidth float64,
	categoryIndex int,
	series []svgSeries,
	yMax float64,
) {
	for seriesIndex, currentSeries := range series {
		value := currentSeries.Values[categoryIndex]
		barHeight := (value / yMax) * layout.PlotHeight
		xValue := groupStart + float64(seriesIndex)*barWidth
		yValue := layout.Top + layout.PlotHeight - barHeight

		_, _ = fmt.Fprintf(
			writer,
			"  <rect x=\"%.2f\" y=\"%.2f\" width=\"%.2f\" "+
				"height=\"%.2f\" fill=\"%s\"/>\n",
			xValue,
			yValue,
			barWidth,
			barHeight,
			currentSeries.Color,
		)
	}
}

func drawLegend(
	writer *bufio.Writer,
	series []svgSeries,
	startX float64,
	startY float64,
) {
	legendY := startY

	for _, currentSeries := range series {
		_, _ = fmt.Fprintf(
			writer,
			"  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" "+
				"stroke=\"%s\" stroke-width=\"3\"/>\n",
			startX,
			legendY,
			startX+legendLineWidth,
			legendY,
			currentSeries.Color,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4\" fill=\"%s\"/>\n",
			startX+legendCircleOffset,
			legendY,
			currentSeries.Color,
		)
		_, _ = fmt.Fprintf(
			writer,
			"  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" fill=\"#333\" "+
				"font-family=\"%s\">%s</text>\n",
			startX+legendTextOffset,
			legendY+gridOffsetY,
			svgFontFamily,
			currentSeries.Name,
		)

		legendY += legendItemHeight
	}
}

func chartYMax(series []svgSeries) float64 {
	maxValue := maxSeriesValue(series) * yAxisPaddingFactor
	if maxValue <= zeroFloat {
		return defaultChartMax
	}

	return maxValue
}

func maxSeriesValue(series []svgSeries) float64 {
	maxValue := zeroFloat

	for _, currentSeries := range series {
		for _, value := range currentSeries.Values {
			maxValue = math.Max(maxValue, value)
		}
	}

	if maxValue == zeroFloat {
		return defaultChartMax
	}

	return maxValue
}

func writeReport(path string, metrics map[string]metric) error {
	err := ensureDir(filepath.Dir(path))
	if err != nil {
		return fmt.Errorf("create report dir: %w", err)
	}

	content := strings.Join(reportLines(metrics), "\n") + "\n"

	err = writeFile(path, []byte(content))
	if err != nil {
		return fmt.Errorf("write report: %w", err)
	}

	return nil
}

func reportLines(metrics map[string]metric) []string {
	lines := make([]string, zeroInt, reportLineCap)
	lines = append(lines, reportIntroSection()...)
	lines = append(lines, reportChartSection()...)
	lines = append(lines, reportKeyNumbersSection(metrics)...)
	lines = append(lines, reportAnalysisSection()...)

	return lines
}

func reportIntroSection() []string {
	return []string{
		reportTitle,
		reportBlankLine,
		reportIntro,
		reportBlankLine,
		"## How To Reproduce",
		reportBlankLine,
		"```sh",
		"go run ./scripts/benchreport.go",
		"```",
		reportBlankLine,
		"## Scope",
		reportBlankLine,
		"- Core constructors: `DateTime`, `ParentIdTag`, `StartTransactionReq`",
		"- Scaling path #1: `SendLocalListReq` (1 to 1000 entries)",
		"- Scaling path #2: `GetConfigurationReq` (1 to 1000 keys)",
		"- Metrics: `ns/op`, `B/op`, `allocs/op`",
		reportBlankLine,
	}
}

func reportChartSection() []string {
	return []string{
		"## Charts",
		reportBlankLine,
		"### 1) SendLocalListReq Scaling (ns/op)",
		reportBlankLine,
		"![SendLocalListReq ns/op](img/sendlocallist_ns.svg)",
		reportBlankLine,
		"### 2) SendLocalListReq Scaling (B/op)",
		reportBlankLine,
		"![SendLocalListReq B/op](img/sendlocallist_bytes.svg)",
		reportBlankLine,
		"### 3) SendLocalListReq Scaling (allocs/op)",
		reportBlankLine,
		"![SendLocalListReq allocs/op](img/sendlocallist_allocs.svg)",
		reportBlankLine,
		"### 4) GetConfigurationReq Scaling (ns/op)",
		reportBlankLine,
		"![GetConfigurationReq ns/op](img/getconfiguration_ns.svg)",
		reportBlankLine,
		"### 5) Core Constructors and Message Path (ns/op)",
		reportBlankLine,
		"![Core constructors ns/op](img/core_constructors_ns.svg)",
		reportBlankLine,
		"### 6) Custom / PrimitiveValidated Ratio (ns/op)",
		reportBlankLine,
		"![Custom vs PrimitiveValidated ratio]" +
			"(img/custom_vs_validated_ratio.svg)",
		reportBlankLine,
	}
}

func reportKeyNumbersSection(metrics map[string]metric) []string {
	return []string{
		"## Key Numbers",
		reportBlankLine,
		"| Case | Custom ns/op | PrimitiveValidated ns/op | Ratio |",
		"| ---- | -----------: | -----------------------: | ----: |",
		formatRatioRow(
			"SendLocalListReq_1000",
			lookupMetric(metrics, "BenchmarkSendLocalListReq_Custom_1000").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkSendLocalListReq_PrimitiveValidated_1000",
			).NsOp,
		),
		formatRatioRow(
			"GetConfigurationReq_1000",
			lookupMetric(
				metrics,
				"BenchmarkGetConfigurationReq_Custom_1000",
			).NsOp,
			lookupMetric(
				metrics,
				"BenchmarkGetConfigurationReq_PrimitiveValidated_1000",
			).NsOp,
		),
		formatRatioRow(
			benchStartTransaction,
			lookupMetric(metrics, "BenchmarkStartTransactionReq_Custom").NsOp,
			lookupMetric(
				metrics,
				"BenchmarkStartTransactionReq_PrimitiveValidated",
			).NsOp,
		),
		reportBlankLine,
	}
}

func reportAnalysisSection() []string {
	return []string{
		"## Analysis",
		reportBlankLine,
		reportAnalysisA,
		reportAnalysisB,
		reportAnalysisC,
		reportAnalysisD,
		reportBlankLine,
		"## Conclusion",
		reportBlankLine,
		reportConclusionA,
		reportConclusionB,
	}
}

func formatRatioRow(name string, customNS float64, validatedNS float64) string {
	return fmt.Sprintf(
		"| %s | %.2f | %.2f | %.2fx |",
		name,
		customNS,
		validatedNS,
		ratio(customNS, validatedNS),
	)
}

func benchmarkVariants() []string {
	return []string{
		variantCustom,
		variantPrimitiveDirect,
		variantPrimitiveValidated,
	}
}

func benchmarkColors() []string {
	return []string{"#d1495b", "#2a9d8f", "#1f77b4"}
}

func dateTimeBenchmarkName(variant string) string {
	switch variant {
	case variantPrimitiveDirect:
		return "BenchmarkDateTime_PrimitiveDirect"
	case variantPrimitiveValidated:
		return "BenchmarkDateTime_PrimitiveValidated"
	default:
		return "BenchmarkDateTime_CustomType"
	}
}

func parentTagBenchmarkName(variant string) string {
	switch variant {
	case variantPrimitiveDirect:
		return "BenchmarkParentIdTag_PrimitiveDirect"
	case variantPrimitiveValidated:
		return "BenchmarkParentIdTag_PrimitiveValidated"
	default:
		return "BenchmarkParentIdTag_CustomChain"
	}
}

func ensureDir(path string) error {
	cleanPath := filepath.Clean(path)

	err := os.MkdirAll(cleanPath, fileModeDir)
	if err != nil {
		return fmt.Errorf("mkdir all %s: %w", cleanPath, err)
	}

	return nil
}

func openFileRead(path string) (*os.File, error) {
	cleanPath := filepath.Clean(path)

	// #nosec G304 -- local maintenance script reads an explicit CLI path.
	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", cleanPath, err)
	}

	return file, nil
}

func createFileWrite(path string) (*os.File, error) {
	err := ensureDir(filepath.Dir(path))
	if err != nil {
		return nil, err
	}

	cleanPath := filepath.Clean(path)

	// #nosec G304 -- local maintenance script writes an explicit CLI path.
	file, err := os.OpenFile(
		cleanPath,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		fileModeRead,
	)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", cleanPath, err)
	}

	return file, nil
}

func writeFile(path string, content []byte) error {
	file, err := createFileWrite(path)
	if err != nil {
		return err
	}

	defer closeFile(file)

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}

	return nil
}

func closeFile(file *os.File) {
	_ = file.Close()
}
