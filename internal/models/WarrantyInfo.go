package models

type WarrantyInfo struct {
	Required             bool   `bson:"required"`
	Description          string `bson:"description"`
	ManufacturerWarranty string `bson:"manufacturer_warranty"`
	Period               string `bson:"period"`
}
