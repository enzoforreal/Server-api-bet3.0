package auth

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/enzof/server-app-bet3.0/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := new(DbMock)
	mockDB.On("CreateUser", mock.AnythingOfType("*auth.User")).Return(nil)

	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {
		RegisterUser(c, mockDB)
	})

	userData := map[string]interface{}{
		"email":    "test@example.com",
		"password": "strongpassword123",
	}
	data, _ := json.Marshal(userData)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should return 200 for successful registration")
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Should not error when parsing response")
	value, exists := response["message"]
	assert.True(t, exists, "Response should contain 'message'")
	assert.Equal(t, "Inscription réussie!", value, "Response message should be 'Inscription réussie!'")
	mockDB.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := new(DbMock)
	hashedPassword, _ := util.HashPassword("strongpassword123")
	mockUser := &User{Email: "test@example.com", Password: hashedPassword}
	mockDB.On("GetUserByEmail", "test@example.com").Return(mockUser, nil)

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		LoginUser(c, mockDB)
	})

	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "strongpassword123",
	}
	data, _ := json.Marshal(loginData)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should return 200 for successful login")

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response["message"].(string)

	assert.Nil(t, err, "Should not error when parsing response")
	assert.True(t, exists, "Response should contain 'message'")
	assert.Equal(t, "Login successful", value, "Response message should be 'Login successful'")

	mockDB.AssertExpectations(t)
}
