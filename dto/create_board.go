package dto

type CreateBoardRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
