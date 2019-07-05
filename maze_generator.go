package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type cell struct {
	x         int
	y         int
	connected []*cell
}

func main() {
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	// Create the maze
	maze := make([]cell, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y*width+x].x = x
			maze[y*width+x].y = y
		}
	}

	// Implementing stack
	var stack []*cell

	// Choose rand cell
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(width * height)
	curCell := &maze[r]
	fmt.Println(r)
	stack = append(stack, curCell)
	unvisited := width*height - 1

	for unvisited != 0 {
		// outside walls
		prtWidth, prtHeight := 2*width-1, 2*height-1
		prtMaze := make([]string, prtWidth*prtHeight)
		for y := 0; y < prtHeight; y++ {
			for x := 0; x < prtWidth; x++ {
				prtMaze[y*prtWidth+x] = "*"
			}
		}
		// connected cells
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				prtMaze[y*2*prtWidth+x*2] = "-"
				for _, curCell := range maze[y*width+x].connected {
					dx, dy := curCell.x-x, curCell.y-y
					prtMaze[(y*2+dy)*prtWidth+x*2+dx] = "+"
				}
			}
		}

		for y := 0; y < prtHeight; y++ {
			for x := 0; x < prtWidth; x++ {
				fmt.Print(prtMaze[y*prtWidth+x])
			}
			fmt.Print("\n")
		}
		fmt.Println()

		// Find directions
		availCells := make([]*cell, 0, 4)
		if curCell.x-1 != -1 && len(maze[curCell.y*width+curCell.x-1].connected) == 0 {
			availCells = append(availCells, &maze[curCell.y*width+curCell.x-1])
		}
		if curCell.x+1 != width && len(maze[curCell.y*width+curCell.x+1].connected) == 0 {
			availCells = append(availCells, &maze[curCell.y*width+curCell.x+1])
		}
		if curCell.y-1 != -1 && len(maze[(curCell.y-1)*width+curCell.x].connected) == 0 {
			availCells = append(availCells, &maze[(curCell.y-1)*width+curCell.x])
		}
		if curCell.y+1 != height && len(maze[(curCell.y+1)*width+curCell.x].connected) == 0 {
			availCells = append(availCells, &maze[(curCell.y+1)*width+curCell.x])
		}

		// pop from stack
		if len(availCells) < 1 {
			curCell, stack = stack[len(stack)-2], stack[:len(stack)-2]
			continue
		}

		// connect
		conCell := availCells[rand.Intn(len(availCells))]
		curCell.connected = append(curCell.connected, conCell)
		conCell.connected = append(conCell.connected, curCell)
		stack = append(stack, conCell)
		curCell = conCell
		unvisited -= 1
	}

	prtWidth, prtHeight := 2*width-1, 2*height-1
	prtMaze := make([]string, prtWidth*prtHeight)
	for y := 0; y < prtHeight; y++ {
		for x := 0; x < prtWidth; x++ {
			prtMaze[y*prtWidth+x] = "*"
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			prtMaze[y*2*prtWidth+x*2] = " " // -
			for _, curCell := range maze[y*width+x].connected {
				dx, dy := curCell.x-x, curCell.y-y
				prtMaze[(y*2+dy)*prtWidth+x*2+dx] = " " // +
			}
		}
	}

	f, err := os.Create("maze")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(strings.Repeat("*", prtWidth+2) + "\n")
	for y := 0; y < prtHeight; y++ {
		f.WriteString("*")
		for x := 0; x < prtWidth; x++ {
			fmt.Print(prtMaze[y*prtWidth+x])
			f.WriteString(prtMaze[y*prtWidth+x])
		}
		fmt.Print("\n")
		f.WriteString("*\n")
	}
	f.WriteString(strings.Repeat("*", prtWidth+2) + "\n")

	f.Sync()
}
