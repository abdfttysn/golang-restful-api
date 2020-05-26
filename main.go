package main

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	jwt "github.com/dgrijalva/jwt-go"
)

var db *gorm.DB

type (
	// todoModel describes a todoModel type
	todoModel struct {
		gorm.Model
		Title string `json:"title"`
		Completed int `json:"completed"`
	}
	// transformedTodo represents a formatted todo
	transformedTodo struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}
	// 
	Credential struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

var access = Credential{
	Username: "root",
	Password: "root123",
}

func init() {
	// open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:root123@/go_api_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema
	db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	fmt.Println("Hi gaes!")
	// })


	router.POST("/api/v1/login", loginHandler)

	v1 := router.Group("/api/v1/todos") 
	{
		v1.POST("/", auth, createTodo)
		v1.GET("/", auth, fetchAllTodo)
		v1.GET("/:id", auth, fetchSingleTodo)
		v1.PUT("/:id", auth, updateTodo)
		v1.DELETE("/:id", auth, deleteTodo)
	}

	router.Run()
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != access.Username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != access.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
	
// createTodo add new todo
func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "No todo found!",
		})
		return
	}

	// transform the todos for building a good rsponse
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, transformedTodo{
			ID: item.ID,
			Title: item.Title,
			Completed: completed,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": _todos,
	})
}

// fetchingSingleTodo fetch a single todo
func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID :=  c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "No todo found!",
		})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := transformedTodo{
		ID: todo.ID,
		Title: todo.Title,
		Completed: completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": _todo,
	})
}

// updateTodo update a todo
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "No todo found!",
		})
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Todo updated successfully!",
	})
}

// deteleTodo remove a todo
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "No todo found!",
		})
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Todo deleted successfully!",
	})
}