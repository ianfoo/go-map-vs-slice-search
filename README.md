# Searching maps vs searching slices

For small sets, it is typically faster to scan a slice in Go than to use a map lookup.
This benchmark exercises searching for a particular int value in both a map of empty structs,
and a slice of ints.

Below are the results from a 13-inch, Early 2015 MacBook Pro (2.7 GHz Intel Core i5)
```
goos: darwin
goarch: amd64
pkg: gitlab.corp.tune.com/ian/map-vs-slice
BenchmarkMapSearch/beginning_(5_elements)-4         	300000000	         5.05 ns/op
BenchmarkMapSearch/middle_(5_elements)-4            	200000000	         5.54 ns/op
BenchmarkMapSearch/end_(5_elements)-4               	200000000	         7.67 ns/op
BenchmarkMapSearch/notfound_(5_elements)-4          	100000000	        11.3 ns/op
BenchmarkMapSearch/beginning_(10_elements)-4        	100000000	        11.1 ns/op
BenchmarkMapSearch/middle_(10_elements)-4           	100000000	        13.3 ns/op
BenchmarkMapSearch/end_(10_elements)-4              	100000000	        13.7 ns/op
BenchmarkMapSearch/notfound_(10_elements)-4         	100000000	        18.4 ns/op
BenchmarkMapSearch/beginning_(20_elements)-4        	100000000	        11.2 ns/op
BenchmarkMapSearch/middle_(20_elements)-4           	100000000	        12.6 ns/op
BenchmarkMapSearch/end_(20_elements)-4              	100000000	        13.7 ns/op
BenchmarkMapSearch/notfound_(20_elements)-4         	50000000	        26.2 ns/op
BenchmarkMapSearch/beginning_(50_elements)-4        	100000000	        10.9 ns/op
BenchmarkMapSearch/middle_(50_elements)-4           	100000000	        12.2 ns/op
BenchmarkMapSearch/end_(50_elements)-4              	100000000	        17.2 ns/op
BenchmarkMapSearch/notfound_(50_elements)-4         	100000000	        16.9 ns/op
BenchmarkMapSearch/beginning_(500_elements)-4       	100000000	        11.0 ns/op
BenchmarkMapSearch/middle_(500_elements)-4          	100000000	        11.7 ns/op
BenchmarkMapSearch/end_(500_elements)-4             	100000000	        13.8 ns/op
BenchmarkMapSearch/notfound_(500_elements)-4        	100000000	        16.8 ns/op
BenchmarkMapSearch/beginning_(2500_elements)-4      	100000000	        11.1 ns/op
BenchmarkMapSearch/middle_(2500_elements)-4         	100000000	        11.0 ns/op
BenchmarkMapSearch/end_(2500_elements)-4            	100000000	        17.2 ns/op
BenchmarkMapSearch/notfound_(2500_elements)-4       	100000000	        16.9 ns/op
BenchmarkSliceSearch/beginning_(5_elements)-4       	500000000	         3.41 ns/op
BenchmarkSliceSearch/middle_(5_elements)-4          	300000000	         4.29 ns/op
BenchmarkSliceSearch/end_(5_elements)-4             	200000000	         6.54 ns/op
BenchmarkSliceSearch/notfound_(5_elements)-4        	200000000	         7.93 ns/op
BenchmarkSliceSearch/beginning_(10_elements)-4      	500000000	         3.44 ns/op
BenchmarkSliceSearch/middle_(10_elements)-4         	200000000	         6.38 ns/op
BenchmarkSliceSearch/end_(10_elements)-4            	200000000	         9.92 ns/op
BenchmarkSliceSearch/notfound_(10_elements)-4       	100000000	        11.2 ns/op
BenchmarkSliceSearch/beginning_(20_elements)-4      	500000000	         3.40 ns/op
BenchmarkSliceSearch/middle_(20_elements)-4         	200000000	        10.1 ns/op
BenchmarkSliceSearch/end_(20_elements)-4            	100000000	        17.0 ns/op
BenchmarkSliceSearch/notfound_(20_elements)-4       	100000000	        18.4 ns/op
BenchmarkSliceSearch/beginning_(50_elements)-4      	500000000	         3.39 ns/op
BenchmarkSliceSearch/middle_(50_elements)-4         	100000000	        20.7 ns/op
BenchmarkSliceSearch/end_(50_elements)-4            	30000000	        41.7 ns/op
BenchmarkSliceSearch/notfound_(50_elements)-4       	30000000	        42.0 ns/op
BenchmarkSliceSearch/beginning_(500_elements)-4     	500000000	         3.42 ns/op
BenchmarkSliceSearch/middle_(500_elements)-4        	10000000	       126 ns/op
BenchmarkSliceSearch/end_(500_elements)-4           	 5000000	       242 ns/op
BenchmarkSliceSearch/notfound_(500_elements)-4      	 5000000	       250 ns/op
BenchmarkSliceSearch/beginning_(2500_elements)-4    	500000000	         3.44 ns/op
BenchmarkSliceSearch/middle_(2500_elements)-4       	 3000000	       550 ns/op
BenchmarkSliceSearch/end_(2500_elements)-4          	 1000000	      1093 ns/op
BenchmarkSliceSearch/notfound_(2500_elements)-4     	 1000000	      1096 ns/op
```
