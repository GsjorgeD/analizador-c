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
	for contentScaner.Scan() {
		if strings.Contains(contentScaner.Text(), "char") {
			if strings.Contains(contentScaner.Text(), " ") {
			  if strings.Contains(contentScaner.Text(), ";") {
						} else {
							fmt.Println("falta el ;")
						}
					} else {
						fmt.Println("asigna valor")
					}
				}
		} else { if string.Contains(contentScaner.Text(), ";"){
			fmt.Println("char mal escrito")
		} else if {

		}
	}


		else if strings.Contains(contentScaner.Text(), "char"){
			fmt.Println("falta el ;")
   }}}
