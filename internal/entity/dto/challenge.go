package dto

type ChallengeResponse struct {
	Challenge       string `json:"challenge"`
	DifficultyBytes int    `json:"difficulty_bytes"`
}

type ChallengeSolutionRequest struct {
	Nonce int `json:"nonce"`
}

type ChallengeQuoteResponse struct {
	Quote string `json:"quote"`
	Error string `json:"error"`
}
