package main

import "fmt"
import "math/rand"
import "time"

type cell struct {
	x         int
	y         int
	connected []*cell
}

func main() {
	width := 3
	height := 3

	maze := make([]cell, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y*width+x].x = x
			maze[y*width+x].y = y
		}
	}

	var stack []*cell

	rand.Seed(time.Now().UnixNano())
	//r := rand.Intn(width * height)
	r := 1
	curCell := &maze[r]
	fmt.Println(r)
	stack = append(stack, curCell)
	unvisited := width*height - 1

	for unvisited > -1 {
		prtWidth, prtHeight := 2*width-1, 2*height-1
		prtMaze := make([]string, prtWidth*prtHeight)
		for y := 0; y < prtHeight; y++ {
			for x := 0; x < prtWidth; x++ {
				prtMaze[y*prtWidth+x] = "*"
			}
		}
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
		fmt.Println("\n", curCell.x, curCell.y)
		if curCell.x-1 != -1 && len(maze[curCell.y*width+curCell.x-1].connected) == 0 {
			curCell.connected = append(maze[curCell.y*width+curCell.x].connected, &maze[curCell.y*width+curCell.x-1])
			curCell = &maze[curCell.y*width+curCell.x-1]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.x+1 != width && len(maze[curCell.y*width+curCell.x+1].connected) == 0 {
			curCell.connected = append(maze[curCell.y*width+curCell.x].connected, &maze[curCell.y*width+curCell.x+1])
			curCell = &maze[curCell.y*width+curCell.x+1]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.y-1 != -1 && len(maze[(curCell.y-1)*width+curCell.x].connected) == 0 {
			curCell.connected = append(maze[curCell.y*width+curCell.x].connected, &maze[(curCell.y-1)*width+curCell.x])
			curCell = &maze[(curCell.y-1)*width+curCell.x]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.y+1 != width && len(maze[(curCell.y+1)*width+curCell.x].connected) == 0 {
			curCell.connected = append(maze[curCell.y*width+curCell.x].connected, &maze[(curCell.y+1)*width+curCell.x])
			curCell = &maze[(curCell.y+1)*width+curCell.x]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		}

		curCell, stack = stack[len(stack)-1], stack[:len(stack)-1]
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
}
