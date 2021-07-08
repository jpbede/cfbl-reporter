package cfbl

import (
	"bytes"
	"fmt"
	gomail "github.com/go-mail/mail"
	"io"
	"net/mail"
	"strings"
)

func ComposeARFReport(originMailBytes *[]byte, reportFullMail bool) ([]byte, error) {
	envelopMsg := gomail.NewMessage()
	envelopMsg.SetCustomMultipartType("report")
	envelopMsg.SetAddressHeader("From", "fbl@example.com", "FBL Sender")

	msg, err := mail.ReadMessage(bytes.NewReader(*originMailBytes))
	if err != nil {
		return nil, err
	}

	// add a plain text part for email clients
	rp := strings.Trim(msg.Header.Get("Return-Path"), "<>")
	envelopMsg.AddAlternative("text/plain",
		fmt.Sprintf("This is an email abuse report for an email message from %s on %s", rp, msg.Header.Get("Date")))

	// add the report itself
	envelopMsg.AddAlternative("message/feedback-report", getReportFields(msg))

	// add the headers of the original mail, or if given full original mail
	if reportFullMail {
		envelopMsg.AddAlternativeWriter("text/rfc822", func(writer io.Writer) error {
			_, err := writer.Write(*originMailBytes)
			return err
		})
	} else {
		envelopMsg.AddAlternative("text/rfc822-headers", getHeaderForReport(msg))
	}

	// now write the complete report to buffer
	msgBuffer := new(bytes.Buffer)
	if _, err := envelopMsg.WriteTo(msgBuffer); err != nil {
		return nil, err
	}

	return msgBuffer.Bytes(), nil
}

func getReportFields(msg *mail.Message) string {
	var fields string

	fields += "Feedback-Type: abuse\r\n"
	fields += "User-Agent: CFBL-REPORTER/0.1\r\n"
	fields += "Version: 0.1\r\n"
	fields += fmt.Sprintf("Original-Mail-From: %s\r\n", strings.Trim(msg.Header.Get("Return-Path"), "<>"))
	fields += fmt.Sprintf("Arrival-Date: %s\r\n", msg.Header.Get("Date"))

	return fields
}

func getHeaderForReport(msg *mail.Message) string {
	var fields string

	fields += fmt.Sprintf("Message-ID: %s\r\n", strings.Trim(msg.Header.Get("Message-ID"), "<>"))
	if cfblFeedbackID := msg.Header.Get("CFBL-Feedback-ID"); cfblFeedbackID != "" {
		fields += fmt.Sprintf("CFBL-Feedback-ID: %s\r\n", cfblFeedbackID)
	}

	return fields
}
