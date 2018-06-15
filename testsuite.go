package main

import (
	"time"
)

func newTestSuite() TestSuite {
    testSuite := TestSuite{}
    testSuite.TestCalls = map[int]TestCalls{}

    return testSuite
}

func (ts TestSuite) addTestCall(tc TestCall) {
    second := int(tc.To / int64(time.Second))
    _, ok := ts.TestCalls[second]
    if ok == false {
        ts.TestCalls[second] = TestCalls{}
    }
    ts.TestCalls[second] = append(ts.TestCalls[second], tc)
}

func (ts TestSuite) getAvailableSeconds() []int {
    keys := make([]int, len(ts.TestCalls))

    i := 0
    for k := range ts.TestCalls {
        keys[i] = k
        i++
    }

    return keys
}

func (ts TestSuite) getSecondTestCalls(second int) TestCalls {
    return ts.TestCalls[second]
}

func (tcs TestCalls) buildReport() TestCallsReport {
    n := len(tcs)
    nWithMemory := 0
    AllTime := int64(0)
    MinTime := int64(-1)
    MaxTime := int64(0)
    AllMemory := 0
    for _, tc := range tcs {
        DurationNanos := tc.To - tc.From
        AllTime = AllTime+DurationNanos
        if (
            MinTime > DurationNanos ||
            MinTime < int64(0)) {
            MinTime = DurationNanos
        }
        if (MaxTime < DurationNanos){
            MaxTime = DurationNanos
        }
        if (tc.Resp.Process.MemoryUsed > 0) {
            nWithMemory++
            AllMemory = AllMemory+tc.Resp.Process.MemoryUsed
        }
    }

    if (n==0) {
        return TestCallsReport{
            N: n,
            AverageTime: 0,
            MinimumTime: 0,
            MaximumTime: 0,
            AverageMemory: 0,
        }
    }

    averageMemory := 0
    if (nWithMemory > 0) {
        averageMemory = AllMemory/nWithMemory
    }
    return TestCallsReport{
        N: n,
        AverageTime: int(AllTime/int64(n)/int64(time.Millisecond)),
        MinimumTime: int(MinTime/int64(time.Millisecond)),
        MaximumTime: int(MaxTime/int64(time.Millisecond)),
        AverageMemory: averageMemory,
    }
}

func (ts TestSuite) buildReport() map[int]TestCallsReport {
    tcr := map[int]TestCallsReport{}

    for seconds, tc := range ts.TestCalls {
        tcr[seconds] = tc.buildReport()
    }
    
    return tcr
}