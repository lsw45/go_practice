package algorithm

import "testing"

type Node struct {
	value int
	next  *Node
	pre   *Node
}

func (head *Node) Insert(i, n int) {
	j := 0
	h := head
	for h != nil && j < i {
		h = h.next
		j++
	}

	if j < i || h == nil {
		return
	}
	tail := &Node{value: n}
	tail.next = nil
	h.next = tail

}

func TestLink(t *testing.T) {
	var head = &Node{
		value: 0,
		next:  nil,
	}

	for i := 0; i < 10; i++ {
		head.Insert(i, i)
	}

	t.Run("reserve link", func(t *testing.T) {
		cur := head
		pre := &Node{}
		for cur != nil {
			cur.next = pre
			pre = cur
			cur = cur.next

			//可合并成一句
			//cur.next, pre, cur = pre, cur, cur.next
		}
	})

	t.Run("find circle", func(t *testing.T) {

	})
}
