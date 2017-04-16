package goneutine

type IWPair interface {
	Eval() float64
}

type IWPairs []IWPair

func (pairs IWPairs) Sum() float64 {
	var result float64
	for _, p := range pairs {
		result += p.Eval()
	}
	return result
}
