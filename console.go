package main

import (
	"fmt"
)

func (ts TestSuite) paint(second int) {
	ts.TestCalls[second].buildReport().paint()
}

func (tcr TestCallsReport) paint() {
	var memory [50]string
	memoryUsedInSlots := int(tcr.AverageTime)
	for i:=0; i<memoryUsedInSlots; i++ {
		memory[i] = "="
	}
	memory[tcr.MinimumTime] = "{"
	if tcr.MaximumTime >= 50 {
		memory[48] = ">"
		memory[49] = ">"
	} else {
		memory[tcr.MaximumTime] = "}"
	}
	

	fmt.Printf("[%2dms] [", tcr.AverageTime)
	for _,val := range(memory) {
		if val=="" {
			val = "-"
		}
		fmt.Print(val)
	}
	fmt.Printf("] [%3d u] [mem:%7d]", tcr.N, tcr.AverageMemory)
	fmt.Println()
}