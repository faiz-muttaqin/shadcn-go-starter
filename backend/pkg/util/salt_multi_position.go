package util

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"

	"github.com/sirupsen/logrus"
)

type Salt struct {
	Salt     string
	Position int
}

// Function to insert random strings at multiple positions
func InsertStringAtPositions(original string, salts ...Salt) string {
	// Sort positions in ascending order
	// sort.Ints(positions[int])
	sort.Slice(salts, func(i, j int) bool {
		return salts[i].Position < salts[j].Position
	})

	for i, salt := range salts {
		salt.Position = salt.Position + len(salt.Salt)*i

		original = original[:salt.Position] + salt.Salt + original[salt.Position:]
	}

	return original
}

// Function to insert random strings at multiple positions
func InsertRandomStringAtPositions(original string, randomStringLength int, positions ...int) string {
	// Sort positions in ascending order
	sort.Ints(positions)

	for i, position := range positions {
		position = position + randomStringLength*i
		randomString := GenerateRandomHexaString(randomStringLength)

		original = original[:position] + randomString + original[position:]
	}

	return original
}

// Function to remove substrings at multiple positions
func RemoveSubstringAtPositions(original string, length int, positions ...int) string {
	// Sort positions in descending order
	sort.Ints(positions)

	// Adjust positions to account for previously inserted strings
	for _, position := range positions {
		original = original[:position] + original[position+length:]
	}
	return original
}

func GenerateSaltedPassword(password string) string {
	salt_a := GenerateRandomHexaString(4)
	salt_b := GenerateRandomHexaString(4)
	salt_c := GenerateRandomHexaString(4)
	salt_d := GenerateRandomHexaString(4)

	salted_password := InsertStringAtPositions(password,
		Salt{Salt: salt_a, Position: 2},
		Salt{Salt: salt_b, Position: 5},
		Salt{Salt: salt_c, Position: 7},
		Salt{Salt: salt_d, Position: 8},
	)

	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write the input data to the hash
	hash.Write([]byte(salted_password))

	// Get the finalized hash result as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashed_password := hex.EncodeToString(hashBytes)

	salted_hashed_password := InsertRandomStringAtPositions(hashed_password, 2, 5, 8, 10, 18)

	salt_with_salted_hash := salt_a + salt_b + salt_c + salt_d + salted_hashed_password

	return salt_with_salted_hash

}

func IsPasswordMatched(password, salt_with_salted_hash string) bool {
	valid_salted_hashed_password := salt_with_salted_hash[16:]
	valid_hashed_password := RemoveSubstringAtPositions(valid_salted_hashed_password, 2, 5, 8, 10, 18)

	// how to parse salt_with_salted_hash get the 16 first string and broke them into every 4 char
	salt_a := salt_with_salted_hash[:4]
	salt_b := salt_with_salted_hash[4:8]
	salt_c := salt_with_salted_hash[8:12]
	salt_d := salt_with_salted_hash[12:16]

	salted_check_password := InsertStringAtPositions(password,
		Salt{Salt: salt_a, Position: 2},
		Salt{Salt: salt_b, Position: 5},
		Salt{Salt: salt_c, Position: 7},
		Salt{Salt: salt_d, Position: 8},
	)

	// Create a new SHA-256 hash
	hash := sha256.New()

	// Write the input data to the hash
	hash.Write([]byte(salted_check_password))

	// Get the finalized hash result as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashed_check_password := hex.EncodeToString(hashBytes)

	return hashed_check_password == valid_hashed_password

}
func TestSalt(testing_password string) {
	password := "password123"

	logrus.Println("default_password")
	logrus.Println(password)
	logrus.Println("testing_password")
	logrus.Println(testing_password)

	// Generate salted and hashed password
	saltedHash := GenerateSaltedPassword(password)

	logrus.Println("saltedHash")
	logrus.Println(saltedHash)
	// Check if the password matches
	isMatched := IsPasswordMatched(testing_password, saltedHash)

	logrus.Println("Password match result:", isMatched)
}
