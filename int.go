
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
		if strings.HasPrefix(contentScaner.Text(), "int") {
 fmt.Println("hay int")
		}

	}
}
