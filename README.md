# app-backend

## Required

- Go
- Firebase 
- PostgreSQL

## Dev
First you need to modify the configuration in the file config/debug.yaml

```yaml
db:
  host: 127.0.0.1
  port: 54320
  user: test
  dbname: edu
  password: test123
  maxcons: 10
  maxidlecons: 2
firebase:
  credentials: /Users/XXXXX/fir-learn-f939d-firebase-adminsdk-hdw4p-8484139d16.json
```
Then set the environment

```bash
export GIN_MODE=debug
```

## RUN
The bee run command will supervise the file system of any Go project using inotify. The results will autocompile and display immediately after any modification in the project folders.

```bash
$ go get github.com/beego/bee
$ bee run

```

