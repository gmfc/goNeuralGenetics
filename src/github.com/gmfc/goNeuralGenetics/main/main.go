package main

import (
    "fmt"
    "encoding/json"
    "neural"
)


func main() {
/*
  n := NewNode()
  n.bias = 0.2
  n.weights[0] = 0.5
  fmt.Println(n)
  
  l := NewLayer()
  l.nodes[0] = n
  
  nn := NewNetwork()
  nn.layers[0] = l
  fmt.Println(nn)
*/
  
  nInA := neural.NewNode()
  nInB := neural.NewNode()
  
  nHA := neural.NewNode()
  nHA.Weights["0"] = 0.1
  nHA.Weights["1"] = 0.2
  nHA.Bias = 0.5
  
  nHB := neural.NewNode()
  nHB.Weights["0"] = 0.11
  nHB.Weights["1"] = 0.22
  nHB.Bias = 0.5
  
  nOA := neural.NewNode()
  nOA.Weights["0"] = 0.111
  nOA.Weights["1"] = 0.222
  nOA.Bias = 0.5
  
  
  layIn := neural.NewLayer()
  layIn.Nodes["0"] = nInA
  layIn.Nodes["1"] = nInB
  
  layH := neural.NewLayer()
  layH.Nodes["0"] = nHA
  layH.Nodes["1"] = nHB
  
  layO := neural.NewLayer()
  layO.Nodes["0"] = nOA
  
  NET := neural.NewNetwork()
  NET.Layers["0"] = layIn
  NET.Layers["1"] = layH
  NET.Layers["2"] = layO


  b, err := json.Marshal(NET)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(b))
  //fmt.Println(NET)

    


  
}
