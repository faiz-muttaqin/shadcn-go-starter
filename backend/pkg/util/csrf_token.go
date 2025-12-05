package util

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func GenerateEncryptedCSRFToken(ip, ua string) string {
	payload := map[string]interface{}{
		"remote_ip":  ip,
		"user_agent": ua,
		"timestamp":  time.Now().Unix(),
	}
	encrypted, _ := GetAESEcryptedURLfromJSON(payload) // implement AES encryption
	return encrypted
}
func CheckCSRFToken(csrfToken, currentIP, currentUA string) error {
	// 1. Decrypt the token
	data, err := GetAESDecryptedURLtoJSON(csrfToken)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("invalid CSRF token format")
	}

	// 2. Extract data
	tokenIP, okIP := data["remote_ip"].(string)
	tokenUA, okUA := data["user_agent"].(string)
	timestampFloat, okTS := data["timestamp"].(float64) // JSON numbers are float64

	if !okIP || !okUA || !okTS {
		return fmt.Errorf("missing fields in CSRF token")
	}

	// 3. Check IP and User-Agent
	if tokenIP != currentIP || tokenUA != currentUA {
		return fmt.Errorf("IP or User-Agent mismatch")
	}

	// 4. Optional: check expiration (e.g., token valid for 10 minutes)
	tokenTime := time.Unix(int64(timestampFloat), 0)
	if time.Since(tokenTime) > 10*time.Minute {
		return fmt.Errorf("CSRF token expired")
	}

	return nil
}
