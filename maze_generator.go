package main

import "fmt"

//import "math/rand"

type cell struct {
	x         int
	y         int
	connected []*cell
}

func main() {
	//width := 5
	//height := 5

	var maze [5 * 5]cell

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			maze[y*5+x].x = x
			maze[y*5+x].y = y
			//maze[y*5+x].walls = make([]cell, 1)
			//maze[y*5+x].connected = append(maze[y*5+x].connected, &maze[y*5+x])
		}
	}

	//curCell := maze[rand.Intn(5 * 5)]

	var stack []*cell

	curCell := &maze[0]
	stack = append(stack, curCell)
	unvisited := 5*5 - 1

	for unvisited != 0 {
		if curCell.x-1 != -1 && len(maze[curCell.y*5+curCell.x-1].connected) == 0 {
			curCell.connected = append(maze[curCell.y*5+curCell.x].connected, &maze[curCell.y*5+curCell.x-1])
			curCell = &maze[curCell.y*5+curCell.x-1]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.x+1 != 5 && len(maze[curCell.y*5+curCell.x+1].connected) == 0 {
			curCell.connected = append(maze[curCell.y*5+curCell.x].connected, &maze[curCell.y*5+curCell.x+1])
			curCell = &maze[curCell.y*5+curCell.x+1]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.y-1 != -1 && len(maze[(curCell.y-1)*5+curCell.x].connected) == 0 {
			curCell.connected = append(maze[curCell.y*5+curCell.x].connected, &maze[(curCell.y-1)*5+curCell.x])
			curCell = &maze[(curCell.y-1)*5+curCell.x]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		} else if curCell.y+1 != 5 && len(maze[(curCell.y+1)*5+curCell.x-1].connected) == 0 {
			curCell.connected = append(maze[curCell.y*5+curCell.x].connected, &maze[(curCell.y+1)*5+curCell.x])
			curCell = &maze[(curCell.y+1)*5+curCell.x]
			stack = append(stack, curCell)
			unvisited -= 1
			continue
		}

		curCell, stack = stack[len(stack)-1], stack[:len(stack)-1]
	}

	fmt.Println(maze)
}
