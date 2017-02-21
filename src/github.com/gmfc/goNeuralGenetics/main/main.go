package main

import (
    "fmt"
    "neural"
    "os"
)

func generateXORNet() neural.Network {
  
  nInA := neural.NewNode()
  nInB := neural.NewNode()
  
  nHA := neural.NewNode()
  nHA.Weights["0"] = 9.945131391752511
  nHA.Weights["1"] = -7.065739412326366
  nHA.Bias = -5.085902468767017
  
  nHB := neural.NewNode()
  nHB.Weights["0"] = 5.465975496452302
  nHB.Weights["1"] = -3.934639404527843
  nHB.Bias = 2.086773671209812
  
  nOA := neural.NewNode()
  nOA.Weights["0"] = 10.26414252538234
  nOA.Weights["1"] = -9.913011252414435
  nOA.Bias = 4.650809367187321
  
  
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
  
  return NET
  
  
}


func main() {
  
  NET := generateXORNet()
  out := neural.RunNetwork(map[string]float64{ "0":0, "1":1 },NET)
  fmt.Println(out["0"])
  
}
