package jobs

func Start() {
	prs := GetPullRequests()
	// fmt.Printf("%+v", prs)
	SendMail(prs)
}
