package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/t0duf0du/AOC2024/cmd/day1"
)

func main() {
	fmt.Println("HOHOHO")
	err := day1.Day1()
	if err != nil {
		pp.Println(err)
	}

}
