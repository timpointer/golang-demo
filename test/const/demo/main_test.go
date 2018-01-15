package main

import "testing"

func Benchmark_getScanCodePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScanCodePath("urn:epc:id:sgcn:69478909.0092.108634966514")
	}
}
