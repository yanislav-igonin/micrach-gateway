package dto

type CreateBoardRequest struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
