package mailer

import (
	brevo "github.com/getbrevo/brevo-go/lib"
)

type BrevoMailer struct {
	apiKey     string
	partnerKey string
	cfg        *brevo.Configuration
	client     *brevo.APIClient
}

func NewBrevo(apiKey, partnerKey string) *BrevoMailer {
	cfg := brevo.NewConfiguration()

	cfg.AddDefaultHeader("api-key", "YOUR_API_KEY")
	cfg.AddDefaultHeader("partner-key", "YOUR_API_KEY")

	br := brevo.NewAPIClient(cfg)

	return &BrevoMailer{
		apiKey:     apiKey,
		partnerKey: partnerKey,
		cfg:        cfg,
		client:     br,
	}
}

func (br BrevoMailer) Send() error {
	return nil
}
