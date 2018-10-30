package main

import (
	"fmt"
)

func (ts TestSuite) paint(second int) {
	ts.TestCalls[second].buildReport().paint()
}

func (tcr TestCallsReport) paint() {
	var memory [100]string
	memoryUsedInSlots := int(tcr.AverageTime/10)
	for i:=0; i<memoryUsedInSlots && i<100; i++ {
		memory[i] = "="
	}
	minimumSlots := int(tcr.MinimumTime/10);
	maximumSlots := int(tcr.MaximumTime/10);
	memory[minimumSlots] = "{"
	if maximumSlots >= 100 {
		memory[98] = ">"
		memory[99] = ">"
	} else {
		memory[maximumSlots] = "}"
	}
	

	fmt.Printf("[%2dms] [", tcr.AverageTime)
	for _,val := range(memory) {
		if val=="" {
			val = "-"
		}
		fmt.Print(val)
	}
	fmt.Printf("] [%3d c] [mem:%7d]", tcr.N, tcr.AverageMemory)
	fmt.Println()
}