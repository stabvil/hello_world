package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("README.md")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)
}