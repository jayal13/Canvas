package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  C w h              - Create a new canvas with width w and height h")
	fmt.Println("  L (x1,y1) (x2-y2)  - Draw a line from (x1,y1) to (x2,y2)")
	fmt.Println("  R (x,y) w h        - Draw a rectangle starting at (x,y) with width w and height h")
	fmt.Println("  F (x,y)            - Fill area starting at (x,y)")
	fmt.Println("  H             	  - Show this help")
	fmt.Println("  Q              	  - Quit")
}

func parsePoint(s string) ([]int, error) {
	s = strings.TrimSpace(s)

	if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
		return nil, fmt.Errorf("invalid point format: %s", s)
	}

	// quitar paréntesis
	inner := s[1 : len(s)-1]

	parts := strings.Split(inner, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid point format: %s", s)
	}

	x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("invalid integers in point: %s", s)
	}

	return []int{x, y}, nil
}
