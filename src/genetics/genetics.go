package genetics

import (
  //"fmt"
  "neural"
  "time"
  "math/rand"
)

type Phenotype struct {
  Fitness float64
  Net neural.Network
}

type Population struct {
  pop []Phenotype
}

type Genetic struct {
  pop Population
  popSize int
}
/*
func mutate(phen Phenotype) Phenotype{
  
}*/

func MutateNet(net neural.Network) neural.Network{
  
  ran :=rand.New(rand.NewSource(time.Now().UnixNano()))
  for layId := range net.Layers {
    if layId!= "0" {
      
      for nodeID, node := range net.Layers[layId].Nodes {
        nNode := neural.NewNode()
        nNode.Bias = node.Bias + (ran.Float64() * 2) - 1

        for weiID, wei := range node.Weights {
          nWei := wei + (ran.Float64() * 2) - 1
          nNode.Weights[weiID] = nWei
        }
        
        net.Layers[layId].Nodes[nodeID] = nNode
        
      }
      
    }
    
  }
  return net
}

/*
func MutateNet(net neural.Network) neural.Network {
  ran :=rand.New(rand.NewSource(time.Now().UnixNano()))

  for layId, _ := range net.Layers {
    if layId!= "0" {

      for nodID, _ := range net.Layers[layId].Nodes {
        net.Layers[layId].Nodes[nodID].Bias += (ran.Float64() * 2) - 1
        
        // ASYNC STUFF!!
        done := make(chan bool) // FLAG CHANEL
        
        for wID, _ := range net.Layers[layId].Nodes[nodID].Weights {
          go func(ww *float64) {
            *ww += (ran.Float64() * 2) - 1
            done <- true
          }(&net.Layers[layId].Nodes[nodID].Weights[wID])
        }
        // Wait complete
        for _ = range net.Layers[layId].Nodes[nodID].Weights {
          <-done
        }
      }
    }
  }
  return net
}*/

func (this *Genetic) evolve() {
  
}

func (this *Genetic) populate() {
  
}

/*
Dado um grafo, checar se G é um cirquito
*/