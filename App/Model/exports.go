package Model

import (
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Utility"
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
		escapeExport.TotalPrice = Utility.FloatToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}
func GetAllExportsByPaginate(offset int, limit int) []Boot.EscapeExport {
	var Export []Boot.EscapeExport
	var EscapeExport []Boot.EscapeExport
	Boot.DB().Model(&Boot.Export{}).Select("*").Offset(offset).Limit(limit).Scan(&Export)
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
		fmt.Println(value.InventoryNumber)
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = value.TotalPrice
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}
func GetAllExportsByPhoneAndName(searchTerm string) []Boot.EscapeExport {
	var Export []Boot.Export
	var EscapeExport []Boot.EscapeExport

	Boot.DB().Model(&Boot.Export{}).Where("name LIKE ? OR phonenumber LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&Export)
	// Boot.DB().Model(&Boot.Export{}).Where("name = ?", searchTerm).Limit(3).Find(&Export)
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
		escapeExport.TotalPrice = Utility.FloatToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}

func GetCountOfExports() int64 {
	var count int64
	Boot.DB().Model(&[]Boot.Export{}).Find(&[]Boot.Export{}).Count(&count)
	return count
}
