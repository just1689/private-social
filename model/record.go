package model

type Record struct {
	RecordType                 RecordType      `json:"recordType"`
	KeyID                      string          `json:"keyID"`
	Payload                    string          `json:"payload"`
	EncryptedPayloadPrivateKey string          `json:"encryptedPayloadPrivateKey"`
	Signature                  string          `json:"signature"`
	ProfileIdentity            ProfileIdentity `json:"profileIdentity"`
}

type RecordType string

var (
	ServerIdentity RecordType = "ServerIdentity"
	ThisServer     RecordType = "ThisServer"
	Profile        RecordType = "Profile"
	Image          RecordType = "Image"
	Post           RecordType = "Post"
)
