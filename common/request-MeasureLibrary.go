package common

type PageSearchMeasureLibraryDto struct {
	PaginationDto
	Name string `form:"name" json:"name"`
}

type MeasureLibraryCreateDto struct {
	HomeWork string `form:"home_work" json:"home_work"`
	RiskType string `form:"risk_type" json:"risk_type"`
	Name     string `form:"name" json:"name"`
	Risk     string `form:"risk" json:"risk"`
	Measures string `form:"measures" json:"measures"`
}

type MeasureLibraryUpdateDto struct {
	HomeWork string `form:"home_work" json:"home_work"`
	RiskType string `form:"risk_type" json:"risk_type"`
	Name     string `form:"name" json:"name"`
	Risk     string `form:"risk" json:"risk"`
	Measures string `form:"measures" json:"measures"`
}
