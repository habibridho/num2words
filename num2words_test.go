package num2words_test

import (
	"fmt"
	"testing"

	"github.com/habibridho/num2words"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convert := num2words.Convert
	Convey("Should convert correctly", t, func() {
		Convey("Small numbers should convert correctly", func() {
			So(Convert(0), ShouldEqual, "nol")
			So(Convert(1), ShouldEqual, "satu")
			So(Convert(5), ShouldEqual, "lima")
			So(Convert(10), ShouldEqual, "sepuluh")
			So(Convert(11), ShouldEqual, "sebelas")
			So(Convert(12), ShouldEqual, "dua belas")
			So(Convert(17), ShouldEqual, "tujuh belas")
		})
		Convey("Tens should convert correctly", func() {
			So(Convert(20), ShouldEqual, "dua puluh")
			So(Convert(30), ShouldEqual, "tiga puluh")
			So(Convert(40), ShouldEqual, "empat puluh")
			So(Convert(50), ShouldEqual, "lima puluh")
			So(Convert(60), ShouldEqual, "enam puluh")
			So(Convert(90), ShouldEqual, "sembilan puluh")
		})
		Convey("Combined numbers should convert correctly", func() {
			So(Convert(21), ShouldEqual, "dua puluh satu")
			So(Convert(34), ShouldEqual, "tiga puluh empat")
			So(Convert(49), ShouldEqual, "empat puluh sembilan")
			So(Convert(53), ShouldEqual, "lima puluh tiga")
			So(Convert(68), ShouldEqual, "enam puluh delapan")
			So(Convert(99), ShouldEqual, "sembilan puluh sembilan")
		})
		Convey("Big numbers should convert correctly", func() {
			So(Convert(100), ShouldEqual, "seratus")
			So(Convert(200), ShouldEqual, "dua ratus")
			So(Convert(500), ShouldEqual, "lima ratus")
			So(Convert(123), ShouldEqual, "seratus dua puluh tiga")
			So(Convert(666), ShouldEqual, "enam ratus enam puluh enam")
			So(Convert(1024), ShouldEqual, "seribu dua puluh empat")
			So(Convert(1000000), ShouldEqual, "satu juta")
			So(Convert(1500000), ShouldEqual, "satu juta lima ratus ribu")
		})
		Convey("Negative numbers should convert correclty", func() {
			So(Convert(-123), ShouldEqual, "minus seratus dua puluh tiga")
		})
	})
}

func ExampleConvert() {
	fmt.Println(num2words.Convert(11))
	fmt.Println(num2words.Convert(123))
	fmt.Println(num2words.Convert(-99))
	// Output: sebelas
	// seratus dua puluh tiga
	// minus sembilan puluh sembilan
}
