package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func linesToJson() {
	file, err := os.Open("facebook.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := ""
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		out += "{\"source\": " + split[0] + ", \"target\":" + split[1] + "},\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(out)

	f, err := os.Create("connections-processed.json")
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
