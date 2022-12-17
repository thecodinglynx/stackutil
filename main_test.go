package stackutil

import (
	"testing"
)

func TestPushAndPopFromStackInts(t *testing.T) {
	s := Stack[int]{[]int{}}
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(3)

	expected := []int{3, 7, 6, 5}
	res := make([]int, 0)
	for {
		if s.IsEmpty() {
			break
		}
		v, err := s.Pop()
		if err != nil {
			t.Fatalf("Unable to Pop from stack, %v", err)
		}
		res = append(res, v)
	}

	if !ArrsEqual(res, expected) {
		t.Errorf("stack: %v - expected to pop values %v, got %v", s, expected, res)
	}
}

type Test struct {
	times int
	str   string
}

func TestPushAndPopFromStackStructs(t *testing.T) {
	s := Stack[Test]{[]Test{}}
	s.Push(Test{5, "a"})
	s.Push(Test{2, "c"})
	s.Push(Test{56, "b"})
	s.Push(Test{1, "d"})

	expected := []Test{{1, "d"}, {56, "b"}, {2, "c"}, {5, "a"}}
	res := make([]Test, 0)
	for {
		if s.IsEmpty() {
			break
		}
		v, err := s.Pop()
		if err != nil {
			t.Fatalf("Unable to Pop from stack, %v", err)
		}
		res = append(res, v)
	}

	if !ArrsEqual(res, expected) {
		t.Errorf("stack: %v - expected to pop values %v, got %v", s, expected, res)
	}
}

func TestPeek(t *testing.T) {
	s := Stack[int]{[]int{}}

	curVal := 5
	s.Push(curVal)

	curVal = 7
	s.Push(curVal)

	curVal = 99
	s.Push(curVal)

	curVal = -15
	s.Push(curVal)

	if v, err := s.Peek(); err != nil || v != curVal {
		t.Fatalf("Peek() failed. Expected %d, got %d", curVal, v)
	}
}

func ArrsEqual[T comparable](l1, l2 []T) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
