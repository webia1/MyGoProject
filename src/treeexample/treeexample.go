package treeexample

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Has(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Has(val)
	case val > it.val:
		return it.right.Has(val)
	default:
		return true
	}
}
