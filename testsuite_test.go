package main

import (
	"fmt"
	"io/ioutil"
	"testing"
    "gopkg.in/yaml.v2"
)

func TestTestSuiteReport(t *testing.T) {
    scenarios := []string{"1","2","3"}
    for _, filename := range scenarios {
        yamlFile, _ := ioutil.ReadFile("scenarios/scenario" + filename + ".yml")
        scenario := Scenario{}
        yaml.Unmarshal(yamlFile, &scenario)
        testSuite := newTestSuite()
        for _, tc := range scenario.Calls {
            testSuite.addTestCall(tc)
        }
        testCallReports := testSuite.buildReport()
        if (len(testCallReports) == 0) {
            t.Error("Empty TestCalls Reports found")
        }
        for second, tcr := range testCallReports {
            if (
                scenario.CallReports[second].N != tcr.N ||
                scenario.CallReports[second].AverageTime != tcr.AverageTime ||
                scenario.CallReports[second].MinimumTime != tcr.MinimumTime ||
                scenario.CallReports[second].MaximumTime != tcr.MaximumTime ||
                scenario.CallReports[second].AverageMemory != tcr.AverageMemory) {
                    t.Error("Bad TestCalls Report")
                    fmt.Println(tcr)
                }
        }

    }
}