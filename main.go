package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	blocks map[int8]string = map[int8]string{
		0: "⬜",
		1: "⬛",
	}
	clearScreen string = "\033[H\033[2J"
)

type Lives [][]int8

func (l Lives) shouldCellLive(i, j int) bool {
	lengthI, lengthJ := len(l), len(l[0])
	distances := [3]int8{-1, 0, 1}
	liveNeighbors := 0
	val := l[i][j]

	for _, distanceI := range distances {
		for _, distanceJ := range distances {
			neighborI, neighborJ := i+int(distanceI), j+int(distanceJ)

			if (distanceI == 0) && (distanceJ == 0) {
				continue
			} else if neighborI >= 0 && neighborI < lengthI && neighborJ >= 0 && neighborJ < lengthJ {
				neighbor := l[neighborI][neighborJ]
				if neighbor == 1 {
					liveNeighbors++
				}
			}
		}
	}
	if val == 1 && (liveNeighbors == 2 || liveNeighbors == 3) {
		return true
	} else if liveNeighbors == 3 {
		return true
	}
	return false
}

func (l Lives) GetNextGen() Lives {
	lengthI, lengthJ := len(l), len(l[0])
	nextLives := make([][]int8, lengthI)
	i := 0
	for i < lengthI {
		nextLives[i] = make([]int8, lengthJ)
		i++
	}

	for i, arr := range l {
		for j := range arr {
			if l.shouldCellLive(i, j) {
				nextLives[i][j] = 1
			} else {
				nextLives[i][j] = 0
			}
		}
	}

	return nextLives
}

func (l Lives) String() string {
	repr := ""
	for _, arr := range l {
		for j, val := range arr {
			if j == 0 {
				repr += "\n"
			}
			repr += blocks[val]
		}
	}

	return repr
}

func From(file string) (Lives, error) {
	if fil, err := os.Open(file); err == nil {
		csvReader := csv.NewReader(fil)
		if data, err := csvReader.ReadAll(); err == nil {
			live := make([][]int8, len(data))

			for i := range data {
				live[i] = make([]int8, len(data[i]))
				for j := range data[i] {
					if num, err := strconv.Atoi(data[i][j]); err == nil {
						live[i][j] = int8(num)
					} else {
						return nil, err
					}
				}
			}
			return live, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: Pattern filename required")
		os.Exit(1)
	}

	file := os.Args[1]
	if lives, err := From(file); err == nil {
		fmt.Print(clearScreen)

		i := 1

		for {
			fmt.Print(lives)
			fmt.Printf("\nGeneration: %d\n", i)
			lives = lives.GetNextGen()
			time.Sleep(250 * time.Millisecond)
			fmt.Print(clearScreen)
			i++
		}
	} else {
		fmt.Println("Attempted to process an unknown file...")
		os.Exit(1)
	}

}
