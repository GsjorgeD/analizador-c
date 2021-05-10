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

	var contadorDo int = 0
	var contadorWhi int = 0

	if err != nil {
		log.Fatal(err)
	}
	contentScaner := bufio.NewScanner(content)
	//fmt.Println(string(content))
	for contentScaner.Scan() {
		if strings.Contains(contentScaner.Text(), "do") {
			contadorDo++
		}
		if strings.Contains(contentScaner.Text(), "while") {
			contadorWhi++
		}
	}
	if contadorDo == contadorWhi {
	} else {
		fmt.Println("corregir ciclo")
	}

}
