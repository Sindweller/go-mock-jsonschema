## go-mock-json-schema

mock data from json schema

按照给定的符合json schema的json来生成对应的mock数据

使用方法：

```shell
make
```

```shell
./sbin/mock --input ./example/input.json

{"data":{"name":"Tradercarnelian Chiller","time":"2022-11-24 21:45:03","value":"this is value"},"list":[3.2842394160417325,16.772807263838583,51.30307557965067,98.79251379081647,3.815431958590972]}
```

example

指定json

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "data": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "mock": {
            "mock": "@name"
          }
        },
        "time": {
          "type": "string",
          "mock": {
            "mock": "@datetime"
          }
        },
        "value": {
          "type": "string",
          "mock": {
            "mock": "this is value"
          }
        }
      }
    },
    "list": {
      "type": "array",
      "items": {
        "type": "number",
        "mock": {
          "mock": "@integer"
        }
      }
    }
  }
}
```

调用函数

```go
GenerateMockData(input)
```

获得数据

```json
{
  "data": {
    "name": "Shriekcherry Dolphin",
    "time": "2022-11-24 19:58:22",
    "value": "this is value"
  },
  "list": [
    0.2138860079711616,
    11.127363432561284,
    13.136448830050695
  ]
}
```


### 简介

练手小工具

项目结构参考 [go project layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

### ref

mock部分：[Mock](https://github.com/nuysoft/Mock)

生成一些名字： [sillyname-go](https://github.com/Pallinder/sillyname-go)

mock功能设计： [yapi](https://github.com/YMFE/yapi)

### 已实现的功能

tbc