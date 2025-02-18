package slimuuid

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/spaolacci/murmur3"
)

/*
    This function utilizes murmur3 hash function to generate a 128 bit hash of the input string
    this takes two hashes and combines them to generate a 128 bit hash , and it takes input a
     uinque string to generate a hash , always call the function like => 
     uuid := "your_unique_string"
     hash := DoubleHashGenerator(uuid)
     fmt.Println(hash)
    This return 12 characters which is 64 base character encoding
*/
func DoubleHashGenerator(uuid string) string {
    hash1, hash2 := murmur3.Sum128([]byte(uuid))
    var fullHash [16]byte 
    binary.BigEndian.PutUint64(fullHash[:8], hash1)
    binary.BigEndian.PutUint64(fullHash[8:], hash2)
    truncated := hex.EncodeToString(fullHash[:6])
    return truncated 
}

/*
    This function utilizes murmur3 hash function to generate a 128 bit hash of the input string
    this takes two hashes and combines them to generate a 128 bit hash . It is same as 
    DoubleHashGenerator but with a seed , always call the function like => 
    uuid := "your_unique_string"
    seed := uint32(time.Now().UnixNano())
    hash := DoubleHashGeneratorWithSeed(uuid, seed)
    fmt.Println(hash)
    This return 12 characters which is 64 base character encoding
*/
func DoubleHashGeneratorWithSeed(uuid string, seed uint32) string {
    hash1, hash2 := murmur3.Sum128WithSeed([]byte(uuid), seed)
    var fullHash [16]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint64(fullHash[:8], hash1)
    binary.BigEndian.PutUint64(fullHash[8:], hash2)
    truncated := hex.EncodeToString(fullHash[:6])
    return truncated 
}

/*
    This function utilizes murmur3 hash function to generate a 32 bit hash of the input string
    always call the function like => 
    uuid := "your_unique_string"
    hash := SingleHashGenerator(uuid)
    fmt.Println(hash)
    This return 8 characters which is 64 base character encoding , but we in code utilizes some
    different algos , additional strings to avoid collision .
*/
func SingleHashGenerator(tt string) string {
    hash1 := murmur3.Sum64([]byte(tt))  // 64-bit instead of 32-bit
    var fullHash [8]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint64(fullHash[:8], hash1)
    truncated := hex.EncodeToString(fullHash[:6])
    return truncated 
}

/*
    Same as SingleHashGenerator but with a seed , always call the function like => 
    uuid := "your_unique_string"
    seed := uint32(time.Now().UnixNano())
    hash := SingleHashGeneratorWithSeed(uuid, seed)
    fmt.Println(hash)
*/
func SingleHashGeneratorWithSeed(tt string, seed uint32) string {
    hash1 := murmur3.Sum32WithSeed([]byte(tt),seed)
    var fullHash [4]byte // stack-allocated for efficiency
    binary.BigEndian.PutUint32(fullHash[0:], hash1)
    truncated := hex.EncodeToString(fullHash[0:])
    return truncated 
}