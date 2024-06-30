package namings

import (
	"fmt"
	"testing"
)

func TestAlphabetical(t *testing.T) {
	type tc struct {
		i int
		o string
	}
	tcs := []tc{{0, "A"},
		{25, "Z"},
		{26, "AA"},
		{51, "AZ"},
		{52, "BA"},
		{702, "AAA"}}

	for _, tc := range tcs {
		g := Bijective26(tc.i)
		if g != tc.o {
			t.Fatal("for:", tc.i, "expected:", tc.o, "got:", g)
		}
	}
}

func ExampleBijective26() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, Bijective26(i))
	}
	// Output:
	// 0 A
	// 1 B
	// 2 C
	// 3 D
	// 4 E
	// 5 F
	// 6 G
	// 7 H
	// 8 I
	// 9 J
	// 10 K
	// 11 L
	// 12 M
	// 13 N
	// 14 O
	// 15 P
	// 16 Q
	// 17 R
	// 18 S
	// 19 T
	// 20 U
	// 21 V
	// 22 W
	// 23 X
	// 24 Y
	// 25 Z
	// 26 AA
	// 27 AB
	// 28 AC
	// 29 AD
	// 30 AE
	// 31 AF
	// 32 AG
	// 33 AH
	// 34 AI
	// 35 AJ
	// 36 AK
	// 37 AL
	// 38 AM
	// 39 AN
	// 40 AO
	// 41 AP
	// 42 AQ
	// 43 AR
	// 44 AS
	// 45 AT
	// 46 AU
	// 47 AV
	// 48 AW
	// 49 AX
	// 50 AY
	// 51 AZ
	// 52 BA
	// 53 BB
	// 54 BC
	// 55 BD
	// 56 BE
	// 57 BF
	// 58 BG
	// 59 BH
	// 60 BI
	// 61 BJ
	// 62 BK
	// 63 BL
	// 64 BM
	// 65 BN
	// 66 BO
	// 67 BP
	// 68 BQ
	// 69 BR
	// 70 BS
	// 71 BT
	// 72 BU
	// 73 BV
	// 74 BW
	// 75 BX
	// 76 BY
	// 77 BZ
	// 78 CA
	// 79 CB
	// 80 CC
	// 81 CD
	// 82 CE
	// 83 CF
	// 84 CG
	// 85 CH
	// 86 CI
	// 87 CJ
	// 88 CK
	// 89 CL
	// 90 CM
	// 91 CN
	// 92 CO
	// 93 CP
	// 94 CQ
	// 95 CR
	// 96 CS
	// 97 CT
	// 98 CU
	// 99 CV
}
