package main

import (
	"container/list"
	"fmt"
	"math/big"
)

// to be used int calculating positive ints in func factorNumIntoPositiveInts
type PositiveInt struct {
	x int64
	y int64
}

func pow(base int64, exponent int64) int64 {
	var output int64 = 1.0
	for exponent > 0 {
		output *= base
		exponent -= 1.0
	}

	return output
}

func superPow(base int64, exponent int64) big.Int {
	var limit big.Int
	limit.Exp(big.NewInt(base), big.NewInt(exponent), nil)

	return limit
}

func isPrime(num int64) bool {
	// all numbers less than 2 are not primes
	if num <= 1 {
		return false
	}

	for i := int64(2); i < num; i++ {
		if (num % i) == 0 {
			return false
		}
	}
	return true
}

func primesThatDivideNum(num int64) *list.List {
	primes := list.New()

	for i := int64(1); i <= num; i++ {
		if ((num % i) == 0) && (isPrime(i)) {
			primes.PushFront(i)
		}
	}

	return primes
}

func calcPhi(num int64) int64 {
	/*
		phi(12) = 12 * ((2-1)*(3-1)) / (2*3) = 12 * 2 / 6 = 4

		a = 12
		b = ((2-1)*(3-1))
		c = (2*3)

		phi(a) = a * b / c
	*/
	a := num

	primes := primesThatDivideNum(num)
	if primes.Len() == 0 {
		return 1
	}

	// b = (2-1)*(3-1)
	var b int64 = 1
	_each(primes, func(i interface{}) { b *= i.(int64) - 1 })

	// c = 2*3
	var c int64 = 1
	_each(primes, func(i interface{}) { c *= i.(int64) })

	return a * b / c
}

func factorNumIntoPositiveInts(num int64) *list.List {
	factors := list.New()

	for i := int64(1); i <= num; i++ {
		if (num % i) == 0 {
			factors.PushFront(PositiveInt{x: i, y: num / i})
		}
	}

	return factors
}

/*
  necklaces(k, n) = 1/n * Sum of (phi(a) k^b)
  1/10 (phi(1) 3^10 + phi(2) 3^5 + phi(5) 3^2 + phi(10) 3^1)
  1/10 (1 * 59049 + 1 * 243 + 4 * 9 + 4 * 3)
  necklaces(k, n) = x * y
*/
func calcNeclaces(k int64, n int64) int64 {
	// (phi(1) 3^10 + phi(2) 3^5 + phi(5) 3^2 + phi(10) 3^1)
	factors := factorNumIntoPositiveInts(n)

	var y int64 = 0.0
	_each(factors, func(v interface{}) {
		pi := v.(PositiveInt)

		y += calcPhi(pi.y) * pow(k, pi.x)

		// pow operation for very big numbers does not give us the correct output
		fmt.Printf("%d  %d  %d %d %d \n", calcPhi(pi.y), calcPhi(pi.y), pow(k, pi.x), k, pi.x)
	})

	return y / n
}

func main() {
	// returns a pointer to the list
	//var primes *list.List = primesThatDivideNum(12)
	//fmt.Println(calcPhi((12)))
	//_each(factorNumIntoPositiveInts(10), func(v interface{}) { fmt.Println(v) })
	/*
		fmt.Printf("calc_neclaces(2, 12)         %d  => 352 \n", calcNeclaces(2, 12))
		fmt.Printf("calc_neclaces(3, 7)          %d  => 315 \n", calcNeclaces(3, 7))
		fmt.Printf("calc_neclaces(9, 4)          %d  => 1665 \n", calcNeclaces(9, 4))
		fmt.Printf("calc_neclaces(21, 3)         %d  => 3101 \n", calcNeclaces(21, 3))
		fmt.Printf("calc_neclaces(99, 2)         %d => 4950 \n", calcNeclaces(99, 2))
		fmt.Printf("calc_neclaces(12345678910, 3) %d => 627225458787209496560873442940 \n", calcNeclaces(12345678910, 3))
	*/

	base := big.NewInt(int64(1))
	base.Mul(base, big.NewInt(int64(12345678910)))
	base.Mul(base, big.NewInt(int64(12345678910)))
	base.Mul(base, big.NewInt(int64(12345678910)))
	fmt.Println(base)
	// 1881676376361628489657928971000 (correct value)
	fmt.Println(base.Int64())
	// 5468804790586766072 (bad value)
	fmt.Println(base.Uint64())
	// 5468804790586766072 (bad value)

	// for this to work we need to change int64 to bit.Int in big number operations

}
