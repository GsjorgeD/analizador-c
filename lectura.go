package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // read the whole file at once
    b, err := ioutil.ReadFile("ejem.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(err)

    // write the whole body at once
    err = ioutil.WriteFile("ejem.txt", b, 0644)
    if err != nil {
        panic(err)
    }
}
