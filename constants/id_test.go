package constants_test

import (
	"strings"
	"testing"

	"github.com/goauthnz/pkg/constants"
	"github.com/oklog/ulid/v2"
)

func TestGenerateDataPrefixWithULID_Format(t *testing.T) {
	prefixes := []constants.DataPrefix{
		constants.Account,
		constants.EmailIdentity,
		constants.PhoneIdentity,
		constants.SSOIdentity,
		constants.Session,
		constants.ResetPassword,
		constants.VerifyEmail,
		constants.VerifyPhone,
	}
	for _, p := range prefixes {
		t.Run(p.String(), func(t *testing.T) {
			id := constants.GenerateDataPrefixWithULID(p)
			wantLen := len(p.String()) + ulid.EncodedSize
			if len(id) != wantLen {
				t.Errorf("len(%q) = %d, want %d", id, len(id), wantLen)
			}
			if !strings.HasPrefix(id, p.String()) {
				t.Errorf("%q does not start with prefix %q", id, p.String())
			}
		})
	}
}

func TestGenerateDataPrefixWithULID_IsUnique(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id := constants.GenerateDataPrefixWithULID(constants.Account)
		if seen[id] {
			t.Fatalf("duplicate ID generated: %q", id)
		}
		seen[id] = true
	}
}

func TestIsValid_AcceptsGeneratedID(t *testing.T) {
	prefixes := []constants.DataPrefix{
		constants.Account,
		constants.EmailIdentity,
		constants.Session,
	}
	for _, p := range prefixes {
		t.Run(p.String(), func(t *testing.T) {
			id := constants.GenerateDataPrefixWithULID(p)
			if !p.IsValid(id) {
				t.Errorf("IsValid(%q) = false for a freshly generated ID", id)
			}
		})
	}
}

func TestIsValid_RejectsWrongPrefix(t *testing.T) {
	id := constants.GenerateDataPrefixWithULID(constants.Account)
	if constants.Session.IsValid(id) {
		t.Errorf("Session.IsValid accepted an Account-prefixed ID: %q", id)
	}
}

func TestIsValid_RejectsTooShort(t *testing.T) {
	short := constants.Account.String() + "01J"
	if constants.Account.IsValid(short) {
		t.Errorf("IsValid accepted a too-short string: %q", short)
	}
}

func TestIsValid_RejectsTooLong(t *testing.T) {
	id := constants.GenerateDataPrefixWithULID(constants.Account)
	long := id + "X"
	if constants.Account.IsValid(long) {
		t.Errorf("IsValid accepted a too-long string: %q", long)
	}
}

func TestIsValid_RejectsEmpty(t *testing.T) {
	if constants.Account.IsValid("") {
		t.Error("IsValid accepted empty string")
	}
}

func TestIsValid_RejectsULIDWithoutPrefix(t *testing.T) {
	rawULID := ulid.Make().String()
	if constants.Account.IsValid(rawULID) {
		t.Errorf("IsValid accepted a bare ULID without prefix: %q", rawULID)
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		prefix constants.DataPrefix
		want   string
	}{
		{constants.Account, "acnt_"},
		{constants.EmailIdentity, "emli_"},
		{constants.PhoneIdentity, "phni_"},
		{constants.SSOIdentity, "ssoi_"},
		{constants.Session, "sess_"},
		{constants.ResetPassword, "rpsw_"},
		{constants.VerifyEmail, "veml_"},
		{constants.VerifyPhone, "vphn_"},
	}
	for _, tc := range cases {
		if tc.prefix.String() != tc.want {
			t.Errorf("String() = %q, want %q", tc.prefix.String(), tc.want)
		}
	}
}
