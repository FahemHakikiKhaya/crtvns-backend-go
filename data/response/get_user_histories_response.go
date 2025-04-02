package response

type GetUserHistoriesResponse struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Text      string  `json:"text"`
	Rate      float64 `json:"rate"`
	Pitch     float64 `json:"pitch"`
	Volume    float64 `json:"volume"`
	VoiceName string  `json:"voiceName"`
	CreatedAt string  `json:"createdAt"`
}