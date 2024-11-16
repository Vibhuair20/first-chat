package model

type Chat struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Msg       string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
type ContactList struct {
	Username     string `json:"username"`
	LastActivity int64  `json:"last_activity"`
}

type BackGround struct {
	Branch string  `json:"id"`
	Reg_no string  `json:"regn_no"`
	Hostel string  `json:"block"`
	CGPA   float64 `json:"cgpa"`
}
