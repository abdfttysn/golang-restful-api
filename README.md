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

## Create Todo
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

## Fetch All Todo
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

## Fetch a Single Todo
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
## Update Todo
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
## Delete Todo
```http
DELETE /api/v1/todos/:id
```
### Response
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `message` | `string` | |
| `status` | `integer` | 200 - DELETED |