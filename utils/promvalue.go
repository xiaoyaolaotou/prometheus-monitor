package utils

import (
	"github.com/prometheus/common/model"
)

func ParseVectorValue(value model.Vector) (map[string]float64, error) {
	result := make(map[string]float64)

	for _, v := range value {
		result[string(v.Metric["instance"])] = float64(v.Value)
	}

	return result, nil
}

func ParseLabelValue(value model.Vector, key string) (map[string]string, error) {
	result := make(map[string]string)

	labelName := model.LabelName(key) // 根据标签取信息

	for _, v := range value {
		result[string(v.Metric["instance"])] = string(v.Metric[labelName])
	}

	return result, nil
}

func ParseMultiVectorValue(value model.Vector, key string) (map[string]map[string]float64, error) {
	result := make(map[string]map[string]float64)
	labelName := model.LabelName(key)
	for _, v := range value {
		instance := string(v.Metric["instance"])
		if _, ok := result[instance]; !ok {
			result[instance] = make(map[string]float64)

		}
		result[instance][string(v.Metric[labelName])] = float64(v.Value)
	}
	return result, nil
}
