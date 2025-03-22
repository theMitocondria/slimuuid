package slimuuid

import (
	"encoding/binary"
    "github.com/cespare/xxhash"
    "sync/atomic"
    "strconv"
    "encoding/base64"
)

var counter uint64 = 0  // atomic counter range of counter 0->2^64(0-18446744073709551615)
/*
    This function utilizes xxhash hash function to generate a 64 bit hash of the input string
    always call the function like => 
    uuid := "your_unique_string"
    hash := SingleHashGenerator(uuid)
    fmt.Println(hash)
    This return 8 characters which is 64 base character encoding , but we in code utilizes some
    different algos , additional strings to avoid collision .
*/
func SingleHashGenerator(tt string) string {
    // Combine input with counter atomically
    counter := atomic.AddUint64(&counter, 1)
    data := append([]byte(tt), []byte(strconv.FormatUint(counter, 10))...)    
    // Generate 64-bit hash
    hash := xxhash.Sum64(data)
    // Convert directly to base64
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, hash)
    return base64.RawURLEncoding.EncodeToString(bytes)[:8]
}

