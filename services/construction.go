package service

import (
	"go-template/common"
	"go-template/models"
)

func AddConstruction(params common.ConstructionCreateDto) (int, error) {
	return models.AddConstruction(params)
}
