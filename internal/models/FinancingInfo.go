package models

type FinancingInfo struct {
	Total    float64 `bson:"total"`
	Year2025 float64 `bson:"year_2025"`
	Year2026 float64 `bson:"year_2026"`
	Year2027 float64 `bson:"year_2027"`
	Later    float64 `bson:"later_years"`
}
