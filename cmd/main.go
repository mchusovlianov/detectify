package main

import "github.com/mchusovlianov/detectify/internal/detector"

func main() {
	det := detector.NewService(1)

	err := det.Run()
	if err != nil {
		panic(err)
	}
}
