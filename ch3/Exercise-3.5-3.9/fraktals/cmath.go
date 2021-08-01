package fraktals

import (
	"log"
	"math"
)

func Mul(real1, img1, real2, img2 interface{}) (real, img interface{}) {
	_, ok := real1.(float32)
	if ok {
		p := real1.(float32)
		q := img1.(float32)
		r := real2.(float32)
		s := img2.(float32)
		real = p*r - q*s
		img = p*s + q*r
		return
	}
	_, ok = real1.(float64)
	if ok {
		p := real1.(float64)
		q := img1.(float64)
		r := real2.(float64)
		s := img2.(float64)
		real = p*r - q*s
		img = p*s + q*r
		return
	}
	log.Println("something wrong in multiplication")
	return
}

func Sub(real1, img1, real2, img2 interface{}) (real, img interface{}) {
	// need refactoring
	f32r1, ok := real1.(float32)
	if ok {
		f32i1 := img1.(float32)
		f32r2 := real2.(float32)
		f32i2 := img2.(float32)
		real = f32r1 - f32r2
		img = f32i1 - f32i2
		return
	}
	f64r1, ok := real1.(float64)
	if ok {
		f64i1 := img1.(float64)
		f64r2 := real2.(float64)
		f64i2 := img2.(float64)
		real = f64r1 - f64r2
		img = f64i1 - f64i2
		return
	}
	log.Println("something wrong in substraction")
	return
}

func Div(real1, img1, real2, img2 interface{}) (real, img interface{}) {
	_, ok := real1.(float32)
	if ok {
		p := real1.(float32)
		q := img1.(float32)
		r := real2.(float32)
		s := img2.(float32)
		div := float32(math.Sqrt(float64(r*r + s*s)))
		real = (p*r + q*s) / div
		img = (q*r - p*s) / div
		return
	}
	_, ok = real1.(float64)
	if ok {
		p := real1.(float64)
		q := img1.(float64)
		r := real2.(float64)
		s := img2.(float64)
		div := math.Sqrt(r*r + s*s)
		real = (p*r + q*s) / div
		img = (q*r - p*s) / div
		return
	}
	log.Println("something wrong in division")
	return
}

func Abs(real, img interface{}) interface{} {
	_, ok := real.(float32)
	if ok {
		x := real.(float32)
		y := real.(float32)
		return float32(math.Sqrt(float64(x*x + y*y)))
	}
	_, ok = real.(float64)
	if ok {
		x := real.(float64)
		y := real.(float64)
		return math.Sqrt(x*x + y*y)
	}
	return nil
}

func Pow(ireal, iimg interface{}, pow int) (real, img interface{}) {
	if pow == 1 {
		real, img = ireal, iimg
		return
	}
	if pow <= 0 {
		log.Println("negative or zero power not supported")
		real, img = nil, nil
		return
	}
	real, img = ireal, iimg
	for pow > 1 {
		pow--
		real, img = Mul(real, img, ireal, iimg)
	}
	return
}
