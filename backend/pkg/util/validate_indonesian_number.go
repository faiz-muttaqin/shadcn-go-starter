package util

import (
	"fmt"
	"strings"
	"unicode"
)

var validPhonePre = map[string]bool{
	// Telkomsel
	"080": true, "0810": true, "0811": true, "0812": true, "0813": true, "082": true, "0821": true, "0822": true, "0823": true, "0850": true, "0851": true, "0852": true, "0853": true,
	// Indosat IM3
	"0814": true, "0815": true, "0816": true, "0854": true, "0855": true, "0856": true, "0857": true, "0858": true,
	// Tri
	"089": true, "0895": true, "0896": true, "0897": true, "0898": true, "0899": true,
	// XL Axiata XL
	"0817": true, "0818": true, "0819": true, "0859": true, "087": true, "0877": true, "0878": true, "0879": true,
	// XL Axiata AXIS
	"083": true, "0831": true, "0832": true, "0833": true, "0838": true,
	// Smartfren
	"0881": true, "0882": true, "0883": true, "084": true, "0884": true, "0885": true, "0886": true, "0887": true, "0888": true, "0889": true,
	// Unknown
	"086": true,
	// Landline
	"021":  true, // Jakarta
	"022":  true, // Bandung
	"024":  true, // Semarang
	"0251": true, // Bogor
	"0252": true, // Sukabumi
	"0254": true, // Serang
	"0261": true, // Cirebon
	"0271": true, // Solo / Surakarta
	"0274": true, // Yogyakarta
	"0281": true, // Purwokerto
	"031":  true, // Surabaya
	"0341": true, // Malang
	"0351": true, // Madiun
	"0361": true, // Denpasar / Bali
	"0380": true, // Kupang
	"0411": true, // Makassar
	"0431": true, // Manado
	"061":  true, // Medan
	"0621": true, // Pematangsiantar
	"0631": true, // Kisaran
	"0641": true, // Rantauprapat
	"0711": true, // Palembang
	"0721": true, // Bandar Lampung
	"0751": true, // Padang
	"0761": true, // Pekanbaru
	"0778": true, // Batam
	"0911": true, // Balikpapan
	"0541": true, // Samarinda
	"0521": true, // Banjarmasin
	"0951": true, // Sorong
	"0967": true, // Manokwari
}

// SanitizePhoneNumber sanitizes the phone number and checks its validity.
func SanitizePhoneNumber(phone string) (string, error) {
	// Remove all unwanted characters and keep only digits & '+' signs
	phone = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r // Keep only digits and '+'
		}
		return -1 // Remove everything else
	}, phone)

	// Handle the case where phone starts with '62' (should be '0')
	if strings.HasPrefix(phone, "62") {
		phone = "0" + phone[2:] // Replace '62' with '0'
	}

	// If the phone starts with '+62', replace it with '0'
	if strings.HasPrefix(phone, "+62") {
		phone = "0" + phone[3:] // Replace '+62' with '0'
	}

	// If the phone starts with "8" (e.g., "85173207755"), prepend "0"
	if strings.HasPrefix(phone, "8") {
		phone = "0" + phone
	}

	// Ensure the phone number starts with '0'
	if !strings.HasPrefix(phone, "0") {
		return "", fmt.Errorf("invalid phone number: %v must start with '0' or '+62'", phone)
	}

	// Validate the phone number prefix
	if !IsValidPhoneNumber(phone) {
		return "", fmt.Errorf("invalid phone number: %v prefix is not valid try using indonesian number that start with '0' or '+62'", phone)
	}

	// Return the sanitized phone number without the leading '0'
	return phone[1:], nil
}

// IsValidPhoneNumber checks if the phone number has a valid prefix
func IsValidPhoneNumber(phone string) bool {
	// Check if the phone number is long enough and if the prefix is valid
	if len(phone) < 4 {
		return false
	}

	// Check the first 4 digits
	prefix := phone[:4]
	if validPhonePre[prefix] {
		return true
	}

	// Check the first 3 digits
	prefix = phone[:3]
	return validPhonePre[prefix]
}
