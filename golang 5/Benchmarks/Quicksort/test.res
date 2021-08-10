D:\Nextcloud\Documents\Programming\Golang 5>go test -bench . QuickSort_test.go //Почти отсортированный
goos: windows
goarch: amd64
BenchmarkInit-12                             559           2003572 ns/op
BenchmarkSort0-12                             16          68185562 ns/op
BenchmarkSort1-12                              1        5364779600 ns/op
BenchmarkSort2-12                              1        9154137300 ns/op
BenchmarkSort2RandomPivot-12                  26          42460777 ns/op
BenchmarkSort1Sedgewick-12                     1        12230874700 ns/op
BenchmarkSort2Sedgewick-12                     1        8850868200 ns/op
BenchmarkSort2Recursiveless-12                 1        14443841000 ns/op
BenchmarkMySort-12                            32          40004325 ns/op
PASS
ok      command-line-arguments  55.216s

D:\Nextcloud\Documents\Programming\Golang 5\Benchmarks\Quicksort>go test -bench . QuickSort_test.go //Рандомный
goos: windows
goarch: amd64
BenchmarkInit-12                              62          17402705 ns/op
BenchmarkSort0-12                              6         188000317 ns/op
BenchmarkSort1-12                             12          94351283 ns/op
BenchmarkSort2-12                             12          90581042 ns/op
BenchmarkSort2RandomPivot-12                  10         109697360 ns/op
BenchmarkSort1Sedgewick-12                    14          82787871 ns/op
BenchmarkSort2Sedgewick-12                    14          76925600 ns/op
BenchmarkSort2Recursiveless-12                13          90708100 ns/op
BenchmarkMySort-12                             1        185781207700 ns/op
BenchmarkMySortOld-12                          9         111765967 ns/op
PASS
ok      command-line-arguments  196.800s