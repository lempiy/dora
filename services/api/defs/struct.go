package defs

type PlayerIdRequest struct {
	PlayerID uint64 `json:"player_id"`
}

type MatchIdRequest struct {
	MatchID uint64 `json:"match_id"`
}

type ParseReplayRequest struct {
	ReplayUrl  string `json:"replay_url"`
	PlayerId   uint64 `json:"player_id"`
	MatchId    uint64 `json:"match_id"`
	ReplaySalt uint64 `json:"replay_salt"`
}

type StatusResponse struct {
	Status string `json:"status"`
	Info string `json:"info,omitempty"`
}
