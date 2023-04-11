package common

type FileCreateDto struct {
	Name           string `form:"name" json:"name"`
	Url            string `form:"url" json:"url"`
	ConstructionID int    `form:"construct_id" json:"construct_id"`
}
