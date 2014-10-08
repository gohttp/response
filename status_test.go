package response

import "github.com/bmizerany/assert"
import "net/http/httptest"
import "testing"

func TestStatusFunctions(t *testing.T) {
	res := httptest.NewRecorder()
	NotFound(res)
	assert.Equal(t, 404, res.Code)
	assert.Equal(t, "Not Found\n", string(res.Body.Bytes()))
	assert.Equal(t, ContentText, res.HeaderMap[ContentType][0])
}

func TestStatusFunctionsMessage(t *testing.T) {
	res := httptest.NewRecorder()
	NotFound(res, "can't find that")
	assert.Equal(t, 404, res.Code)
	assert.Equal(t, "can't find that\n", string(res.Body.Bytes()))
	assert.Equal(t, ContentText, res.HeaderMap[ContentType][0])
}

func TestStatusFunctionsJSON(t *testing.T) {
	res := httptest.NewRecorder()
	Unauthorized(res, map[string]string{"error": "token_expired", "message": "Token expired!"})
	assert.Equal(t, 401, res.Code)
	assert.Equal(t, `{"error":"token_expired","message":"Token expired!"}`, string(res.Body.Bytes()))
	assert.Equal(t, ContentJSON, res.HeaderMap[ContentType][0])
}
