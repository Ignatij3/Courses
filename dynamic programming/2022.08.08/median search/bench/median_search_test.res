goos: windows
goarch: amd64
cpu: AMD Ryzen 5 2600X Six-Core Processor           
BenchmarkQuickSearch/10^3-12   	  182750	      7371 ns/op	    8192 B/op	       1 allocs/op
BenchmarkQuickSearch/10^6-12   	     100	  10782672 ns/op	 8003584 B/op	       1 allocs/op
BenchmarkQuickSearch/10^9-12   	       1	96248004900 ns/op	8000004096 B/op	       1 allocs/op
BenchmarkMedianOfMedians/10^3-12         	   19161	     82997 ns/op	   29376 B/op	     403 allocs/op
BenchmarkMedianOfMedians/10^6-12         	      20	  57691125 ns/op	28812800 B/op	  400003 allocs/op
BenchmarkMedianOfMedians/10^9-12         	       1	339105585100 ns/op	28800012480 B/op	400000005 allocs/op
PASS
ok  	command-line-arguments	464.732s
