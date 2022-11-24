package mock

import "encoding/json"

type ObjectOrArray struct {
	//Schema     string `json:"$schema"`
	Type       string               `json:"type"`
	Properties map[string]BasicMock `json:"properties"`
	Items      BasicMock            `json:"items"`
	BasicMock
}

type MockJson struct {
	Type       string                   `json:"type"`
	Properties map[string]ObjectOrArray `json:"properties"`
}

func (m *MockJson) ToString() string {
	res, _ := json.Marshal(m)
	return string(res)
}

func (m *ObjectOrArray) ToString() string {
	res, _ := json.Marshal(m)
	return string(res)
}

func (m *ObjectOrArray) ToTypeAndMock() BasicMock {
	return BasicMock{Mock: m.Mock, Type: m.Type}
}

type BasicMock struct {
	Type             string `json:"type"`
	Mock             Mock   `json:"mock"`
	ExclusiveMinimum bool   `json:"exclusiveMinimum"`
	ExclusiveMaximum bool   `json:"exclusiveMaximum"`
	Default          string `json:"default"`
	Minimum          int    `json:"minimum"`
	Maximum          int    `json:"maximum"`
	Enum             []int  `json:"enum"`
	EnumDesc         string `json:"enumDesc"`
}

type Mock struct {
	Mock string `json:"mock"`
}
