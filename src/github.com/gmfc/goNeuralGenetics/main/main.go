package main

import (
    "fmt"
    "encoding/json"
)

type node struct {
  Bias float32
  Weights map[string]float32
}

type layer struct {
  Nodes map[string]node
}

type network struct {
  Layers map[string]layer
}

func NewNode() node {
  var n node
  n.Weights = make(map[string]float32)
  return n
}

func NewLayer() layer {
  var l layer
  l.Nodes = make(map[string]node)
  return l
}

func NewNetwork() network {
  var n network
  n.Layers = make(map[string]layer)
  return n
}

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
  
  nInA := NewNode()
  nInB := NewNode()
  
  nHA := NewNode()
  nHA.Weights["0"] = 0.1
  nHA.Weights["1"] = 0.2
  nHA.Bias = 0.5
  
  nHB := NewNode()
  nHB.Weights["0"] = 0.11
  nHB.Weights["1"] = 0.22
  nHB.Bias = 0.5
  
  nOA := NewNode()
  nOA.Weights["0"] = 0.111
  nOA.Weights["1"] = 0.222
  nOA.Bias = 0.5
  
  
  layIn := NewLayer()
  layIn.Nodes["0"] = nInA
  layIn.Nodes["1"] = nInB
  
  layH := NewLayer()
  layH.Nodes["0"] = nHA
  layH.Nodes["1"] = nHB
  
  layO := NewLayer()
  layO.Nodes["0"] = nOA
  
  NET := NewNetwork()
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
