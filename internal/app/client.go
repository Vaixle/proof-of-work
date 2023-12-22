package app

import (
	"encoding/json"
	"fmt"
	"github.com/Vaixle/proof-of-work/internal/entity/dto"
	"github.com/Vaixle/proof-of-work/pkg/pow"
	"net"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Run(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	// Get challenge from the server
	decoder := json.NewDecoder(conn)
	var challenge dto.ChallengeResponse
	err = decoder.Decode(&challenge)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// Solve the Proof of Work challenge
	solution := pow.SolveChallenge(challenge.Challenge, challenge.DifficultyBytes)

	// Send the response to the server
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(dto.ChallengeSolutionRequest{Nonce: solution})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// Read and print the quote from the server
	var quoteResponse dto.ChallengeQuoteResponse
	err = decoder.Decode(&quoteResponse)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	if quoteResponse.Error != "" {
		fmt.Println(quoteResponse.Error)
	} else {
		fmt.Println("Received Quote:", quoteResponse.Quote)
	}

}
