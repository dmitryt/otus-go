package list

import (
	"errors"
)

// Item - List item element
type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

// List - DoubleLinked List
type List struct {
	Len   int
	First *Item
	Last  *Item
}

// is useful for testing
func NewList(values []interface{}) *List {
	var items []Item
	list := List{Len: len(values)}
	for i := 0; i < list.Len; i++ {
		item := Item{Value: values[i]}
		if i > 0 {
			// i - 1 should reference to previous el before append
			item.Prev = &items[i-1]
			items[i-1].Next = &item
		}
		items = append(items, item)
	}
	if len(items) == 0 {
		return &list
	}
	list.First = &items[0]
	list.Last = &items[len(items)-1]
	return &list
}

// PushFront - insert element to the beginning of the list. Return the new length.
func (l *List) PushFront(value interface{}) int {
	newItem := Item{Value: value, Next: l.First}
	// In case, when we add the first item
	if l.Last == nil {
		l.Last = &newItem
	}
	if l.First == nil {
		l.First = &newItem
	}
	l.First.Prev = &newItem
	l.First = &newItem
	l.Len++
	return l.Len
}

// PushBack - insert element to the end of the list. Return the new length.
func (l *List) PushBack(value interface{}) int {
	newItem := Item{Value: value, Prev: l.Last}
	// In case, when we add the first item
	if l.First == nil {
		l.First = &newItem
	}
	if l.Last == nil {
		l.Last = &newItem
	}
	l.Last.Next = &newItem
	l.Last = &newItem
	l.Len++
	return l.Len
}

// Remove - Remove element from the list by provided index. Return true if item was removed successfully and false otherwise.
func (l *List) Remove(index int) error {
	// Does it make sense to make index as uint? On the other hand it's inconvenient to use uint everywhere
	if index < 0 || index > l.Len-1 {
		return errors.New("index is put of range")
	}
	item := l.First
	for i := 0; i <= index; i++ {
		if i != index {
			item = item.Next
			continue
		}
		if i == 0 {
			// Set element as first and update the reference to the first element
			l.First = l.First.Next
			// check if element exists in case, when on previous step it was the last one
			if l.First != nil {
				l.First.Prev = nil
			}
		}
		if i == l.Len-1 {
			// Set element as last and update the reference to the last element
			l.Last = l.Last.Prev
			// check if element exists in case, when on previous step it was the first one
			if l.Last != nil {
				l.Last.Next = nil
			}
		}
		if i > 0 && i < l.Len-1 {
			// As far as I understood, item should be removed by GC from the memory, right?
			item.Prev = item.Next
		}
		l.Len--
		return nil
	}
	// actually, impossible case
	return errors.New("element with provided index was not found")
}
