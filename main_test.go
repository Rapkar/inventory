package main

import (
	"encoding/json"
	"fmt"
	auth "inventory/App/Auth"
	boot "inventory/App/Boot"
	model "inventory/App/Model"
	"inventory/App/Utility"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// Mock برای تابع GetAllPaymentsByAttribiute
type MockPaymentModel struct {
	mock.Mock
}

func (m *MockPaymentModel) GetAllPaymentsByAttribiute(term string) []string {
	args := m.Called(term)
	return args.Get(0).([]string)
}
func TestUniqueId() {
	uniqueString := Utility.MakeRandValue()
	if !model.CheckExportNumberFound(uniqueString) {
		uniqueString = Utility.MakeRandValue()
	} else {
		return
	}
}
func TestLoginEndpoint(t *testing.T) {
	// Setup router
	r := gin.Default()
	v1 := r.Group("/auth")
	{
		v1.POST("/login", func(c *gin.Context) {
			var login boot.Login
			login.Email = c.PostForm("email")
			login.Password = c.PostForm("pass")
			authorized, _, _ := auth.CheckAuth(login)

			if authorized {
				c.JSON(http.StatusOK, gin.H{"authenticated": true})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"authenticated": false})
			}
		})
	}

	// Test cases
	tests := []struct {
		name       string
		email      string
		password   string
		wantAuth   bool
		wantStatus int
	}{
		{"Valid credentials", "hosseinbidar7@gmail.com", "0000", true, http.StatusOK},
		{"Invalid credentials", "invalid@example.com", "wrongpass", false, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(fmt.Sprintf("email=%s&pass=%s", tt.email, tt.password)))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Record response
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}

			// Parse response
			var response map[string]bool
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("could not parse response: %v", err)
			}

			// Check authentication result
			if response["authenticated"] != tt.wantAuth {
				t.Errorf("got auth %v, want %v", response["authenticated"], tt.wantAuth)
			}
		})
	}

}
