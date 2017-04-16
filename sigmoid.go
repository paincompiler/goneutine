package goneutine

import "math"

type Sigmoid struct {
	bias float64
}

func NewSigmoid(bias float64) Sigmoid {
	return Sigmoid{
		bias,
	}
}

func (sig Sigmoid) Bias() float64 {
	return sig.bias
}

type sigmoidInput struct {
	Input, Weight float64
}

func NewSigmoidInput(i, w float64) sigmoidInput {
	return sigmoidInput{i, w}
}

func (si sigmoidInput) Eval() float64 {
	return float64(si.Input * si.Weight)
}

func (sig Sigmoid) Eval(inputs IWPairs) float64 {
	return σ(inputs.Sum() + sig.Bias())
}

func σ(z float64) float64 {
	return 1.0 / (1 + math.Exp(-z))
}
