package common

type PageSearchMeasureLibraryDto struct {
	PaginationDto
	Name string `form:"name" json:"name"`
}

type MeasureLibraryDto struct {
	ID       int      `json:"id"`
	HomeWork string   `json:"home_work"`
	RiskType string   `json:"risk_type"`
	Name     string   `json:"name"`
	Risk     string   `json:"risk"`
	Measures []string `json:"measures"`
}

type MeasureLibraryCreateDto struct {
	HomeWork string   `form:"home_work" json:"home_work"`
	RiskType string   `form:"risk_type" json:"risk_type"`
	Name     string   `form:"name" json:"name"`
	Risk     string   `form:"risk" json:"risk"`
	Measures []string `form:"measures" json:"measures"`
}

type MeasureLibraryUpdateDto struct {
	HomeWork string   `form:"home_work" json:"home_work"`
	RiskType string   `form:"risk_type" json:"risk_type"`
	Name     string   `form:"name" json:"name"`
	Risk     string   `form:"risk" json:"risk"`
	Measures []string `form:"measures" json:"measures"`
}
