package neural

import (
   //"fmt"
  "math"
)

type Node struct {
  Bias float64
  Weights map[string]float64
}

type Layer struct {
  Nodes map[string]Node
}

type Network struct {
  Layers map[string]Layer
}

func NewNode() Node {
  var n Node
  n.Weights = make(map[string]float64)
  return n
}

func NewLayer() Layer {
  var l Layer
  l.Nodes = make(map[string]Node)
  return l
}

func NewNetwork() Network {
  var n Network
  n.Layers = make(map[string]Layer)
  return n
}

func (net *Network)RunNetwork(input map[string]float64) map[string]float64 {
  // m := map[string]string{ "key1":"val1", "key2":"val2" };
  var output map[string]float64
  for layId, lay := range net.Layers {
    if layId!= "0" {
      output =  make(map[string]float64)
      for nodId, nod := range lay.Nodes {
        sum := nod.Bias
        
        // ASYNC STUFF!!
        done := make(chan bool) // FLAG CHANEL
        
        for from, weigh := range nod.Weights {
          go func(ww float64, in float64) {
            sum += ww * in
            done <- true
          }(weigh,input[from])
        }
        // Wait complete
        for _ = range nod.Weights {
          <-done
        }
        output[nodId] = 1 / (1 + math.Exp(-sum)) 
      }
      input = output
    }
  }
  return output
}