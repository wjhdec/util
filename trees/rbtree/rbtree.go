// 实现红黑树
package rbtree

import "log"

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

func (rbtree *RBTree) Insert(data int64) {
	if rbtree.root == nil {
		rootnode := NewRBNode(data)
		rootnode.color = BLACK
		rbtree.root = rootnode
	} else {
		rbtree.insertNode(rbtree.root, data)
	}
}

func (rbtree *RBTree) insertNode(pnode *RBNode, data int64) {
	if pnode.value >= data {
		// 插入数据不大于父节点，插入左节点
		if pnode.left != nil {
			rbtree.insertNode(pnode.left, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.left = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	} else {
		// 插入数据大于父节点，插入右节点
		if pnode.right != nil {
			rbtree.insertNode(pnode.right, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.right = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	}
}

func (rbtree *RBTree) rotateLeft(node *RBNode) {
	if tmproot, err := node.rotate(LEFTROTATE); err == nil {
		// 旋转时需验证是否有root的改动
		if tmproot != nil {
			rbtree.root = tmproot
		}
	} else {
		log.Printf(err.Error())
	}
}

func (rbtree *RBTree) rotateRight(node *RBNode) {
	if tmproot, err := node.rotate(RIGHTROTATE); err == nil {
		// 旋转时需验证是否有root的改动
		if tmproot != nil {
			rbtree.root = tmproot
		}
	} else {
		log.Printf(err.Error())
	}
}

func (rbtree *RBTree) insertCheck(node *RBNode) {
	if node.parent == nil {
		// 检查1：若插入节点没有父节点，则该节点为root
		rbtree.root = node
		// 设置根节点为black
		rbtree.root.color = BLACK
		return
	}

	// 父节点是黑色的话直接添加，红色则进行处理
	if node.parent.color == RED {
		if node.getUncle() != nil && node.getUncle().color == RED {
			// 父节点及叔父节点都是红色，则转为黑色
			node.parent.color = BLACK
			node.getUncle().color = BLACK
			// 祖父节点改成红色
			node.getGrandParent().color = RED
			// 递归处理
			rbtree.insertCheck(node.getGrandParent())
		} else {
			// 父节点红色，父节点的兄弟节点不存在或为黑色
			isleft := node == node.parent.left
			isparentleft := node.parent == node.getGrandParent().left
			if !isleft && isparentleft {
				rbtree.rotateLeft(node.parent)
				rbtree.rotateRight(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if isleft && !isparentleft {
				rbtree.rotateRight(node.parent)
				rbtree.rotateLeft(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if isleft && isparentleft {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				rbtree.rotateRight(node.getGrandParent())
			} else if !isleft && !isparentleft {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				rbtree.rotateLeft(node.getGrandParent())
			}
		}
	}
}

func (rbtree *RBTree) Delete(data int64) {
	rbtree.delete_child(rbtree.root, data)
}

func (rbtree *RBTree) delete_child(n *RBNode, data int64) bool {
	if data < n.value {
		if n.left == nil {
			return false
		}
		return rbtree.delete_child(n.left, data)
	}
	if data > n.value {
		if n.right == nil {
			return false
		}
		return rbtree.delete_child(n.right, data)
	}

	if n.right == nil || n.left == nil {
		// 两个都为空或其中一个为空，转为删除一个子树节点的问题
		rbtree.delete_one(n)
		return true
	}

	//两个都不为空，转换成删除只含有一个子节点节点的问题
	mostLeftChild := n.right.getLeftMostChild()
	tmpval := n.value
	n.value = mostLeftChild.value
	mostLeftChild.value = tmpval

	rbtree.delete_one(mostLeftChild)

	return true
}

// 删除只有一个子节点的节点
func (rbtree *RBTree) delete_one(n *RBNode) {
	var child *RBNode
	isadded := false
	if n.left == nil {
		child = n.right
	} else {
		child = n.left
	}

	if n.parent == nil && n.left == nil && n.right == nil {
		n = nil
		rbtree.root = n
		return
	}
	if n.parent == nil {
		n = nil
		child.parent = nil
		rbtree.root = child
		rbtree.root.color = BLACK
		return
	}

	if n.color == RED {
		if n == n.parent.left {
			n.parent.left = child

		} else {
			n.parent.right = child
		}
		if child != nil {
			child.parent = n.parent
		}
		n = nil
		return
	}

	if child != nil && child.color == RED && n.color == BLACK {
		if n == n.parent.left {
			n.parent.left = child

		} else {
			n.parent.right = child
		}
		child.parent = n.parent
		child.color = BLACK
		n = nil
		return
	}

	// 如果没有孩子节点，则添加一个临时孩子节点
	if child == nil {
		child = NewRBNode(0)
		child.parent = n
		isadded = true
	}

	if n.parent.left == n {
		n.parent.left = child
	} else {
		n.parent.right = child
	}

	child.parent = n.parent

	if n.color == BLACK {
		if !isadded && child.color == RED {
			child.color = BLACK
		} else {
			rbtree.deleteCheck(child)
		}
	}

	// 如果孩子节点是后来加的，需删除
	if isadded {
		if child.parent.left == child {
			child.parent.left = nil
		} else {
			child.parent.right = nil
		}
		child = nil
	}
	n = nil
}

func (rbtree *RBTree) deleteCheck(n *RBNode) {
	if n.parent == nil {
		n.color = BLACK
		return
	}
	if n.getSibling().color == RED {
		n.parent.color = RED
		n.getSibling().color = BLACK
		if n == n.parent.left {
			rbtree.rotateLeft(n.parent)
		} else {
			rbtree.rotateRight(n.parent)
		}
	}
	//注意：这里n的兄弟节点发生了变化，不再是原来的兄弟节点
	is_parent_red := n.parent.color
	is_sib_red := n.getSibling().color
	is_sib_left_red := BLACK
	is_sib_right_red := BLACK
	if n.getSibling().left != nil {
		is_sib_left_red = n.getSibling().left.color
	}
	if n.getSibling().right != nil {
		is_sib_right_red = n.getSibling().right.color
	}

	if !is_parent_red && !is_sib_red && !is_sib_left_red && !is_sib_right_red {
		n.getSibling().color = RED
		rbtree.deleteCheck(n.parent)
		return
	}
	if is_parent_red && !is_sib_red && !is_sib_left_red && !is_sib_right_red {
		n.getSibling().color = RED
		n.parent.color = BLACK
		return
	}
	if n.getSibling().color == BLACK {
		if n.parent.left == n && is_sib_left_red && !is_sib_right_red {
			n.getSibling().color = RED
			n.getSibling().left.color = BLACK
			rbtree.rotateRight(n.getSibling())
		} else if n.parent.right == n && !is_sib_left_red && is_sib_right_red {
			n.getSibling().color = RED
			n.getSibling().right.color = BLACK
			rbtree.rotateLeft(n.getSibling())
		}
	}
	n.getSibling().color = n.parent.color
	n.parent.color = BLACK
	if n.parent.left == n {
		n.getSibling().right.color = BLACK
		rbtree.rotateLeft(n.parent)
	} else {
		n.getSibling().left.color = BLACK
		rbtree.rotateRight(n.parent)
	}
}

// log输出树
func printTreeInLog(n *RBNode, front string) {
	if n != nil {
		var colorstr string
		if n.color == RED {
			colorstr = "红"
		} else {
			colorstr = "黑"
		}
		log.Printf(front+"%d,%s\n", n.value, colorstr)
		printTreeInLog(n.left, front+"-(l)|")

		printTreeInLog(n.right, front+"-(r)|")
	}
}
