package jobs

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/Prithvipal/sailpoint/config"
)

func SendMail(prs []PullRequest) {
	cfg := config.GetConfig()
	from := cfg.Mail.From
	password := cfg.Mail.Pass

	toEmailAddress := cfg.Mail.To
	to := []string{toEmailAddress}

	host := cfg.Mail.Smtp
	port := cfg.Mail.Port
	address := host + ":" + port

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
