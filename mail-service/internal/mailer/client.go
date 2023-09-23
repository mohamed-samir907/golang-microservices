package mailer

import (
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func newClient(m *Mail) (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	return server.Connect()
}

func sendMessage(client *mail.SMTPClient, msg *Message, formattedMessage, plainMessage string) error {
	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject).
		SetBody(mail.TextPlain, plainMessage).
		AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, file := range msg.Attachments {
			email.AddAttachment(file)
		}
	}

	return email.Send(client)
}

func getEncryption(enc string) mail.Encryption {
	switch enc {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
