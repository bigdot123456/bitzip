package bitzip

const (
	WordWidth    uint = 64
	LogWordWidth uint = 6
)

type BitZip struct {
	raws [][]uint64 // 六十四进制数长度：六十四进制数字
}

func NewBitZip() *BitZip {
	b := new(BitZip)
	return b
}

func NewBitZip2(l int) *BitZip {
	b := new(BitZip)
	b.extand(l)
	return b
}

// 保存的最大位数
func (b *BitZip) Num() int {
	return len(b.raws) + 1
}

func IndexNum(i int) int {
	return i - 1
}

func (b *BitZip) extand(l int) *BitZip {
	newraws := make([][]uint64, l)
	for i := 0; i < len(b.raws); i++ {
		newraws[i] = b.raws[i]
	}
	newraws[IndexNum(l)] = make([]uint64, l)
	b.raws = newraws
	return b
}

func (b *BitZip) Set(i uint) *BitZip {
	bits := Dec2H64Bits(i)

	if len(bits) >= b.Num() {
		b.extand(len(bits))
	}

	for i, bit := range bits {
		b.raws[IndexNum(len(bits))][i] |= bit
	}

	return b
}

func (b *BitZip) Exists(x uint) bool {
	bits := Dec2H64Bits(x)
	if len(bits) >= b.Num() {
		return false
	}

	for i, bit := range bits {
		if b.raws[IndexNum(len(bits))][i]&bit == 0 {
			return false
		}
	}

	return true
}

func (b *BitZip) Raw() [][]uint64 {
	return b.raws
}

/*
        e....d
66/64 = 1....2
1 /64 = 0....1
*/
// 十进制转六十四进制并按位保存
func Dec2H64Bits(i uint) (h64s []uint64) {
	h64s = make([]uint64, 0, 1)
	for i != 0 {
		e := i >> LogWordWidth
		d := i - e<<LogWordWidth
		i = e

		h64s = append(h64s, 1<<(d&(WordWidth-1)))
	}
	return h64s
}

func H64Bits2Dec(h64s []uint64) (dec uint) {
	if len(h64s) == 0 {
		return
	}

	dec = Log2(h64s[0])
	for i := 1; i < len(h64s); i++ {
		dec += Log2(h64s[i]) << (LogWordWidth * uint(i))
	}
	return dec
}

func Exp2(x uint) uint64 {
	return 1 << (x & (WordWidth - 1))
}

func Log2(x uint64) (y uint) {
	for x != 1 {
		x = x >> 1
		y++
	}
	return y
}
