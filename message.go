package nchats

type message struct {
	Time     string `json:"time,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Message  string `json:"message"`
}
