package tasks

import "testing"

func TestCheckPossibility(t *testing.T) {
	var res bool
	var wanted bool
	var v []int

	v = []int{}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{4, 2, 3}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{4, 2, 1}
	res = checkPossibility(v)
	wanted = false
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{1, 2, 5}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{2, 2, 2, 2, 2}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{-1, 4, 2, 3}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}

	v = []int{3, 4, 2, 3}
	res = checkPossibility(v)
	wanted = false
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}
	v = []int{1, 2, 4, 5, 3}
	res = checkPossibility(v)
	wanted = true
	if res != wanted {
		t.Errorf("<%v>: wanted %v, got %v\n", v, wanted, res)
	}
}
