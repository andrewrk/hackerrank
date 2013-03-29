package main

import (
	"os"
	"bufio"
	"io"
	"strconv"
	"fmt"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	size_line, err := reader.ReadString('\n')
	if err != nil { panic(err) }
	_, err = reader.ReadString('\n')
	if err != nil { panic(err) }
	m64, err := strconv.ParseInt(strings.TrimSpace(size_line), 10, 32)
	if err != nil { panic(err) }
	m := int(m64)
	var mario_pos pos
	var princess_pos pos
	for y := 0; y < m; y++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF { panic(err) }
		for x, char := range line {
			if char == 'm' {
				mario_pos.x = x
				mario_pos.y = y
			} else if char == 'p' {
				princess_pos.x = x
				princess_pos.y = y
			}
		}
	}
	if mario_pos.x < princess_pos.x {
		fmt.Println("RIGHT")
	} else if mario_pos.x > princess_pos.x {
		fmt.Println("LEFT")
	} else if mario_pos.y < princess_pos.y {
		fmt.Println("DOWN")
	} else if mario_pos.y > princess_pos.y {
		fmt.Println("UP")
	}
}
