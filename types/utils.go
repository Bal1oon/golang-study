package types

type ApiResponse struct {
	Result      int64       `json:"result"` // 작은 따옴표(' ')가 아니라 백틱(` `)이므로 주의
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"` // 에러 코드를 받기 위함
}

func NewApiResponse(description string, result int64, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrCode:     errCode,
	}
}
