[![CircleCI](https://circleci.com/gh/moemoe89/go-unit-test-sql.svg?style=svg)](https://circleci.com/gh/moemoe89/go-unit-test-sql)
[![codecov](https://codecov.io/gh/moemoe89/go-unit-test-sql/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/go-unit-test-sql)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-unit-test-sql)](https://goreportcard.com/report/github.com/moemoe89/go-unit-test-sql)

# GO-UNIT-TEST-SQL #

Example Mock Unit Test for SQL in Golang

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ go-unit-test-sql/
  |     |
  |     +--+ main.go
  |        + repository/
  |        |
  |        +--+ repository.go
  |        |
  |        +--+ mysql
  |        |  |
  |        |  +--+ mysql.go
  |        |     + mysql_test.go
  |        |
  |        +--+ postgres
  |           |
  |           +--+ postgres.go
  |              + postgres_test.go
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup

* Setup Golang <https://golang.org>
* Setup Docker <https://www.docker.com>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> go-unit-test-sql
```

## How to Run Test
```
$ go test ./...
```

## License

MIT