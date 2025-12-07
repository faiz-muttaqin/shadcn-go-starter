package version

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
)

type Info struct {
	Version     string   `json:"version"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Host        string   `json:"host,omitempty"`
	BasePath    string   `json:"basepath,omitempty"`
	Schemes     []string `json:"schemes,omitempty"`
	Datetime    string   `json:"datetime,omitempty"`
}

func getExecutablePath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	return exe
}

func VersionFilePath() string {
	return filepath.Join(os.Getenv("APP_DIR"), "package.json")
}

// --- Logic pembentukan versi ---
func nextPatchVersion(v string) string {
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return "0.0.1"
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "0.0.1"
	}
	return fmt.Sprintf("%s.%s.%d", parts[0], parts[1], patch+1)
}

// updateJSONField updates a specific field in JSON while preserving order and formatting
func updateJSONField(jsonData []byte, key string, value interface{}) ([]byte, error) {
	// Convert value to JSON string
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	valueStr := string(valueBytes)

	// Create regex pattern to find the key
	// Matches: "key": "value" or "key": {...} or "key": [...]
	pattern := regexp.MustCompile(`("` + regexp.QuoteMeta(key) + `"\s*:\s*)([^,\n\}]+(,|\s*\}|\s*\]))`)

	// Check if key exists
	if pattern.Match(jsonData) {
		// Replace existing value
		result := pattern.ReplaceAll(jsonData, []byte(`${1}`+valueStr+`${3}`))
		return result, nil
	}

	// Key doesn't exist, add it after the first opening brace
	// Find first { and add key after it
	openBrace := bytes.Index(jsonData, []byte("{"))
	if openBrace == -1 {
		return jsonData, fmt.Errorf("invalid JSON structure")
	}

	// Check if there's already content (not empty object)
	nextNonSpace := openBrace + 1
	for nextNonSpace < len(jsonData) && (jsonData[nextNonSpace] == ' ' || jsonData[nextNonSpace] == '\n') {
		nextNonSpace++
	}

	var newField string
	if nextNonSpace < len(jsonData) && jsonData[nextNonSpace] != '}' {
		// Has content, add with comma
		newField = fmt.Sprintf("\n  \"%s\": %s,", key, valueStr)
	} else {
		// Empty object, add without comma
		newField = fmt.Sprintf("\n  \"%s\": %s\n", key, valueStr)
	}

	result := make([]byte, 0, len(jsonData)+len(newField))
	result = append(result, jsonData[:openBrace+1]...)
	result = append(result, []byte(newField)...)
	result = append(result, jsonData[openBrace+1:]...)

	return result, nil
}

// --- Logic utama pembuatan & update VERSION (hanya untuk go run) ---
func Generate(path string) (*Info, error) {
	embedPath := filepath.Join(os.Getenv("APP_DIR"), "package.json")
	if path == "" || func() bool { _, err := os.Stat(path); return os.IsNotExist(err) }() {
		path = embedPath
	}
	now := time.Now()
	var info Info
	info.BasePath = util.Getenv("VITE_BASE_PATH", "")
	info.Host = util.Getenv("APP_LOCAL_HOST", "")

	var originalData []byte
	var packageData map[string]interface{}

	// Check if file exists
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, create new version info
			dirName := filepath.Base(os.Getenv("APP_DIR"))
			info = Info{
				Name:        dirName,
				Version:     "0.0.1",
				Description: "Auto-generated VERSION file",
				Datetime:    now.Format("2006-01-02 15:04:05"),
			}
			packageData = make(map[string]interface{})
		} else {
			return nil, fmt.Errorf("failed to access VERSION file: %v", err)
		}
	} else {
		// File exists, read it
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read package.json: %v", err)
		}
		originalData = data

		// Parse to extract values
		if err := json.Unmarshal(data, &packageData); err != nil {
			fmt.Printf("Warning: Invalid JSON format in package.json file, creating new: %v\n", err)
			dirName := filepath.Base(os.Getenv("APP_DIR"))
			info = Info{
				Version:     "0.0.1",
				Name:        dirName,
				Description: "Auto-generated package.json file (recreated due to invalid format)",
				Datetime:    now.Format("2006-01-02 15:04:05"),
			}
			originalData = nil
		} else {
			// Extract version and increment it
			if v, ok := packageData["version"].(string); ok {
				info.Version = nextPatchVersion(v)
			} else {
				info.Version = "0.0.1"
			}

			if n, ok := packageData["name"].(string); ok {
				info.Name = n
			}

			if d, ok := packageData["description"].(string); ok {
				info.Description = d
			}

			if h, ok := packageData["host"].(string); ok {
				info.Host = h
			}

			if bp, ok := packageData["basepath"].(string); ok {
				info.BasePath = bp
			}

			if schemes, ok := packageData["schemes"].([]interface{}); ok {
				info.Schemes = make([]string, len(schemes))
				for i, scheme := range schemes {
					if s, ok := scheme.(string); ok {
						info.Schemes[i] = s
					}
				}
			}

			info.Datetime = now.Format("2006-01-02 15:04:05")

			// Ensure Name matches folder name
			dirName := filepath.Base(os.Getenv("APP_DIR"))
			if info.Name == "" {
				info.Name = dirName
			}
		}
	}

	// Set environment variables
	os.Setenv("APP_NAME", info.Name)
	os.Setenv("APP_VERSION", info.Version)
	os.Setenv("APP_DESCRIPTION", info.Description)

	var finalData []byte
	var err error

	// 🔧 PRESERVE ORDER: Update fields in original JSON structure
	if len(originalData) > 0 {
		// Start with original data
		finalData = originalData

		// Update only the fields we manage, preserving order
		finalData, err = updateJSONField(finalData, "version", info.Version)
		if err != nil {
			return nil, err
		}

		finalData, err = updateJSONField(finalData, "name", info.Name)
		if err != nil {
			return nil, err
		}

		if info.Description != "" {
			finalData, err = updateJSONField(finalData, "description", info.Description)
			if err != nil {
				return nil, err
			}
		}

		finalData, err = updateJSONField(finalData, "datetime", info.Datetime)
		if err != nil {
			return nil, err
		}

		// if info.Host != "" {
		// 	finalData, err = updateJSONField(finalData, "host", info.Host)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// }

		if info.BasePath != "" {
			finalData, err = updateJSONField(finalData, "basepath", info.BasePath)
			if err != nil {
				return nil, err
			}
		}

		if len(info.Schemes) > 0 {
			finalData, err = updateJSONField(finalData, "schemes", info.Schemes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// No original data, create new JSON with our fields
		packageData["version"] = info.Version
		packageData["name"] = info.Name
		packageData["description"] = info.Description
		packageData["datetime"] = info.Datetime

		// if info.Host != "" {
		// 	packageData["host"] = info.Host
		// }
		if info.BasePath != "" {
			packageData["basepath"] = info.BasePath
		}
		if len(info.Schemes) > 0 {
			packageData["schemes"] = info.Schemes
		}

		finalData, _ = json.MarshalIndent(packageData, "", "  ")
	}

	// Write the file
	if err := os.WriteFile(path, finalData, 0644); err != nil {
		return nil, fmt.Errorf("failed to write package.json file: %v", err)
	}

	// Copy to embed path if different
	if path != embedPath {
		input, err := os.ReadFile(path)
		if err == nil {
			_ = os.WriteFile(embedPath, input, 0644)
		}
	}

	return &info, nil
}

// --- Load versi (bisa dari embed atau file) ---
func Load(embedded []byte) (*Info, error) {

	if len(embedded) > 0 {
		var info Info
		if err := json.Unmarshal(embedded, &info); err != nil {
			return nil, fmt.Errorf("invalid embedded VERSION: %v", err)
		}
		return &info, nil
	}

	data, err := os.ReadFile(VersionFilePath())
	if err != nil {
		return nil, fmt.Errorf("VERSION not found or embedded")
	}
	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("invalid VERSION: %v", err)
	}
	return &info, nil
}

// --- Print ke stdout ---
func Print(info *Info) {
	fmt.Printf("\n📦 %s\n", info.Name)
	fmt.Printf("Version     : %s\n", info.Version)
	fmt.Printf("Description : %s\n", info.Description)
	// if info.Host != "" {
	// 	fmt.Printf("Host        : %s\n", info.Host)
	// }
	// if info.BasePath != "" {
	// 	fmt.Printf("BasePath    : %s\n", info.BasePath)
	// }
	// if len(info.Schemes) > 0 {
	// 	fmt.Printf("Schemes     : %v\n", info.Schemes)
	// }
	fmt.Printf("Datetime    : %s\n", info.Datetime)
}
