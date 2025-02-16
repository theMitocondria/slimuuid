package slimuuid
import (
	"encoding/binary"
    "encoding/hex"
    "github.com/spaolacci/murmur3"
    "time"
)

func hashGenerator(uuid string) string {
   
    hash1, hash2 := murmur3.Sum128([]byte(uuid))
    
    var fullHash [16]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint64(fullHash[:8], hash1)
    binary.BigEndian.PutUint64(fullHash[8:], hash2)
    
    truncated := hex.EncodeToString(fullHash[:6])
    return truncated 
}

func hashGeneratorFast(tt string) string {
    hash1 := murmur3.Sum32WithSeed([]byte(tt),uint32(time.Now().UnixNano()))
    var fullHash [4]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint32(fullHash[0:], hash1)
    truncated := hex.EncodeToString(fullHash[0:])
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

func GenerateFast(unique string) string {
    timePart := TimePartFast()
	hashedPart :=  hashGeneratorFast(unique +timePart)
	return timePart + hashedPart 
}

//kis machine ne ise call kia ( ye to hum default add kr skte h hr bnde ka jo ki hogi hr bnde ki seed alg hogi) + () 
