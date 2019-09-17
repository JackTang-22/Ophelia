package test

import (
	"Ophelia/go/src/internal/service"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/17 15:51
 * @Description:
 */

func TestStringSliceEqual(t *testing.T) {
	convey.Convey("TestStringSliceEqual should return true when a != nil && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		convey.So(service.StringSliceEqual(a, b), convey.ShouldBeTrue)
	})

	convey.Convey("TestStringSliceEqual should return true when a != nil && b != nil", t, func() {
		convey.So(service.StringSliceEqual(nil, nil), convey.ShouldBeTrue)
	})

	convey.Convey("TestStringSliceEqual should return false when a != nil && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		convey.So(service.StringSliceEqual(a, b), convey.ShouldBeFalse)
	})

	convey.Convey("TestStringSliceEqual should return false when a != nil && b != nil", t, func() {
		a := []string{"hello", "world"}
		b := []string{"hello", "goconvey"}
		convey.So(service.StringSliceEqual(a, b), convey.ShouldBeFalse)
	})

	convey.Convey("TestStringSliceEqual", t, func() {
		convey.Convey("should return true when a != nil && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			convey.So(service.StringSliceEqual(a, b), convey.ShouldBeTrue)
		})

		convey.Convey("should return true when a != nil && b != nil", func() {
			convey.So(service.StringSliceEqual(nil, nil), convey.ShouldBeTrue)
		})

		convey.Convey("should return false when a != nil && b != nil", func() {
			a := []string(nil)
			b := []string{}
			convey.So(service.StringSliceEqual(a, b), convey.ShouldBeFalse)
		})

		convey.Convey("should return false when a != nil && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			convey.So(service.StringSliceEqual(a, b), convey.ShouldBeFalse)
		})
	})
}
