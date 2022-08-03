package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push only 1 element", func(t *testing.T) {
		// у одного элемента не должно быть следующих и предыдущих элементов
		l := NewList()

		l.PushFront(10)

		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Prev)
		require.Nil(t, l.Back().Next)
		require.Equal(t, l.Front(), l.Back())
	})

	t.Run("`erase` list", func(t *testing.T) {
		l := NewList()

		l.PushFront(25)
		l.Remove(l.Front())

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front last element", func(t *testing.T) {
		l := NewList()

		l.PushFront(2)
		l.PushFront(1)
		l.PushBack(3)

		l.MoveToFront(l.Back())

		require.Equal(t, 3, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("move to front random element", func(t *testing.T) {
		l := NewList()

		movedItem := l.PushFront(1)
		l.PushFront(2)
		l.PushBack(3)

		l.MoveToFront(movedItem)

		require.Equal(t, 3, l.Back().Value)
		require.Equal(t, 1, l.Front().Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
