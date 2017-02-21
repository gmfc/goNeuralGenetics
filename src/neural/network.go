package neural

type Node struct {
  Bias float32
  Weights map[string]float32
}

type Layer struct {
  Nodes map[string]Node
}

type Network struct {
  Layers map[string]Layer
}

func NewNode() Node {
  var n Node
  n.Weights = make(map[string]float32)
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