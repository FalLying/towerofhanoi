package main

import "fmt"

type Column struct {
	topo  int
	array []int
}

type Tower struct {
	startColumn  Column
	middleColumn Column
	endColumn    Column
}

func validMoviments(tower Tower) {
	if tower.startColumn.topo < tower.middleColumn.topo {

	} else if tower.startColumn.topo < tower.endColumn.topo {

	} else if tower.middleColumn.topo < tower.startColumn.topo {

	} else if tower.middleColumn.topo < tower.endColumn.topo {

	} else if tower.endColumn.topo < tower.startColumn.topo {

	} else if tower.endColumn.topo < tower.middleColumn.topo {
	}
}

func isWin(tower Tower) bool {
	won := (tower.endColumn.array[0] == 3) && (tower.endColumn.array[1] == 2) && (tower.endColumn.array[2] == 1)
	return won
}

func printWorld(tower Tower) {
	for i := 0; i < len(tower.startColumn.array); i++ {
		fmt.Print(tower.startColumn.array[i])
	}
	fmt.Println()

	for i := 0; i < len(tower.middleColumn.array); i++ {
		fmt.Print(tower.middleColumn.array[i])
	}
	fmt.Println()

	for i := 0; i < len(tower.endColumn.array); i++ {
		fmt.Print(tower.endColumn.array[i])
	}
}

func initGame(tower *Tower) {
	*tower = Tower{Column{1, []int{3, 2, 1}}, Column{9, []int{0, 0, 0}}, Column{9, []int{0, 0, 0}}}
}

func main() {
	tower := Tower{}

	initGame(&tower)
	printWorld(tower)
}
