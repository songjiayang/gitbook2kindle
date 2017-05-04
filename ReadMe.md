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
--> Downloading: gnuhpc/redis-all-about
--> Downloading: yeasy/docker_practice
--> Downloading: 0xax/linux-insides
--> Downloading: zhaoda/webpack
--> Downloading: frontendmasters/front-end-handbook
--> Downloaded: gnuhpc/redis-all-about
--> Downloaded: frontendmasters/front-end-handbook
--> Downloaded: yeasy/docker_practice
--> Downloaded: zhaoda/webpack
--> Downloaded: 0xax/linux-insides
--> Syncing: yeasy/docker_practice, zhaoda/webpack, 0xax/linux-insides, gnuhpc/redis-all-about, frontendmasters/front-end-handbook
--> Synced: yeasy/docker_practice, zhaoda/webpack, 0xax/linux-insides, gnuhpc/redis-all-about, frontendmasters/front-end-handbook
```
