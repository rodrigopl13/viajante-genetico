package genetico

import (
	"github.com/rodrigopl13/viajante-genetico/plano"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Generation struct {
	Population [][]int
	Cities     plano.Cities
	Distance   []float64
}

func NewGeneration(population, sizeChromosome int, distribution plano.Cities) *Generation {
	ng := &Generation{
		Population: make([][]int, population),
		Cities:     distribution,
		Distance:   make([]float64, population),
	}
	generateRandom(ng, sizeChromosome)
	generateDistance(ng)
	return ng
}

func NextGeneration(currentGeneration *Generation) *Generation {
	population := len(currentGeneration.Population)
	ng := &Generation{
		Population: make([][]int, population),
		Cities:     currentGeneration.Cities,
		Distance:   make([]float64, population),
	}

	generateCompete(currentGeneration, ng)

	generateOperations(ng)

	generateDistance(ng)
	return ng
}

func generateRandom(ng *Generation, sizeChromosome int) {
	var wg sync.WaitGroup
	for i := range ng.Population {
		wg.Add(1)
		go randomChromosome(ng.Population, i, sizeChromosome, &wg)
	}

	wg.Wait()
}

func generateCompete(currentGeneration, newGeneration *Generation) {
	population := len(currentGeneration.Population)
	var wg sync.WaitGroup

	for i := 0; i < population; i++ {
		wg.Add(1)
		go competeChromosomes(currentGeneration, newGeneration, i, 0.5, &wg)
	}

	wg.Wait()
}

func generateOperations(newGeneration *Generation) {
	var wg sync.WaitGroup
	countInversion := 0
	var r int
	for i := range newGeneration.Population {
		wg.Add(1)
		rand.Seed(time.Now().UnixNano())
		r = rand.Intn(2)
		if r == 0 && countInversion <= 50 {
			go Inversion(newGeneration.Population[i], &wg)
			countInversion++
		} else {
			go Intercambio(newGeneration.Population[i], &wg)
		}
	}
	wg.Wait()
}

func generateDistance(newGeneration *Generation) {
	var wg sync.WaitGroup
	for i := range newGeneration.Population {
		wg.Add(1)
		go calculateDistanceChromosome(
			newGeneration.Population[i],
			newGeneration.Distance,
			i,
			newGeneration.Cities,
			&wg,
		)
	}
	wg.Wait()
}

func competeChromosomes(
	currentGeneration,
	newGeneration *Generation,
	position int,
	percentaje float32,
	wg *sync.WaitGroup,
) {
	population := len(currentGeneration.Population)
	p := float32(population) * percentaje
	minDistance := math.MaxFloat64
	var randomIndex, minIndex int
	for i := 0; i < int(p); i++ {
		rand.Seed(time.Now().UnixNano())
		randomIndex = rand.Intn(population)
		if currentGeneration.Distance[randomIndex] < minDistance {
			minDistance = currentGeneration.Distance[randomIndex]
			minIndex = randomIndex
		}
	}
	sizeChromosome := len(currentGeneration.Population[minIndex])
	newChromosome := make([]int, sizeChromosome)
	copy(newChromosome, currentGeneration.Population[minIndex])
	newGeneration.Population[position] = newChromosome
	wg.Done()
}

func randomChromosome(p [][]int, chromosome, sizeChromosome int, wg *sync.WaitGroup) {
	c := make(map[int]bool)
	var random int
	p[chromosome] = make([]int, sizeChromosome)
	for i := 0; i < len(p[chromosome]); i++ {
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(sizeChromosome) + 1
		for c[random] {
			rand.Seed(time.Now().UnixNano())
			random = rand.Intn(sizeChromosome) + 1
		}
		c[random] = true
		p[chromosome][i] = random
	}
	wg.Done()
}

func calculateDistanceChromosome(
	chromosome []int,
	distance []float64,
	index int,
	cities plano.Cities,
	wg *sync.WaitGroup,
) {
	if len(chromosome) < 2 {
		distance[index] = 0
		return
	}
	var d float64
	d = calculateDistance(
		cities[chromosome[0]].X,
		cities[chromosome[0]].Y,
		cities[chromosome[1]].X,
		cities[chromosome[1]].Y,
	)

	if len(chromosome) > 2 {
		for i := 1; i < len(chromosome)-1; i++ {
			d += calculateDistance(
				cities[chromosome[i]].X,
				cities[chromosome[i]].Y,
				cities[chromosome[i+1]].X,
				cities[chromosome[i+1]].Y,
			)
		}
	}
	distance[index] = d
	wg.Done()
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x2-x1), 2) + math.Pow(math.Abs(y2-y1), 2))
}

func Inversion(a []int, wg *sync.WaitGroup) {
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
	wg.Done()
}

func Intercambio(a []int, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn((len(a) / 2)) + 1

	rand.Seed(time.Now().UnixNano())
	pos1 := rand.Intn((len(a)) - (size * 2) + 1)

	rand.Seed(time.Now().UnixNano())
	pos2 := rand.Intn(len(a)-pos1-(size*2)+1) + pos1 + size

	var tmp int
	for i := 0; i < size; i++ {
		tmp = a[pos1+i]
		a[pos1+i] = a[pos2+i]
		a[pos2+i] = tmp
	}
	wg.Done()
}
