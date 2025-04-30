package go_myfatoorah

import "fmt"

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/jason",
	}
}

func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/jason",
		"Authorization": fmt.Sprintf("Bearer %s", auth),
	}
}
