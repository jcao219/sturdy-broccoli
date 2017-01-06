package main

import (
	"bufio"
	"fmt"
	_ "github.com/gonum/graph"
	_ "github.com/gonum/graph/simple"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("One connected graph per line:")
		fmt.Println("1 2  1 4  1 3  1 5  1 6")
		fmt.Println("The above is a claw graph with 6 vertices, 5 edges.")
		fmt.Println("Notice the two spaces between edges.")
		fmt.Println("Node number must start with 1 and go up to N")
		fmt.Println("Need file with graphs in the above special format as argument.")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
