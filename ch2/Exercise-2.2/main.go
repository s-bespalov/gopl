package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/s-bespalov/gopl/ch2/Exercise-2.2/lenconv"
	"github.com/s-bespalov/gopl/ch2/Exercise-2.2/tempconv"
	"github.com/s-bespalov/gopl/ch2/Exercise-2.2/wconv"
)

func printT(o io.Writer, v float64) {
	c := tempconv.Celsius(v)
	f := tempconv.Fahrenheit(v)
	fmt.Fprintf(o, "%v = %v, %v = %v\n", c, tempconv.CToF(c), f, tempconv.FToC(f))
}

func printL(o io.Writer, v float64) {
	m := lenconv.Meter(v)
	f := lenconv.Feet(v)
	fmt.Fprintf(o, "%v = %v, %v = %v\n", m, lenconv.MToF(m), f, lenconv.FToM(f))
}

func printW(o io.Writer, v float64) {
	k := wconv.Kilogram(v)
	p := wconv.Pound(v)
	fmt.Fprintf(o, "%v = %v, %v = %v\n", k, wconv.KToP(k), p, wconv.PToK(p))
}

func printer(o io.Writer, v float64) {
	printT(o, v)
	printL(o, v)
	printW(o, v)
}

func main() {
	var input []float64
	if len(os.Args[1:]) == 0 {
		var v float64
		l, err := fmt.Scan(&v)
		log.Println("here", l, err)
		for err == nil && l > 0 {
			input = append(input, v)
			l, err = fmt.Scan(&v)
		}
	} else {
		input = make([]float64, len(os.Args)-1)
		for i, arg := range os.Args[1:] {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Fatalf("Exercise-2.2: %v\n", err)
			}
			input[i] = v
		}
	}

	for _, v := range input {
		printer(os.Stdout, v)
	}
}
