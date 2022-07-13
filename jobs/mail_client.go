package jobs

import "net/smtp"

func SendMail() {
	from := "prithvirathore.learn@gmail.com"
	password := ""

	toEmailAddress := "prithvirathore99@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	// auth := smtp.CRAMMD5Auth(from, password)
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}
