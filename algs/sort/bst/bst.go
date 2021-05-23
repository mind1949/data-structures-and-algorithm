package bst

func Sort(s []int) {
	if len(s) <= 1 {
		return
	}
	var b bst
	for _, v := range s {
		b.insert(v)
	}
	s = s[:0]
	b.inorder(func(v int) {
		s = append(s, v)
	})
}

// bst 二叉树
type bst struct {
	root *node
}

func (b *bst) insert(v int) {
	n := &node{v: v}
	if b.root == nil {
		b.root = n
		return
	}

	p := b.root
	c := p
	for c != nil {
		p = c
		if n.v < c.v {
			c = c.l
		} else {
			c = c.r
		}
	}
	n.p = p
	if n.v < p.v {
		p.l = n
	} else {
		p.r = n
	}
}

// inorder 中序遍历
func (b *bst) inorder(handle func(v int)) {
	b.root.inorder(handle)
}

type node struct {
	v int
	p *node // 父节点
	l *node // 左子树节点
	r *node // 右子树节点
}

func (n *node) inorder(handle func(v int)) {
	if n == nil {
		return
	}
	n.l.inorder(handle)
	handle(n.v)
	n.r.inorder(handle)
}
