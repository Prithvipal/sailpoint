package config

import (
	"flag"
	"sync"
)

var (
	cfg        *Config
	onceConfig sync.Once
)

type Config struct {
	Git  GitConfig
	Mail MailConfig
}

type GitConfig struct {
	URL   string
	Repo  string
	Owner string
}

func (cfg *GitConfig) DeclFlag() {
	flag.StringVar(&cfg.URL, "cfg.git.url", "https://api.github.com/repos/%v/%v/pulls?state=all&sort=updated&direction=desc", "Git URL to get pull request")
	flag.StringVar(&cfg.Repo, "cfg.git.repo", "test", "Git repo name from which we want to get PRs")
	flag.StringVar(&cfg.Owner, "cfg.git.owner", "Prithvipal", "Git Owner/account name where we have git repo")
}

type MailConfig struct {
	From string
	Pass string
	To   string
	Smtp string
	Port string
}

func (cfg *MailConfig) DeclFlag() {
	flag.StringVar(&cfg.From, "cfg.mail.from", "prithvirathore.learn@gmail.com", "Email ID which will send mail")
	flag.StringVar(&cfg.Pass, "cfg.mail.pass", "123", "Password of from Email id")
	flag.StringVar(&cfg.To, "cfg.mail.to", "prithvirathore99@gmail.com", "The mail id where want to send mail")
	flag.StringVar(&cfg.Smtp, "cfg.mail.smtp", "smtp.gmail.com", "The SMTP service name")
	flag.StringVar(&cfg.Port, "cfg.mail.port", "587", "The SMTP port")

}
func (cfg *Config) DeclFlag() {
	cfg.Git.DeclFlag()
	cfg.Mail.DeclFlag()
	flag.Parse()
}

func GetConfig() *Config {
	onceConfig.Do(func() {
		cfg = &Config{}
	})
	return cfg
}
