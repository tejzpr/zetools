package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io/ioutil"
)

// HMAC is a helper function that returns a new HMAC hash.
func HMAC(text string, fileName string, key string, hashType string) (string, error) {
	if text == "" && fileName == "" {
		return "", errors.New("You must provide a text or a file to be hashed")
	}

	var data []byte
	var err error
	if text != "" {
		data = []byte(text)
	} else {
		data, err = ioutil.ReadFile(fileName)
		if err != nil {
			return "", err
		}
	}

	var h hash.Hash
	if hashType == "sha512" {
		h = hmac.New(sha512.New, []byte(key))
	} else if hashType == "sha256" {
		h = hmac.New(sha256.New, []byte(key))
	} else {
		return "", errors.New("Invalid hash type")
	}
	h.Write(data)
	sha := hex.EncodeToString(h.Sum(nil))
	return string(sha), nil
}
