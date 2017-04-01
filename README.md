# bitzip
Compress bitmap

`go test -bench=.* -benchmem`

BenchmarkBitZip-2        5000000               241 ns/op              72 B/op          5 allocs/op
BenchmarkMap-2          30000000                44.3 ns/op             0 B/op          0 allocs/op
BenchmarkBitset-2       1000000000               2.29 ns/op            0 B/op          0 allocs/op
BenchmarkRoaring-2       3000000               607 ns/op             152 B/op          6 allocs/op