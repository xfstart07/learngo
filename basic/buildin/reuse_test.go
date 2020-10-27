// Author: xufei
// Date: 2019/3/21

package buildin

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReuseSlice(t *testing.T) {
	Convey("slice 复用", t, func() {
		Convey("对一个 slice 执行 append 操作", func() {
			si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			si2 := si1
			si2 = append(si2, 0) // 因为 si1 声明的时候长度和容量已经固定，这时要在 si2 后面添加元素，就会重新分配内存，这样 si1 和 si2 就不一样了

			Convey("重新分配内存", func() {
				header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
				header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
				fmt.Println(header1.Data)
				fmt.Println(header2.Data)
				So(header1.Data, ShouldNotEqual, header2.Data)
			})
		})

		Convey("一个 slice 和它的切片", func() {
			si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			si2 := si1[:7]

			Convey("不重新分配内存", func() {
				header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
				header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
				fmt.Println(header1.Data)
				fmt.Println(header2.Data)
				So(header1.Data, ShouldEqual, header2.Data)
			})
		})
	})
}
