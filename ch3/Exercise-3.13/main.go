package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("KIB:\t%d\nMiB:\t%d\nGiB:\t%d\nTiB:\t%d\n", KiB, MiB, GiB, TiB)
	fmt.Printf("PIB:\t%d\nEiB:\t%d\nZiB:\t%.0f\nYiB:\t%.0f\n", PiB, EiB, float64(ZiB), float64(YiB))
	fmt.Printf("KB:\t%d\nMB:\t%d\nGB:\t%d\nTB:\t%d\n", KB, MB, GB, TB)
	fmt.Printf("PB\t%d\nEB:\t%d\nZB:\t%.0f\nYB:\t%.0f\n", PB, EB, float64(ZB), float64(YB))

}
