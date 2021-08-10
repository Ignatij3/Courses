D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //10.000 (разница в 11.72 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12            184449              7034 ns/op
BenchmarkHeapSort-12               14948             82492 ns/op
PASS
ok      command-line-arguments  3.603s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //100.000 элементов (разница в 11.41 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12             17787             64824 ns/op
BenchmarkHeapSort-12                1736            740289 ns/op
PASS
ok      command-line-arguments  3.406s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //1.000.000 элементов (разница в 11.23 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12              1666            745276 ns/op
BenchmarkHeapSort-12                 152           8372719 ns/op
PASS
ok      command-line-arguments  4.552s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //50.000.000 элементов (разница в 5.17 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                15          72731367 ns/op
BenchmarkHeapSort-12                   3         376071300 ns/op
PASS
ok      command-line-arguments  7.670s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //75.000.000 элементов (разница в 5.64 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                12         102924775 ns/op
BenchmarkHeapSort-12                   2         581016150 ns/op
PASS
ok      command-line-arguments  9.056s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //90.000.000 элементов (разница в 7.11 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                13          87894177 ns/op
BenchmarkHeapSort-12                   2         625410600 ns/op
PASS
ok      command-line-arguments  9.304s


D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //90.000.000 элементов (разница в 1.58 раз (не в мою пользу))
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                 1        1006346300 ns/op
BenchmarkHeapSort-12                   2         634843850 ns/op
PASS
ok      command-line-arguments  7.387s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //91.000.000 элементов (разница в 7.78 раз)
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                14          79928193 ns/op
BenchmarkHeapSort-12                   2         622008450 ns/op
PASS
ok      command-line-arguments  9.000s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //91.000.000 элементов (разница в 1.60 раз (не в мою пользу))
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                 1        1019911400 ns/op
BenchmarkHeapSort-12                   2         635763850 ns/op
PASS
ok      command-line-arguments  7.347s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //100.000.000 элементов (разница в 1.52 раз (не в мою пользу))
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                 1        1271189500 ns/op
BenchmarkHeapSort-12                   2         834093050 ns/op
PASS
ok      command-line-arguments  9.445s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Heapsort>go test -bench . Heapsort_test.go //900.000.000 элементов (разница в 1.87 раз (не в мою пользу))
goos: windows
goarch: amd64
BenchmarkMyHeapSort-12                 1        12314464000 ns/op
BenchmarkHeapSort-12                   1        6551678100 ns/op
PASS
ok      command-line-arguments  108.299s //Мой комп рожал это полторы минуты, а миллиард вообще не потянул