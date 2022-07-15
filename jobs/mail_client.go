package jobs

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/Prithvipal/sailpoint/config"
	"github.com/sirupsen/logrus"
)

func SendMail(prs []PullRequest) error {
	cfg := config.GetConfig()
	from := cfg.Mail.From
	password := cfg.Mail.Pass

	toEmailAddress := cfg.Mail.To
	to := []string{toEmailAddress}

	host := cfg.Mail.Smtp
	port := cfg.Mail.Port
	address := host + ":" + port

	body, err := getMsg(prs)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", from, password, host)

	return smtp.SendMail(address, auth, from, to, body.Bytes())

}

func getMsg(prs []PullRequest) (bytes.Buffer, error) {
	var msg bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	_, err := msg.Write([]byte(fmt.Sprintf("Subject: Pull Request status \n%s\n\n", mimeHeaders)))
	if err != nil {
		logrus.Error("error while writing subject to buffer", err)
		return msg, err
	}
	body, err := getContent(prs)
	if err != nil {
		return msg, err
	}
	_, err = msg.Write(body.Bytes())
	return msg, err
}

func getContent(prs []PullRequest) (bytes.Buffer, error) {
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
	t, err := tmpl.Parse(tmplTxt)
	if err != nil {
		logrus.Error("error while parsing mail template", err)
	}
	var body bytes.Buffer
	err = t.Execute(&body, prs)
	return body, err
}
