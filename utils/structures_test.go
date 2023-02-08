package utils

import "testing"

func TestReformatDates(t *testing.T) {
	var cat CatStruct = CatStruct{
		"id",
		[]string{"cute", "lovely"},
		"Alex",
		"Sat Jun 04 2022 03:40:20 GMT+0000 (Coordinated Universal Time)",
		"Tue Oct 11 2022 07:52:32 GMT+0000 (Coordinated Universal Time)"}

	cat.ReformatDates()
	if (cat.CreatedAt != "Sat Jun 04 2022 03:40:20 GMT+0000") || (cat.UpdatedAt != "Tue Oct 11 2022 07:52:32 GMT+0000") {
		t.Errorf("The formats of CreatedAt and UpdatedAt are '%s' and '%s', respectively. Expacted without ' (Coordinated Universal Time)'", cat.CreatedAt, cat.UpdatedAt)
	}

}
