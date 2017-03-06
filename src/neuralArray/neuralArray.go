package neuralArray

import (
  // "fmt"
  "math"
)

func GenerateEmptyNet(setup []int) [][][]float64 {
  numberOfLayers := len(setup)
  N := make([][][]float64, numberOfLayers)// create neuralNet
  for l := 0; l < numberOfLayers; l++ {// fore layer...
    neuronsOnLayer := setup[l]//see how many neurons it should have
    N[l] = make([][]float64, neuronsOnLayer)// create layer in network
    for neu := 0; neu < neuronsOnLayer; neu++ {// fore neuron on layer...
      // Size of neuron array is bias + weights, Ex.: L0 has 3 neurons, neuron on L1 will have setup[0](3) + 1 (bias) elements in array (total 4)
      // ALL NEURONS ON LAYER 0 ARE EMPTY!!
      elements := 0
      if l>0 {
        elements = setup[l-1]+1 // if neuron not in L0 setup array size
      }
      N[l][neu] = make([]float64, elements)// create neuron array
      for elem := 0; elem < elements; elem++ {// fore element in neuron array...
        N[l][neu][elem] = 1 // initiate as zero
      }
    }
  }
  return N
}


func Mutate(net [][][]float64,setup []int) [][][]float64 {
  NewNet := GenerateEmptyNet(setup)
  for l:=0;l<len(net);l++{// fore layer...
    for n:=0;n<len(net[l]);n++{// fore neuron...
      for a:=0;a<len(net[l][n]);a++ {// fore element in neuron array...
        NewNet[l][n][a] = net[l][n][a] + 1// set new value
      }
    }
  }
  return NewNet
}

// Parece que funciona, nao sei bem como...
func Run(net [][][]float64, input []float64) []float64 {
  //var output []float64
  for l:=1;l<len(net);l++{// fore layer... not layer 0. jump over input
    output := make([]float64, len(net[l])) // reset output layer buffer 
    for n:=0;n<len(net[l]);n++{// fore neuron...
      sum := net[l][n][0] // set sum as bias
      for a:=1;a<len(net[l][n]);a++ {// fore element in neuron array... not element 0. jump over bias
        sum += net[l][n][a] * input[a-1] // sum and multiply by input at position a
      }
      output[n] = 1 / (1 + math.Exp(-sum)) // Sigmoid here
    }
    input = output
  }
  return input
}