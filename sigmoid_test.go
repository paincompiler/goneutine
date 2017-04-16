package goneutine

import "testing"

func TestSigmoid_SimulatePerceptron(t *testing.T) {
	bitwiseSumCases := [][]int64{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
	}
	for i, c := range bitwiseSumCases {
		sum, carry := sBitwiseSum(float64(c[0]), float64(c[1]))
		if sum != c[2] || carry != c[3] {
			t.Logf("expect sum: %d and get %d", c[2], sum)
			t.Logf("expect carry bit: %d and get %d", c[3], carry)
			t.Errorf("failed at %d", i)
		}
	}
}

const infiniteConst = 99999

func sBitwiseSum(x1, x2 float64) (sum, carry int64) {
	var (
		weight float64 = -2 * infiniteConst
		bias   float64 = 3 * infiniteConst
	)
	if x1*weight+bias == 0 || x2*weight+bias == 0 {
		panic("invalid inputs")
	}
	pt0 := NewSigmoid(bias)
	input1 := NewSigmoidInput(x1, weight)
	input2 := NewSigmoidInput(x2, weight)

	interOutput := pt0.Eval(IWPairs{input1, input2})

	return int64(pt0.Eval(
			IWPairs{
				NewSigmoidInput(
					pt0.Eval(IWPairs{
						NewSigmoidInput(interOutput, weight),
						NewSigmoidInput(x1, weight),
					}), weight),
				NewSigmoidInput(
					pt0.Eval(IWPairs{
						NewSigmoidInput(interOutput, weight),
						NewSigmoidInput(x2, weight),
					}), weight),
			})),
		int64(pt0.Eval(IWPairs{NewSigmoidInput(interOutput, weight*2)}))
}
