package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "enter file name")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("unable to open")
	}
	r := csv.NewReader(file)
	lines, err := r.readAll()

}
