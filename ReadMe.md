### gitbook2kindle

A driver help to sync your stared GitBook with kindle.


#### Config App

run `gitbook2kindle -h` to see all config options

```
-config.gitbook string
    config gitbook cookie
-config.kindle string
    config kindle account
-config.smtp.email string
    config send smtp email
-config.smtp.host string
    config send smtp server host
-config.smtp.password string
    config send smtp email
-config.smtp.server string
    config send smtp server address

```

Run `gitbook2kindle -config` to see current config

#### Run Command

Run `gitbook2kindle -run` to start app and then you can see the logs,

```
start download book:  gnuhpc/redis-all-about
start download book:  zhaoda/webpack
start download book:  0xax/linux-insides
start download book:  yeasy/docker_practice
start download book:  frontendmasters/front-end-handbook
end download book:  zhaoda/webpack
end download book:  frontendmasters/front-end-handbook
end download book:  gnuhpc/redis-all-about
end download book:  yeasy/docker_practice
end download book:  0xax/linux-insides
start send 5 books to kindle: songjiayang@qiniu.com
end send 5 books to kindle: songjiayang@qiniu.com
```
