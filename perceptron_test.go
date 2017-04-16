package goneutine

import "testing"

func TestPerceptron_Evaluate(t *testing.T) {
	bitwiseSumCases := [][]int64{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
	}
	for i, c := range bitwiseSumCases {
		sum, carry := bitwiseSum(c[0], c[1])
		if sum != c[2] || carry != c[3] {
			t.Logf("expect sum: %d and get %d", c[2], sum)
			t.Logf("expect carry bit: %d and get %d", c[3], carry)
			t.Errorf("failed at %d", i)
		}
	}
}

func bitwiseSum(x1, x2 int64) (sum, carry int64) {
	var (
		weight int64 = -2
		bias   int64 = 3
	)
	pt0 := NewPerceptron(bias)
	input1 := NewPerceptronInput(x1, weight)
	input2 := NewPerceptronInput(x2, weight)

	interOutput := pt0.Eval(IWPairs{input1, input2})

	return pt0.Eval(
		IWPairs{
			NewPerceptronInput(
				pt0.Eval(IWPairs{
					NewPerceptronInput(interOutput, weight),
					NewPerceptronInput(x1, weight),
				}), weight),
			NewPerceptronInput(
				pt0.Eval(IWPairs{
					NewPerceptronInput(interOutput, weight),
					NewPerceptronInput(x2, weight),
				}), weight),
		}),
		pt0.Eval(IWPairs{NewPerceptronInput(interOutput, weight * 2)})
}
