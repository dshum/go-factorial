package factorial

import (
	"math/big"
	"testing"
)

func TestGetFactorial(t *testing.T) {
	expectedInvalidArgResult := 1
	invalidArgResult := GetFactorial(0)

	if invalidArgResult != expectedInvalidArgResult {
		t.Errorf("GetFactorial(\"0\") failed; expected %d, got %d", expectedInvalidArgResult, invalidArgResult)
	}

	expectedResult := 120
	result := GetFactorial(5)

	if result != expectedResult {
		t.Errorf("GetFactorial(\"5\") failed; expected %d, got %d", expectedResult, result)
	}
}

func TestGetBigFactorial(t *testing.T) {
	expectedInvalidArgResult := big.NewInt(int64(1))
	invalidArgResult := GetBigFactorial(-10)

	if expectedInvalidArgResult.Cmp(invalidArgResult) != 0 {
		t.Errorf("GetFactorial(\"-10\") failed; expected %d, got %d", expectedInvalidArgResult, invalidArgResult)
	}

	var expectedResult, _ = new(big.Int).SetString("30414093201713378043612608166064768844377641568960512000000000000", 10)
	result := GetBigFactorial(50)

	if expectedResult.Cmp(result) != 0 {
		t.Errorf("GetFactorial(\"50\") failed; expected %v, got %v", expectedResult, result)
	}
}
