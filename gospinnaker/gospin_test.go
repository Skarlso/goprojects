package main

import "testing"

//BenchmarkExecuteProcesses does a benchmark test of the ExecuteProcesses function
func BenchmarkExecuteProcesses(b *testing.B) {
	for i := 0; i < 30; i++ {
		ExecuteProcesses()
	}
}
