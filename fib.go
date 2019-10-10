/*Package fib created to calculate Fibonacci numbers
This just a test package to deal with benchmarks and tests
*/
package fib

import (
	"math/big"
)

// Fib choose method of calculating Fibonacci number according to 'n' number and calculates it
// See the benchmark results, maybe slower then other functions because of big.Int allocation
func Fib(n uint) *big.Int {
	switch {
	case n < 92:
		return big.NewInt(int64(Matrix(n)))
	case n < 177:
		return IterBig(n)
	default:
		return MatrixBig(n)
	}
}

// Iter calculates Fibonacci number iteratively
func Iter(n uint) uint {
	var a, b uint = 0, 1
	for i := uint(0); i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// IterBig calculates Fibonacci number iteratively using big.Int
func IterBig(n uint) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for i := uint(0); i < n; i++ {
		a, b = b, a.Add(a, b)
	}
	return a
}

// Matrix calculates Fibonacci number with matrix multiplication method
func Matrix(n uint) uint {
	// A matrix
	var a, b, c, d uint = 1, 1, 1, 0
	// Temp A ,matrix for Multiplication
	var ta, tb, tc uint
	// R vector
	var rc, rd uint = 0, 1

	for n > 0 {
		// If pow is even
		if (n & 1) != 0 {
			// Multiply R vector with A matrix
			tc = rc
			rc = rc*a + rd*c
			rd = tc*b + rd*d
		}

		// Multiply A matrix by itself
		ta = a
		tb = b
		tc = c
		a = a*a + b*c
		b = ta*b + b*d
		c = c*ta + d*c
		d = tc*tb + d*d

		// Decreasing pow
		n >>= 1
	}

	return rc
}

// MatrixBig calculates Fibonacci number with matrix multiplication method using big.Int
func MatrixBig(n uint) *big.Int {
	// A matrix
	a, b, c, d := big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0)
	// Temp A matrix for Multiplication
	ta, tb, tc := new(big.Int), new(big.Int), new(big.Int)
	// Temp variables for memory allocation
	t1, t2 := new(big.Int), new(big.Int)
	// R vector
	rc, rd := big.NewInt(0), big.NewInt(1)

	for n > 0 {
		// If pow is even
		if (n & 1) != 0 {
			// Multiply R vector with A matrix
			tc.Set(rc)
			rc.Add(t1.Mul(rc, a), t2.Mul(rd, c)) // rc = rc*a + rd*c
			rd.Add(t1.Mul(tc, b), t2.Mul(rd, d)) //rd = tc*b + rd*d
		}

		// Multiply A matrix by itself
		ta.Set(a)
		tb.Set(b)
		tc.Set(c)
		a.Add(t1.Mul(a, a), t2.Mul(b, c))   // a = a*a + b*c
		b.Add(t1.Mul(ta, b), t2.Mul(b, d))  // b = ta*b + b*d
		c.Add(t1.Mul(c, ta), t2.Mul(d, c))  // c = c*ta + d*c
		d.Add(t1.Mul(tc, tb), t2.Mul(d, d)) // d = tc*tb + d*d

		// Decreasing pow
		n >>= 1
	}

	return rc
}
