# gendao
generate dao code from database

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hapoon/gendao/master/LICENSE)

gendao is a tool that generate dao(data access object) code from database for Go.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

`go get gopkg.in/hapoon/gendao.v1`

## Requirement

- [gorp.v2](https://github.com/go-gorp/gorp)
- [squirrel](https://github.com/Masterminds/squirrel)

## Usage

`gendao`

```
-h, -host arg 
    Set host name to access database.
-p, -port arg
    Set port number to access database.
-u, -user arg
    Set user name to access database.
-pass, -password arg
    Set password to access database.
-n, -name arg
    Set database name to access database.
-verbose
    Enable verbose.
-o, -output arg
    Save the generated dao in output. Default is current directory.
-e, -exclude arg
    Exclude files. To specify multiple files, files are separated by comma.
```

## License

[MIT License](LICENSE)
