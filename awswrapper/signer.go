package awswrapper

import (
	"crypto/rsa"
	"errors"
	url_ "net/url"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

// Signer is a wrapper over aws's signers
type Signer struct {
	urlSigner *sign.URLSigner
}

var (
	signers = make(map[string]*Signer)
)

// GetURLSigner gets URL signer by given keyID and private key
func GetURLSigner(keyID string, privateKey *rsa.PrivateKey) *Signer {
	if signer, ok := signers[keyID]; ok {
		return signer
	}
	urlSigner := sign.NewURLSigner(keyID, privateKey)
	signer := &Signer{urlSigner}
	signers[keyID] = signer
	return signer
}

// SignURL generates a signed cloudfront URL
// domain: the cloudfront domain
// objectPath: the S3 path to the object
// keyID: the cloudfront key ID
// privKey: the private key
// validityTime: the validity time of this url in seconds
func (o *Signer) SignURL(domain, objectPath string, validityTime time.Duration, fileName ...string) (string, error) {
	if o.urlSigner == nil {
		return "", errors.New("no url signer")
	}
	url := "https://" + domain + objectPath
	if len(fileName) > 0 {
		disposition := "filename=\"" + fileName[0] + "\""
		encodedHeaders := "response-content-disposition=attachment%3B" + url_.QueryEscape(disposition)
		url += "?" + encodedHeaders
	}
	return o.urlSigner.Sign(url, time.Now().UTC().Add(validityTime*time.Second))
}
