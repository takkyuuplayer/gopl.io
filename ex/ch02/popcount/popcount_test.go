package popcount

import "testing"

func TestPopCount(t *testing.T) {
	if PopCount(0) != 0 { // 0
		t.Errorf(`PopCount(0) = %#v, want %#v`, PopCount(0), 0)
	}

	if PopCount(1) != 1 { // 1
		t.Errorf(`PopCount(10) = %#v, want %#v`, PopCount(1), 2)
	}

	if PopCount(10) != 2 { // 1010
		t.Errorf(`PopCount(10) = %#v, want %#v`, PopCount(10), 2)
	}

	if PopCount(100) != 3 { // 11010
		t.Errorf(`PopCount(10) = %#v, want %#v`, PopCount(100), 3)
	}

}

func TestPopCountByLoop(t *testing.T) {
	if PopCountByLoop(0) != 0 { // 0
		t.Errorf(`PopCountByLoop(0) = %#v, want %#v`, PopCountByLoop(0), 0)
	}

	if PopCountByLoop(1) != 1 { // 1
		t.Errorf(`PopCountByLoop(10) = %#v, want %#v`, PopCountByLoop(1), 2)
	}

	if PopCountByLoop(10) != 2 { // 1010
		t.Errorf(`PopCountByLoop(10) = %#v, want %#v`, PopCountByLoop(10), 2)
	}

	if PopCountByLoop(100) != 3 { // 11010
		t.Errorf(`PopCountByLoop(10) = %#v, want %#v`, PopCountByLoop(100), 3)
	}
}

func TestPopCountByShift(t *testing.T) {
	if PopCountByShift(0) != 0 { // 0
		t.Errorf(`PopCountByShift(0) = %#v, want %#v`, PopCountByShift(0), 0)
	}

	if PopCountByShift(1) != 1 { // 1
		t.Errorf(`PopCountByShift(10) = %#v, want %#v`, PopCountByShift(1), 2)
	}

	if PopCountByShift(10) != 2 { // 1010
		t.Errorf(`PopCountByShift(10) = %#v, want %#v`, PopCountByShift(10), 2)
	}

	if PopCountByShift(100) != 3 { // 11010
		t.Errorf(`PopCountByShift(10) = %#v, want %#v`, PopCountByShift(100), 3)
	}

}

func TestPopCountByAndOperator(t *testing.T) {
	if PopCountByAndOperator(0) != 0 { // 0
		t.Errorf(`PopCountByAndOperator(0) = %#v, want %#v`, PopCountByAndOperator(0), 0)
	}

	if PopCountByAndOperator(1) != 1 { // 1
		t.Errorf(`PopCountByAndOperator(10) = %#v, want %#v`, PopCountByAndOperator(1), 2)
	}

	if PopCountByAndOperator(10) != 2 { // 1010
		t.Errorf(`PopCountByAndOperator(10) = %#v, want %#v`, PopCountByAndOperator(10), 2)
	}

	if PopCountByAndOperator(100) != 3 { // 11010
		t.Errorf(`PopCountByAndOperator(10) = %#v, want %#v`, PopCountByAndOperator(100), 3)
	}

}
