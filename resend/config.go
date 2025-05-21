package resend

// Config is the configuration for the Resend client.
type Config struct {
	// APIKey is provided by Resend.
	APIKey string `env:"RESEND_API_KEY"`
}
