package helpers

import "fmt"

func SendEmailWarning(email string, newIP string) {
	fmt.Printf("Warning: suspicious activity detected for user %s. New IP: %s\n", email, newIP)
}
