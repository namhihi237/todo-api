# TODO LIST

## User stories

- Service not needed authentication
- User can create a item TODO
- User can edit a item TODO
- User can delete a item TODO
- User can change status of a item TODO from doing to done
- User can see all items TODO (pagination)

## Golang

- Gin framework
- GORM ORM (for mysql)

## Design system

todo_items
<code>
CREATE TABLE `todo_items` (
`id` int NOT NULL AUTO_INCREMENT,
`title` varchar(150) CHARACTER SET utf8 NOT NULL,
`status` enum('Doing','Finished') DEFAULT 'Doing',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
</code>

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

- Automatically reload server when code change

* Using air

- go install github.com/cosmtrek/air@latest

- run `air init`
- run `air`

## RUN

- Start the server: `air`
