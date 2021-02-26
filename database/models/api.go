package models

type ReqBody struct {
	Voice  string `json:"voice" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Speed  string `json:"speed" binding:"required"`
	Format string `json:"format" binding:"required"`
}
type ResBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Text    string `json:"text"`
	Link    string `json:"link"`
}
type ErrBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
