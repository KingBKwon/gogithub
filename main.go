package main

import (
   "fmt"
   "time"
)

func main(){
   for{
      fmt.Println("Hello GO1")
      time.Sleep(10 * time.Second)
   }
}