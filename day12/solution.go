package main

import (
	"aoc/util"
	"fmt"
	"math"
)

type Coord struct {
	x  int
	y  int
	el int
}

func canCrossTo(orig Coord, dest Coord, flipped bool) bool {
	if flipped {
		return orig.el <= dest.el+1
	}
	return orig.el >= dest.el-1
}

type Board struct {
	contents [][]Coord
}

func (board *Board) getCoords() []Coord {
	result := make([]Coord, 0)
	for _, row := range board.contents {
		result = append(result, row...)
	}
	return result
}

func (board *Board) getNeighbors(coord Coord, flipped bool) []Coord {
	result := make([]Coord, 0)
	if coord.x > 0 {
		candidate := board.contents[coord.y][coord.x-1]
		if canCrossTo(coord, candidate, flipped) {
			result = append(result, candidate)
		}
	}
	if coord.x < len(board.contents[0])-1 {
		candidate := board.contents[coord.y][coord.x+1]
		if canCrossTo(coord, candidate, flipped) {
			result = append(result, candidate)
		}
	}
	if coord.y > 0 {
		candidate := board.contents[coord.y-1][coord.x]
		if canCrossTo(coord, candidate, flipped) {
			result = append(result, candidate)
		}
	}
	if coord.y < len(board.contents)-1 {
		candidate := board.contents[coord.y+1][coord.x]
		if canCrossTo(coord, candidate, flipped) {
			result = append(result, candidate)
		}
	}
	return result
}

func parse(lines []string) (Board, Coord, Coord) {
	start := Coord{}
	end := Coord{}
	board := Board{}

	for yC, line := range lines {
		row := make([]Coord, len(line))
		if line != "" {
			for xC, char := range line {
				if char == 'S' {
					start = Coord{xC, yC, 0}
					row[xC] = start
				} else if char == 'E' {
					end = Coord{xC, yC, 25}
					row[xC] = end
				} else {
					row[xC] = Coord{xC, yC, int(char) - 97}

				}
			}
		}
		board.contents = append(board.contents, row)
	}

	return board, start, end
}

func dijk(board Board, start Coord, finishCheck func(coord Coord) bool, flippedNeighbors bool) int {
	dist := make(map[Coord]int)
	dist[start] = 0
	prev := make(map[Coord]Coord)
	unvisited := make(map[Coord]struct{})
	for _, coord := range board.getCoords() {
		unvisited[coord] = util.EMPTY_STRUCT
	}

	for len(unvisited) > 0 {
		u := Coord{}
		uDist := math.MaxInt
		for coord := range unvisited {
			unvisDist, ok := dist[coord]

			if ok && unvisDist < uDist {
				u = coord
				uDist = unvisDist
			}
		}
		if finishCheck(u) {
			return uDist
		}

		delete(unvisited, u)

		for _, n := range board.getNeighbors(u, flippedNeighbors) {
			nDist, nDistPresent := dist[n]
			newDist := uDist + 1
			if !nDistPresent || newDist < nDist {
				dist[n] = newDist
				prev[n] = u
			}
		}
	}

	return -1
}

func part1(board Board, start Coord, end Coord) int {
	return dijk(board, start, func(coord Coord) bool { return coord == end }, false)
}

func part2(board Board, end Coord) int {
	return dijk(board, end, func(coord Coord) bool { return coord.el == 0 }, true)
}

func main() {
	lines := util.FileAsLines("input")
	board, start, end := parse(lines)

	fmt.Println(part1(board, start, end))
	fmt.Println(part2(board, end))

}
