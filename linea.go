package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	////////////lectura del archivo
	content, err := os.Open("ejem.c")

	if err != nil {
		log.Fatal(err)
	}
	contentScaner := bufio.NewScanner(content)
	//fmt.Println(string(content))
	for contentScaner.Scan() {
		if strings.Contains(contentScaner.Text(), "#include") {
			if strings.Contains(contentScaner.Text(), "<") {
				if strings.Contains(contentScaner.Text(), ">") {
				} else {
					fmt.Println("falta el >")
				}
			} else {
				fmt.Println("falta el <")
			}
		} else {

			if strings.Contains(contentScaner.Text(), "<") && strings.Contains(contentScaner.Text(), ".h") && strings.Contains(contentScaner.Text(), ">") {
				fmt.Println("falta el include")
			} else if strings.Contains(contentScaner.Text(), "#include") && strings.Contains(contentScaner.Text(), ".h") && strings.Contains(contentScaner.Text(), ">") {
				fmt.Println("falta el <")
			} else if strings.Contains(contentScaner.Text(), "#include") && strings.Contains(contentScaner.Text(), "<") && strings.Contains(contentScaner.Text(), ".h") {
				fmt.Println("falta el >")
			}
		}

	}
}
