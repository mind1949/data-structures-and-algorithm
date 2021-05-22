package list

func New() List {
	var n Node
	n.pre = &n
	n.next = &n

	l := &list{
		root: n,
	}
	n.list = l
	return l
}

// List 双向链表
type List interface {
	Front() *Node
	Back() *Node
	Remove(n *Node) interface{}
	PushFront(v interface{}) *Node
	PushBack(v interface{}) *Node
	InsertAfter(v interface{}, at *Node) *Node
	InsertBefore(v interface{}, at *Node) *Node
	MoveFront(n *Node)
	MoveBack(n *Node)
	MoveBefore(n *Node, at *Node)
	MoveAfter(n *Node, at *Node)
	PushFrontList(ol *list)
	PushBackList(ol *list)
}

var _ List = (*list)(nil)
type list struct {
	root Node // 哨兵节点
	len int
}

func (l *list) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *list) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.pre
}

func (l *list) Remove(n *Node) interface{} {
	if n.list == l {
		n.pre.next = n.next
		n.next.pre = n.pre
		n.next = nil
		n.pre = nil
		n.list = nil
		l.len--
	}
	return n.Value
}

func (l *list) PushFront(v interface{}) *Node {
	n := &Node{
		Value: v, 
		pre: &l.root,
		next: l.root.next,
		list: l,
	}

	l.root.next.pre = n
	l.root.next = n
	l.len++
	return n
}

func (l *list) PushBack(v interface{}) *Node {
	n := &Node{
		Value: v,
		pre: l.root.pre,
		next: &l.root,
		list: l,
	}

	l.root.pre.next = n
	l.root.pre = n
	l.len++
	return n
}

func (l *list) InsertAfter(v interface{}, at *Node) *Node {
	if at.list != l {
		return nil
	}

	n := &Node{
		Value: v,
		pre: at,
		next: at.next,
	}
	at.next.pre = n
	at.next = n
	l.len++
	return n
}

func (l *list) InsertBefore(v interface{}, at *Node) *Node {
	if at.list != l {
		return nil
	}

	n := &Node {
		Value: v,
		pre: at.pre,
		next: at,
	}
	n.pre.next = n
	n.next.pre = n
	l.len++
	return n
}

func (l *list) MoveFront(n *Node) {
	if n.list != l || l.root.next == n {
		return
	}
	n.pre.next = n.next
	n.next.pre = n.pre

	n.pre = &l.root
	n.next = l.root.next
	n.pre.next = n
	n.next.pre = n
}

func (l *list) MoveBack(n *Node) {
	if n.list != l || l.root.pre == n {
		return
	}
	n.pre.next = n.next
	n.next.pre = n.pre

	n.pre = l.root.pre.pre
	n.next = &l.root
	n.pre.next = n
	n.next.pre = n
}

func (l *list) MoveBefore(n *Node, at *Node) {
	if n.list != l || at.list != l || n == at {
		return
	}

	n.pre.next = n.next
	n.next.pre = n.pre
	n.pre = at.pre
	n.next = at
	n.pre.next = n
	n.next.pre = n
}

func (l *list) MoveAfter(n *Node, at *Node) {
	if n.list != l || at.list != l || n == at {
		return
	}

	n.pre.next = n.next
	n.next.pre = n.pre

	n.pre = at
	n.next = at.next
	n.pre.next = n
	n.next.pre = n
}

func (l *list) PushFrontList(ol *list) {
	if ol.len == 0 {
		return
	}

	lFrontNode := l.Front()
	for n := ol.Front(); n != &ol.root; n = n.next {
		l.InsertBefore(n.Value, lFrontNode)
	}
}

func (l *list) PushBackList(ol *list) {
	if ol.len == 0 {
		return
	}
	for n := ol.Front(); n != &ol.root; n = n.next {
		l.InsertBefore(n.Value, &l.root)
	}
}

type Node struct {
	Value interface{}
	pre *Node
	next *Node
	list *list
}

func (n *Node) Pre() *Node {
	return n.pre
}

func (n *Node) Next() *Node {
	return n.next
}
