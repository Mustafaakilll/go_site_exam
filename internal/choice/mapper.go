package choice

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateChoiceRequest struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`

	QuestionID int `json:"question_id"`
}

type UpdateChoiceRequest struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type ChoiceResponseDTO struct {
	Data []ChoiceDTO `json:"data"`
	BaseResponse
}

type ChoiceDTO struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
	QuestionID int    `json:"question_id"`
}

type QuestionDTO struct {
	ID      int         `json:"ID"`
	Text    string      `json:"Text"`
	Type    int         `json:"Type"`
	Point   int         `json:"Point"`
	Choices []ChoiceDTO `json:"choices"`
}

type Data struct {
	Question QuestionDTO `json:"question"`
}
