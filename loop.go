package main

import "fmt"
import "time"

func loop() {
	var s float32
	for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}