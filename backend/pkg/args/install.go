package args

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var serviceName string

func Install(envFiles ...string) error {
	if len(os.Args) > 1 && os.Args[1] == "install" {

		useEnv := true
		if len(os.Args) > 2 && os.Args[2] == "no-env" {
			useEnv = false
		}
		// Determine environment file name
		var envFileName string
		if len(envFiles) > 0 && envFiles[0] != "" {
			envFileName = envFiles[0]
		} else {
			envFileName = ".env" // Default
		}
		// Get the current directory
		currentDir, err := os.Getwd()
		if err != nil {
			logrus.Error(err)
			fmt.Printf("Error getting current directory: %v\n", err)
			return err
		}

		// Get executable path
		exePath, err := os.Executable()
		if err != nil {
			logrus.Error(err)
			fmt.Printf("Error getting executable path: %v\n", err)
			return err
		}
		exePath = filepath.Clean(exePath)
		// exeName := filepath.Base(exePath)

		// Determine service name
		_, folderName := filepath.Split(currentDir)
		folderName = strings.Trim(folderName, "/\\")
		serviceName = folderName
		if os.Getenv("SERVICE_NAME") != "" {
			serviceName = os.Getenv("SERVICE_NAME")
		}

		// Find .env file if useEnv is true
		var envPath string
		if useEnv {
			envPath = findEnvFile(currentDir, filepath.Dir(exePath), envFileName)
			if envPath == "" {
				return fmt.Errorf(".env file not found in current directory or executable directory")
			}
			fmt.Printf("Using .env file: %s\n", envPath)
			currentDir = filepath.Dir(envPath)
		}

		if err := installService(serviceName, currentDir, exePath, envPath, useEnv); err != nil {
			fmt.Printf("Error installing service: %v\n", err)
			return err
		}
		return errors.New("App Installed As Service")
	}
	return nil
}

// findEnvFile searches for .env file in multiple locations
func findEnvFile(currentDir, exeDir, envFileName string) string {
	// Priority 1: Current working directory
	envPath := filepath.Join(currentDir, envFileName)
	if _, err := os.Stat(envPath); err == nil {
		return envPath
	}

	// Priority 2: Executable directory
	envPath = filepath.Join(exeDir, envFileName)
	if _, err := os.Stat(envPath); err == nil {
		return envPath
	}
	if envFileName == ".env" {
		// Priority 3: Check for .env.production
		envPath = filepath.Join(currentDir, ".env.production")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		envPath = filepath.Join(exeDir, ".env.production")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
	}

	return ""
}

func installService(serviceDesc, currentDir, exePath, envPath string, useEnv bool) error {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error getting current user: %v", err)
	}

	// Check if executable exists
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		return fmt.Errorf("executable not found at: %s", exePath)
	}

	// Define the service file path
	serviceFilePath := fmt.Sprintf("/etc/systemd/system/%s.service", serviceName)

	// Create the service file content
	var serviceFileContent string

	if useEnv {
		// With environment file
		serviceFileContent = fmt.Sprintf(`[Unit]
Description=%s
After=network.target

[Service]
Type=simple
User=%s
Group=%s
WorkingDirectory=%s
ExecStart=%s
Restart=always
RestartSec=3
EnvironmentFile=%s
LimitNOFILE=1000000

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=%s

[Install]
WantedBy=multi-user.target
`, serviceDesc, currentUser.Username, currentUser.Username, currentDir, exePath, envPath, serviceName)
	} else {
		// Without environment file
		serviceFileContent = fmt.Sprintf(`[Unit]
Description=%s
After=network.target

[Service]
Type=simple
User=%s
Group=%s
WorkingDirectory=%s
ExecStart=%s
Restart=always
RestartSec=3
LimitNOFILE=1000000

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=%s

[Install]
WantedBy=multi-user.target
`, serviceDesc, currentUser.Username, currentUser.Username, currentDir, exePath, serviceName)
	}

	fmt.Printf("Creating service file at: %s\n", serviceFilePath)
	fmt.Printf("Executable path: %s\n", exePath)
	fmt.Printf("Working directory: %s\n", currentDir)
	if useEnv {
		fmt.Printf("Environment file: %s\n", envPath)
	} else {
		fmt.Println("No environment file will be used")
	}

	// Check if we have sudo access
	if os.Geteuid() != 0 {
		fmt.Println("\n⚠️  This command requires root privileges.")
		fmt.Println("Please run with sudo or as root user.")
		fmt.Println("\nExample:")
		fmt.Printf("  sudo %s install\n", exePath)
		if !useEnv {
			fmt.Printf("  sudo %s install no-env\n", exePath)
		}
		return fmt.Errorf("insufficient permissions: need root to install systemd service")
	}

	// Write the service file
	if err := os.WriteFile(serviceFilePath, []byte(serviceFileContent), 0644); err != nil {
		return fmt.Errorf("error writing service file: %v", err)
	}

	fmt.Printf("✓ Service file created at %s\n", serviceFilePath)

	// Reload systemd daemon
	fmt.Println("Reloading systemd daemon...")
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return fmt.Errorf("error reloading systemd daemon: %v", err)
	}
	fmt.Println("✓ Systemd daemon reloaded")

	// Enable the service
	fmt.Printf("Enabling service %s...\n", serviceName)
	if err := exec.Command("systemctl", "enable", fmt.Sprintf("%s.service", serviceName)).Run(); err != nil {
		return fmt.Errorf("error enabling service: %v", err)
	}
	fmt.Printf("✓ Service %s enabled\n", serviceName)

	// Start the service
	fmt.Printf("Starting service %s...\n", serviceName)
	if err := exec.Command("systemctl", "start", fmt.Sprintf("%s.service", serviceName)).Run(); err != nil {
		return fmt.Errorf("error starting service: %v", err)
	}
	fmt.Printf("✓ Service %s started\n", serviceName)

	// Show service status
	fmt.Printf("\n🎉 Service installed successfully!\n\n")
	fmt.Println("Useful commands:")
	fmt.Printf("  sudo systemctl status %s    - Check service status\n", serviceName)
	fmt.Printf("  sudo systemctl restart %s   - Restart service\n", serviceName)
	fmt.Printf("  sudo systemctl stop %s      - Stop service\n", serviceName)
	fmt.Printf("  sudo journalctl -u %s -f    - View logs\n", serviceName)
	fmt.Printf("  sudo systemctl disable %s   - Disable service\n", serviceName)

	return nil
}
