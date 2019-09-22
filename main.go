package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/rodrigopl13/viajante-genetico/genetico"
	"github.com/rodrigopl13/viajante-genetico/plano"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"strconv"
	"time"
)

type application struct {
	path           *canvas.Image
	history        *canvas.Image
	window         fyne.Window
	generacion     *genetico.Generation
	bestChromosome []int
	bestDistance   float64
	countX         int
	xy             plotter.XYs
	plot           *plot.Plot
	labelBest      *widget.Label
}

func main() {
	a := app.New()
	g := application{
		path: &canvas.Image{
			FillMode: canvas.ImageFillOriginal,
		},
		window: a.NewWindow("Problema Viajero"),
		history: &canvas.Image{
			FillMode: canvas.ImageFillOriginal,
		},
		bestDistance: math.MaxFloat64,
		labelBest:    widget.NewLabel(fmt.Sprintf("%.3f", 0.0)),
	}
	g.window.SetContent(
		widget.NewVBox(
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				g.path,
				g.history,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(3),
				widget.NewButton("Iniciar Applicacion", g.startApp),
				widget.NewLabel("Best Distance"),
				g.labelBest,
			),
		),
	)
	g.window.ShowAndRun()
}

func (a *application) startApp() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	a.plot = p

	p.Title.Text = "Historic"
	p.Y.Label.Text = "distance"
	p.X.Label.Text = "Generation"

	a.path.File = "image/graph.png"
	a.history.File = "image/historic.png"

	a.xy = plotter.XYs{}
	a.generacion = genetico.NewGeneration(100, 20, plano.CreateCities())
	time.Sleep(1 * time.Second)
	a.createGraph(0)
	for i := 0; i < 100; i++ {
		a.generacion = genetico.NextGeneration(a.generacion)
		time.Sleep(1 * time.Second)
		a.createGraph(i + 1)
	}
	time.Sleep(1 * time.Second)
	a.createGraphBestAll()
}

func (a *application) getBestChromosome() []int {
	minDistance := math.MaxFloat64
	index := 0
	for i, v := range a.generacion.Distance {
		if minDistance > v {
			index = i
			minDistance = v
		}
	}
	if a.bestDistance > minDistance {
		a.bestChromosome = a.generacion.Population[index]
		a.bestDistance = minDistance
	}

	fmt.Println("BEST >>>>>>>>>>>>", a.generacion.Population[index], "distance: "+fmt.Sprintf("%.3f", minDistance))
	return a.generacion.Population[index]
}

func (a *application) createGraph(i int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Generacion " + strconv.Itoa(i)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	err = plotutil.AddLinePoints(p, a.createPoints(a.getBestChromosome()))
	if err != nil {
		panic(err)
	}
	points := plotter.XY{}
	points.X = float64(i)
	points.Y = a.bestDistance

	a.xy = append(a.xy, points)
	err = plotutil.AddLinePoints(a.plot, a.xy)
	if err != nil {
		panic(err)
	}

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "image/graph.png"); err != nil {
		panic(err)
	}

	if err = a.plot.Save(7*vg.Inch, 7*vg.Inch, "image/historic.png"); err != nil {
		panic(err)
	}

	canvas.Refresh(a.path)
	canvas.Refresh(a.history)
	a.labelBest.SetText(fmt.Sprintf("%.3f", a.bestDistance))
}

func (a *application) createGraphBestAll() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Best of all Generations "
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	err = plotutil.AddLinePoints(p, a.createPoints(a.bestChromosome))
	if err != nil {
		panic(err)
	}
	if err := p.Save(7*vg.Inch, 7*vg.Inch, "image/graph.png"); err != nil {
		panic(err)
	}
	canvas.Refresh(a.path)

}

func (a *application) createPoints(chromosome []int) plotter.XYs {
	pts := make(plotter.XYs, len(chromosome))
	for i := range pts {
		pts[i].X = a.generacion.Cities[chromosome[i]].X
		pts[i].Y = a.generacion.Cities[chromosome[i]].Y

	}
	return pts
}
