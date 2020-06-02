package bitmap

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitmap(t *testing.T) {
	Convey("Bitmap", t, func() {
		bitmap := New()
		So(bitmap.Exists(0), ShouldBeFalse)
		So(bitmap.Exists(1), ShouldBeFalse)
		So(bitmap.Exists(31), ShouldBeFalse)
		So(bitmap.Exists(32), ShouldBeFalse)
		bitmap.Add(0)
		So(bitmap.Exists(0), ShouldBeTrue)
		So(bitmap.Exists(32), ShouldBeFalse)
		bitmap.Add(32)
		So(bitmap.Exists(32), ShouldBeTrue)
		bitmap.Add(33)
		So(bitmap.Exists(33), ShouldBeTrue)
	})
}
