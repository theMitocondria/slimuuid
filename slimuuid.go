package slimuuid

import (
	"encoding/binary"
    "encoding/hex"
    "strings"
    "github.com/spaolacci/murmur3"
)

func hashGenerator(uuid string) string {
    trimmed := strings.ReplaceAll(uuid, "-", "")
    hash1, hash2 := murmur3.Sum128([]byte(trimmed))
    
    var fullHash [16]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint64(fullHash[:8], hash1)
    binary.BigEndian.PutUint64(fullHash[8:], hash2)
    
    truncated := hex.EncodeToString(fullHash[:6])
    return truncated 
}

func Generate() string {
	uuid , err := ID()
    if err != nil {
        return "" 
    }

    timePart := TimePart()
	hashedPart :=  hashGenerator(uuid)
	slimId := timePart + hashedPart 

	return slimId
}
