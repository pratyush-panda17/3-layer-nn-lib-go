package main

import (
	"math"
)

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func dsigmoid(y float64) float64 {
	return y * (1 - y)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func square(x float64) float64 {
	return x * x
}

type NeuralNetwork struct {
	in_nodes  int
	hid_nodes int
	out_nodes int

	weights_ih *Matrix
	weights_ho *Matrix
	bias_ih    *Matrix
	bias_ho    *Matrix

	lr float64
}

func newNetwork(numI int, numH int, numO int) *NeuralNetwork {
	net := new(NeuralNetwork)
	net.in_nodes = numI
	net.hid_nodes = numH
	net.out_nodes = numO

	net.weights_ih = newMatrix(numH, numI)
	net.weights_ho = newMatrix(numO, numH)

	net.weights_ih.randomize()
	net.weights_ih.round(3)

	net.weights_ho.randomize()
	net.weights_ho.round(3)

	net.bias_ih = newMatrix(numH, 1)
	net.bias_ho = newMatrix(numO, 1)

	net.bias_ih.randomize()
	net.bias_ih.round(3)

	net.bias_ho.randomize()
	net.bias_ho.round(3)

	net.lr = 0.1

	return net
}

func (net *NeuralNetwork) feedForward(input []float64) []float64 {
	INPUT := arrToColumnMat(input)
	A := Multiply(net.weights_ih, INPUT)
	A.add(net.bias_ih)
	A.Map(sigmoid)

	OUTPUT := Multiply(net.weights_ho, A)
	OUTPUT.add(net.bias_ho)
	OUTPUT.Map(sigmoid)
	OUTPUT.round(3)
	return flatten(OUTPUT)
}

func (net *NeuralNetwork) backPropogate(OUTPUT_ERRORS *Matrix) *Matrix {
	return Multiply(Transpose(net.weights_ho), OUTPUT_ERRORS)
}

func (net *NeuralNetwork) train(input []float64, target []float64) {
	INPUT := arrToColumnMat(input)

	OUTPUT_FROM_HIDDEN := Multiply(net.weights_ih, INPUT)
	OUTPUT_FROM_HIDDEN.add(net.bias_ih)
	OUTPUT_FROM_HIDDEN.Map(sigmoid)

	OUTPUT := Multiply(net.weights_ho, OUTPUT_FROM_HIDDEN)
	OUTPUT.add(net.bias_ho)
	OUTPUT.Map(sigmoid)

	TARGET := arrToColumnMat(target)

	OUTPUT_ERRORS := Subtract(TARGET, OUTPUT)

	A := ScaleByColMatrix(Map(OUTPUT, dsigmoid), OUTPUT_ERRORS)
	A.scale(net.lr)

	UPDATE_WEIGHTS_HO := Multiply(A, Transpose(OUTPUT_FROM_HIDDEN))
	net.weights_ho.add(UPDATE_WEIGHTS_HO)
	net.weights_ho.round(3) //Hidden to output weights updated

	net.bias_ho.add(A)
	net.bias_ho.round(3) //Hidden to Output biases Updated

	HIDDEN_ERRORS := net.backPropogate(OUTPUT_ERRORS)
	B := ScaleByColMatrix(Map(OUTPUT_FROM_HIDDEN, dsigmoid), HIDDEN_ERRORS)
	B.scale(net.lr)

	UPDATE_WEIGHTS_IH := Multiply(B, Transpose(INPUT))
	net.weights_ih.add(UPDATE_WEIGHTS_IH)
	net.weights_ih.round(3) //Input to Hidden weights updated

	net.bias_ih.add(B)
	net.bias_ih.round(3) //Input to Bias weights updated

}

func (net *NeuralNetwork) mse(input []float64, target []float64) float64 {
	output := net.feedForward(input)
	var mse float64 = 0
	for i := 0; i < net.out_nodes; i++ {
		mse += square(target[i] - output[i])
	}
	return mse
}

func (net *NeuralNetwork) avg_mse(inputs [][]float64, targets [][]float64) float64 {
	var mse float64 = 0
	for i := 0; i < len(inputs); i++ {
		mse += net.mse(inputs[i], targets[i])
	}
	return mse / float64(len(inputs))
}

func main() {

}
