package fib

import (
	"fmt"
	"math/big"
	"testing"
)

const BenchDepth = uint(500)

// Type for benchmarks to hide all fnuctions behind one type
type testFunc func(uint) interface{}

func tFibIter(a uint) interface{} {
	return Iter(a)
}

func tFibMatrix(a uint) interface{} {
	return Matrix(a)
}

func tFib(a uint) interface{} {
	return Fib(a)
}

func tFibIterBig(a uint) interface{} {
	return IterBig(a)
}

func tFibMatrixBig(a uint) interface{} {
	return MatrixBig(a)
}

func GenericBench(fibFunc testFunc, b *testing.B) {
	for d := uint(8); d < BenchDepth; d = (d / 2) * 3 {
		b.Run(fmt.Sprintf("%4d Fibonacci number", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibFunc(d)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	GenericBench(tFib, b)
}

func BenchmarkFibIter(b *testing.B) {
	GenericBench(tFibIter, b)
}

func BenchmarkFibMatrix(b *testing.B) {
	GenericBench(tFibMatrix, b)
}

func BenchmarkFibIterBig(b *testing.B) {
	GenericBench(tFibIterBig, b)
}

func BenchmarkFibMatrixBig(b *testing.B) {
	GenericBench(tFibMatrixBig, b)
}

// Test function returning a slice of Fibonacci numbers form 0 to 90
func FibNumbers() []uint {
	return []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
		55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181,
		6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229,
		832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986,
		102334155, 165580141, 267914296, 433494437, 701408733,
		1134903170, 1836311903, 2971215073, 4807526976, 7778742049,
		12586269025, 20365011074, 32951280099, 53316291173, 86267571272,
		139583862445, 225851433717, 365435296162, 591286729879, 956722026041,
		1548008755920, 2504730781961, 4052739537881, 6557470319842, 10610209857723,
		17167680177565, 27777890035288, 44945570212853, 72723460248141, 117669030460994,
		190392490709135, 308061521170129, 498454011879264, 806515533049393, 1304969544928657,
		2111485077978050, 3416454622906707, 5527939700884757, 8944394323791464, 14472334024676221,
		23416728348467685, 37889062373143906, 61305790721611591, 99194853094755497, 160500643816367088,
		259695496911122585, 420196140727489673, 679891637638612258, 1100087778366101931, 1779979416004714189,
	}
}

// types only for testing
type fibFunc func(uint) uint
type fibFuncBig func(uint) *big.Int

func GenericTest(f fibFunc, t *testing.T) {
	for i, fibN := range FibNumbers() {
		if f(uint(i)) != fibN {
			t.Fatalf("Calculating #%d Fibonacci number. Want %d, but got %d\n",
				i, FibNumbers()[i], f(uint(i)),
			)
		}
	}
}

func TestFibIter(t *testing.T) {
	GenericTest(Iter, t)
}

func TestFibMatrix(t *testing.T) {
	GenericTest(Matrix, t)
}

func GenericTestBig(f fibFuncBig, t *testing.T) {
	for i, fibN := range FibNumbers() {
		if r := f(uint(i)); r.Cmp(big.NewInt(int64(fibN))) != 0 {
			t.Fatalf("Calculating #%d Fibonacci number. Want %d, but got %d\n",
				i, FibNumbers()[i], r,
			)
		}
	}
}

func TestFibIterBig(t *testing.T) {
	GenericTestBig(IterBig, t)
}

func TestFibMatrixBig(t *testing.T) {
	GenericTestBig(MatrixBig, t)
}

func TestFib(t *testing.T) {
	GenericTestBig(Fib, t)
}
