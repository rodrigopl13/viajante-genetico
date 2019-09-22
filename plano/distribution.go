package plano

type Coordenada struct {
	X float64
	Y float64
}

type Cities map[int]Coordenada

func CreateCities() Cities {
	d := map[int]Coordenada{
		1: {
			X: 3,
			Y: 4,
		},
		2: {
			X: 5,
			Y: 6,
		},
		3: {
			X: 9,
			Y: 7,
		},
		4: {
			X: 6,
			Y: 2,
		},
		5: {
			X: 9,
			Y: 1,
		},
		6: {
			X: 2,
			Y: 7,
		},
		7: {
			X: 3,
			Y: 1,
		},
		8: {
			X: 7,
			Y: 5,
		},
		9: {
			X: 1,
			Y: 5,
		},
		10: {
			X: 7,
			Y: 3,
		},
		11: {
			X: 9,
			Y: 5,
		},
		12: {
			X: 5,
			Y: 9,
		},
		13: {
			X: 4,
			Y: 2,
		},
		14: {
			X: 9,
			Y: 3,
		},
		15: {
			X: 2,
			Y: 3,
		},
		16: {
			X: 7,
			Y: 8,
		},
		17: {
			X: 3,
			Y: 8,
		},
		18: {
			X: 1,
			Y: 10,
		},
		19: {
			X: 9,
			Y: 9,
		},
		20: {
			X: 5,
			Y: 4,
		},
	}
	return d
}
