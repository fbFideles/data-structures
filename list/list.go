package list

import "fmt"

// NewInt64List -
func NewInt64List() (l Int64List) {
	return
}

// Int64List -
type Int64List struct {
	Begin *Int64ListValue
	End   *Int64ListValue
	Size  int
}

// Int64ListValue -
type Int64ListValue struct {
	Next  *Int64ListValue
	Prev  *Int64ListValue
	Value int64
}

// PushBack -
func (l *Int64List) PushBack(v int64) {
	l.Size++
	aux := new(Int64ListValue)
	aux.Value = v

	if l.Begin == nil {
		l.Begin = aux
		l.End = aux
		l.Begin.Value = v
		l.End.Value = v

		return
	}

	aux.Prev = l.End
	l.End.Next = aux
	l.End = aux
}

// PushFront -
func (l *Int64List) PushFront(v int64) {
	l.Size++
	aux := new(Int64ListValue)
	aux.Value = v

	if l.Begin == nil {
		l.Begin = aux
		l.End = aux
		l.Begin.Value = v
		l.End.Value = v

		return
	}

	aux.Next = l.Begin
	l.Begin.Prev = aux
	l.Begin = aux
}

// PrintList -
func (l *Int64List) PrintList() {
	v := l.Begin
	for v != nil {
		fmt.Printf("%v ", v.Value)
		v = v.Next
	}
	fmt.Printf("\n")
}

// RPrintList -
func (l *Int64List) RPrintList() {
	v := l.End
	for v != nil {
		fmt.Printf("%v ", v.Value)
		v = v.Prev
	}
	fmt.Printf("\n")
}
