package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton afer calling GetInstance, not nil")
	}

	expectCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Error()
	}

	counter2 := GetInstance()
	if counter2 != expectCounter {
		t.Fail()
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Fail()
	}



}
