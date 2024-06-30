package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix struct {
	matrix [][]float64
	rows   int
	cols   int
}

func round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

func initDoubleArr(rows int, cols int) [][]float64 {
	mat := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, cols)
	}
	return mat
}

func newMatrix(rows int, cols int) *Matrix {
	var mat *Matrix = new(Matrix)
	mat.rows = rows
	mat.cols = cols
	mat.matrix = initDoubleArr(rows, cols)
	return mat
}

func (mat *Matrix) insert(data [][]float64) {
	if len(data) == mat.rows && len(data[0]) == mat.cols {
		for i := 0; i < mat.rows; i++ {
			mat.matrix[i] = data[i]
		}
	}
}

func (mat *Matrix) change(rows int, cols int, data [][]float64) {
	if rows-mat.rows > 0 {
		for i := 0; i < rows-mat.rows; i++ {
			mat.matrix = append(mat.matrix, []float64{})
		}
	}

	if rows-mat.rows < 0 {
		mat.matrix = mat.matrix[0:rows]
	}
	mat.rows = rows
	mat.cols = cols

	for i := 0; i < mat.rows; i++ {
		mat.matrix[i] = data[i]
	}
}

func (mat *Matrix) randomize() {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mat.matrix[i][j] = (rand.Float64() * 2) - 1
		}
	}
}

func (mat *Matrix) scale(n float64) {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mat.matrix[i][j] = mat.matrix[i][j] * n
		}
	}
}

func Scale(mat *Matrix, n float64) *Matrix {
	scaled := newMatrix(mat.rows, mat.cols)
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			scaled.matrix[i][j] = mat.matrix[i][j] * n
		}
	}
	return scaled
}

func (mat *Matrix) shift(n float64) {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mat.matrix[i][j] = mat.matrix[i][j] + n
		}
	}
}

func Shift(mat *Matrix, n float64) *Matrix {
	shifted := newMatrix(mat.rows, mat.cols)

	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			shifted.matrix[i][j] = mat.matrix[i][j] + n
		}
	}
	return shifted
}

func (mat *Matrix) add(mat2 *Matrix) {
	if mat2.rows == mat.rows && mat2.cols == mat.cols {
		for i := 0; i < mat.rows; i++ {
			for j := 0; j < mat.cols; j++ {
				mat.matrix[i][j] = mat.matrix[i][j] + mat2.matrix[i][j]
			}
		}
	}
}

func Add(mat1 *Matrix, mat2 *Matrix) *Matrix {
	sum := newMatrix(mat1.rows, mat1.cols)
	if mat1.rows == mat2.rows && mat1.cols == mat2.cols {
		for i := 0; i < sum.rows; i++ {
			for j := 0; j < sum.cols; j++ {
				sum.matrix[i][j] = mat1.matrix[i][j] + mat2.matrix[i][j]
			}
		}
	}
	return sum
}

func (mat *Matrix) subtract(mat2 *Matrix) {
	if mat2.rows == mat.rows && mat2.cols == mat.cols {
		for i := 0; i < mat.rows; i++ {
			for j := 0; j < mat.cols; j++ {
				mat.matrix[i][j] = mat.matrix[i][j] - mat2.matrix[i][j]
			}
		}
	}
}

func Subtract(mat1 *Matrix, mat2 *Matrix) *Matrix {
	diff := newMatrix(mat1.rows, mat1.cols)
	if mat1.rows == mat2.rows && mat1.cols == mat2.cols {
		for i := 0; i < diff.rows; i++ {
			for j := 0; j < diff.cols; j++ {
				diff.matrix[i][j] = mat1.matrix[i][j] - mat2.matrix[i][j]
			}
		}
	}
	return diff
}

func (mat *Matrix) multiply(mat2 *Matrix) {
	product := initDoubleArr(mat.rows, mat2.cols)
	if mat.cols == mat2.rows {
		var p float64
		for i := 0; i < mat.rows; i++ {
			for j := 0; j < mat2.cols; j++ {
				p = 0
				for k := 0; k < mat.cols; k++ {
					p = p + (mat.matrix[i][k] * mat2.matrix[k][j])
				}
				product[i][j] = p
			}
		}
		mat.cols = mat2.cols
		mat.insert(product)
	}
}

func Multiply(mat1 *Matrix, mat2 *Matrix) *Matrix {
	prod := newMatrix(mat1.rows, mat2.cols)
	if mat1.cols == mat2.rows {
		var p float64
		for i := 0; i < prod.rows; i++ {
			for j := 0; j < prod.cols; j++ {
				p = 0
				for k := 0; k < mat1.cols; k++ {
					p = p + (mat1.matrix[i][k] * mat2.matrix[k][j])
				}
				prod.matrix[i][j] = p
			}
		}
	}
	return prod
}

func (mat *Matrix) transpose() {
	trans := initDoubleArr(mat.cols, mat.rows)
	for i := 0; i < mat.cols; i++ {
		for j := 0; j < mat.rows; j++ {
			trans[i][j] = mat.matrix[j][i]
		}
	}
	mat.change(mat.cols, mat.rows, trans)
}

func Transpose(mat *Matrix) *Matrix {
	trans := newMatrix(mat.cols, mat.rows)
	for i := 0; i < trans.rows; i++ {
		for j := 0; j < trans.cols; j++ {
			trans.matrix[i][j] = mat.matrix[j][i]
		}
	}
	return trans
}

func (mat *Matrix) Map(f func(x float64) float64) {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mat.matrix[i][j] = f(mat.matrix[i][j])
		}
	}
}

func Map(mat *Matrix, f func(x float64) float64) *Matrix {
	mapped := newMatrix(mat.rows, mat.cols)
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mapped.matrix[i][j] = f(mat.matrix[i][j])
		}
	}
	return mapped
}

func flatten(mat *Matrix) []float64 {
	var arr []float64 = make([]float64, 0, mat.rows*mat.cols)
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			arr = append(arr, mat.matrix[i][j])
		}
	}
	return arr
}

func identity(n int) *Matrix {
	ident := newMatrix(n, n)
	for i := 0; i < n; i++ {
		ident.matrix[i][i] = 1
	}
	return ident
}

func (mat *Matrix) scaleRow(row int, c float64) {
	for i := 0; i < mat.cols; i++ {
		mat.matrix[row][i] = c * mat.matrix[row][i]
	}
}

func (mat *Matrix) addRow(row1 int, c float64, row2 int) {
	for i := 0; i < mat.cols; i++ {
		mat.matrix[row2][i] = (c * mat.matrix[row1][i]) + mat.matrix[row2][i]
	}
}

func (mat *Matrix) swapRow(row1 int, row2 int) {
	mat.matrix[row1], mat.matrix[row2] = mat.matrix[row2], mat.matrix[row1]
}

func arrToColumnMat(arr []float64) *Matrix {
	mat := newMatrix(len(arr), 1)
	for i := 0; i < mat.rows; i++ {
		mat.matrix[i][0] = arr[i]
	}
	return mat
}

func arrToRowMat(arr []float64) *Matrix {
	mat := newMatrix(1, len(arr))
	mat.matrix[0] = arr
	return mat
}

func ScaleByColMatrix(mat *Matrix, col_mat *Matrix) *Matrix {
	scaled := newMatrix(mat.rows, mat.cols)
	scaled.insert(mat.matrix)
	if scaled.rows == col_mat.rows && col_mat.cols == 1 {
		for i := 0; i < scaled.rows; i++ {
			scaled.scaleRow(i, col_mat.matrix[i][0])
		}
	}
	return scaled
}

func (mat *Matrix) round(prec int) {
	for i := 0; i < mat.rows; i++ {
		for j := 0; j < mat.cols; j++ {
			mat.matrix[i][j] = round(mat.matrix[i][j], prec)
		}
	}
}

func (mat *Matrix) print() {
	fmt.Println(mat.matrix)
}
