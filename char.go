package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
  content, err:=os.Open("ejem.c")

  if err != nil {
		log.Fatal(err)
	}
  contentScaner := bufio.NewScanner(content)
  for contentScaner.Scan(){
    if strings.Contains(contentScaner.Text(), "char"){
      if  strings.Contains(contentScaner.Text(), " ") {
        if strings.Contains(contentScaner.Text(), ";") {
        }else {
          fmt.Println("Falta el ;")
          }
       }else {
        fmt.Println("Falta asignar valor")
      }
  } else {

    if strings.Contains(contentScaner.Text(), ";") && strings.Contains(contentScaner.Text(), "cha"){
      fmt.Println("char esta mal escrito")
    } else if strings.Contains(contentScaner.Text(), ";")&& strings.Contains(contentScaner.Text(), "chat"){
      fmt.Println("char esta mal escrito")
    } else if strings.Contains(contentScaner.Text(), ";")&& strings.Contains(contentScaner.Text(), "har"){
      fmt.Println("char esta mal escrito")
    }else if strings.Contains(contentScaner.Text(), ";")&& strings.Contains(contentScaner.Text(), "chart"){
      fmt.Println("char esta mal escrito")
    }else if strings.Contains(contentScaner.Text(), ";")&& strings.Contains(contentScaner.Text(), "chr"){
      fmt.Println("char esta mal escrito")
    }else if strings.Contains(contentScaner.Text(), "char"){
      fmt.Println("falta el ;")
    }
    }
  }
}
