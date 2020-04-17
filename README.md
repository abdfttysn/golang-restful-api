# Build RESTful API service in golang using gin-gonic framework

I’m going to use golang simplest/fastest framework gin-gonic and a beautiful ORM gorm for our database work. To install these packages go to your workspace $GOPATH/src and run these command below:
```
$ go get gopkg.in/gin-gonic/gin.v1
$ go get -u github.com/jinzhu/gorm
$ go get github.com/go-sql-driver/mysql
```
In generic crud application we need the API’s as follows:
1. POST todos/
2. GET todos/
3. GET todos/{id}
4. PUT todos/{id}
5. DELETE todos/{id}
