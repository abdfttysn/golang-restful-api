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

## CREATE TODO
```http
POST /api/v1/todos
```
### Request
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `title` | `string` | **Required**. Your Title |
| `completed` | `boolean` |  |
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `message` | `string` | |
| `resourceId` | `integer` | Returned ID |
| `status` | `integer` | 201 - Created |

## FETCH ALL TODO
```http
GET /api/v1/todos
```
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `integer` | |
| `title` | `string` | |
| `completed` | `boolean` | |
| `status` | `integer` | |

## FETCH A SINGLE TODO
```http
GET /api/v1/todos/:id
```
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `integer` | |
| `title` | `string` | |
| `completed` | `boolean` | |
| `status` | `integer` | |
## UPDATE A TODO
```http
PUT /api/v1/todos/:id
```
### Request
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `title` | `string` | **Required**. Your Title |
| `completed` | `boolean` |  |
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `message` | `string` | |
| `status` | `integer` | 200 - Updated |
## DELETE A TODO
```http
DELETE /api/v1/todos/:id
```
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `message` | `string` | |
| `status` | `integer` | 200 - DELETED |