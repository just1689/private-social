package model

type ServerInstance struct {
	InstanceIdentity InstanceIdentity `json:"instanceIdentity"`
	Features         []Feature        `json:"features"`
}
