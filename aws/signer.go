package aws

import (
	"errors"
	url_ "net/url"
	"time"

	"tecgit01.tectusdreamlab.com/TDS/common-utils-backend/convertor"
	"tecgit01.tectusdreamlab.com/TDS/common-utils-backend/encryption"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

// Signer is a wrapper over aws's signers
type Signer struct {
	urlSigner *sign.URLSigner
}

var (
	signers = make(map[string]*Signer)
)

// GetURLSigner gets URL signer by given keyID and privKey
func GetURLSigner(keyID, privKeyString string) *Signer {
	if signer, ok := signers[keyID]; ok {
		return signer
	}
	privKeyBytes, err := convertor.Base64ToBytes(privKeyString)
	if err != nil {
		return &Signer{nil}
	}
	privKey, err := encryption.UnMarshalPrivateKey(privKeyBytes)
	if err != nil {
		return &Signer{nil}
	}
	urlSigner := sign.NewURLSigner(keyID, privKey)
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
