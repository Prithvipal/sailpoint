package jobs

func Start() {
	prs := GetPullRequests()
	SendMail(prs)
}
