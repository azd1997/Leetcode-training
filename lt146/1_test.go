package lt146

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {

	lru := Constructor(2)
	lru.printLRU()

	lru.Put(1,1)
	lru.printLRU()
	lru.Put(2,2)
	lru.printLRU()

	fmt.Println("key=1, val=", lru.Get(1))  // 1
	fmt.Println("key=3, val=", lru.Get(3))  // -1

	lru.Put(3, 3)  // 挤掉（2,2），插到头部之后
	lru.printLRU()

	fmt.Println("key=2, val=", lru.Get(2))  // -1

	lru.Put(4, 4)  // (1,1)被淘汰
	lru.printLRU()
	fmt.Println("key=1, val=", lru.Get(1))	// -1

	fmt.Println("key=3, val=", lru.Get(3))  // 3
	fmt.Println("key=4, val=", lru.Get(4))  // 4

	lru.Put(3, 5)
	lru.printLRU()
}
