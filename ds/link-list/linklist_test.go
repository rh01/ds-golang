package linklist

import (
	"errors"
	"testing"
)

func TestPush(t *testing.T) {
	q := New()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	if q.Len() != 10 {
		t.Error()
	}

	events, _ := q.Select(isEvent)

	if events.Len() != 5 {
		t.Error()
	}

	for e := range events.Iter() {
		n, _ := e.(int)
		if n%2 != 0 {
			t.Error()
		}
	}

	next, _ := events.Next()
	i, _ := next()
	n, _ := i.(int)
	if n != 0 {
		t.Error()
	}

	var val interface{}

	var sum = 0
	for it, has_next := events.Next(); has_next; val, has_next = it() {
		n, _ := val.(int)
		sum += n
	}
	// var list *List
	// for list := q.Iter(); list.Next()

	t.Logf("sum is %d", sum)

	if sum == 20 {
		t.Error()
		// t.Logf(n)
	}



}

func isEvent(a interface{}) (bool, error) {
	n, ok := a.(int)
	if ok != true {
		return false, errors.New("Failed type assertion on a")
	}

	return n%2 == 0, nil
}
