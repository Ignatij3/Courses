goos: windows
goarch: amd64
cpu: AMD Ryzen 5 2600X Six-Core Processor           
Benchmark_Linear_1e3-12         	 4431280	       267.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_Linear_1e9-12         	       4	   253478925 ns/op	       0 B/op	       0 allocs/op
Benchmark_Linear_1e3_Mod-12        	 2289460	       522.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_Linear_1e9_Mod-12        	       2	   521953250 ns/op	       0 B/op	       0 allocs/op

Benchmark_MathPow_1e3-12        	40649580	       30.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_MathPow_1e9-12        	36362974	       34.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_MathPow_1e18-12       	39189303	       30.49 ns/op	       0 B/op	       0 allocs/op

Benchmark_Log2_1e3-12           	11607278	       113.2 ns/op	      96 B/op	       2 allocs/op
Benchmark_Log2_1e9-12          	     5681252	       219.9 ns/op	     272 B/op	       2 allocs/op
Benchmark_Log2_1e18-12         	     3060478	       412.6 ns/op	     544 B/op	       2 allocs/op
Benchmark_Log2_1e3_Mod-12      	     9673050	       124.3 ns/op	      96 B/op	       2 allocs/op
Benchmark_Log2_1e9_Mod-12      	     4272849	       289.5 ns/op	     272 B/op	       2 allocs/op
Benchmark_Log2_1e18_Mod-12     	     2279528	       540.8 ns/op	     544 B/op	       2 allocs/op

Benchmark_Log2_Ver2_1e3-12     	   278397339	       4.256 ns/op	       0 B/op	       0 allocs/op
Benchmark_Log2_Ver2_1e9-12      	73365776	       16.81 ns/op	       0 B/op	       0 allocs/op
Benchmark_Log2_Ver2_1e18-12     	39804163	       28.81 ns/op	       0 B/op	       0 allocs/op
Benchmark_Log2_Ver2_1e3_Mod-12 	   196476874	       6.280 ns/op	       0 B/op	       0 allocs/op
Benchmark_Log2_Ver2_1e9_Mod-12     	49929681	       23.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_Log2_Ver2_1e18_Mod-12 	28728682	       41.19 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	command-line-arguments	32.231s
