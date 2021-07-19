package stack

import "testing"

func TestEmptyStack(t *testing.T) {
	mystack := New()
	if mystack.Size() > 0 {
		t.FailNow()
	}
}

func TestOnePush(t *testing.T) {
	mystack := New()
	element := 100
	mystack.Push(element)
	if mystack.Size() != 1 {
		t.FailNow()
	}
}

func TestPopEmpty(t *testing.T) {
	mystack := New()
	_, ok := mystack.Pop()
	if ok {
		t.FailNow()
	}
}

func TestOnePushOnePop(t *testing.T){
	mystack := New()
	element := 100
	mystack.Push(element)
	popped, ok := mystack.Pop()
	
	if !ok {
		t.FailNow()
	}

	val, ok := popped.(int)
	if !ok || val != element {
		t.FailNow()
	}	
}

func TestTwoPushTwoPop(t *testing.T) {
	mystack := New()
	element1 := 100
	element2 := 200
	mystack.Push(element1)
	mystack.Push(element2)
	popped1, ok := mystack.Pop()
	if !ok {
		t.FailNow()
	}
	val1, ok := popped1.(int)
	if !ok || val1 != element2 {
		t.FailNow()
	}

	popped2, ok := mystack.Pop()
	if !ok {
		t.FailNow()
	}
	val2, ok := popped2.(int)
	if !ok || val2 != element1 {
		t.FailNow()
	}
}

func TestTopEmpty(t *testing.T) {
	mystack := New()
	top := mystack.Top()
	if top >= 0 {
		t.FailNow()
	}
}

func TestTopAfterTwoPush(t *testing.T) {
	mystack := New()
	mystack.Push(100)
	mystack.Push(200)
	top := mystack.Top()
	if top != 1{
		t.FailNow()
	}
}

func BenchmarkNewStack(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		New()
	}
}

func BenchmarkStackPushPop(b *testing.B) {
	mystack := New()
	for i := 1; i <= b.N; i++ {
		mystack.Push(i)
		mystack.Pop()
	}
}
