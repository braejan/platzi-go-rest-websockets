package entities

type Home struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type HomeResponse struct {
	*Home        `json:"home"`
	ResponseCode uint8 `json:"response_code"`
}
