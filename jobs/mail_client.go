package jobs

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendMail(prs []PullRequest) {
	from := "prithvirathore.learn@gmail.com"
	password := ""

	toEmailAddress := "prithvirathore99@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// subject := "Subject: This is the subject of the mail\n"
	// body := "This is the body of the mail"
	// message := []byte(subject + body)
	body := getMsg(prs)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, body.Bytes())
	if err != nil {
		panic(err)
	}
}

func getMsg(prs []PullRequest) bytes.Buffer {
	var msg bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg.Write([]byte(fmt.Sprintf("Subject: Pull Request status \n%s\n\n", mimeHeaders)))
	body := getContent(prs)
	msg.Write(body.Bytes())
	return msg
}

func getContent(prs []PullRequest) bytes.Buffer {
	tmplTxt := `
	<!DOCTYPE html>
	<html>
	<body>
		<p>Hi All,</p>
		<p>Please find below pull request status:</p>
		
    	<table>
  			<tr>
    			<th>PR</th>
    			<th>State</th>
    			<th>Created At</th>
				<th>Updated At</th>
  			</tr>
			  {{range $val := .}}
			  <tr>
    			<td>{{$val.URL}}</td>
    			<td>{{$val.State}}</td>
    			<td>{{$val.CreatedAt}}</td>
				<td>{{$val.UpdatedAt}}</td>
  			</tr>
			  {{end}}
		</table>
	</body>
	</html>
	`

	tmpl := template.New("mail_tmpl")
	t, _ := tmpl.Parse(tmplTxt)
	var body bytes.Buffer
	t.Execute(&body, prs)

	return body
}
