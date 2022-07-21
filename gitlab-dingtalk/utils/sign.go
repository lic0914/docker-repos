package utils

import (
	"net/url"
    "encoding/base64"
    "crypto/hmac"
    "fmt"
    "crypto/sha256"
)



func hmac256(msg string,secret string) string{    
    // Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(secret))
    // Write Data to it
    h.Write([]byte(msg))
    // Get result and encode as hexadecimal string
    sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
    return sha
}

func Sign(ts string,secret string) string{
    
    s := ts + "\n" + secret
    sign := url.QueryEscape(hmac256(s,secret))
    fmt.Printf("sign: %s ts: %s\n", sign, s)

    return sign
} 