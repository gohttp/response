
# response

 HTTP response utilities.

 View the [docs](http://godoc.org/github.com/gohttp/response).

 ```go
response.Pretty = false
response.JSON(w, user)
response.XML(w, user)
response.NotFound(w)
response.Unauthorized(w)
 ```

## Provides

 - JSON responses
 - XML responses
 - Error responses
 - Status code responses
 - Toggling of "pretty" responses
