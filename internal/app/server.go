package app

import (
	"encoding/json"
	"fmt"
	"github.com/Vaixle/proof-of-work/internal/entity/dto"
	"github.com/Vaixle/proof-of-work/pkg/pow"
	"net"
)

type Server struct {
	difficulty int
}

func NewServer(difficulty int) *Server {
	return &Server{difficulty: difficulty}
}

func (s *Server) Run(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Generate a challenge
	challenge := pow.GenerateChallenge(s.difficulty)

	// Send challenge to the client
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(challenge)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Get solution from the client
	decoder := json.NewDecoder(conn)
	var receivedMessage dto.ChallengeSolutionRequest
	err = decoder.Decode(&receivedMessage)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Validate solution from the client
	if pow.ValidatePoW(challenge.Challenge, receivedMessage.Nonce, s.difficulty) {
		// PoW is valid, send a random quote
		randomQuote := pow.GetRandomQuote()
		err = encoder.Encode(randomQuote)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	} else {
		err = encoder.Encode(dto.ChallengeQuoteResponse{Error: "PoW is invalid"})
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}
}
