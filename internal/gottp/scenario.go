package gottp

type Scenario struct {
	Name    string
	Method  string
	URL     string
	Headers []Header
	Body    []byte
}

type Header struct {
	Key   string
	Value string
}
