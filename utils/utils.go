package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// (i interface{}) wil get ANYTHING
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

// It restores byte data to golang data
func FromBytes(i interface{}, data []byte) {
	// it replace the byte version of b
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}

// It hashes data
func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}

func ToJson(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)
	return r
}
