package apperrors

type ErrCode string

const (
	Unkown ErrCode = "U000"
	InsertDetailFailed ErrCode = "S001"
	GetDataFailed      ErrCode = "S002"
  NAData             ErrCode = "S003"
)
