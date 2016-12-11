package main

import (
    "fmt"
    "time"
)

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}
func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}
func printer1(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg,"1")
        time.Sleep(time.Second * 1)
    }
}
func printer2(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg,"2")
        time.Sleep(time.Second * 1)
    }
}
func main() {
    var c chan string = make(chan string)

    go pinger(c)
	go ponger(c)
    go printer1(c)
	go printer2(c)

    var input string
    fmt.Scanln(&input)
}