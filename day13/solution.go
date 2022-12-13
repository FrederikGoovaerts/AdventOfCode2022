package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	isList bool
	cont   []Packet
	val    int
}

func parsePacket(line []string, start int) (Packet, int) {
	packet := Packet{false, make([]Packet, 0), 0}
	for index := start; index < len(line); index++ {
		val := line[index]
		if val == "]" {
			packet.isList = true
			return packet, index
		} else if val == "[" {
			packet.isList = true
			if line[index+1] == "]" {
				return packet, index + 1
			} else {
				subPacket, subEnd := parsePacket(line, index+1)
				packet.cont = append(packet.cont, subPacket)
				index = subEnd // Will go through index++
			}
		} else if val == "," {
			subPacket, subEnd := parsePacket(line, index+1)
			packet.cont = append(packet.cont, subPacket)
			index = subEnd // Will go through index++
		} else {
			if line[index+1] == "0" {
				packet.val = 10
				return packet, index + 1
			} else {
				numberValue, _ := strconv.Atoi(val)
				packet.val = numberValue
				return packet, index
			}
		}
	}
	return packet, len(line)
}

func areSorted(left Packet, right Packet) int {
	if !left.isList && !right.isList {
		return util.ClampToOne(left.val - right.val)
	}
	if left.isList && right.isList {
		for i := 0; i < len(left.cont); i++ {
			if i == len(right.cont) {
				return 1
			}
			sortResult := areSorted(left.cont[i], right.cont[i])
			if sortResult != 0 {
				return sortResult
			}
		}
		if len(left.cont) < len(right.cont) {
			return -1
		} else {
			return 0
		}
	}
	container := Packet{true, make([]Packet, 1), 0}
	if left.isList {
		container.cont[0] = right
		return areSorted(left, container)
	} else {
		container.cont[0] = left
		return areSorted(container, right)
	}
}

func part1(packets []Packet) int {
	result := 0
	for i := 0; i < len(packets); i += 2 {
		if areSorted(packets[i], packets[i+1]) <= 0 {
			result += (i / 2) + 1
			// fmt.Println((i/2)+1, "sorted")
		}
	}
	return result
}

func main() {
	lines := util.FileAsLines("input")
	packets := make([]Packet, 0)

	for _, line := range lines {
		if line != "" {
			splitLine := strings.Split(line, "")
			packet, _ := parsePacket(splitLine, 0)
			packets = append(packets, packet)
		}
	}

	// for _, pack := range packets {
	// 	fmt.Println(pack)
	// }
	fmt.Println(part1(packets))
}
