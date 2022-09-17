package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(li *ListItem)
	MoveToFront(li *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List           // Remove me after realization.
	root *ListItem // sentinel list element, only &root, root.prev, and root.next are used
	len  int       // current list length excluding (this) sentinel element

	// Place your code here.
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Next
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Prev
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := new(ListItem)
	li.Value = v
	li.Next = l.root.Next
	li.Prev = l.root
	l.root.Next = li
	if l.len == 0 {
		l.root.Prev = li
	} else {
		li.Next.Prev = li
	}
	l.len++
	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := new(ListItem)
	li.Value = v
	li.Next = l.root
	li.Prev = l.root.Prev
	l.root.Prev = li
	if l.len == 0 {
		l.root.Next = li
	} else {
		li.Prev.Next = li
	}
	l.len++
	return li
}

func (l *list) Remove(li *ListItem) {
	li.Next.Prev = li.Prev
	li.Prev.Next = li.Next
	l.len--
}

func (l *list) MoveToFront(li *ListItem) {
	if li == l.root.Next {
		return
	}
	li.Next.Prev = li.Prev
	li.Prev.Next = li.Next
	li.Next = l.root.Next
	li.Prev = l.root
	l.root.Next.Prev = li
	l.root.Next = li
}

func NewList() List {
	l := new(list)
	l.root = new(ListItem)
	l.root.Next = l.root
	l.root.Prev = l.root
	l.len = 0
	return l
}
