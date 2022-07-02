# TODO LIST

## User stories

- Service needed authentication
- User can create a item TODO
- User can edit a item TODO
- User can delete a item TODO
- User can change status of a item TODO from doing to done
- User can see all items TODO (pagination)

## Golang

- Gin framework
- GORM ORM (for mysql)

## Design system

api_todo_items

- [post] /v1/items
- [get] /v1/items
- [get] /v1/items/:id
- [put] /v1/items/:id
- [delete] /v1/items/:id

## Process

### Install library

- go get -u github.com/gin-gonic/gin
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql

#### validator
- https://pkg.go.dev/github.com/astaxie/beego/validation
### bcrypt
- https://pkg.go.dev/golang.org/x/crypto/bcrypt

- Automatically reload server when code change

* Using air

- go install github.com/cosmtrek/air@latest

- run `air init`
- run `air`



## RUN

- Start the server: `air`
