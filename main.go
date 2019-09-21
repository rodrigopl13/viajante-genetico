package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
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
	image      *canvas.Image
	window     fyne.Window
	generacion *genetico.Generation
}

func main() {

	a := app.New()

	g := application{
		image: &canvas.Image{
			FillMode: canvas.ImageFillOriginal,
		},
		window: a.NewWindow("Problema Viajero"),
	}

	g.window.SetContent(widget.NewVBox(g.image, widget.NewButton("Iniciar Applicacion", g.startApp)))

	g.window.ShowAndRun()

}

func (a *application) startApp() {
	a.generacion = genetico.NewGeneration(100, 20, plano.CreateCities())
	time.Sleep(1 * time.Second)
	a.createGraph(1)
	//printDistance(a.generacion.Distance)
	for i := 1; i <= 5; i++ {
		a.generacion = genetico.NextGeneration(a.generacion)
		time.Sleep(2 * time.Second)
		a.createGraph(i + 1)
		//printDistance(a.generacion.Distance)
	}
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
	fmt.Println("BEST >>>>>>>>>>>>", a.generacion.Population[index], "distance after: "+fmt.Sprintf("%.3f", minDistance))
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

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "image/graph.png"); err != nil {
		panic(err)
	}
	a.image.File = "image/graph.png"
	canvas.Refresh(a.image)
}

func (a *application) createPoints(chromosome []int) plotter.XYs {
	pts := make(plotter.XYs, len(chromosome))
	for i := range pts {
		pts[i].X = a.generacion.Cities[chromosome[i]].X
		pts[i].Y = a.generacion.Cities[chromosome[i]].Y

	}
	return pts
}

//func printPopulation(p [][]int) {
//	for i := range p {
//		fmt.Printf("%3d :: ", i+1)
//		for j := range p[i] {
//			fmt.Printf("%d \t", p[i][j])
//		}
//		fmt.Println()
//	}
//}
//
func printDistance(distance []float64) {
	for i := range distance {
		fmt.Printf("%3d :: %.3f\n", i+1, distance[i])

	}
}
