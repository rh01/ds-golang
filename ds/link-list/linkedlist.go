package linklist

import "errors"

type List struct {
	head *node
	tail *node
}

type node struct {
	data interface{} //generic type.
	next *node
	prev *node
}

func (list *List) Push(data interface{}) {
	node := &node{data: data, next: nil, prev: nil}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}
}

func (list *List) Pop() (interface{}, error) {
	if list.head == nil {
		return nil, emptyList()
	}
	if list.head == list.tail {
		data := list.head.data
		list.head = nil
		list.tail = nil
		return data, nil
	} else {
		current := list.tail
		data := current.data
		list.tail = current.prev
		list.tail.next = nil
		current = nil
		return data, nil
	}
}

func emptyList() error {
	return errors.New("Empty list")
}

func New() *List {
	return &List{nil, nil}
}

func FromArgs(args ...interface{}) *List {
	list := New()
	for _, arg := range args {
		list.Push(arg)
	}
	return list
}

// build list from array
func FromArray(args []interface{}) *List {
	list := New()
	for i := 0; i < len(args); i++ {
		list.Push(args[i])
	}
	return list
}

// get a copy versionn of list
func (list *List) Clone() *List {
	newList := New()

	for e := range list.Iter() {
		newList.Push(e)
	}

	return newList
}

// iterate through the list and get the element.
func (list *List) Iter() chan interface{} {
	ch := make(chan interface{})

	go func() {
		current := list.head

		for current != nil {
			ch <- current.data
			current = current.next
		}
		close(ch)

	}()
	return ch
}

// get length of list
func (list *List) Len() uint {
	var count uint = 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

//
func (list *List) Select(grep func(a interface{}) (bool, error)) (*List, error) {
	newList := New()

	currrent := list.head
	for currrent != nil {
		b, err := grep(currrent.data)
		if err != nil {
			return nil, err
		}
		if b == true {
			newList.Push(currrent.data)
		}

		currrent = currrent.next
	}

	return newList, nil
}

// An alias of Iter
func (list *List) Each() <-chan interface{} {
	return list.Iter()
}

//
// func (list *List) Next() (func() (interface{}, bool), bool) {
// 	current := list.head
// 	return func() (interface{}, bool) {
// 		// prev_curr := current
// 		// temp := prev_curr.data
// 		list.head = current.next
// 		return current.data, true
// 	}, true
// }

// func (l *List) Next() *List {
	
// }