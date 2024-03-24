package types

type APIResponse struct {
	Result      int64  `json:"result"` // 작은 따옴표(' ')가 아니라 백틱(` `)이므로 주의
	Description string `json:"description"`
}
