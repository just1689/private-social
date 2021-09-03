package model

type ProfileIdentity struct {
	IdentityVersion  string           `json:"identityVersion"`
	PublicKey        string           `json:"publicKey"`
	Proof            string           `json:"proof"` //public key signed with private key
	InstanceIdentity InstanceIdentity `json:"instanceIdentity"`
}

type InstanceIdentity struct {
	IdentityVersion string `json:"identityVersion"`
	ApiURL          string `json:"apiURL"`
	PublicKey       string `json:"publicKey"`
	Proof           string `json:"proof"` //public key signed with private key
}
