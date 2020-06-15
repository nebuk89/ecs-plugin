package types

import "encoding/json"

type TaskStatus struct {
	Name             string
	State            string
	Service          string
	NetworkInterface string
	PublicIP         string
	Ports            []string
}

const (
	StackCreate = iota
	StackDelete
)

type LogConsumer interface {
	Log(service, container, message string)
}

type Secret struct {
	ID          string            `json:"ID"`
	Name        string            `json:"Name"`
	Labels      map[string]string `json:"Labels"`
	Description string            `json:"Description"`
	username    string
	password    string
}

func NewSecret(name, username, password, description string) Secret {
	return Secret{
		Name:        name,
		username:    username,
		password:    password,
		Description: description,
	}
}

func (s Secret) ToJSON() (string, error) {
	b, err := json.MarshalIndent(&s, "", "\t")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s Secret) GetCredString() (string, error) {
	creds := map[string]string{
		"username": s.username,
		"password": s.password,
	}
	b, err := json.Marshal(&creds)
	if err != nil {
		return "", err
	}
	return string(b), nil
}