package problems

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var obj LRUCache

func TestConstructor(t *testing.T) {
	Convey("测试通过capacity初始化LRU缓存", t, func() {
		obj = Constructor(12)
		So(obj, ShouldNotBeNil)
	})
}

func TestLRUCache_Get(t *testing.T) {
	Convey("测试获取LRU缓存中的数据", t, func() {

		obj = Constructor(12)
		obj.Put(1, 1)

		Convey("获取缓存中存在的数据", func() {
			ret := obj.Get(1)
			So(ret, ShouldEqual, 1)
		})

		Convey("测试缓存中不存在的数据", func() {
			ret := obj.Get(2)
			So(ret, ShouldEqual, -1)
		})
	})
}

func TestLRUCache_Put(t *testing.T) {
	Convey("测试向缓存中写入数据", t, func() {

		// 初始化缓存大小
		obj = Constructor(2)

		Convey("缓存未满，向其中写入数据", func() {
			obj.Put(1, 1)
			So(obj.size, ShouldEqual, 1)

			obj.Put(2, 2)
			So(obj.size, ShouldEqual, 2)
		})

		Convey("缓存已满，向其中写入数据", func() {
			obj.Put(1, 1)
			So(obj.size, ShouldEqual, 1)

			obj.Put(2, 2)
			So(obj.size, ShouldEqual, 2)

			obj.Put(3, 3)
			So(obj.size, ShouldEqual, 2)

			ret := obj.Get(1)
			// 缓存已满，插入数据3后，数据1要出缓存，因此get(1)=-1
			So(ret, ShouldEqual, -1)
		})

	})
}

// Main  .go 中的main函数
func Main(m *testing.M) {
	obj := Constructor(2)

	fmt.Println("put 1, 0: ")
	obj.Put(1, 0)
	obj.PrintList()

	fmt.Println("put 2, 2: ")
	obj.Put(2, 2)
	obj.PrintList()

	fmt.Println("get 1: ", obj.Get(1))
	obj.PrintList()

	fmt.Println("put 3, 3: ")
	obj.Put(3, 3)
	obj.PrintList()

	fmt.Println("get 2: ", obj.Get(2))
	obj.PrintList()

	fmt.Println("put 4, 4: ")
	obj.Put(4, 4)
	obj.PrintList()

	fmt.Println("get 1: ", obj.Get(1))
	obj.PrintList()

	fmt.Println("get 3: ", obj.Get(3))
	obj.PrintList()

	fmt.Println("get 4: ", obj.Get(4))
	obj.PrintList()
}
