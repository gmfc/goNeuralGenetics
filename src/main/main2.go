package main

import (
  "fmt"
  "neuralArray"
)


func main() {  
  setup := []int{2,3,2,1}
  net := neuralArray.GenerateEmptyNet(setup)
  

  fmt.Println("created net: " ,net)
  net2 := neuralArray.Mutate(net,setup)
  net3 := neuralArray.Mutate(net2,setup)
  fmt.Println("net:" ,net)
  fmt.Println("net2:", net2)
  fmt.Println("net3:", net3)
}