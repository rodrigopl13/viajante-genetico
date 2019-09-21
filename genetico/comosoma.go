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
	var wg sync.WaitGroup
	p := make([][]int, population)
	for i := range p {
		wg.Add(1)
		go randomChromosome(p, i, sizeChromosome, &wg)
	}

	wg.Wait()
	//var wg2 sync.WaitGroup
	d := make([]float64, population)
	for i := range p {
		//wg2.Add(1)
		go calculateDistanceChromosome(p[i], d, i, distribution)
	}
	//wg2.Wait()
	return &Generation{
		Population: p,
		Cities:     distribution,
		Distance:   d,
	}
}

func NextGeneration(currentGeneration *Generation) *Generation {
	population := len(currentGeneration.Population)
	var wg sync.WaitGroup
	ng := &Generation{
		Population: make([][]int, population),
		Cities:     currentGeneration.Cities,
		Distance:   make([]float64, population),
	}

	for i := 0; i < population; i++ {
		wg.Add(1)
		go competeChromosomes(currentGeneration, ng, i, 0.5, &wg)
	}

	wg.Wait()
	for i := range ng.Population {
		if i%2 == 0 {
			Intercambio(ng.Population[i])
		} else {
			Inversion(ng.Population[i])
		}
	}
	//var wg2 sync.WaitGroup
	for i := range ng.Population {
		//wg2.Add(1)
		go calculateDistanceChromosome(ng.Population[i], ng.Distance, i, ng.Cities)
	}
	//wg2.Wait()

	return ng
}

func competeChromosomes(
	currentGeneration,
	newGeneration *Generation,
	position int,
	percentaje float32,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
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
}

func randomChromosome(p [][]int, chromosome, sizeChromosome int, wg *sync.WaitGroup) {
	defer wg.Done()
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
}

func calculateDistanceChromosome(
	chromosome []int,
	distance []float64,
	index int,
	cities plano.Cities,
	//wg *sync.WaitGroup,
) {
	//defer wg.Wait()
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
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x2-x1), 2) + math.Pow(math.Abs(y2-y1), 2))
}

func Inversion(a []int) {
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

func Intercambio(a []int) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn((len(a) / 2))

	rand.Seed(time.Now().UnixNano())
	pos1 := rand.Intn((len(a) / 2) - size)

	rand.Seed(time.Now().UnixNano())
	pos2 := rand.Intn((len(a)/2)-size) + (len(a)/2 - 1)

	var tmp int
	for i := 0; i <= size; i++ {
		tmp = a[pos1+i]
		a[pos1+i] = a[pos2+i]
		a[pos2+i] = tmp
	}
}
