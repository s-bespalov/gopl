// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/s-bespalov/gopl/ch2/lenconv"
	"github.com/s-bespalov/gopl/ch2/tempconv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			process(arg)
		}
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	process(text)
}

func process(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	// convert temperature
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", k, tempconv.KToC(k), c, tempconv.CToK(c))
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToK(f), k, tempconv.KToF(k))

	// convert length
	m := lenconv.Meter(t)
	fe := lenconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), fe, lenconv.FToM(fe))
}
