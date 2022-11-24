package random

var REGION = []string{"东北", "华北", "华东", "华中", "华南", "西南", "西北"}

// Region 随机地区
func Region() string {
	return REGION[randIndex(len(REGION))]
}
