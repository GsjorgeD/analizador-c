package main
import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)
func main() {
  file, err := os.Open( "ejem.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  nombres := "while{ }   while{}ihdxhshdfjhdsofodsjoif;"#agarrr //agarrra tofdo menos while
	nombresComoArreglo := strings.Split(nombres, "while")
	for _, nombre := range nombresComoArreglo {
		fmt.Println(nombre)
}
}
//limpiar el codigo y separa por puntos y comaspones un if esls si encvunetra un punto y coma y sino o separe por llaves
//saber cuantos elemntos tiene el archivo
//leer linea por linea y cuando pase de linea sumar uno al putero
//pasar funciones que verifiquen si tiene < debe terminar en > y
//otra funcion que compruebe si que hay un main
//que cuente lass llaves {    deben sier el mismo numero de estas}
//que al finalo tengan ;
