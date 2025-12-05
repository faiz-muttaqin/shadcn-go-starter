package util

import (
	"encoding/base64"
	"sort"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/argon2"
)

type SaltArgon struct {
	Salt     string
	Position int
}

// ===== Utility Functions =====

func InsertStringAtPositionsArgon2(original string, salts ...SaltArgon) string {
	sort.Slice(salts, func(i, j int) bool {
		return salts[i].Position < salts[j].Position
	})

	for i, salt := range salts {
		salt.Position = salt.Position + len(salt.Salt)*i
		original = original[:salt.Position] + salt.Salt + original[salt.Position:]
	}

	return original
}

func InsertRandomStringAtPositionsArgon2(original string, randomStringLength int, positions ...int) string {
	sort.Ints(positions)
	for i, position := range positions {
		position = position + randomStringLength*i
		randomString := GenerateRandomHexaString(randomStringLength)
		original = original[:position] + randomString + original[position:]
	}
	return original
}

func RemoveSubstringAtPositionsArgon2(original string, length int, positions ...int) string {
	sort.Ints(positions)
	for i := len(positions) - 1; i >= 0; i-- {
		pos := positions[i]
		original = original[:pos] + original[pos+length:]
	}
	return original
}

// ===== Argon2 Password Hashing =====

func GenerateSaltedPasswordArgon2(password string) string {
	// Generate 4 salt parts
	saltParts := []string{
		GenerateRandomHexaString(4),
		GenerateRandomHexaString(4),
		GenerateRandomHexaString(4),
		GenerateRandomHexaString(4),
	}
	fullSalt := saltParts[0] + saltParts[1] + saltParts[2] + saltParts[3]

	// Salted password
	saltedPassword := InsertStringAtPositionsArgon2(password,
		SaltArgon{Salt: saltParts[0], Position: 2},
		SaltArgon{Salt: saltParts[1], Position: 5},
		SaltArgon{Salt: saltParts[2], Position: 7},
		SaltArgon{Salt: saltParts[3], Position: 8},
	)

	// Argon2 hashing
	hash := argon2.IDKey([]byte(saltedPassword), []byte(fullSalt), 1, 64*1024, 4, 32)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// Obfuscate hash with inserted strings
	obfuscatedHash := InsertRandomStringAtPositionsArgon2(encodedHash, 2, 5, 8, 10, 18)

	// Return salt + obfuscated hash
	return fullSalt + obfuscatedHash
}

func IsPasswordMatchedArgon2(password, full string) bool {
	if len(full) < 16 {
		return false
	}

	// Extract salt parts
	saltParts := []string{
		full[0:4],
		full[4:8],
		full[8:12],
		full[12:16],
	}
	fullSalt := saltParts[0] + saltParts[1] + saltParts[2] + saltParts[3]

	// Extract and de-obfuscate hash
	obfuscatedHash := full[16:]
	realHash := RemoveSubstringAtPositions(obfuscatedHash, 2, 5, 8, 10, 18)

	// Reconstruct salted password
	saltedPassword := InsertStringAtPositionsArgon2(password,
		SaltArgon{Salt: saltParts[0], Position: 2},
		SaltArgon{Salt: saltParts[1], Position: 5},
		SaltArgon{Salt: saltParts[2], Position: 7},
		SaltArgon{Salt: saltParts[3], Position: 8},
	)

	// Re-hash input
	hash := argon2.IDKey([]byte(saltedPassword), []byte(fullSalt), 1, 64*1024, 4, 32)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return encodedHash == realHash
}

// ===== Testing =====

func TestSaltArgon2(testingPassword string) {
	password := "password123"

	logrus.Println("Input password       :", password)
	logrus.Println("Testing against      :", testingPassword)

	// Generate hash
	saltedHash := GenerateSaltedPassword(password)
	logrus.Println("Generated saltedHash:", saltedHash)

	// Validate
	match := IsPasswordMatched(testingPassword, saltedHash)
	logrus.Println("Password match result:", match)
	// if match {
	// 	logrus.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	// 	logrus.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	// 	logrus.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	// } else {
	// 	logrus.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	// 	logrus.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	// 	logrus.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	// 	logrus.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	// }
}
