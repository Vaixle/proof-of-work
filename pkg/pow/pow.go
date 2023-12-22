package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Vaixle/proof-of-work/internal/entity/dto"
	"github.com/dchest/uniuri"
	"math"
	"math/big"
	"math/rand"
	"strconv"
)

var quotes = [4]string{"One’s first step in wisdom is to question everything—and one’s last is to come to terms with everything.",
	"It’s the simple things in life that are the most extraordinary; only wise men are able to understand them.",
	"Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.",
	"To forgive is wisdom, to forget is genius. And easier. Because it's true. It's a new world every heart beat.",
}

// GenerateChallenge Generate a random challenge
func GenerateChallenge(difficulty int) dto.ChallengeResponse {
	return dto.ChallengeResponse{Challenge: uniuri.NewLen(16), DifficultyBytes: difficulty}
}

func ValidatePoW(challenge string, nonce, difficulty int) bool {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	hash := CalculateHash(challenge + strconv.Itoa(nonce))

	hashInt := new(big.Int)
	hashInt.SetString(hash, 16)

	isValid := hashInt.Cmp(target) == -1
	return isValid
}

// GetRandomQuote get quotes
func GetRandomQuote() dto.ChallengeQuoteResponse {
	return dto.ChallengeQuoteResponse{Quote: quotes[rand.Intn(4)]}
}

// CalculateHash calculates the hash of a block(in our case just record)
func CalculateHash(record string) string {
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// SolveChallenge solve challenge from the server
func SolveChallenge(challenge string, difficulty int) int {
	var nonce int

	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	for nonce < math.MaxInt64 {
		hash := CalculateHash(challenge + strconv.Itoa(nonce))

		hashInt := new(big.Int)
		hashInt.SetString(hash, 16)

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce
}
