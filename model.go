package main

type Status struct {
	Elasticsearch string `json:"elasticsearch" yaml:"redis"`
	Redis bool `json:"redis" yaml:"redis"`
}

type Process struct {
	MemoryUsed int `json:"memory_used" yaml:"memory_used"`
}

type CheckHealthResponse struct {
    Status Status `json:"status" yaml:"status"`
    Process Process `json:"process" yaml:"process"`
}

type TestSuite struct {
    TestCalls map[int]TestCalls
}

type TestCalls []TestCall

type TestCallsReport struct {
    N int `yaml:"n"`
    AverageTime int `yaml:"average_time"`
    MinimumTime int `yaml:"min_time"`
    MaximumTime int `yaml:"max_time"`
    AverageMemory int `yaml:"average_memory"`
}

type TestCall struct {
    User User
    Resp CheckHealthResponse `yaml:"resp"`
    From int64 `yaml:"from"`
    To int64 `yaml:"to"`
}

type User struct {
    Id int
    ActiveFrom int
    ActiveTo int 
}

type TestDuration struct {
    Duration int
}

type Scenario struct {
    Calls []TestCall `yaml:"calls"`
	CallReports map[int]TestCallsReport `yaml:"report"`
}