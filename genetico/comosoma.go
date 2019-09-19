package genetico

import (
	"github.com/rodrigopl13/viajante-genetico/plano"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Generation struct {
	Population  [][]int32
	Coordenates plano.Distribution
	Distance    []float64
}

func NewGeneration(population, sizeChromosome int32, distribution plano.Distribution) *Generation {
	var wg sync.WaitGroup
	p := make([][]int32, population)
	wg.Add(len(p))
	for i := range p {
		go randomChromosome(p, int32(i), sizeChromosome, &wg)
	}

	wg.Wait()
	d := make([]float64, population)
	for i := range p {
		go calculateDistanceChromosome(p[i], d, i, distribution)
	}
	//wg.Wait()
	return &Generation{
		Population:  p,
		Coordenates: distribution,
		Distance:    d,
	}
}

func NextGeneration(current *Generation) {

}

func randomChromosome(p [][]int32, chromosome, sizeChromosome int32, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(map[int32]bool)
	var random int32
	p[chromosome] = make([]int32, sizeChromosome)
	for i := 0; i < len(p[chromosome]); i++ {
		rand.Seed(time.Now().UnixNano())
		random = rand.Int31n(sizeChromosome) + 1
		for c[random] {
			rand.Seed(time.Now().UnixNano())
			random = rand.Int31n(sizeChromosome) + 1
		}
		c[random] = true
		p[chromosome][i] = random
	}
}

func calculateDistanceChromosome(
	chromosome []int32,
	distance []float64,
	index int,
	distribution plano.Distribution,
) {
	if len(chromosome) < 2 {
		distance[index] = 0
		return
	}
	var d float64
	d = calculateDistance(
		distribution.Cities[chromosome[0]].X,
		distribution.Cities[chromosome[0]].Y,
		distribution.Cities[chromosome[1]].X,
		distribution.Cities[chromosome[1]].Y,
	)

	if len(chromosome) > 2 {
		for i := 1; i < len(chromosome)-1; i++ {
			d += calculateDistance(
				distribution.Cities[chromosome[i]].X,
				distribution.Cities[chromosome[i]].Y,
				distribution.Cities[chromosome[i+1]].X,
				distribution.Cities[chromosome[i+1]].Y,
			)
		}
	}
	distance[index] = d
}

func calculateDistance(x1, y1, x2, y2 int32) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(x2-x1)), 2) + math.Pow(math.Abs(float64(y2-y1)), 2))
}

func inversion(a []int) {
	rand.Seed(time.Now().UnixNano())
	start := rand.Intn(len(a))
	size := rand.Intn(len(a) - 2)
	i := 0
	var j, k, tmp int
	for i <= (size / 2) {
		j = start + i
		if j > len(a)-1 {
			j = j - len(a)
		}
		k = start + size - i
		if k > len(a)-1 {
			k = k - len(a)
		}
		tmp = a[k]
		a[k] = a[j]
		a[j] = tmp
		i++
	}
}

func intercambio() {

}
