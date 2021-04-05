package main

import (
	"github.com/java-akademie/myutils"
)

type (
	apple int
	pear  int
	fruit int
)

var (
	i, j               int
	i8, j8             int8
	i16, j16           int16
	i32, j32           int32
	i64, j64           int64
	u                  uint
	u8                 uint8
	u16                uint16
	u32                uint32
	u64                uint64
	e32, f32, g32, h32 float32
	e64, f64, g64, h64 float64
	b1, b2             byte // uint8
	r1, r2             rune // int32
	ok1, ok2           bool
	text1, text2       string
	apples             apple // int
	pears              pear  // int
	fruits             fruit // int
)

func initVariables() {
	i = 9_000_000_000_000_000_000   // 9 Trillionen
	i8 = 120                        //
	i16 = 32_000                    //
	i32 = 2_000_000_000             // 2 Milliarden
	i64 = 9_000_000_000_000_000_000 // 9 Trillionen

	j = -9_000_000_000_000_000_000   // -9 Trillionen
	j8 = -120                        //
	j16 = -32_000                    //
	j32 = -2_000_000_000             // -2 Milliarden
	j64 = -9_000_000_000_000_000_000 // -9 Trillionen

	u = 18_000_000_000_000_000_000   // 18 Trillionen
	u8 = 250                         //
	u16 = 32_000                     //
	u32 = 4_000_000_000              // 4 Milliarden
	u64 = 18_000_000_000_000_000_000 // 18 Trillionen

	e32 = 2_000_000_000  // + 2 Milliarden
	e64 = 4.7e+300       // very large positive number
	f32 = -2_000_000_000 // - 2 Milliarden
	f64 = -4.7e+300      // very large negative number
	g32 = 0.000_000_001  // 1/Milliardstel
	g64 = 4.7e-300       // very small positive number
	h32 = -0.000_000_001 // -1/Milliardstel
	h64 = -4.7e-300      // very small negative number

	b1 = 'A'    // uint8 or byte
	b2 = 'X'    // uint8 or byte
	r1 = 64_000 // int32 or rune negative would be possible but for runes not reasonable
	r2 = 22_000 // int32 or rune negative would be possible but for runes not reasonable

	ok1 = true
	ok2 = false
	text1 = "dies ist text 11"
	text2 = "dies ist text 22"

	apples = 5
	pears = 7
}

func showAllVariables() {
	myutils.H2("showAllVariables")
	myutils.ShowObjects(i, i8, i16, i32, i64)
	myutils.ShowObjects(j, j8, j16, j32, j64)
	myutils.ShowObjects(u, u8, u16, u32, u64)
	myutils.ShowObjects(e32, e64)
	myutils.ShowObjects(f32, f64)
	myutils.ShowObjects(g32, g64)
	myutils.ShowObjects(h32, h64)
	myutils.ShowObjects(ok1, ok2)
	myutils.ShowObjects(text1, text2)

	myutils.Comment("show byte's")
	myutils.ShowObjects(
		b1, string(b1),
		b2, string(b2),
	)
	myutils.Comment("show rune's")
	myutils.ShowObjects(
		r1, string(r1),
		r2, string(r2),
	)
	myutils.Wait()
}
