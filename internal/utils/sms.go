package utils

import "fmt"

func SendSMS(to, message string) error {
	fmt.Printf("[SMS] To: %s | Message: %s\n", to, message)
	// TODO: Integrate Twilio
	return nil
}