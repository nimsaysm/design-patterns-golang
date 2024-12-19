package main

import "fmt"

// 4. create factory struct that creates cars of desired type based on argument

func getCar(carType string) (ICar, error) {
	switch carType {
	case "Bmw":
		return newBmw(), nil
	case "Ferrari":
		return newFerrari(), nil
	default:
		return nil, fmt.Errorf("this option does not exist")
	}
}
