package mock

import (
	"encoding/json"
	random2 "github.com/Sindweller/go-mock-jsonschema/pkg/mock/random"
	lorem "github.com/drhodes/golorem"
	"math/rand"
)

type MockSource struct {
	name string `json:"name"`
	mock string `json:"mock"`
}

var MOCK_SOURCE = []MockSource{
	{name: "字符串", mock: "@string"},
	{name: "自然数", mock: "@natural"},
	{name: "浮点数", mock: "@float"},
	{name: "字符", mock: "@character"},
	{name: "布尔", mock: "@boolean"},
	{name: "url", mock: "@url"},
	{name: "域名", mock: "@domain"},
	{name: "ip地址", mock: "@ip"},
	{name: "id", mock: "@id"},
	{name: "guid", mock: "@guid"},
	{name: "当前时间", mock: "@now"},
	{name: "时间戳", mock: "@timestamp"},
	{name: "日期", mock: "@date"},
	{name: "时间", mock: "@time"},
	{name: "日期时间", mock: "@datetime"},
	{name: "图片连接", mock: "@image"},
	{name: "图片data", mock: "@imageData"},
	{name: "颜色", mock: "@color"},
	{name: "颜色hex", mock: "@hex"},
	{name: "颜色rgba", mock: "@rgba"},
	{name: "颜色rgb", mock: "@rgb"},
	{name: "颜色hsl", mock: "@hsl"},
	{name: "整数", mock: "@integer"},
	{name: "email", mock: "@email"},
	{name: "大段文本", mock: "@paragraph"},
	{name: "句子", mock: "@sentence"},
	{name: "单词", mock: "@word"},
	{name: "大段中文文本", mock: "@cparagraph"},
	{name: "中文标题", mock: "@ctitle"},
	{name: "标题", mock: "@title"},
	{name: "姓名", mock: "@name"},
	{name: "中文姓名", mock: "@cname"},
	{name: "中文姓", mock: "@cfirst"},
	{name: "中文名", mock: "@clast"},
	{name: "英文姓", mock: "@first"},
	{name: "英文名", mock: "@last"},
	{name: "中文句子", mock: "@csentence"},
	{name: "中文词组", mock: "@cword"},
	{name: "地址", mock: "@region"},
	{name: "省份", mock: "@province"},
	{name: "城市", mock: "@city"},
	{name: "地区", mock: "@county"},
	{name: "转换为大写", mock: "@upper"},
	{name: "转换为小写", mock: "@lower"},
	{name: "挑选（枚举）", mock: "@pick"},
	{name: "打乱数组", mock: "@shuffle"},
	{name: "协议", mock: "@protocol"},
}

// 最外层 根据json schema生成mock
func GenerateMockData(resbody string) map[string]interface{} {
	// 结果字典
	res := make(map[string]interface{})
	// 从jsonschema中构建
	mockjson := MockJson{}
	// 将resbody复制给mockjson
	json.Unmarshal([]byte(resbody), &mockjson)
	// 遍历properties并且生成随机mock结果
	for k, v := range mockjson.Properties {
		if v.Type == "object" {
			res[k] = GenerateMockData(v.ToString())
		} else if v.Type == "array" {
			res[k] = generateArrayByValue(BasicMock{Type: v.Items.Type, Mock: v.Items.Mock})
		} else {
			res[k] = generateByValue(v.ToBasicMock())
		}
	}
	return res
}

func GetMock(input string) string {
	res := GenerateMockData(input)
	output, _ := json.Marshal(res)
	return string(output)
}

// mock 数组数据
func generateArrayByValue(mock BasicMock) []interface{} {
	// TODO 先随机出数组长度
	arrlen := rand.Intn(5) + 1
	// 根据items的type来选择
	arr := make([]interface{}, arrlen)
	for i := range arr {
		arr[i] = generateByValue(mock)
	}
	return arr
}

// 根据最最基础的value来mock
func generateByValue(value BasicMock) interface{} {
	// value 也是字典
	//{
	//      "type": "array",
	//      "items": {
	//        "type": "string",
	//		  "mock": {
	//           "mock": "@name"
	//		  }
	//      }
	//    }
	// 先转换为字典
	switch value.Type {
	case "string":
		return generateString(value)
	case "integer":
		return rand.Intn(1024)
	case "number":
		return float64(rand.Intn(100)) * rand.Float64()
	case "boolean":
		rd := rand.Intn(1)
		if rd == 0 {
			return false
		} else {
			return true
		}
	default:
		return nil
	}
}

// 处理string的mock
func generateString(value BasicMock) string {
	if value.Mock.Mock != "" {
		switch value.Mock.Mock {
		case "@name":
			return random2.GenerateStupidName()
		case "@cname":
			return random2.Cname()
		case "@datetime":
			return random2.Datetime()
		case "@last":
			return random2.Last()
		case "@first":
			return random2.First()
		case "@cparagraph":
			// TODO 上下限
			return random2.RandZN(50, 100)
		case "@timestamp":
			return string(random2.Timestamp())
		default:
			// 如果没有指定 则返回该值
			if value.Mock.Mock[0] == '@' {
				sentence := lorem.Sentence(1, rand.Intn(10))
				return sentence[0 : len(sentence)-1]
			}
			return value.Mock.Mock
		}
	} else {
		sentence := lorem.Sentence(1, rand.Intn(10))
		return sentence[0 : len(sentence)-1]
	}
}
