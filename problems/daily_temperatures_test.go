package problems

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	Convey("TestDailyTemperatures", t, func() {
		res := DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
		So(res, ShouldResemble, []int{1, 1, 4, 2, 1, 1, 0, 0})
	})
}
