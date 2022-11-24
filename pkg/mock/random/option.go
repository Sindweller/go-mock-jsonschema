package random

import "math/rand"

type ArrayOptions struct {
	Type      string        `json:"type"`       // 数组item的类型
	Mock      string        `json:"mock"`       // @name 或其他mock选项
	Enum      []interface{} `json:"enum"`       // 从中枚举
	MinLength int           `json:"min_length"` // 数组最小长度
	MaxLength int           `json:"max_length"` // 数组最大长度
	Pattern   string        `json:"pattern"`    // 数组元素所应匹配的正则表达式
}

func randLen(min, max int) int {
	// 如果没有填写max 则默认为10
	if max == 0 {
		max = min + 10
	}
	return rand.Intn(max-min) + min
}

func randIndex(len int) int {
	return rand.Intn(len)
}
