package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var ODOO_LOGIN_COOKIE []*http.Cookie
var ODOO_CSNA_LOGIN_COOKIE []*http.Cookie

func LoginOdoo() ([]*http.Cookie, error) {
	loginURL := os.Getenv("ODOO_WEB_URL") + "/web/session/authenticate"
	loginPayload := map[string]interface{}{
		"jsonrpc": "2.0",
		"params": map[string]string{
			"db":       os.Getenv("ODOO_WEB_DB"),
			"login":    os.Getenv("ODOO_WEB_USER"),
			"password": os.Getenv("ODOO_WEB_PASS"),
		},
	}

	loginJSON, err := json.Marshal(loginPayload)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error creating login payload: %w", err)
	}

	loginReq, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(loginJSON))
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error creating login request: %w", err)
	}

	loginReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	loginResp, err := client.Do(loginReq)
	if err != nil {
		logrus.Error("Info : "+"ODOO_WEB_DB :"+os.Getenv("ODOO_WEB_DB"), "ODOO_WEB_USER :"+os.Getenv("ODOO_WEB_USER"), "ODOO_WEB_PASS :"+os.Getenv("ODOO_WEB_PASS"), err) // Line 44
		return nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer loginResp.Body.Close()
	ODOO_LOGIN_COOKIE = loginResp.Cookies()
	return loginResp.Cookies(), nil
}
func SendDataToOdoo(jsonData []byte) (string, error) {
	// Prepare log file
	currentDate := time.Now().Format("2006_01_02") // YYYY_MM_DD

	// Log folder
	logDir := "./log/send_data_to_odoo"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create log directory: %w", err)
	}

	// Create log file name
	logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error(err)
		return "", fmt.Errorf("error opening log file: %w", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	// Log the payload being sent
	logger.Printf("%s Request Payload:\n%s\n", time.Now().Format(time.RFC3339), string(jsonData))
	// Create the HTTP request to Odoo with the session ID
	req, err := http.NewRequest("POST", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage/post_data", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error(err)
		logger.Printf("Error creating request: %v\n", err)
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Add cookies from the previous response to the new request
	for _, cookie := range ODOO_LOGIN_COOKIE {
		req.AddCookie(cookie) // Use AddCookie to append cookies to the request
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		time.Sleep(10 * time.Second)
		logger.Printf("Error sending request: %v\n", err)
		LoginOdoo()
		SendDataToOdoo(jsonData)
		logger.Printf("Error sending request: %v\n", err)
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		logger.Printf("Error reading response body: %v\n", err)
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return string(responseBody), fmt.Errorf("received non-200 response from Odoo: %s\n%s", resp.Status, string(responseBody))
	}
	// Log the response
	logger.Printf("%s Response:\n%s\n ----- END ----- \n\n", time.Now().Format(time.RFC3339), string(responseBody))
	return string(responseBody), nil
}
func UpdateDataToOdooByEmail(email string, newEmail, phoneNumber, accHolderName, accHolderNumber, partnerName, street string) (string, error) {
	// Prepare log file
	// currentDate := time.Now().Format("2006_01_02") // YYYY_MM_DD
	// Use StringBuilder for logging
	var logBuffer strings.Builder

	// Defer: write log to file at the end
	defer func() {
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " \n\n ----- END ----- \n\n")
		currentDate := time.Now().Format("2006_01_02")
		logDir := "./log/update_data_to_odoo"
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			logrus.Error("[ERROR][ODOO_DIST]failed to create log directory:", err)
			return
		}
		logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Error("[ERROR][ODOO_DIST]error opening log file:", err)
			return
		}
		defer logFile.Close()
		if _, err := logFile.WriteString(logBuffer.String()); err != nil {
			logrus.Error("[ERROR][ODOO_DIST]error writing log file:", err)
		}
	}()
	// Log folder
	// if err := os.MkdirAll(logDir, 0755); err != nil {
	// 	return "", fmt.Errorf("failed to create log directory: %w", err)
	// }

	// Create log file name
	// logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)

	// logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return "", fmt.Errorf("error opening log file: %w", err)
	// }
	// defer logFile.Close()

	// logger := log.New(logFile, "", log.LstdFlags)

	client := &http.Client{}

	// Step 1: Get res.users data by email (login field)
	getUserPayload := map[string]interface{}{
		"model":  "res.users",
		"domain": [][]interface{}{{"login", "=", email}},
		"fields": []string{"id", "name"},
	}

	getUserJSON, err := json.Marshal(getUserPayload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + fmt.Sprintf("error marshalling payload: %v\n", err))
		return "", fmt.Errorf("error marshalling get user payload: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Get User Request Payload GET /iid_api_manage :\n%s\n", string(getUserJSON)))

	// Create GET request for res.users
	getUserReq, err := http.NewRequest("GET", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage", bytes.NewBuffer(getUserJSON))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error creating GET /iid_api_manage user request: %v\n", err))
		return "", fmt.Errorf("error creating get user request: %w", err)
	}

	getUserReq.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_LOGIN_COOKIE {
		getUserReq.AddCookie(cookie)
	}

	getUserResp, err := client.Do(getUserReq)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error sending GET /iid_api_manage user request: %v\n", err))
		return "", fmt.Errorf("error sending get user request: %w", err)
	}
	defer getUserResp.Body.Close()

	getUserResponseBody, err := io.ReadAll(getUserResp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error reading GET /iid_api_manage user response body: %v\n", err))
		return "", fmt.Errorf("error reading get user response body: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("GET /iid_api_manage User Response:\n%s\n", string(getUserResponseBody)))

	if getUserResp.StatusCode != http.StatusOK {
		return string(getUserResponseBody), fmt.Errorf("received non-200 response from Odoo GET /iid_api_manage user: %s\n%s", getUserResp.Status, string(getUserResponseBody))
	}

	// Parse user response to get user ID
	var userResponse map[string]interface{}
	if err := json.Unmarshal(getUserResponseBody, &userResponse); err != nil {
		return "", fmt.Errorf("error parsing user response: %w", err)
	}

	// Extract user ID (assuming the response contains an array of users)
	// var userID string
	// if data, ok := userResponse["data"].([]interface{}); ok && len(data) > 0 {
	// 	if user, ok := data[0].(map[string]interface{}); ok {
	// 		if id, ok := user["id"].(float64); ok {
	// 			userID = fmt.Sprintf("%.0f", id)
	// 		}
	// 	}
	// }

	// if userID == "" {
	// 	return "", fmt.Errorf("user not found with email: %s", email)
	// }
	var userID string
	if result, ok := userResponse["result"].([]interface{}); ok && len(result) > 0 {
		if resultData, ok := result[0].(map[string]interface{}); ok {
			if records, ok := resultData["records"].([]interface{}); ok && len(records) > 0 {
				if user, ok := records[0].(map[string]interface{}); ok {
					if id, ok := user["id"].(float64); ok {
						userID = fmt.Sprintf("%.0f", id)
					}
				}
			}
		}
	}

	if userID == "" {
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("GET /iid_api_manage User not found with email: %s\n", email))
		return "", fmt.Errorf("user not found with email: %s", email)
	}

	// Step 2: Update res.users record
	updateUserPayload := map[string]interface{}{
		"model": "res.users",
		"records": map[string]interface{}{
			userID: map[string]interface{}{
				"login":             newEmail,
				"phone":             phoneNumber,
				"acc_holder_number": accHolderNumber,
				"acc_holder_name":   accHolderName,
				"partner_name":      partnerName,
				"street":            street,
			},
		},
	}

	updateUserJSON, err := json.Marshal(updateUserPayload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error marshalling update user payload: %v\n", err))
		return "", fmt.Errorf("error marshalling update user payload: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("%s Update User Request Payload PUT /iid_api_manage :\n%s\n", time.Now().Format(time.RFC3339), string(updateUserJSON)))

	// Create PUT request for res.users
	updateUserReq, err := http.NewRequest("PUT", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage", bytes.NewBuffer(updateUserJSON))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error creating update PUT /iid_api_manage user request: %v\n", err))
		return "", fmt.Errorf("error creating update user request: %w", err)
	}

	updateUserReq.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_LOGIN_COOKIE {
		updateUserReq.AddCookie(cookie)
	}

	updateUserResp, err := client.Do(updateUserReq)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error sending update PUT /iid_api_manage user request: %v\n", err))
		return "", fmt.Errorf("error sending update user request: %w", err)
	}
	defer updateUserResp.Body.Close()

	updateUserResponseBody, err := io.ReadAll(updateUserResp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error reading update PUT /iid_api_manage user response body: %v\n", err))
		return "", fmt.Errorf("error reading update user response body: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf(" Update User PUT /iid_api_manage Response:\n%s\n", string(updateUserResponseBody)))

	if updateUserResp.StatusCode != http.StatusOK {
		return string(updateUserResponseBody), fmt.Errorf("received non-200 response from Odoo PUT /iid_api_manage update user: %s\n%s\n", updateUserResp.Status, string(updateUserResponseBody))
	}

	// Step 3: Get res.partner data by email
	getPartnerPayload := map[string]interface{}{
		"model":  "res.partner",
		"domain": [][]interface{}{{"email", "=", newEmail}}, // Use newEmail since we just updated it
		"fields": []string{"id", "name"},
	}

	getPartnerJSON, err := json.Marshal(getPartnerPayload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error marshalling get partner payload: %v\n", err))
		return "", fmt.Errorf("error marshalling get partner payload: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Get Partner Request Payload GET /iid_api_manage:\n%s\n", string(getPartnerJSON)))

	// Create GET request for res.partner     >>>>>>>>>>>>>>>> RES PARTNER <<<<<<<<<<<<<<<<<
	getPartnerReq, err := http.NewRequest("GET", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage", bytes.NewBuffer(getPartnerJSON))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error creating GET /iid_api_manage partner request: %v\n", err))
		return "", fmt.Errorf("error creating get partner request: %w", err)
	}

	getPartnerReq.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_LOGIN_COOKIE {
		getPartnerReq.AddCookie(cookie)
	}

	getPartnerResp, err := client.Do(getPartnerReq)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error sending GET /iid_api_manage partner request: %v\n", err))
		return "", fmt.Errorf("error sending get partner request: %w", err)
	}
	defer getPartnerResp.Body.Close()

	getPartnerResponseBody, err := io.ReadAll(getPartnerResp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error reading get partner response body: %v\n", err))
		return "", fmt.Errorf("error reading get partner response body: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("GET /iid_api_manage Partner Response:\n%s\n", string(getPartnerResponseBody)))

	if getPartnerResp.StatusCode != http.StatusOK {
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("received non-200 response from GET /iid_api_manage partner: %s\n%s\n", getPartnerResp.Status, string(getPartnerResponseBody)))
		return string(getPartnerResponseBody), fmt.Errorf("received non-200 response from GET /iid_api_manage partner: %s\n%s\n", getPartnerResp.Status, string(getPartnerResponseBody))
	}

	// Parse partner response to get partner ID
	var partnerResponse map[string]interface{}
	if err := json.Unmarshal(getPartnerResponseBody, &partnerResponse); err != nil {
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("error parsing partner response GET /iid_api_manage: %s\n", err.Error()))
		return "", fmt.Errorf("error parsing partner response: %w", err)
	}
	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Partner Response Body GET /iid_api_manage: %s\n", string(getPartnerResponseBody)))
	// logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Partner Response Parsed: %v\n", partnerResponse))
	// Extract partner ID
	var partnerID string
	// if result, ok := partnerResponse["result"].([]interface{}); ok && len(result) > 0 {
	// 	if resultData, ok := result[0].(map[string]interface{}); ok {
	// 		if records, ok := resultData["records"].([]interface{}); ok && len(records) > 0 {
	// 			if partner, ok := records[0].(map[string]interface{}); ok {
	// 				if id, ok := partner["id"].(float64); ok {
	// 					partnerID = fmt.Sprintf("%.0f", id)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	response := strings.ReplaceAll(string(getPartnerResponseBody), " ", "")
	for part := range strings.SplitSeq(response, `[{"id":`) {
		if len(part) > 0 && part[0] >= '0' && part[0] <= '9' {
			// Extract the ID from this part
			idEnd := strings.Index(part, ",")
			if idEnd > 0 {
				// partnerID = part[:idEnd]
				partnerID = strings.Split(part, `,"name"`)[0]
				break
			}
		}
	}

	if partnerID == "" {
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error Parsing partnerID GET /iid_api_manage: %s\n", string(getPartnerResponseBody)))
		return "", fmt.Errorf("partner not found with email: %s", newEmail)
	}
	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Partner ID :%s\n", partnerID))

	// Step 4: Update res.partner record
	updatePartnerPayload := map[string]interface{}{
		"model": "res.partner",
		"records": map[string]interface{}{
			partnerID: map[string]interface{}{
				"email":             newEmail,
				"acc_holder_number": accHolderNumber,
				"acc_holder_name":   accHolderName,
				"partner_name":      partnerName,
				"street":            street,
			},
		},
	}

	updatePartnerJSON, err := json.Marshal(updatePartnerPayload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error marshalling PUT /iid_api_manage partner payload: %v\n", err))
		return "", fmt.Errorf("error marshalling update partner payload: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Update Partner Request Payload:\n%s\n", string(updatePartnerJSON)))

	// Create PUT request for res.partner
	updatePartnerReq, err := http.NewRequest("PUT", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage", bytes.NewBuffer(updatePartnerJSON))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error creating PUT /iid_api_manage partner request: %v\n", err))
		return "", fmt.Errorf("error creating update partner request: %w", err)
	}

	updatePartnerReq.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_LOGIN_COOKIE {
		updatePartnerReq.AddCookie(cookie)
	}

	updatePartnerResp, err := client.Do(updatePartnerReq)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error sending PUT /iid_api_manage partner request: %v\n", err))
		return "", fmt.Errorf("error sending update partner request: %w", err)
	}
	defer updatePartnerResp.Body.Close()

	updatePartnerResponseBody, err := io.ReadAll(updatePartnerResp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Error reading PUT /iid_api_manage partner response body: %v\n", err))
		return "", fmt.Errorf("error reading update partner response body: %w", err)
	}

	logBuffer.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("Update PUT /iid_api_manage Response:\n%s", string(updatePartnerResponseBody)))

	if updatePartnerResp.StatusCode != http.StatusOK {
		return string(updatePartnerResponseBody), fmt.Errorf("received non-200 response from Odoo PUT /iid_api_manage update partner: %s\n%s", updatePartnerResp.Status, string(updatePartnerResponseBody))
	}

	// Return final response
	return string(updatePartnerResponseBody), nil
}

func SendRemoveDataToOdooByEmail(email string) (string, error) {
	// Prepare log file
	currentDate := time.Now().Format("2006_01_02") // YYYY_MM_DD

	// Log folder
	logDir := "./log/send_data_to_odoo"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create log directory: %w", err)
	}

	// Create log file name
	logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error(err)
		return "", fmt.Errorf("error opening log file: %w", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	// Prepare JSON payload
	payload := map[string]string{
		"email": email,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logrus.Error(err)
		logger.Printf("Error marshalling payload: %v\n", err)
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	// Log the payload being sent
	logger.Printf("%s Request Payload:\n%s\n", time.Now().Format(time.RFC3339), string(jsonData))

	// Create the HTTP request to Odoo to delete user by email
	req, err := http.NewRequest("POST", os.Getenv("ODOO_WEB_URL")+"/iid_api_manage/delete_user", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error(err)
		logger.Printf("Error creating request: %v\n", err)
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Add cookies from the previous login to the new request
	for _, cookie := range ODOO_LOGIN_COOKIE {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		time.Sleep(10 * time.Second)
		logger.Printf("Error sending request: %v\n", err)
		// try to re-login and retry once
		if _, lerr := LoginOdoo(); lerr == nil {
			return SendRemoveDataToOdooByEmail(email)
		}
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		logger.Printf("Error reading response body: %v\n", err)
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Log the response
	logger.Printf("%s Response (status: %s):\n%s\n ----- END ----- \n\n", time.Now().Format(time.RFC3339), resp.Status, string(responseBody))

	if resp.StatusCode != http.StatusOK {
		return string(responseBody), fmt.Errorf("received non-200 response from Odoo: %s\n%s", resp.Status, string(responseBody))
	}

	return string(responseBody), nil
}

func LoginOdooCS() ([]*http.Cookie, error) {
	loginURL := os.Getenv("ODOO_CSNA_RSS_URL") + "/web/session/authenticate"
	loginPayload := map[string]interface{}{
		"jsonrpc": "2.0",
		"params": map[string]string{
			"db":       os.Getenv("ODOO_CSNA_RSS_DB"),
			"login":    os.Getenv("ODOO_CSNA_RSS_USER"),
			"password": os.Getenv("ODOO_CSNA_RSS_PASS"),
		},
	}

	loginJSON, err := json.Marshal(loginPayload)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error creating login payload: %w", err)
	}

	loginReq, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(loginJSON))
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error creating login request: %w", err)
	}

	loginReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	loginResp, err := client.Do(loginReq)
	if err != nil {
		logrus.Error("Info : "+string(loginJSON), err) // Line 44
		return nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer loginResp.Body.Close()
	ODOO_CSNA_LOGIN_COOKIE = loginResp.Cookies()
	return loginResp.Cookies(), nil
}
func CreateDataToOdooCS(tid, mid, email, phone, serial_number, address, edc_type string) (string, error) {
	// Use StringBuilder for logging
	var logBuffer strings.Builder

	// Defer: write log to file at the end
	defer func() {
		currentDate := time.Now().Format("2006_01_02")
		logDir := "./log/create_data_to_odoo_cs"
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			logrus.Error("failed to create log directory:", err)
			return
		}
		logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Error("error opening log file:", err)
			return
		}
		defer logFile.Close()
		if _, err := logFile.WriteString(logBuffer.String()); err != nil {
			logrus.Error("error writing log file:", err)
		}
	}()

	payload := map[string]any{
		"params": map[string]any{
			"model":           "res.partner",
			"name":            tid + mid,
			"email":           email,
			"x_serial_number": serial_number,
			"phone":           phone,
			"street":          address,
			"x_product":       edc_type,
			"customer_rank":   1,
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("error marshalling payload: %v\n", err))
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	logBuffer.WriteString(fmt.Sprintf(
		"%s Request Payload:\n%s\n",
		time.Now().Format(time.RFC3339),
		string(jsonData),
	))

	req, err := http.NewRequest("POST", os.Getenv("ODOO_CSNA_RSS_URL")+"/api/createdata", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("Error creating request: %v\n", err))
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_CSNA_LOGIN_COOKIE {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		time.Sleep(10 * time.Second)
		logBuffer.WriteString(fmt.Sprintf("Error sending request: %v\n", err))
		LoginOdooCS()
		CreateDataToOdooCS(tid, mid, email, phone, serial_number, address, edc_type)
		logBuffer.WriteString(fmt.Sprintf("Error sending request: %v\n", err))
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("Error reading response body: %v\n", err))
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Non-OK HTTP status: %s", resp.Status)
		logBuffer.WriteString(fmt.Sprintf("Non-OK HTTP status: %s\n%s\n", resp.Status, string(responseBody)))
		return "", fmt.Errorf("received non-200 response from Odoo: %s\n%s", resp.Status, string(responseBody))
	}
	logBuffer.WriteString(fmt.Sprintf("Response:\n%s\n", string(responseBody)))
	return string(responseBody), nil
}
func UpdateDataToOdooCS(tid, mid, email, phone, serial_number, address, edc_type string) (string, error) {

	// Use StringBuilder for logging
	var logBuffer strings.Builder

	// Defer: write log to file at the end
	defer func() {
		currentDate := time.Now().Format("2006_01_02")
		logDir := "./log/update_data_to_odoo_cs"
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			logrus.Error("failed to create log directory:", err)
			return
		}
		logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Error("error opening log file:", err)
			return
		}
		defer logFile.Close()
		if _, err := logFile.WriteString(logBuffer.String()); err != nil {
			logrus.Error("error writing log file:", err)
		}
	}()

	payload := map[string]any{
		"jsonrpc": "2.0",
		"params": map[string]any{
			"model":           "res.partner",
			"name":            tid + mid,
			"email":           email,
			"x_serial_number": serial_number,
			"phone":           phone,
			"street":          address,
			"x_product":       edc_type,
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("error marshalling payload: %v\n", err))
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	logBuffer.WriteString(fmt.Sprintf(
		"%s Request Payload:\n%s\n",
		time.Now().Format(time.RFC3339),
		string(jsonData),
	))

	req, err := http.NewRequest("POST", os.Getenv("ODOO_CSNA_RSS_URL")+"/api/updatename", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("Error creating request: %v\n", err))
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_CSNA_LOGIN_COOKIE {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		time.Sleep(10 * time.Second)
		logBuffer.WriteString(fmt.Sprintf("Error sending request: %v\n", err))
		LoginOdooCS()
		UpdateDataToOdooCS(tid, mid, email, phone, serial_number, address, edc_type)
		logBuffer.WriteString(fmt.Sprintf("Error sending request: %v\n", err))
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("Error reading response body: %v\n", err))
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Non-OK HTTP status: %s", resp.Status)
		logBuffer.WriteString(fmt.Sprintf("Non-OK HTTP status: %s\n%s\n", resp.Status, string(responseBody)))
		return "", fmt.Errorf("received non-200 response from Odoo: %s\n%s", resp.Status, string(responseBody))
	}
	logBuffer.WriteString(fmt.Sprintf("Response:\n%s\n", string(responseBody)))
	return string(responseBody), nil
}
func DeleteDataToOdooCS(tid, mid string) (string, error) {
	// Use StringBuilder for logging
	var logBuffer strings.Builder

	// Defer: write log to file at the end
	defer func() {
		currentDate := time.Now().Format("2006_01_02")
		logDir := "./log/delete_data_to_odoo_cs"
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			logrus.Error("[ERROR][ODOO_CS]failed to create log directory:", err)
			return
		}
		logFileName := fmt.Sprintf("%s/request_%s.log", logDir, currentDate)
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Error("[ERROR][ODOO_CS]error opening log file:", err)
			return
		}
		defer logFile.Close()
		if _, err := logFile.WriteString(logBuffer.String()); err != nil {
			logrus.Error("[ERROR][ODOO_CS]error writing log file:", err)
		}
	}()

	payload := map[string]any{
		"jsonrpc": "2.0",
		"params": map[string]any{
			"name": tid + mid,
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("[ERROR][ODOO_CS]error marshalling payload: %v\n", err))
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	logBuffer.WriteString(fmt.Sprintf(
		"%s Request Payload:\n%s\n",
		time.Now().Format(time.RFC3339),
		string(jsonData),
	))

	req, err := http.NewRequest("POST", os.Getenv("ODOO_CSNA_RSS_URL")+"/api/delete_partner", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("[ERROR][ODOO_CS]Error creating request: %v\n", err))
		return "", fmt.Errorf("[ERROR][ODOO_CS]error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for _, cookie := range ODOO_CSNA_LOGIN_COOKIE {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		time.Sleep(10 * time.Second)
		logBuffer.WriteString(fmt.Sprintf("[ERROR][ODOO_CS]Error sending request: %v\n", err))
		LoginOdooCS()
		DeleteDataToOdooCS(tid, mid)
		logBuffer.WriteString(fmt.Sprintf("[ERROR][ODOO_CS]Error sending request: %v\n", err))
		return "", fmt.Errorf("[ERROR][ODOO_CS]error sending request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		logBuffer.WriteString(fmt.Sprintf("[ERROR][ODOO_CS]Error reading response body: %v\n", err))
		return "", fmt.Errorf("[ERROR][ODOO_CS]error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Non-OK HTTP status: %s", resp.Status)
		logBuffer.WriteString(fmt.Sprintf("Non-OK HTTP status: %s\n%s\n", resp.Status, string(responseBody)))
		return "", fmt.Errorf("[ERROR][ODOO_CS]received non-200 response from Odoo: %s\n%s", resp.Status, string(responseBody))
	}
	logBuffer.WriteString(fmt.Sprintf("Response:\n%s\n", string(responseBody)))
	return string(responseBody), nil
}
