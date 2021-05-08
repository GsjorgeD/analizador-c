package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "reflect"
)

func main() {

    content, err := ioutil.ReadFile("ejem.txt")

     if err != nil {
          log.Fatal(err)
     }

    //fmt.Println(string(content))

    verifi := strings.Contains(string(content), "#include")//hacer una funcion que cheque argumentos
    //fmt.Println(verifi)
    if verifi==false{
    fmt.Println("error en include")
    }
    pedazos := strings.Split(string(content), " ")

      //fmt.Println(pedazos)//pofdria hacer una funcion con esto y pasarle el argumento que quier
//hacer con while que por cada line cheque hacer para las n ves de lineas que hayya
//si existe sumar a uno y llevar la cuenta y despues comparara
    numpedazos := len(pedazos)
    for index := 0; index < numpedazos; index++ {
    checa("{", string(pedazos[index]))
    }


    checa("#include", string(pedazos[0]))
    fmt.Println(reflect.TypeOf(verifi))//verififica que tpo de datos es


  }
  func checa(palabra string, content string) {
    ver := strings.Contains(string(content), palabra)
    fmt.Println(ver)
  }
