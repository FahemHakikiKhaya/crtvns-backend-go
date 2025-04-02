package request

type CreateUserHistoryRequest struct {
	UserId    int     `validate:"required" json:"userId"`
	Text      string  `validate:"required" json:"text"`
	Rate      float64 `validate:"required" json:"rate"`
	Pitch     float64 `validate:"required" json:"pitch"`
	Volume    float64 `validate:"required" json:"volume"`
	VoiceName string  `validate:"required" json:"voiceName"`
}
