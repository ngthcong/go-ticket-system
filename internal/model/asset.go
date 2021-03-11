package model

type (
	Assets struct {
		ID           int
		Name         string
		SerialNumber string
		TypeId       int
		UserId       int
		Description  string
	}
	AssetType struct {
		ID          int
		Name        string
		Description string
	}
)
