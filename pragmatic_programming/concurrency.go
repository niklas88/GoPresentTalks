package main

import (
	"fmt"
)

// START OMIT
func sumPartial(seq []float64, res chan float64) {
	var sum float64 // 0.0 is the zero (uninitalized) value
	for _, v := range seq {
		sum += v
	}
	res <- sum // HLconc
}

func sum(seq []float64) float64 {
	resultChan := make(chan float64) // HLconc
	defer close(resultChan)
	go sumPartial(seq[:len(seq)/2], resultChan) // HLconc
	go sumPartial(seq[len(seq)/2:], resultChan) // HLconc
	return <-resultChan + <-resultChan // HLconc
}

// END OMIT

func main() {
	sequence := []float64{1.0, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	fmt.Println("Sum:", sum(sequence))
}
