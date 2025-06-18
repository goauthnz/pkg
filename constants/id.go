package constants

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

// DataPrefix is a 4-character prefix for the data. Using this prefix, we can easily identify the type of the data.
// Using it with ULID, we can generate a unique ID for the data.
type DataPrefix string

const (
	// account related prefixes
	Account DataPrefix = "acnt_"

	// identity related prefixes
	EmailIdentity DataPrefix = "emli_"
	PhoneIdentity DataPrefix = "phni_"
	SSOIdentity   DataPrefix = "ssoi_"

	// session related prefixes
	Session DataPrefix = "sess_"

	// password and verification related prefixes
	ResetPassword DataPrefix = "rpsw_"
	VerifyEmail   DataPrefix = "veml_"
	VerifyPhone   DataPrefix = "vphn_"
)

// String returns the string representation of the DataPrefix.
// It returns the string representation of the DataPrefix.
func (dp DataPrefix) String() string {
	return string(dp)
}

// IsValid checks if the string is a valid ULID with the given prefix.
// It returns true if the string is a valid ULID with the given prefix, false otherwise.
func (dp DataPrefix) IsValid(s string) bool {
	return strings.HasPrefix(s, string(dp)) && len(s) == len(string(dp))+ulid.EncodedSize
}

// GenerateDataPrefixWithULID generates a ULID with the given prefix.
// It returns a string that is a concatenation of the prefix and the ULID.
func GenerateDataPrefixWithULID[T DataPrefix](prefixType T) string {
	return string(prefixType) + ulid.Make().String()
}
