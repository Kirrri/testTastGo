package services

import "OneGO/models"

var measures []models.Measure
var currentMeasureID int

func GetMeasures() []models.Measure {
	return measures
}

func GetMeasureByID(id int) (models.Measure, bool) {
	for _, item := range measures {
		if item.MeasureID == id {
			return item, true
		}
	}
	return models.Measure{}, false
}

func CreateMeasure(measure *models.Measure) {
	currentMeasureID++
	measure.MeasureID = currentMeasureID
	measures = append(measures, *measure)
}

func UpdateMeasure(id int, measure *models.Measure) (models.Measure, bool) {
	for index, item := range measures {
		if item.MeasureID == id {
			measures[index] = *measure
			measure.MeasureID = id
			return *measure, true
		}
	}
	return models.Measure{}, false
}

func DeleteMeasure(id int) bool {
	for index, item := range measures {
		if item.MeasureID == id {
			measures = append(measures[:index], measures[index+1:]...)
			return true
		}
	}
	return false
}
