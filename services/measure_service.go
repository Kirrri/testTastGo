package services

import "OneGO/models"

var measures []models.Measure
var currentMeasureID int

func GetMeasures() []models.Measure {
	return measures
}

func GetMeasureByID(id int) (models.Measure, bool) {
	for _, item := range measures {
		if item.Id == id {
			return item, true
		}
	}
	return models.Measure{}, false
}

func CreateMeasure(measure *models.Measure) {
	currentMeasureID++
	measure.Id = currentMeasureID
	measures = append(measures, *measure)
}

func UpdateMeasure(id int, measure *models.Measure) (models.Measure, bool) {
	for index, item := range measures {
		if item.Id == id {
			measures = append(measures[:index], measures[index+1:]...)
			measure.Id = id
			measures = append(measures, *measure)
			return *measure, true
		}
	}
	return models.Measure{}, false
}

func DeleteMeasure(id int) {
	for index, item := range measures {
		if item.Id == id {
			measures = append(measures[:index], measures[index+1:]...)
			break
		}
	}
}
