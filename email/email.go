package email

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// ValidateFormat checks whether the email address is valid in format
func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("invalid format")
	}
	return nil
}

// ValidateReachability checks whether the email address is reachable
func ValidateReachability(email string, timeout time.Duration) error {
	i := strings.LastIndexByte(email, '@')
	host := email[i+1:]

	// The following operation will have a timeout.
	deadline := time.After(timeout)

	c := make(chan (error))
	go func() {
		mx, err := net.LookupMX(host)
		if err != nil {
			c <- errors.New("unresolvable host")
			return
		}
		fmt.Println("host", mx[0].Host, "reachable")
		// client, err := smtp.Dial(fmt.Sprintf("%s:%d", mx[0].Host, 25))
		// if err != nil {
		// 	fmt.Println("host", mx[0].Host, "dial failed")
		// 	c <- err
		// 	return
		// }
		// fmt.Println("host", mx[0].Host, "dial done")
		// defer client.Close()
		// err = client.Hello("checkmail.me")
		// if err != nil {
		// 	c <- err
		// 	fmt.Println("host", mx[0].Host, "hello not sent")
		// 	return
		// }
		// fmt.Println("host", mx[0].Host, "hello sent")
		// err = client.Mail("wumuxian1988@gmail.com")
		// if err != nil {
		// 	c <- err
		// 	fmt.Println("host", mx[0].Host, "mail command not successful")
		// 	return
		// }
		// fmt.Println("host", mx[0].Host, "mail command successful")
		// err = client.Rcpt(email)
		// if err != nil {
		// 	c <- err
		// 	fmt.Println("host", mx[0].Host, "recpt not done")
		// 	return
		// }
		// fmt.Println("host", mx[0].Host, "recpt done")
		c <- nil
	}()

	select {
	case <-deadline:
		return errors.New("check reachiability timeout")
	case err := <-c:
		return err
	}
}

// Send send emails
// to - the receiver's email
// subject - the subject
// content - the email content body
// authName - SMTP auth name
// authPass - SMTP auth password
// authAddr - SMTP auth address
// senderName - The sender's name
// senderAddr - The sender's email address
func SendSMTP(to, subject, content, authName, authPass, authAddr, senderName, senderAddr string) (err error) {

	err = ValidateFormat(to)
	if err != nil {
		return
	}

	err = ValidateReachability(to, time.Second*3)
	if err != nil {
		return
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", authName, authPass, authAddr)

	header := make(map[string]string)

	fromAddr := mail.Address{
		Name:    senderName,
		Address: senderAddr,
	}

	toAddr := mail.Address{
		Name:    "",
		Address: to,
	}

	header["From"] = fromAddr.String()
	header["To"] = toAddr.String()
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(content))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(authAddr+":25", auth, fromAddr.Address, []string{to}, []byte(message))

	return
}
