package mailer

import (
	"context"

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
	// cfg.AddDefaultHeader("partner-key", "YOUR_API_KEY")

	br := brevo.NewAPIClient(cfg)

	return &BrevoMailer{
		apiKey:     apiKey,
		partnerKey: partnerKey,
		cfg:        cfg,
		client:     br,
	}
}

func (m BrevoMailer) Send() error {
	var ctx context.Context
	_, r, err := m.client.TransactionalEmailsApi.SendTransacEmail(ctx, brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Email: "rime.platform@gmail.com",
			Name:  "Rime Platform",
		},
		To:          []brevo.SendSmtpEmailTo{{Email: "manuel.mtzv816@gmail.com"}},
		Subject:     "Hello, World!",
		TextContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err != nil {
		return err
	}

	if r.StatusCode != 201 {
		return err
	}

	return nil
}
