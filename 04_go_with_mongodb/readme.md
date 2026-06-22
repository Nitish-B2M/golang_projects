Q: how to check race condition in golang? 
Again, I'm not sure if this is the right place for a question. If it isn't please let me know and I will delete it. Thanks.
I have an application that reads from two different channels simultaneously (one channel to read data in real time while another one to get some statistics). The problem is when reading, sometimes there are race conditions where the first reader gets a value before the second reader has finished writing its values into memory and then it starts overwriting them.
I'm using golang channels for this purpose (see code below) but I don't know if that's the right way to do it or not? Is there any other way of doing this in go? Thanks.
package main
import "fmt"

func read(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Println("Reading: ", <-ch) // this is where the race condition happens
    }
}

func write(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        ch <-i // this is where the race condition happens
     }
}

func main(){
    var wg sync.WaitGroup
    ch := make(chan int)
    for i:= 0; i < 10; i++ {
        wg.Add(2) // this is where the race condition happens
        go read(ch, &wg)
        go write(ch, &wg)
    }
    wg.Wait()
}
Run: $ go run main.go
Check is race condition: $ go run -race main.go
Check is race condition (with output): $ go run -race -v main.go
package main

A: You can use a mutex to ensure that only one goroutine is reading from the channel at any given time. This will prevent race conditions in your code. Here's an example of how you might do this:
package main

import (
    "fmt"
    "sync"
)

func read(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Println("Reading: ", <-ch)
    }
}

func write(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        ch <- i
    }
}

func main(){
    var wg sync.WaitGroup
    ch := make(chan int)
    for i:=0; i < 10; i++ {
        wg.Add(2)
        go read(ch, &wg)
        go write(ch, &wg)
    }
    wg.Wait()
}