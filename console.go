package main

import (
	"fmt"
)

func (ts TestSuite) paint(second int) {
	ts.TestCalls[second].buildReport().paint()
}

func (tcr TestCallsReport) paint() {
	var memory [100]string
	memoryUsedInSlots := int(tcr.AverageTime)
	for i:=0; i<memoryUsedInSlots; i++ {
		memory[i] = "="
	}
	memory[tcr.MinimumTime] = "{"
	if tcr.MaximumTime >= 100 {
		memory[98] = ">"
		memory[99] = ">"
	} else {
		memory[tcr.MaximumTime] = "}"
	}
	

	fmt.Printf("Mem: %4dms [", tcr.AverageTime)
	for _,val := range(memory) {
		if val=="" {
			val = "-"
		}
		fmt.Print(val)
	}
	fmt.Printf("] [%4d u] [mem: %8d]", tcr.N, tcr.AverageMemory)
	fmt.Println()
}