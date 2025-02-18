package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func nodesParsing() {
	groups := [][]int{
		{71, 215, 54, 61, 298, 229, 81, 253, 193, 97, 264, 29, 132, 110, 163, 259, 183, 334, 245, 222},
		{173},
		{155, 99, 327, 140, 116, 147, 144, 150, 270},
		{51, 83, 237},
		{125, 344, 295, 257, 55, 122, 223, 59, 268, 280, 84, 156, 258, 236, 250, 239, 69},
		{23},
		{337, 289, 93, 17, 111, 52, 137, 343, 192, 35, 326, 310, 214, 32, 115, 321, 209, 312, 41, 20},
		{225, 46},
		{282},
		{336, 204, 74, 206, 292, 146, 154, 164, 279, 73},
		{42, 14, 216, 2},
		{324, 265, 54, 161, 298, 76, 165, 199, 203, 13, 66, 113, 97, 252, 313, 238, 158, 240, 331, 332, 134, 218, 118, 235, 311, 151, 308, 212, 70, 211},
		{278},
		{138, 131, 68, 143, 86},
		{175, 227},
		{108, 208, 251, 125, 325, 176, 133, 276, 198, 271, 288, 316, 96, 246, 347, 121, 7, 3, 170, 323, 56, 338, 23, 109, 141, 67, 345, 55, 114, 122, 50, 304, 318, 65, 15, 45, 317, 322, 26, 31, 168, 124, 285, 255, 129, 40, 172, 274, 95, 207, 128, 339, 233, 1, 294, 280, 224, 269, 256, 60, 328, 189, 146, 77, 196, 64, 286, 89, 22, 39, 190, 281, 117, 38, 213, 135, 197, 291, 21, 315, 261, 47, 36, 186, 169, 342, 49, 9, 16, 185, 219, 123, 72, 309, 103, 157, 277, 105, 139, 148, 248, 341, 62, 98, 63, 297, 242, 10, 152, 236, 308, 82, 87, 136, 200, 183, 247, 290, 303, 319, 6, 314, 104, 127, 25, 69, 171, 119, 79, 340, 301, 188, 142},
		{251, 94, 330, 5, 34, 299, 254, 24, 180, 194, 281, 101, 266, 135, 197, 173, 36, 9, 85, 57, 37, 258, 309, 80, 139, 202, 187, 249, 58, 127, 48, 92},
		{90, 52, 172, 126, 294, 179, 145, 105, 210},
		{177},
		{93, 33, 333, 17, 137, 44, 343, 326, 214, 115, 312, 41, 20},
		{244, 282, 262, 293, 220, 174},
		{12},
		{267},
		{28, 149, 162},
	}

	file, err := os.Open("facebook.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodesArray := make([]string, 0)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if !slices.Contains(nodesArray, split[0]) {
			nodesArray = append(nodesArray, split[0])
		}
		if !slices.Contains(nodesArray, split[1]) {
			nodesArray = append(nodesArray, split[1])
		}
	}

	out := ""

	for i := range nodesArray {
		groupId := 0

		for y := range groups {
			for x := range groups[y] {
				if nodesArray[i] == strconv.Itoa(groups[y][x]) {
					groupId = y
				}
			}
		}

		out += "{\"id\": " + nodesArray[i] + ", \"group\": " + strconv.Itoa(groupId) + "},\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("nodes.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(out)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
