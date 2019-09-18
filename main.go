package main

import (
	"fmt"
	"github.com/rodrigopl13/viajante-genetico/genetico"
	"github.com/rodrigopl13/viajante-genetico/plano"
)

func main() {
	p := genetico.NewPopulation(100, 20, plano.CreateCities())
	printPopulation(p.Population)
}

func printPopulation(p [][]int32){
	for i := range p {
		fmt.Printf("%3d :: ",i+1)
		for j := range p[i]{
			fmt.Printf("%d \t", p[i][j])
		}
		fmt.Println()
	}
}