package neural

type node struct {
  bias float32
  weights make(map[int]float32)
}