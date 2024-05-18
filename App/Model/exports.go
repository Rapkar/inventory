package Model

import (
	"encoding/json"

	"inventory/App/Boot"
)

func GetAllExports() []Boot.EscapeExport {
	var Export []Boot.Export
	var EscapeExport []Boot.EscapeExport
	Boot.DB().Model(&Boot.Export{}).Select("*").Scan(&Export)
	// if len(Export) == 0 {
	// 	return []Boot.EscapeExport{}
	// }

	EscapeExport = make([]Boot.EscapeExport, len(Export))

	for i, value := range Export {
		var escapeExport Boot.EscapeExport
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
