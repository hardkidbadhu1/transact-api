package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transact-api/constants"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	h := NewController()

	h.Healthz(c)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	}{
		Status:  "ok",
		Version: constants.AppVersion,
	}

	jsonStr, _ := json.Marshal(expectedBody)
	assert.JSONEq(t, string(jsonStr), w.Body.String())
}