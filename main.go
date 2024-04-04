package main

import (
   "fmt"
   "time"
)

func main(){
   for{
      fmt.Println("Hello GO")
      time.Sleep(10 * time.Second)
   }
}