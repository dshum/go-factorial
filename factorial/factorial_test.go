package factorial

import "testing"

func TestGetFactorial(t *testing.T) {
	zeroResult := GetFactorial(0)

	if zeroResult != 1 {
		t.Errorf("GetFactorial(\"0\") failed; expected %d, got %d", 1, zeroResult)
	}

	result := GetFactorial(5)

	if result != 120 {
		t.Errorf("GetFactorial(\"5\") failed; expected %d, got %d", 120, result)
	}
}
