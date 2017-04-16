package goneutine

type Perceptron struct {
	bias int64 `json:"bias"`
}

type IWPair struct {
	Input, Weight int64
}

func (iwp IWPair) Eval() int64 {
	return iwp.Input * iwp.Weight
}

func NewPerceptron(bias int64) Perceptron {
	return Perceptron{
		bias: bias,
	}
}

func (p Perceptron) Eval(inputs ...IWPair) int64 {
	var sum int64
	for _, i := range inputs {
		sum += i.Eval()
	}
	if sum+p.Bias() > 0 {
		return 1
	}
	return 0
}

func (p Perceptron) Bias() int64 {
	return p.bias
}
