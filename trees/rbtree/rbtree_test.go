package rbtree

import (
	"log"
	"testing"
)

func Test_rbtree(test *testing.T) {
	rbtree := NewRBTree()

	int64arr := [...]int64{1, 2, 3, 4, 5, 6, 7, 8}

	for _, num := range int64arr {
		rbtree.Insert(num)
	}

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 50; i++ {
	// 	rbtree.Insert(r.Int63n(100))
	// }
	log.Print("输出红黑树@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	printTreeInLog(rbtree.root, "(root)")

	log.Print("删除节点@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	rbtree.Delete(int64(1))
	rbtree.Delete(int64(2))
	rbtree.Delete(int64(3))
	printTreeInLog(rbtree.root, "(root)")
}
