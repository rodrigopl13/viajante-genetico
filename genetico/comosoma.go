package genetico

import (
	"github.com/rodrigopl13/viajante-genetico/plano"
	"math/rand"
)

type Genetic struct{
	Population [][]int32
	Coordenates plano.Distribution
}


func NewPopulation(population, sizeChromosome int32, distribution plano.Distribution) *Genetic {

	p := make([][]int32, population)
	for i := range p {
		go randomChromosome(p, int32(i), sizeChromosome)
	}

	for i := range p{
		go calculateDistance(p[i], distribution)
	}

	return &Genetic{
		Population:  p,
		Coordenates: distribution,
	}
}

func randomChromosome(p [][]int32, chromosome, sizeChromosome int32) {
	c := make(map[int32]bool)
	var random int32
	p[chromosome] = make([]int32, sizeChromosome+1)
	for i := 0; i<len(p[chromosome])-1; i++{
		random = rand.Int31n(sizeChromosome) + 1
		for c[random] {
			random = rand.Int31n(sizeChromosome) + 1
		}
		c[random] = true
		p[chromosome][i] = random
	}
}

func calculateDistance(chromosome []int32, distribution plano.Distribution){
	//var distance int32
	for i:=0; i<len(chromosome) -1; i++ {
		// d := mat.Sqrt(math.Pow())
	}

}



