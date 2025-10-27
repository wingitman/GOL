package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"math/rand"
	"time"
)

const (
	gridSize = 32 
)

func main() {
	// Create an empty slice of slices
	rows := make([][]int, 0)

	for r := 0; r < gridSize; r++ {
		col := make([]int, 0)
		for c := 0; c < gridSize; c++ {
			n := rand.Intn(10) // Intn(2) gives 0 or 1 (never 2)		r := 
			if n > 2 {
				n = 0
			} else {
				n = 1
			}
			col = append(col, n)
		}
		rows = append(rows, col)
	}

	for i := 100; i > 0; i-- {
		update(rows);
		render(rows)
		fmt.Printf("Iteration: {v}", i)
    time.Sleep(100 * time.Millisecond) 
	}
}

// GAME LOGIC //
func update(rows [][]int) {
	newRows := make([][]int, len(rows))
	for i := range rows {
		newRows[i] = make([]int, len(rows[i]))
		copy(newRows[i], rows[i]) // start with a copy of current grid
	}
	/*
		Any live cell with fewer than two live neighbours dies, as if by underpopulation.
		Any live cell with two or three live neighbours lives on to the next generation.
		Any live cell with more than three live neighbours dies, as if by overpopulation.
		Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	*/
	for r,row := range rows {
		for c,col := range row {
			l := 0

			// check up/down
			if r-1 >= 0 && rows[r-1][c] == 1 {
				l++
			}
			if r+1 < gridSize && rows[r+1][c] == 1 {
				l++
			}

			// check left/right
			if c-1 >= 0 && row[c-1] == 1 {
				l++
			}
			if c+1 < gridSize && row[c+1] == 1 {
				l++
			}

			// (Optional) check diagonals — standard Conway’s Game of Life uses these too:
			if r-1 >= 0 && c-1 >= 0 && rows[r-1][c-1] == 1 {
				l++
			}
			if r-1 >= 0 && c+1 < gridSize && rows[r-1][c+1] == 1 {
				l++
			}
			if r+1 < gridSize && c-1 >= 0 && rows[r+1][c-1] == 1 {
				l++
			}
			if r+1 < gridSize && c+1 < gridSize && rows[r+1][c+1] == 1 {
				l++
			}

			if col == 1 {
				if l < 2 || l > 3 {
					newRows[r][c] = 0
				}
			} else if l == 3 {
				newRows[r][c] = 1
			}		
		}

		for r := range rows {
			copy(rows[r], newRows[r])
		}
	}
}

// HELPERS //
func render(rows [][]int) {
	clearConsole()
	for _, row := range rows {
		for _,col := range row {
			if (col == 1) {
				fmt.Printf(" ■")
			} else {
				fmt.Printf(" □")
			}
		}
		fmt.Println("")
	}
}

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
