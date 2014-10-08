package response

import (
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError(t *testing.T) {
	res := httptest.NewRecorder()
	Error(res, http.StatusBadRequest)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "Bad Request\n", string(res.Body.Bytes()))
	assert.Equal(t, ContentText, res.HeaderMap[ContentType][0])
}
