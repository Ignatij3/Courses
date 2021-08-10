
F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_list1_1_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   25878	     46520 ns/op
BenchmarkSort512-4    	    7503	    160276 ns/op
BenchmarkSort1024-4   	    2144	    561822 ns/op
BenchmarkSort2048-4   	     532	   2250187 ns/op
PASS
ok  	command-line-arguments	5.899s

F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_list1_2_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   25312	     47195 ns/op
BenchmarkSort512-4    	    7503	    164618 ns/op
BenchmarkSort1024-4   	    2143	    595194 ns/op
BenchmarkSort2048-4   	     530	   2274455 ns/op
PASS
ok  	command-line-arguments	6.032s

F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_list2_1_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   24454	     48281 ns/op
BenchmarkSort512-4    	    7503	    168254 ns/op
BenchmarkSort1024-4   	    2070	    577762 ns/op
BenchmarkSort2048-4   	     528	   2356868 ns/op
PASS
ok  	command-line-arguments	6.172s

F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_list2_2_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   24655	     50042 ns/op
BenchmarkSort512-4    	    7504	    165193 ns/op
BenchmarkSort1024-4   	    2106	    620535 ns/op
BenchmarkSort2048-4   	     481	   2321832 ns/op
PASS
ok  	command-line-arguments	6.022s

F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_containerlist_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   15064	     97751 ns/op
BenchmarkSort512-4    	    5101	    215204 ns/op
BenchmarkSort1024-4   	    1225	    962729 ns/op
BenchmarkSort2048-4   	     232	   5114139 ns/op
PASS
ok  	command-line-arguments	7.914s

F:\Work. GO\go_V_work\06\samples\benchmark>go test -bench . benchmark_slice_test.go  
goos: windows
goarch: amd64
BenchmarkSort256-4    	   36589	     32307 ns/op
BenchmarkSort512-4    	   10000	    142712 ns/op
BenchmarkSort1024-4   	    3625	    333253 ns/op
BenchmarkSort2048-4   	    1010	   1194813 ns/op
PASS
ok  	command-line-arguments	5.820s
