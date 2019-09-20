package main

import (
	"fmt"
	"github.com/rodrigopl13/viajante-genetico/genetico"
	"github.com/rodrigopl13/viajante-genetico/plano"
)

func main() {
	p := genetico.NewGeneration(100, 20, plano.CreateCities())
	printPopulation(p.Population)
	printDistance(p.Distance)
}

func printPopulation(p [][]int) {
	for i := range p {
		fmt.Printf("%3d :: ", i+1)
		for j := range p[i] {
			fmt.Printf("%d \t", p[i][j])
		}
		fmt.Println()
	}
}

func printDistance(distance []float64) {
	for i := range distance {
		fmt.Printf("%3d :: %.3f\n", i+1, distance[i])

	}
}
