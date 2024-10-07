# merkle-vs-verkle
benchmark against verkle and merkle  
proofs are encoded to json  

# parameters
chunksize: 32 (verkle tree chunk size)

# Usage
test all benchmarks:  
```
go test run=nil -bench . -benchmem
```  
bench serialization:  
```
go test run=nil -bench "Serialize.erkleProof" -benchmem
```  
bench serialization result:  
```
goos: linux
goarch: amd64
pkg: github.com/dahn510/merkle-vs-verkle
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkCreateAndSerializeMerkleProof10000000Byte-16    	 425661	     3138 ns/op	   3349 B/op	     14 allocs/op
--- BENCH: BenchmarkCreateAndSerializeMerkleProof10000000Byte-16
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
    main.go:186: proof size: 1751
	... [output truncated]
BenchmarkCreateAndSerializeVerkleProof10000000Byte-16    	     52	 26236626 ns/op	12803984 B/op	 189018 allocs/op
--- BENCH: BenchmarkCreateAndSerializeVerkleProof10000000Byte-16
    main.go:234: verkle proof marshalled size: 1568
    main.go:235: stateDiff marshalled size: 206
    main.go:234: verkle proof marshalled size: 1568
    main.go:235: stateDiff marshalled size: 206
    main.go:234: verkle proof marshalled size: 1568
    main.go:235: stateDiff marshalled size: 206
    main.go:234: verkle proof marshalled size: 1568
    main.go:235: stateDiff marshalled size: 206
    main.go:234: verkle proof marshalled size: 1568
    main.go:235: stateDiff marshalled size: 206
	... [output truncated]
PASS
ok  	github.com/dahn510/merkle-vs-verkle	20.210s
```
