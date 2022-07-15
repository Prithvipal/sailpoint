# sailpoint

## Run

```
make run <from-email-password>
```

## Args

- `cfg.mail.from`: Email ID which will send mail. **Default:** prithvirathore.learn@gmail.com
- `cfg.mail.pass`: Password of from Email id. **Default:** 123
- `cfg.mail.to` The mail id where want to send mail. **Default:** prithvirathore99@gmail.com
- `cfg.mail.smtp` The SMTP service name. **Default:** smtp.gmail.com
- `cfg.mail.port` The SMTP port. **Default:** 587
- `cfg.git.url` Git URL to get pull request. **Default:** https://api.github.com/repos/%v/%v/pulls?state=all&sort=updated&direction=desc
- `cfg.git.repo` Git repo name from which we want to get PRs. **Default:** test
- `cfg.git.owner`Git Owner/account name where we have git repo. **Default:** Prithvipal
## References

**1. List API Doc:** https://docs.github.com/en/rest/pulls/pulls#list-pull-requests
