package defs

type PlayerIdRequest struct {
	PlayerID uint64 `json:"player_id"`
}

type StatusResponse struct {
	Status string `json:"status"`
	Info string `json:"info,omitempty"`
}
