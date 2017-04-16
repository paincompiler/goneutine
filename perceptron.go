package goneutine

type Perceptron struct {
	bias int64
}

type perceptronInput struct {
	Input, Weight int64
}

func NewPerceptronInput(i, w int64) perceptronInput {
	return perceptronInput{i, w}
}

func (pi perceptronInput) Eval() float64 {
	return float64(pi.Input * pi.Weight)
}

func NewPerceptron(bias int64) Perceptron {
	return Perceptron{
		bias: bias,
	}
}

func (p Perceptron) Eval(inputs IWPairs) int64 {
	if int64(inputs.Sum())+p.Bias() > 0 {
		return 1
	}
	return 0
}

func (p Perceptron) Bias() int64 {
	return p.bias
}
