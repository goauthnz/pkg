package constants

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

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

	ResetPassword DataPrefix = "rpsw_"
	VerifyEmail   DataPrefix = "veml_"
	VerifyPhone   DataPrefix = "vphn_"

	// user related prefixes
	User DataPrefix = "user_"

	// space related prefixes
	Space DataPrefix = "spac_"
	Tile  DataPrefix = "tile_"
)

func (dp DataPrefix) String() string {
	return string(dp)
}

func (dp DataPrefix) IsValid(s string) bool {
	return strings.HasPrefix(s, string(dp)) && len(s) == len(string(dp))+ulid.EncodedSize
}

func GenerateDataPrefixWithULID[T DataPrefix](prefixType T) string {
	return string(prefixType) + ulid.Make().String()
}
