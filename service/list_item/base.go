package list_item

import (
	"github.com/siskinc/srv-name-list/models"
	"strings"
)

func (service *Service) makeMultiValue(fields, values []string) []models.MultiValueItem {
	if len(fields) != len(values) {
		panic("fields len not equal values len")
	}
	var result []models.MultiValueItem
	for i := 0; i < len(fields); i++ {
		result = append(result, models.MultiValueItem{
			Key:   fields[i],
			Value: values[i],
		})
	}
	return result
}

func (service *Service) makeValue(multiValue []models.MultiValueItem) string {
	valueItems := make([]string, len(multiValue))
	for i := 0; i < len(multiValue); i++ {
		valueItems[i] = multiValue[i].Value
	}
	return strings.Join(valueItems, ",")
}
