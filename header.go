package go_myfatoorah

import "fmt"

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
}

func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		REQ_SIGN_HEAD_NAME: fmt.Sprintf("bearer %s", auth),
	}
}
