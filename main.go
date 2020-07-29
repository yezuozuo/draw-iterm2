package draw_iterm2

import (
	"os"
	"io/ioutil"
	"image"
	"bytes"
	"log"
	"github.com/mattn/go-sixel"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"github.com/benoitmasson/plotters/piechart"
)

const WIDTH = 4
const HEIGHT = 4

func Pie(calArr []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.HideAxes()

	values := plotter.Values{}
	values = calArr

	pie, err := piechart.NewPieChart(values)
	if err != nil {
		panic(err)
	}
	pie.Color = color.RGBA{255, 0, 0, 255}
	p.Add(pie)

	fileName := "/tmp/pie.png"
	p.Save(WIDTH*vg.Inch, HEIGHT*vg.Inch, fileName)

	outputCli(fileName)
}

func PieAdvance() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Legend.Top = true
	p.HideAxes()

	pie1, err := piechart.NewPieChart(plotter.Values{1, 2})
	if err != nil {
		panic(err)
	}
	pie1.Color = color.RGBA{255, 0, 0, 255}
	pie1.Total = 12
	pie1.Labels.Nominal = []string{"one", "two"}
	pie1.Labels.Values.Show = true
	pie1.Labels.Values.Percentage = true
	p.Add(pie1)
	p.Legend.Add("sample 1", pie1)

	pie2, err := piechart.NewPieChart(plotter.Values{3, 2})
	if err != nil {
		panic(err)
	}
	pie2.Color = color.RGBA{0, 255, 0, 255}
	pie2.Offset.Value = 3
	pie2.Total = 12
	pie2.Labels.Nominal = []string{"three", "four"}
	pie2.Labels.Values.Show = true
	pie2.Labels.Values.Percentage = true
	p.Add(pie2)
	p.Legend.Add("sample 2", pie2)

	pie3, err := piechart.NewPieChart(plotter.Values{4})
	if err != nil {
		panic(err)
	}
	pie3.Color = color.RGBA{0, 0, 255, 255}
	pie3.Offset.Value = 8
	pie3.Total = 12
	pie3.Offset.X = vg.Length(10)
	pie3.Offset.Y = vg.Length(-15)
	pie3.Labels.Position = 1.1
	pie3.Labels.Nominal = []string{"five"}
	pie3.Labels.Values.Show = true
	pie3.Labels.Values.Percentage = true
	p.Add(pie3)
	p.Legend.Add("sample 3", pie3)

	fileName := "/tmp/pie.png"
	p.Save(WIDTH*vg.Inch, HEIGHT*vg.Inch, fileName)

	outputCli(fileName)
}

func Line(calArr []float64) {
	p, _ := plot.New()

	p.Title.Text = "pic"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	XYS2 := make([]plotter.XY, len(calArr))

	for k, v := range calArr {
		item := new(plotter.XY)
		*item = plotter.XY{float64(k), float64(v)}
		XYS2[k] = *item
	}

	XYS := plotter.XYs{}
	XYS = XYS2

	plotutil.AddLinePoints(p, XYS)

	fileName := "/tmp/line.png"

	p.Save(WIDTH*vg.Inch, HEIGHT*vg.Inch, fileName)

	outputCli(fileName)
}

func Histogram(calArr []float64) {

	p, _ := plot.New()
	p.Title.Text = "pic"

	calArrLen := len(calArr)

	XYS2 := make([]plotter.XY, calArrLen)

	for k, v := range calArr {
		item := new(plotter.XY)
		*item = plotter.XY{float64(k), float64(v)}
		XYS2[k] = *item
	}

	bins := plotter.XYs{}
	bins = XYS2

	h, _ := plotter.NewHistogram(bins, calArrLen)
	p.Add(h)

	fileName := "/tmp/histogram.png"
	p.Save(WIDTH*vg.Inch, HEIGHT*vg.Inch, fileName)

	outputCli(fileName)
}

func outputCli(fileName string) {

	got, _ := ioutil.ReadFile(fileName)
	img, _, err := image.Decode(bytes.NewReader(got))

	if err != nil {
		log.Fatalf("Cannot decode image from stdin: %v", err)
	}

	err = sixel.NewEncoder(os.Stdout).Encode(img)
	if err != nil {
		log.Fatalf("Cannot encode image into sixel format: %v", err)
	}
}
