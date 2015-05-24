// model
package model

const (
	HTTP_ADDR     = ":8080"
	CFG_VOICE     = "name"
	DEFAULT_SHELL = "/bin/sh"
)

type Request struct {
	DeviceName string `json:"deviceName"`
	User       string `json:"user"`
	Text       string `json:"text"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func GetIntroductionWords() []string {
	return []string{"who", "introduct"}
}
