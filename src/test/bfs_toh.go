package main

import (
	"errors"
	"fmt"
)

var (
	uses int
)

type Column struct {
	topo  int
	array []int
}

type Tower struct {
	startColumn  Column
	middleColumn Column
	endColumn    Column
}

type Node struct {
	Value     Tower
	Parent    *Node
	Childrens []*Node
}

func (s *Column) push(e int) {
	s.array = append(s.array, e)
}

func (s *Column) pop() (int, error) {

	if len(s.array) == 0 {
		return -1, errors.New("Empty Stack")
	}

	res := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	return res, nil
}

func copy(tower Tower) (c Tower) {
	c = Tower{}

	c.startColumn.topo = tower.startColumn.topo
	c.middleColumn.topo = tower.middleColumn.topo
	c.endColumn.topo = tower.endColumn.topo

	for i := 0; i < len(tower.startColumn.array); i++ {
		c.startColumn.array = append(c.startColumn.array, tower.startColumn.array[i])
	}
	for i := 0; i < len(tower.middleColumn.array); i++ {
		c.middleColumn.array = append(c.middleColumn.array, tower.middleColumn.array[i])
	}
	for i := 0; i < len(tower.endColumn.array); i++ {
		c.endColumn.array = append(c.endColumn.array, tower.endColumn.array[i])
	}
	return
}

func switchDisk(columnFrom *Column, columnTo *Column) {

	columnTo.topo = columnFrom.topo

	x, err := columnFrom.pop()
	if err != nil {
		panic(err)
	}

	lenghtColumnFrom := len(columnFrom.array)

	if lenghtColumnFrom == 0 {
		columnFrom.topo = 9
	} else {
		columnFrom.topo = columnFrom.array[lenghtColumnFrom-1]
	}

	columnTo.push(x)
}

func bfs(tree *Node, tower *Tower) *Node {
	uses = 0
	queue := []*Node{}
	queue = append(queue, tree)
	return bfsUtil(queue)
}

func bfsUtil(queue []*Node) (res *Node) {
	uses += 1

	fmt.Println("Visão do mundo ", uses, ":")
	printWorld(queue[0].Value)
	fmt.Println()

	if len(queue) == 0 {
		return
	}

	if isWin(queue[0].Value) {
		fmt.Println("Result founded in ", uses, "steps")
		res = queue[0]
		return
	}

	for _, move := range validMoviments(queue[0].Value) {
		newNode := Node{move, queue[0], []*Node{}}
		queue[0].Childrens = append(queue[0].Childrens, &newNode)
		queue = append(queue, &newNode)
	}

	return bfsUtil(queue[1:])
}

func validMoviments(tower Tower) (result []Tower) {
	if len(tower.startColumn.array) > 0 {

		if tower.startColumn.topo < tower.middleColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.startColumn, &aux.middleColumn)
			result = append(result, aux)
		}
		if tower.startColumn.topo < tower.endColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.startColumn, &aux.endColumn)
			result = append(result, aux)
		}
	}
	if len(tower.middleColumn.array) > 0 {
		if tower.middleColumn.topo < tower.startColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.middleColumn, &aux.startColumn)
			result = append(result, aux)
		}
		if tower.middleColumn.topo < tower.endColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.middleColumn, &aux.endColumn)
			result = append(result, aux)
		}
	}
	if len(tower.endColumn.array) > 0 {
		if tower.endColumn.topo < tower.startColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.endColumn, &aux.startColumn)
			result = append(result, aux)
		}
		if tower.endColumn.topo < tower.middleColumn.topo {
			aux := copy(tower)
			switchDisk(&aux.endColumn, &aux.middleColumn)
			result = append(result, aux)
		}
	}

	return
}

func isWin(tower Tower) bool {
	var won bool = false
	if len(tower.endColumn.array) == 3 {
		won = (tower.endColumn.array[0] == 3) && (tower.endColumn.array[1] == 2) && (tower.endColumn.array[2] == 1)
	}
	return won
}

func printWorld(tower Tower) {
	col1 := 3 - len(tower.startColumn.array)
	col2 := 3 - len(tower.middleColumn.array)
	col3 := 3 - len(tower.endColumn.array)

	for i := 0; i < len(tower.startColumn.array); i++ {
		fmt.Print(tower.startColumn.array[i])
	}
	for i := 0; i < col1; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i := 0; i < len(tower.middleColumn.array); i++ {
		fmt.Print(tower.middleColumn.array[i])
	}
	for i := 0; i < col2; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i := 0; i < len(tower.endColumn.array); i++ {
		fmt.Print(tower.endColumn.array[i])
	}
	for i := 0; i < col3; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func initGame(tower *Tower) {
	*tower = Tower{Column{1, []int{3, 2, 1}}, Column{9, []int{}}, Column{9, []int{}}}
	// *tower = Tower{Column{1, []int{1}}, Column{9, []int{}}, Column{2, []int{3, 2}}}
}

func main() {
	tower := Tower{}

	initGame(&tower)
	tree := Node{tower, nil, []*Node{}}

	result := bfs(&tree, &tower)
	printWorld(result.Value)
}
