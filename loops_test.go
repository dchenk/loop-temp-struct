package loop_temp_struct

import "testing"

var s1 = smallStruct{15, []byte(nil)}

var s2 = smallStruct{23425, []byte{'a', 'b', 'c'}}

var s3 = smallStruct{4, []byte{'a', 'b'}}

var a1 = []smallStruct{s1, s2}

var a2 = []smallStruct{s1, s2, s3}

func BenchmarkA1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Val = A(a1)
	}
}

func BenchmarkB1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Val = B(a1)
	}
}

func BenchmarkA2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Val = A(a2)
	}
}

func BenchmarkB2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Val = B(a2)
	}
}
