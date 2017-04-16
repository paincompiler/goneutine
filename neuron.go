package goneutine

type Neuron interface {
	Bias() int64
	Eval() int64
}
