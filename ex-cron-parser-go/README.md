# Cron parser

Parses cron syntax and outputs the times it corresponds to.

For example, the following input argument:
```
$ ./cron-parser-go "*/15 0 1,15 * 1-5 /usr/bin/find"
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
```

## Dev and Testing

Load go dependencies with `go mod tidy`

Run tests with `go test ./...`

Quick Run `go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"`

Generate an executable with `go build`

