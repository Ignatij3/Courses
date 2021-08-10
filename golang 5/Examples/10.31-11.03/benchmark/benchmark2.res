D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>benchmark.bat

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_list1_1_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                38181             31130 ns/op
BenchmarkSort512-12                10000            105594 ns/op
BenchmarkSort1024-12                3154            367479 ns/op
BenchmarkSort2048-12                 831           1519999 ns/op
BenchmarkSort5096-12                  85          14000651 ns/op
BenchmarkSort100000-12                 1        19061506300 ns/op
PASS
ok      command-line-arguments  27.640s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_list1_2_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                41700             28483 ns/op
BenchmarkSort512-12                12771             94210 ns/op
BenchmarkSort1024-12                3746            325168 ns/op
BenchmarkSort2048-12                 958           1297086 ns/op
BenchmarkSort5096-12                  79          13137230 ns/op
BenchmarkSort100000-12                 1        19078186600 ns/op
PASS
ok      command-line-arguments  28.572s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_list2_1_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                42068             28075 ns/op
BenchmarkSort512-12                13089             91366 ns/op
BenchmarkSort1024-12                3988            312464 ns/op
BenchmarkSort2048-12                 981           1242693 ns/op
BenchmarkSort5096-12                  85          13628266 ns/op
BenchmarkSort100000-12                 1        19110552500 ns/op
PASS
ok      command-line-arguments  28.701s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_list2_2_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                39308             30245 ns/op
BenchmarkSort512-12                12385             97180 ns/op
BenchmarkSort1024-12                3746            327691 ns/op
BenchmarkSort2048-12                 874           1337570 ns/op
BenchmarkSort5096-12                  85          13518091 ns/op
BenchmarkSort100000-12                 1        20835215700 ns/op
PASS
ok      command-line-arguments  30.473s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_containerlist_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                26193             44632 ns/op
BenchmarkSort512-12                 9222            140007 ns/op
BenchmarkSort1024-12                1933            615134 ns/op
BenchmarkSort2048-12                 367           3259020 ns/op
BenchmarkSort5096-12                  51          23560727 ns/op
BenchmarkSort100000-12                 1        35323317300 ns/op
PASS
ok      command-line-arguments  44.501s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>go test -bench . benchmark_slice_test.go
goos: windows
goarch: amd64
BenchmarkSort256-12                77852             15715 ns/op
BenchmarkSort512-12                25648             47185 ns/op
BenchmarkSort1024-12                8563            143653 ns/op
BenchmarkSort2048-12                2497            513471 ns/op
BenchmarkSort5096-12                 392           3053735 ns/op
BenchmarkSort100000-12                 1        1268126400 ns/op
PASS
ok      command-line-arguments  10.622s

D:\Nextcloud\Documents\Programming\Golang 5\Examples\10.31-11.03\benchmark>