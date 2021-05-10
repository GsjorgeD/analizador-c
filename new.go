package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "reflect"
)

func main() {
    var llaveApe int = 0
    var llaveEnd int = 0
    var parenApe int = 0
    var parenEnd int = 0
    ////////////lectura del archivo
    content, err := ioutil.ReadFile("ejem.c")

     if err != nil {
          log.Fatal(err)
     }

    //fmt.Println(string(content))
    pedazos := strings.Split(string(content), " ") //separa el archivo en fracmento
    verifi := strings.Contains(string(content), "#include")//hacer una funcion que cheque argumentos
    //fmt.Println(verifi)
    numinclude := len(pedazos[0])//cuenta cuanto elementos hay en include
    if verifi==false{
    fmt.Println("error en include")
    }else if numinclude !=8{
      fmt.Println("error en include")
    }



//////////////////////////////verifica <> en la libreria
    simbMenor:= strings.Contains(string(pedazos[1]), "<")
    simbMayor := strings.Contains(string(pedazos[1]), ">")
    leth := strings.Contains(string(pedazos[1]), ".h")
    fmt.Println((simbMayor))
    fmt.Println((simbMenor))
    if simbMayor==false || simbMenor==false || leth==false{
      fmt.Println("verifica la libreria")
    }
    ////verificar Int
  //  pointer:= 2
    //findInt := strings.Contains(string(pedazos[pointer]), "int")
    //if findInt == true{
      //pointer++
        //findEqual := strings.Contains(string(pedazos[pointer]), "=")
        //no olvidar verifiar el tipo de dato
      //if
    //}

      //***ignorar fmt.Println(pedazos)//pofdria hacer una funcion con esto y pasarle el argumento que quier
//hacer con while que por cada line cheque hacer para las n ves de lineas que hayya
//si existe sumar a uno y llevar la cuenta y despues comparara
    numpedazos := len(pedazos)
    ////////busca llaves de apertura
    for index := 0; index < numpedazos; index++ {
    //checa("{", string(pedazos[index]))
    findllave := strings.Contains(string(pedazos[index]), "{")
      if findllave==true{
        llaveApe = llaveApe+1
      }
    }
    fmt.Println(llaveApe)
    ////////busca llaves de cerradura
    for index := 0; index < numpedazos; index++ {
    findllave := strings.Contains(string(pedazos[index]), "}")
      if findllave==true{
        llaveEnd = llaveEnd+1
      }
    }
    fmt.Println(llaveEnd)
  //////////Verifica si esta la misma cantidad de llaves
    if llaveApe > llaveEnd{
      fmt.Println("Revisa tus llaves de cierre")
    }else if llaveEnd>llaveApe{
      fmt.Println("Revisa tus llaves de apertura")
    }
    ///////////verifica parentesis de apertura
    for index := 0; index < numpedazos; index++ {
    findllave := strings.Contains(string(pedazos[index]), "(")
      if findllave==true{
        parenApe = parenApe+1
      }
    }
    fmt.Println(parenApe)
    ///////////verifica parentesis de apertura
    for index := 0; index < numpedazos; index++ {
    findllave := strings.Contains(string(pedazos[index]), ")")
      if findllave==true{
        parenEnd = parenEnd+1
      }
    }
    fmt.Println(parenEnd)
    //////////Verifica si esta la misma cantidad de parentesis
      if parenApe > parenEnd{
        fmt.Println("Revisa tus perentesis de cierre")
      }else if parenEnd>parenApe{
        fmt.Println("Revisa tus parentesis de apertura")
      }

    checa("#include", string(pedazos[0]))
    fmt.Println(reflect.TypeOf(verifi))//verififica que tpo de datos es




  }
  func checa(palabra string, content string) {
    ver := strings.Contains(string(content), palabra)
    fmt.Println(ver)
  }
