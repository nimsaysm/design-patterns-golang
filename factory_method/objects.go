package main

// 3. create concrete objects (embed struct and implement interface methods)

type Bmw struct {
	Car
}

func newBmw() ICar {
	return &Bmw{
		Car: Car{
			name: "BMW car",
		},
	}
}

type Ferrari struct {
	Car
}

func newFerrari() ICar {
	return &Ferrari{
		Car: Car{
			name: "Ferrari car",
		},
	}
}
