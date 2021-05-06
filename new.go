package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {

    content, err := ioutil.ReadFile("ejem.txt")

     if err != nil {
          log.Fatal(err)
     }

    //fmt.Println(string(content))

    fmt.Println(strings.Contains(string(content), ";"))//hacer una funcion que cheque argumentos

    nombresComoArreglo := strings.Split(string(content), ";")

      fmt.Println(nombresComoArreglo[0])//pofdria hacer una funcion con esto y pasarle el argumento que quier
//hacer con while que por cada line cheque hacer para las n ves de lineas que hayya
//si existe sumar a uno y lklevar la cuenta ty despues comparara

  }
