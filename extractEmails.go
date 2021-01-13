package main

import (
	"github.com/mcnijman/go-emailaddress"
)

func extractEmail(file []byte) (extractEmails []string) {
	for _, email := range emailaddress.Find(file, false) {
		emailStringfied := email.String()

		extractEmails = append(extractEmails, emailStringfied)
	}
	return
}
