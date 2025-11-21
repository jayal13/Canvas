package cmd

import (
	"bufio"
	"canvas/canvas"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewListener() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("|!|JC Canvas|!|")
	fmt.Println("First, create a canvas.")

	c := createCanvas(scanner)
	if c == nil {
		return fmt.Errorf("could not create canvas")
	}

	fmt.Println("Canvas created:")
	fmt.Println(c.Render())

	printHelp()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			return fmt.Errorf("error with scanner")
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		cmd := strings.ToUpper(fields[0])
		args := fields[1:]

		var err error

		switch cmd {
		case "Q":
			fmt.Println("Bye!")
			return nil

		case "H":
			printHelp()
			continue

		case "C":
			if len(args) != 2 {
				fmt.Println("C only accepts two values")
				continue
			}
			w, err1 := strconv.Atoi(args[0])
			h, err2 := strconv.Atoi(args[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Width and height must be integers")
				continue
			}
			c, err = canvas.NewCanva(w, h)
			if err != nil {
				fmt.Println("Error creating canvas:", err)
				continue
			}
			fmt.Println(c.Render())

		case "L":
			if len(args) != 2 {
				fmt.Println("Usage: L (x1,y1) (x2,y2)")
				continue
			}

			p1, err1 := parsePoint(args[0])
			p2, err2 := parsePoint(args[1])

			if err1 != nil || err2 != nil {
				fmt.Println("Invalid points")
				continue
			}

			err = c.NewLine(p1, p2)
			if err != nil {
				fmt.Println("Error drawing line:", err)
				continue
			}

			fmt.Println(c.Render())

		case "R":
			if len(args) != 3 {
				fmt.Println("Usage: R (x,y) width height")
				continue
			}

			p, err := parsePoint(args[0])
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			w, err1 := strconv.Atoi(args[1])
			h, err2 := strconv.Atoi(args[2])

			if err1 != nil || err2 != nil {
				fmt.Println("Width and height must be integers")
				continue
			}

			err = c.NewRectangle(p, w, h)
			if err != nil {
				fmt.Println("Error drawing rectangle:", err)
				continue
			}

			fmt.Println(c.Render())

		case "F":
			if len(args) != 1 {
				fmt.Println("Usage: F (x,y)")
				continue
			}

			p, err := parsePoint(args[0])
			if err != nil {
				fmt.Println("Invalid point")
				continue
			}

			err = c.Fill(p)
			if err != nil {
				fmt.Println("Error filling:", err)
				continue
			}

			fmt.Println(c.Render())

		default:
			fmt.Println("Unknown command:", cmd)
			printHelp()
		}
	}
}

func createCanvas(scanner *bufio.Scanner) *canvas.Canvas {
	for {
		fmt.Print("Enter canvas width and height (e.g. 20 10): ")
		if !scanner.Scan() {
			return nil
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Please enter two integers: width height")
			continue
		}
		w, err1 := strconv.Atoi(parts[0])
		h, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Width and height must be integers")
			continue
		}
		c, err := canvas.NewCanva(w, h)
		if err != nil {
			fmt.Println("Error creating canvas:", err)
			continue
		}
		return c
	}
}
