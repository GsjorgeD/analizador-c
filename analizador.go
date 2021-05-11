package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "os"
    "bufio"
)

func main() {
  var llaveApe int = 0
  var llaveEnd int = 0
  var parenApe int = 0
  var parenEnd int = 0
  var contadorDo int = 0
  var contadorWhile int = 0
  var numMain int = 0
  var archivo string="ejem.c"

    //***lectura del archivo
  content, err := ioutil.ReadFile(archivo)
    if err != nil {
      log.Fatal(err)
      }
    //***
    secciones := strings.Split(string(content), " ") //Separa el archivo
    findInclude := strings.Contains(string(content), "#include")//Busca include y regresa un valor Booleano
    numInclude := len(secciones[0])//Cuenta cuantos elementos hay en include
    if findInclude==false{
    fmt.Println("Error generado: revisa en '#include'")
    }else if numInclude !=8{
      fmt.Println("Error generado: revisa en '#include'")
    }
    //***Verifica los elementos de la libreria
    simbMenor:= strings.Contains(string(secciones[1]), "<")
    simbMayor := strings.Contains(string(secciones[1]), ">")
    letH := strings.Contains(string(secciones[1]), ".h")
    if simbMayor==false || simbMenor==false || letH==false{
      fmt.Println("Error generado en la libreria '<......h>'")
    }
    //***
    ////***Verifica nuevamente las librerias
    content2, err2 := os.Open(archivo) //**archivo**
   	if err2 != nil {
  	  log.Fatal(err2)
  	}
  	contentScaner := bufio.NewScanner(content2)
  	for contentScaner.Scan() {
  		if strings.Contains(contentScaner.Text(), "#include") {
  			if strings.Contains(contentScaner.Text(), "<") {
  				if strings.Contains(contentScaner.Text(), ">") {
  					if strings.Contains(contentScaner.Text(), ".h") {
  					} else {
  						fmt.Println("Falta el .h")
  					}
  				} else {
  					fmt.Println("Falta el >")
  				}
  			} else {
  				fmt.Println("Falta el <")
  			}
  		} else {

  			if strings.Contains(contentScaner.Text(), "<") && strings.Contains(contentScaner.Text(), ".h") && strings.Contains(contentScaner.Text(), ">") {
  				fmt.Println("Falta el '#include'")
  			} else if strings.Contains(contentScaner.Text(), "#include") && strings.Contains(contentScaner.Text(), ".h") && strings.Contains(contentScaner.Text(), ">") {
  				fmt.Println("Falta el <")
  			} else if strings.Contains(contentScaner.Text(), "#include") && strings.Contains(contentScaner.Text(), "<") && strings.Contains(contentScaner.Text(), ".h") {
  				fmt.Println("Falta el >")
  			}
  		}

  	}
    //***
    numSecciones := len(secciones)//***Cuenta cuantas secciones hay***
    //***Verifica main
    findMain := strings.Contains(string(content), "main")//verifica si esta main y regresa un booleano
    if findMain==false{
      fmt.Println("Error generado: main no esta definido")
    }
    for index := 0; index < numSecciones; index++ { //***Cuentacuantos main hay***
    findNumMain := strings.Contains(string(secciones[index]), "main")
      if findNumMain==true{
        numMain = numMain+1
      }
    }
    if numMain >1{
      fmt.Println("Error generado: la funcion main fue definida previamente")
    }
    //***
    //***Verifica ciclo do while
    //conten3 err3 contentscanner2 contador do y contadorWhile
    content3, err3 := os.Open(archivo)//**archivo**
    if err3 != nil {
      log.Fatal(err3)
    }
    contentScaner2 := bufio.NewScanner(content3)
    for contentScaner2.Scan() {
      if strings.Contains(contentScaner2.Text(), "do") {
        contadorDo++
      }
      if strings.Contains(contentScaner2.Text(), "while") && strings.Contains(contentScaner2.Text(), "(") && strings.Contains(contentScaner2.Text(), ")") {
        contadorWhile++
      }
    }
    if contadorDo == contadorWhile {
    } else {
      fmt.Println("Error generado: revisa la estructura del ciclo 'do while'")
    }
    //***
    //Identifica si se declara una variable del tipo int y verifica si ;
     for puntero := 2; puntero < numSecciones; puntero++{
     findInt := strings.Contains(string(secciones[puntero]), "int")
     findPrint := strings.Contains(string(secciones[puntero]), "print")
       if findInt == true && findPrint==false{
         puntero++
         findPunto := strings.Contains(string(secciones[puntero]), ";")
         if findPunto == false{
           puntero++
           findPunto2 := strings.Contains(string(secciones[puntero]), ";")
           if findPunto2 == false{
             fmt.Println("Error generado: falta ';' en la variable int")
           }
         }
        }
       }
      //***
      //***Identifica si se declara una variable del tipo int y verifica si ;
     //conten4 err4 contentscanner3
     content4, err4:=os.Open(archivo)//**archivo**

     if err4 != nil {
   		log.Fatal(err4)
   	}
     contentScaner3 := bufio.NewScanner(content4)
     for contentScaner3.Scan(){
       if strings.Contains(contentScaner3.Text(), "char"){
         if  strings.Contains(contentScaner3.Text(), " ") {
           if strings.Contains(contentScaner3.Text(), ";") {
           }else {
             fmt.Println("Error generado: falta ';' en la variable char")
             }
          }else {
           fmt.Println("Falta asignar valor")
         }
     }
     }
     //***
    //***Busca llaves de apertura
    for index := 0; index < numSecciones; index++ {
    findLlave := strings.Contains(string(secciones[index]), "{")
      if findLlave==true{
        llaveApe = llaveApe+1
      }
    }
    //***
    //Busca llaves de cerradura
    for index := 0; index < numSecciones; index++ {
    findLlave := strings.Contains(string(secciones[index]), "}")
      if findLlave==true{
        llaveEnd = llaveEnd+1
      }
    }
    //***
    //**Verifica si esta la misma cantidad de llaves
    if llaveApe > llaveEnd{
      fmt.Println("Error generado: revisa tus llaves de cierre '}'")
    }else if llaveEnd>llaveApe{
      fmt.Println("Error generado: revisa tus llaves de apertura '{'")
    }
    //***
    //Verifica parentesis de apertura
    for index := 0; index < numSecciones; index++ {
    findLlave := strings.Contains(string(secciones[index]), "(")
      if findLlave==true{
        parenApe = parenApe+1
      }
    }
    //***
    //Verifica parentesis de cierre
    for index := 0; index < numSecciones; index++ {
    findLlave := strings.Contains(string(secciones[index]), ")")
      if findLlave==true{
        parenEnd = parenEnd+1
      }
    }
    //***
    //***Verifica si esta la misma cantidad de parentesis
      if parenApe > parenEnd{
        fmt.Println("Error generado: revisa tus parentesis de cierre ')'")
      }else if parenEnd>parenApe{
        fmt.Println("Error generado: revisa tus parentesis de apertura '('")
      }
    //***
    fmt.Print("\n*** Se mostraran mensajes de error hasta que se solucionen,")
    fmt.Print(" 'compila y ejecuta si no se muestra ninguno' ***\n")
}
