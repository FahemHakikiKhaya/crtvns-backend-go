package model

type UserHistory struct {
	Id        int
	UserId    int
	Text      string
	Rate      float64
	Pitch     float64
	Volume    float64
	VoiceName string
	CreatedAt string
}