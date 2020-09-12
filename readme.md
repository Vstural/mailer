# mailer 

a mail tool for send mail 

```sh
mailer target@mail.com 'title' 'content'
```

default config path
```
HOMEDIR/.mailer_conf
```

config example
```json
{
    "server": "smtp.mail.com",
    "port": 587,
    "user_name": "your-mail@mail.com",
    "password": "your_pw"
}
```

[WIP]flags
```sh
-c path of config 
```