package rbtree

import (
	"log"
	"testing"
)

func addSon(value int64, parent *RBNode, isleft bool) *RBNode {
	son := NewRBNode(value)
	son.parent = parent
	if isleft {
		parent.left = son
	} else {
		parent.right = son
	}
	return son
}

func Test_rbnode_rotate(test *testing.T) {
	root := NewRBNode(1)
	l := addSon(2, root, LEFTROTATE)
	r := addSon(3, root, RIGHTROTATE)
	addSon(4, l, LEFTROTATE)
	addSon(5, l, RIGHTROTATE)
	addSon(6, r, LEFTROTATE)
	addSon(7, r, RIGHTROTATE)
	log.Printf("输入数据")
	printTreeInLog(root, "(root)")
	log.Printf("父节点")
	log.Printf("%d", root.left.right.parent.value)
	log.Printf("父节点的兄弟节点")
	log.Printf("%d", root.left.right.getUncle().value)
	log.Printf("祖父节点")
	log.Printf("%d", root.left.right.getGrandParent().value)
	log.Printf("左旋")
	if tmproot, err := root.right.rotate(LEFTROTATE); err == nil {
		if tmproot != nil {
			root = tmproot
		}
		printTreeInLog(root, "(root)")
	} else {
		log.Printf(err.Error())
	}

	log.Print("右旋")
	if tmproot, err := root.left.rotate(RIGHTROTATE); err == nil {
		if tmproot != nil {
			root = tmproot
		}
		printTreeInLog(root, "(r)")
	} else {
		log.Printf(err.Error())
	}

}
