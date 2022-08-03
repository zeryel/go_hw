package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

func NewList() List {
	return &list{0, nil, nil}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{v, nil, nil}

	if 0 == l.Len() {
		l.last = newItem
	} else {
		first := l.first
		first.Prev = newItem
		newItem.Next = first
	}

	l.first = newItem
	l.len++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{v, nil, nil}

	if 0 == l.Len() {
		l.first = newItem
	} else {
		last := l.last
		last.Next = newItem
		newItem.Prev = last
	}

	l.last = newItem
	l.len++

	return newItem
}

func (l *list) Remove(i *ListItem) {
	if nil != i.Next {
		i.Next.Prev = i.Prev
	}
	if nil != i.Prev {
		i.Prev.Next = i.Next
	}

	l.len--
	if 0 == l.len {
		l.first = nil
		l.last = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	} else if i == l.Back() {
		i.Prev.Next = nil
		l.last = i.Prev
	} else {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}

	l.first.Prev = i
	i.Next = l.first
	l.first = i
	l.first.Prev = nil
}
