package rr

// ReqRoom room create response
type ReqRoom struct {
	Name string `json:"name"`
}

// ResRoom room response
type ResRoom struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
