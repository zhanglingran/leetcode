package problems

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMax(t *testing.T) {
	Convey("测试Max函数", t, func() {
		So(Max(1, 2), ShouldEqual, 2)
	})
}

func TestLengthOfLongestSubstring(t *testing.T) {
	Convey("测试返回一个字符串最长子串的长度", t, func() {
		Convey("字符串不为空的时候", func() {
			So(LengthOfLongestSubstring("asdddsa"), ShouldEqual, 3)
		})

		Convey("字符串为空的时候", func() {
			So(LengthOfLongestSubstring(""), ShouldEqual, 0)
		})
	})
}
