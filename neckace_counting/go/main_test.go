package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	if isPrime(2) != true {
		t.Error(`2 is not a prime?`)
	}

	if isPrime(4) != false {
		t.Error("4 is not a prime?")
	}
}

func TestPrimesThatDivideNum(t *testing.T) {
	primes := primesThatDivideNum(10)
	filteredPrimes := _filter(primes, func(v interface{}) bool {
		if v.(int) == 2 || v.(int) == 5 {
			return true
		}
		return false
	})

	if filteredPrimes.Len() != 2 {
		t.Error("primesThatDivideNum of 10 should return two records 2 and 3.")
	}
}

func TestCalcPhi(t *testing.T) {
	if calcPhi(12) != 4 {
		t.Error("calcPhi of 12 should return 4")
	}
}

func TestFactorNumIntoPositiveInts(t *testing.T) {
	factors := factorNumIntoPositiveInts(10)
	filtededFactors := _filter(factors, func(v interface{}) bool {
		if v.(PositiveInt).x == 1 || v.(PositiveInt).x == 2 || v.(PositiveInt).x == 5 || v.(PositiveInt).x == 10 {
			return true
		}
		return false
	})

	if filtededFactors.Len() != 4 {
		t.Error("factorNumIntoPositiveInts of 10 should return for records 1, 2, 5 and 10.")
	}
}
