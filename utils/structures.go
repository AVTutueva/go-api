package utils

import (
	"strings"
)

type CatStruct struct {
	Id        string   `json:"_id"`
	Tags      []string `json:"tags"`
	Owner     string   `json:"owner"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

func (cat *CatStruct) ReformatDates() {
	cat.CreatedAt = strings.Replace(cat.CreatedAt, " (Coordinated Universal Time)", "", 1)
	cat.UpdatedAt = strings.Replace(cat.UpdatedAt, " (Coordinated Universal Time)", "", 1)
}

func PrepareCatStruct(cats *[]CatStruct) {

	for i := 0; i < len(*cats); i++ {
		(*cats)[i].ReformatDates()
	}
}
