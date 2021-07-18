package rr

// ReqMessage Message Request struct
type ReqMessage struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

// ResMessage Message Response struct
type ResMessage struct {
	ID        string `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
