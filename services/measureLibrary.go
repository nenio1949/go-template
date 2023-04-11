package service

import (
	"go-template/common"
	"go-template/models"
)

func AddMeasureLibrary(params common.MeasureLibraryCreateDto) (int, error) {
	return models.AddMeasureLibrary(params)
}
