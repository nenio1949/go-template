package common

type FileDto struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Url  string `form:"url" json:"url"`
	Type string `form:"type" json:"type"`
}
