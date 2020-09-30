package main

import "log"

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
	log.Printf("\nKB:\t%d\nMB:\t%d\nGB:\t%d\nTB:\t%d\nPB:\t%d\nEB:\t%d\nZB:\t%e\nYB:\t%e", KB, MB, GB, TB, PB, EB, float64(ZB), float64(YB))
}
