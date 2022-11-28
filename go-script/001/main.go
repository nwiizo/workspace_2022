package main

import (
	"github.com/bitfield/script"
)

func main() {
	script.File("access.log").Column(1).Freq().First(10).Stdout()
}
