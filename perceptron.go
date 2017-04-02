package goneutine

type Perceptron struct {
	Bias int64 `json:"bias"`
}

type IWPair struct {
	Input, Weight int64
}

func (iwp IWPair) Eval() int64 {
	return iwp.Input * iwp.Weight
}

func NewPerceptron(bias int64) Perceptron {
	return Perceptron{
		Bias: bias,
	}
}

func (p Perceptron) Output(inputs ...IWPair) int64 {
	var sum int64
	for _, i := range inputs {
		sum += i.Eval()
	}
	if sum+p.Bias > 0 {
		return 1
	}
	return 0
}
