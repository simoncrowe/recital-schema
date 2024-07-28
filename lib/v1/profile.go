package profiles

import (
	"bytes"
	"encoding/json"
	"io"
)

type Profile struct {
	Text     *string                `json:"text"`
	Metadata map[string]interface{} `json:"metadata"`
}

func DecodeProfileJSON(r io.Reader) (Profile, error) {
	var profile Profile
	if err := json.NewDecoder(r).Decode(&profile); err != nil {
		return profile, err
	}
	return profile, nil
}

func EncodeProfileJSON(profile Profile) (string, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buff).Encode(&profile); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func ValidateProfile(profile Profile) []string {
	errMsgs := []string{}
	if profile.Text == nil {
		errMsgs = append(errMsgs, "The field \"text\" is required")
	}
	return errMsgs
}
