package main

import (
	"bufio"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var benchLine = regexp.MustCompile(`^(Benchmark\S+)-\d+\s+\d+\s+([0-9.]+)\s+ns/op`)

type benchResult struct {
	Name string
	NsOp float64
}

type point struct {
	X float64
	Y float64
}

func main() {
	inPath := flag.String("in", "reports/analysis_bench.txt", "input benchmark output")
	outPath := flag.String("out", "reports/analysis_bench_line.svg", "output svg line chart")
	flag.Parse()

	results, err := parseBenchFile(*inPath)
	if err != nil {
		exitErr(err)
	}

	series := buildVariableInputSeries(results)
	if len(series) == 0 {
		exitErr(errors.New("no variable-input benchmark series found"))
	}

	if err := writeSVG(*outPath, series); err != nil {
		exitErr(err)
	}

	fmt.Printf("Created %s\n", *outPath)
}

func parseBenchFile(path string) ([]benchResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input: %w", err)
	}
	defer func() { _ = file.Close() }()

	results := make([]benchResult, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := benchLine.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}

		nsOp, err := strconv.ParseFloat(match[2], 64)
		if err != nil {
			continue
		}

		results = append(results, benchResult{Name: match[1], NsOp: nsOp})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan input: %w", err)
	}

	if len(results) == 0 {
		return nil, errors.New("no benchmark lines parsed")
	}

	return results, nil
}

func buildVariableInputSeries(results []benchResult) map[string]map[int]float64 {
	series := map[string]map[int]float64{
		"Custom":             {},
		"PrimitiveDirect":    {},
		"PrimitiveValidated": {},
	}

	for _, result := range results {
		name := strings.TrimPrefix(result.Name, "Benchmark")
		if !strings.HasPrefix(name, "SendLocalListReq_") {
			continue
		}

		parts := strings.Split(name, "_")
		if len(parts) < 3 {
			continue
		}

		sizeStr := parts[len(parts)-1]
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			continue
		}

		variant := parts[len(parts)-2]
		switch variant {
		case "Custom":
			series["Custom"][size] = result.NsOp
		case "PrimitiveDirect":
			series["PrimitiveDirect"][size] = result.NsOp
		case "PrimitiveValidated":
			series["PrimitiveValidated"][size] = result.NsOp
		}
	}

	for key, values := range series {
		if len(values) == 0 {
			delete(series, key)
		}
	}

	return series
}

func writeSVG(path string, series map[string]map[int]float64) error {
	const (
		width       = 980.0
		height      = 560.0
		leftMargin  = 90.0
		rightMargin = 40.0
		topMargin   = 70.0
		bottom      = 90.0
	)

	xMin := 1.0
	xMax := 100.0
	yMin := 0.0
	yMax := maxY(series) * 1.10
	if yMax <= 0 {
		yMax = 1
	}

	plotW := width - leftMargin - rightMargin
	plotH := height - topMargin - bottom

	xToPx := func(x float64) float64 {
		return leftMargin + ((x-xMin)/(xMax-xMin))*plotW
	}

	yToPx := func(y float64) float64 {
		return topMargin + (1-((y-yMin)/(yMax-yMin)))*plotH
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer func() { _ = file.Close() }()

	writer := bufio.NewWriter(file)
	defer func() { _ = writer.Flush() }()

	_, _ = fmt.Fprintf(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	_, _ = fmt.Fprintf(writer,
		"<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%.0f\" height=\"%.0f\" viewBox=\"0 0 %.0f %.0f\">\n",
		width,
		height,
		width,
		height,
	)
	_, _ = fmt.Fprintln(writer, "  <rect x=\"0\" y=\"0\" width=\"100%\" height=\"100%\" fill=\"white\"/>")
	_, _ = fmt.Fprintln(writer, "  <text x=\"490\" y=\"34\" font-size=\"22\" text-anchor=\"middle\" font-family=\"Helvetica,Arial,sans-serif\">SendLocalListReq: Variable Input Benchmark</text>")
	_, _ = fmt.Fprintln(writer, "  <text x=\"490\" y=\"56\" font-size=\"13\" text-anchor=\"middle\" fill=\"#555\" font-family=\"Helvetica,Arial,sans-serif\">ns/op (lower is better) by input size</text>")

	drawAxes(writer, leftMargin, topMargin, plotW, plotH, xToPx, yToPx, yMax)
	drawSeries(writer, series, xToPx, yToPx)
	drawLegend(writer)

	_, _ = fmt.Fprintln(writer, "</svg>")

	return nil
}

func drawAxes(
	writer *bufio.Writer,
	left,
	top,
	plotW,
	plotH float64,
	xToPx func(float64) float64,
	yToPx func(float64) float64,
	yMax float64,
) {
	_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n", left, top+plotH, left+plotW, top+plotH)
	_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\" stroke-width=\"1.5\"/>\n", left, top, left, top+plotH)

	xTicks := []int{1, 25, 100}
	for _, tick := range xTicks {
		x := xToPx(float64(tick))
		_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#ddd\"/>\n", x, top, x, top+plotH)
		_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#222\"/>\n", x, top+plotH, x, top+plotH+6)
		_, _ = fmt.Fprintf(writer, "  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"middle\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%d</text>\n", x, top+plotH+24, tick)
	}

	yTicks := 6
	for i := 0; i <= yTicks; i++ {
		value := (float64(i) / float64(yTicks)) * yMax
		y := yToPx(value)
		_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"#eee\"/>\n", left, y, left+plotW, y)
		_, _ = fmt.Fprintf(writer, "  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" text-anchor=\"end\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%.0f</text>\n", left-8, y+4, value)
	}

	_, _ = fmt.Fprintf(writer, "  <text x=\"%.2f\" y=\"%.2f\" font-size=\"13\" text-anchor=\"middle\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">Input size (entries)</text>\n", left+plotW/2, top+plotH+54)
	_, _ = fmt.Fprintf(writer, "  <text x=\"26\" y=\"%.2f\" font-size=\"13\" text-anchor=\"middle\" fill=\"#333\" transform=\"rotate(-90 26 %.2f)\" font-family=\"Helvetica,Arial,sans-serif\">ns/op</text>\n", top+plotH/2, top+plotH/2)
}

func drawSeries(
	writer *bufio.Writer,
	series map[string]map[int]float64,
	xToPx func(float64) float64,
	yToPx func(float64) float64,
) {
	colors := map[string]string{
		"Custom":             "#d1495b",
		"PrimitiveDirect":    "#2a9d8f",
		"PrimitiveValidated": "#1f77b4",
	}

	order := []string{"Custom", "PrimitiveDirect", "PrimitiveValidated"}
	sizes := []int{1, 25, 100}

	for _, variant := range order {
		values, ok := series[variant]
		if !ok {
			continue
		}

		points := make([]point, 0, len(sizes))
		for _, size := range sizes {
			value, exists := values[size]
			if !exists {
				continue
			}
			points = append(points, point{X: xToPx(float64(size)), Y: yToPx(value)})
		}

		if len(points) == 0 {
			continue
		}

		_, _ = fmt.Fprintf(
			writer,
			"  <polyline points=\"%s\" fill=\"none\" stroke=\"%s\" stroke-width=\"3\"/>\n",
			polylinePoints(points),
			colors[variant],
		)

		for _, pt := range points {
			_, _ = fmt.Fprintf(writer, "  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4.3\" fill=\"%s\"/>\n", pt.X, pt.Y, colors[variant])
		}
	}
}

func drawLegend(writer *bufio.Writer) {
	entries := []struct {
		Label string
		Color string
	}{
		{Label: "Custom", Color: "#d1495b"},
		{Label: "PrimitiveDirect", Color: "#2a9d8f"},
		{Label: "PrimitiveValidated", Color: "#1f77b4"},
	}

	x := 650.0
	y := 86.0
	for _, entry := range entries {
		_, _ = fmt.Fprintf(writer, "  <line x1=\"%.2f\" y1=\"%.2f\" x2=\"%.2f\" y2=\"%.2f\" stroke=\"%s\" stroke-width=\"3\"/>\n", x, y, x+30, y, entry.Color)
		_, _ = fmt.Fprintf(writer, "  <circle cx=\"%.2f\" cy=\"%.2f\" r=\"4\" fill=\"%s\"/>\n", x+15, y, entry.Color)
		_, _ = fmt.Fprintf(writer, "  <text x=\"%.2f\" y=\"%.2f\" font-size=\"12\" fill=\"#333\" font-family=\"Helvetica,Arial,sans-serif\">%s</text>\n", x+40, y+4, xmlEscape(entry.Label))
		y += 20
	}
}

func polylinePoints(points []point) string {
	parts := make([]string, 0, len(points))
	for _, pt := range points {
		parts = append(parts, fmt.Sprintf("%.2f,%.2f", pt.X, pt.Y))
	}
	return strings.Join(parts, " ")
}

func maxY(series map[string]map[int]float64) float64 {
	max := 0.0
	for _, values := range series {
		for _, v := range values {
			max = math.Max(max, v)
		}
	}
	return max
}

func xmlEscape(s string) string {
	buffer := &strings.Builder{}
	_ = xml.EscapeText(buffer, []byte(s))
	return buffer.String()
}

func exitErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
