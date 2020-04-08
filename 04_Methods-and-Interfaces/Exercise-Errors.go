package main

import (
	"fmt"
	"log"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(2*z))) < 1e-8 {
			return z, nil
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}
}

func main() {
	// fmt.Println(Sqrt(2))
	if i, err := Sqrt(2); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%v\n", i)
	}
	// fmt.Println(Sqrt(-2))
	if i, err := Sqrt(-2); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%v", i)
	}
}
