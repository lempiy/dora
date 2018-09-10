package defs

type PlayerIdRequest struct {
	PlayerID int64 `json:"player_id"`
}

type StatusResponse struct {
	Status string `json:"status"`	
}
