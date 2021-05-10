package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
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

	verifi := (strings.Contains(string(content), "#include") && strings.Contains(string(content), " <") && strings.Contains(string(content), ">")) //hacer una funcion que cheque argumentos
	//fmt.Println(verifi)
	if verifi == false {
		fmt.Println("error en include")
	}

	pedazos := strings.Split(string(content), " ")

	//***ignorar fmt.Println(pedazos)//pofdria hacer una funcion con esto y pasarle el argumento que quier
	//hacer con while que por cada line cheque hacer para las n ves de lineas que hayya
	//si existe sumar a uno y llevar la cuenta y despues comparara
	numpedazos := len(pedazos)
	////////busca llaves de apertura
	for index := 0; index < numpedazos; index++ {
		//checa("{", string(pedazos[index]))
		findllave := strings.Contains(string(pedazos[index]), "{")
		if findllave == true {
			llaveApe = llaveApe + 1
		}
	}
	fmt.Println(llaveApe)
	////////busca llaves de cerradura
	for index := 0; index < numpedazos; index++ {
		findllave := strings.Contains(string(pedazos[index]), "}")
		if findllave == true {
			llaveEnd = llaveEnd + 1
		}
	}
	fmt.Println(llaveEnd)
	//////////Verifica si esta la misma cantidad de llaves
	if llaveApe > llaveEnd {
		fmt.Println("Revisa tus llaves de cierre")
	} else if llaveEnd > llaveApe {
		fmt.Println("Revisa tus llaves de apertura")
	}
	///////////verifica parentesis de apertura
	for index := 0; index < numpedazos; index++ {
		findllave := strings.Contains(string(pedazos[index]), "(")
		if findllave == true {
			parenApe = parenApe + 1
		}
	}
	fmt.Println(parenApe)
	///////////verifica parentesis de apertura
	for index := 0; index < numpedazos; index++ {
		findllave := strings.Contains(string(pedazos[index]), ")")
		if findllave == true {
			parenEnd = parenEnd + 1
		}
	}
	fmt.Println(parenEnd)
	//////////Verifica si esta la misma cantidad de parentesis
	if parenApe > parenEnd {
		fmt.Println("Revisa tus perentesis de cierre")
	} else if parenEnd > parenApe {
		fmt.Println("Revisa tus parentesis de apertura")
	}

	checa("#include", string(pedazos[0]))
  a:=(reflect.TypeOf(22)) //verififica que tpo de datos es
	fmt.Println(a)
}
func checa(palabra string, content string) {
	ver := strings.Contains(string(content), palabra)
	fmt.Println(ver)
}
