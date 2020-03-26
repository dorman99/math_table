package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("must be: ggo run main.go [operation] [range]")
		os.Exit(0)
	}
	s := os.Args[1]

	if head, err := strconv.Atoi(os.Args[2]); err != nil {
		fmt.Println("Error: Must be a number")
	} else {
		switch {
		case s == "multi":
			mult(head)
		case s == "add":
			add(head)
		case s == "min":
			sub(head)
		case s == "div":
			div(head)
		default:
			fmt.Println(`Not Known Operation ["multi", "add", "min", "div"]`)
		}
	}
}

func div(head int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= head; j++ {
			if j != 0 {
				fmt.Printf("%5.1f", float64(i)/float64(j))
			} else {
				fmt.Printf("%5d", 0)
			}
		}
		fmt.Println()
	}
}

func sub(head int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= head; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}
}

func add(head int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= head; j++ {
			fmt.Printf("%5d", i+j)
		}
		fmt.Println()
	}
}

func mult(head int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= head; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= head; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
