
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	freq       = 0  // milliseconds between frames
	iterations = 10000
)

func main() {
	// Get terminal size
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Initialize grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Random initial population (~20% alive)
	rand.Seed(time.Now().UnixNano())
	numCells := height * width / 5
	for i := 0; i < numCells; i++ {
		r := rand.Intn(height)
		c := rand.Intn(width)
		grid[r][c] = 1
	}

	// Main loop
	for iter := 0; iter < iterations; iter++ {
		render(grid)
		grid = update(grid)
		time.Sleep(freq * time.Millisecond)
	}
}

// Render grid in terminal
func render(grid [][]int) {
	fmt.Print("\033[H\033[2J") // fast clear screen
	for _, row := range grid {
		line := make([]rune, len(row))
		for i, col := range row {
			if col == 1 {
				line[i] = '■'
			} else {
				line[i] = ' '
			}
		}
		fmt.Println(string(line))
	}
}

// Count live neighbors
func countNeighbors(grid [][]int, r, c int) int {
	count := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue // skip self
			}
			nr := r + dr
			nc := c + dc
			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
				count += grid[nr][nc]
			}
		}
	}
	return count
}

// Update grid according to Conway's rules
func update(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := range grid {
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	for r := range grid {
		for c := range grid[r] {
			neighbors := countNeighbors(grid, r, c)
			if grid[r][c] == 1 {
				if neighbors < 2 || neighbors > 3 {
					newGrid[r][c] = 0 // dies
				}
			} else {
				if neighbors == 3 {
					newGrid[r][c] = 1 // becomes alive
				}
			}
		}
	}

	return newGrid
}

