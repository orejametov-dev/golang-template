package handler

import "github.com/gin-gonic/gin"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Другие поля пользователя
}

func GetUser(c *gin.Context) {
	// Ваш код для получения пользователя
	// Здесь я просто создам фиктивного пользователя для примера
	user := User{
		ID:   123,
		Name: "John Doe",
	}

	// Отправляем JSON-ответ
	c.JSON(200, gin.H{"user": user})
}
