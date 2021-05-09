///manda un error la funcion

checa("}", content, numpedazos)


func checa(palabra string, content []byte, numpedazos int) {
  var llave int = 0
  for index := 0; index < numpedazos; index++ {
  ver := strings.Contains(string(content[index]), palabra)
  if ver==true{
    llave = llave+1
  }
  }
  fmt.Println(llave)
}
