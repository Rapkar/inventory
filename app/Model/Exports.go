package model

import (
	"encoding/json"

	"inventory/boot"
)

func GetAllExports() []boot.EscapeExport {
	var Export []boot.Export
	var EscapeExport []boot.EscapeExport
	boot.DB().Model(&boot.Export{}).Select("*").Scan(&Export)
	// if len(Export) == 0 {
	// 	return []boot.EscapeExport{}
	// }

	EscapeExport = make([]boot.EscapeExport, len(Export))

	for i, value := range Export {
		var escapeExport boot.EscapeExport
		escapeExport.Name = value.Name
		escapeExport.Address = value.Address
		escapeExport.Number = value.Number
		escapeExport.Phonenumber = value.Phonenumber
		escapeExport.Tax = value.Tax
		escapeExport.InventoryNumber = value.InventoryNumber
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = FloatToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}
func FloatToString(value float64) string {
	val, err := json.Marshal(value)
	vals := ""
	if err == nil {
		vals = string(val)
	}
	return vals
}
