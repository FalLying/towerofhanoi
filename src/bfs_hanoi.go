package main

import (
	"errors"
	"fmt"
	"strconv"
)

var count int = 0

const (
	EMPTY  = 0
	SM     = 1
	MD     = 2
	XL     = 3
	OBJ_SM = 4
	OBJ_MD = 5
	OBJ_XL = 6
)

type Stack struct {
	a []int
}

type Tower struct {
	name  string
	disks Stack
}

type Point struct {
	x int
	y int
}

type Game struct {
	field  [][]int
	sm     Point
	md     Point
	xl     Point
	obj_sm Point
	obj_md Point
	obj_xl Point
}

func (s *Stack) push(e int) {
	s.a = append(s.a, e)
}

func (s *Stack) pop() (int, error) {

	if len(s.a) == 0 {
		return -1, errors.New("Empty Stack")
	}

	res := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]

	return res, nil
}

func validMoviments(n int, source *Tower, dest *Tower, aux *Tower, game Game) {

	if n == 1 {
		r, err := source.disks.pop()
		if err != nil {
			panic(err)
		}
		dest.disks.push(r)
		count++
		fmt.Printf("Mova o disco %d da %v para a %v\n", n, source.name, dest.name)

		if source.name == "coluna 1" {
			if dest.name == "coluna 2" {
				game.field[0] = source.disks.a
				game.field[1] = dest.disks.a
				game.field[2] = aux.disks.a
			} else {
				game.field[0] = source.disks.a
				game.field[1] = aux.disks.a
				game.field[2] = dest.disks.a
			}
		} else if source.name == "coluna 3" {
			game.field[0] = aux.disks.a
			game.field[1] = dest.disks.a
			game.field[2] = source.disks.a
		} else if source.name == "coluna 2" {
			if dest.name == "coluna 1" {
				game.field[0] = dest.disks.a
				game.field[1] = source.disks.a
				game.field[2] = aux.disks.a
			} else {
				game.field[0] = aux.disks.a
				game.field[1] = source.disks.a
				game.field[2] = dest.disks.a
			}
		}

		printWorld(game)
		fmt.Println()

		return
	}
	validMoviments(n-1, source, aux, dest, game)

	r, err := source.disks.pop()
	if err != nil {
		panic(err)
	}
	dest.disks.push(r)
	count++
	fmt.Printf("Mova o disco %d da %v para a %v\n", n, source.name, dest.name)

	if source.name == "coluna 1" {
		if dest.name == "coluna 2" {
			game.field[0] = source.disks.a
			game.field[1] = dest.disks.a
			game.field[2] = aux.disks.a
		} else {
			game.field[0] = source.disks.a
			game.field[1] = aux.disks.a
			game.field[2] = dest.disks.a
		}
	} else if source.name == "coluna 3" {
		game.field[0] = aux.disks.a
		game.field[1] = dest.disks.a
		game.field[2] = source.disks.a
	} else if source.name == "coluna 2" {
		if dest.name == "coluna 1" {
			game.field[0] = dest.disks.a
			game.field[1] = source.disks.a
			game.field[2] = aux.disks.a
		} else {
			game.field[0] = aux.disks.a
			game.field[1] = source.disks.a
			game.field[2] = dest.disks.a
		}
	}

	printWorld(game)
	fmt.Println()

	validMoviments(n-1, aux, dest, source, game)
}

func isWin(game Game) bool {
	var isSmWin bool = (game.sm.x == game.obj_sm.x) && (game.sm.y == game.obj_sm.y)
	var isMdWin bool = (game.md.x == game.obj_md.x) && (game.md.y == game.obj_md.y)
	var isXlWin bool = (game.xl.x == game.obj_xl.x) && (game.xl.y == game.obj_xl.y)

	return isSmWin && isMdWin && isXlWin
}

func printWorld(game Game) {
	for i := 0; i < len(game.field); i++ {
		fmt.Print("coluna " + strconv.Itoa(i+1) + ": { ")
		for j := 0; j < len(game.field[i]); j++ {
			switch game.field[i][j] {
			case 0:
				fmt.Print("|")
			case 1:
				fmt.Print("= ")
			case 2:
				fmt.Print("== ")
			case 3:
				fmt.Print("=== ")
			}
		}
		fmt.Print("}\n")
	}
}

func initGame(game *Game) {
	*game = Game{[][]int{{3, 2, 1}, {}, {}}, Point{2, 0}, Point{1, 0}, Point{0, 0}, Point{2, 2}, Point{1, 2}, Point{0, 2}}
}

func main() {

	var srcPeg Tower
	var auxPeg Tower
	var destPeg Tower

	srcPeg.name = "coluna 1"
	auxPeg.name = "coluna 2"
	destPeg.name = "coluna 3"

	game := Game{}

	initGame(&game)

	srcPeg.disks.a = game.field[0]
	auxPeg.disks.a = game.field[1]
	destPeg.disks.a = game.field[2]

	fmt.Print("\n__________________________\n\tResultado\n__________________________\n\n")

	printWorld(game)
	fmt.Println()

	validMoviments(3, &srcPeg, &destPeg, &auxPeg, game)
}
