package mailer

type Client interface {
	Send() error
}
