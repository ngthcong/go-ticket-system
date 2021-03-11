package model

type (
	ResponseError struct {
		ErrorCode    int
		ErrorMessage string
	}

	AccessDetails struct {
		UserId int
		Role   int
	}
)
