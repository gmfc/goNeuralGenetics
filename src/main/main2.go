package main

import (
  "fmt"
  "neuralArray"
)


func main() {  
  setup := []int{2,2,1}
  net := neuralArray.GenerateEmptyNet(setup)
  

  fmt.Println("created net: " ,net)


  
  inp := []float64{1,1}
  
  
  
  fmt.Println("out:", neuralArray.Run(net,inp))
}