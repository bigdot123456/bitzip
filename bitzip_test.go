package bitzip

import (
	"fmt"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/willf/bitset"
)

func TestBitZip(t *testing.T) {
	fmt.Println(Exp2(0), Exp2(1), Exp2(2))
	fmt.Println(Log2(1), Log2(2), Log2(4))
	fmt.Println(Dec2H64Bits(66), H64Bits2Dec(Dec2H64Bits(66)))

	b := NewBitZip()
	b.Set(2)
	fmt.Println(b.Raw(), b.Exists(2), b.Exists(67))
	b.Set(67)
	fmt.Println(b.Raw(), b.Exists(2), b.Exists(67), b.Exists(1))
}

var (
	bitzip = NewBitZip()
)

func init() {
	bitzip.Set(10).Set(11)
}

func BenchmarkBitZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if bitzip.Exists(1) {

		}
		if bitzip.Exists(11) {

		}
		if bitzip.Exists(1000000) {

		}
	}
}

func BenchmarkMap(b *testing.B) {
	var B = make(map[int]int8, 3)
	B[10] = 1
	B[11] = 1
	for i := 0; i < b.N; i++ {
		if _, exists := B[1]; exists {

		}
		if _, exists := B[11]; exists {

		}
		if _, exists := B[1000000]; exists {

		}
	}
}

func BenchmarkBitset(b *testing.B) {
	var B bitset.BitSet
	B.Set(10).Set(11)
	for i := 0; i < b.N; i++ {
		if B.Test(1) {

		}
		if B.Test(11) {

		}
		if B.Test(1000000) {

		}
	}
}

func BenchmarkRoaring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B := roaring.BitmapOf(10, 11)
		if B.ContainsInt(1) {

		}
		if B.ContainsInt(11) {
		}
		if B.ContainsInt(1000000) {

		}

	}
}
