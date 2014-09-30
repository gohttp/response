package response

import "net/http"
import "fmt"

// message returns the message string given to
// one of the status functions or falls back on
// default status text.
func message(msg []string, code int) string {
	if len(msg) > 0 {
		return msg[0]
	} else {
		return http.StatusText(code)
	}
}

// Continue response.
func Continue(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusContinue)
	fmt.Fprintln(w, message(msg, http.StatusContinue))
}

// SwitchingProtocols response.
func SwitchingProtocols(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusSwitchingProtocols)
	fmt.Fprintln(w, message(msg, http.StatusSwitchingProtocols))
}

// OK response.
func OK(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, message(msg, http.StatusOK))
}

// Created response.
func Created(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, message(msg, http.StatusCreated))
}

// Accepted response.
func Accepted(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, message(msg, http.StatusAccepted))
}

// NonAuthoritativeInfo response.
func NonAuthoritativeInfo(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNonAuthoritativeInfo)
	fmt.Fprintln(w, message(msg, http.StatusNonAuthoritativeInfo))
}

// NoContent response.
func NoContent(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, message(msg, http.StatusNoContent))
}

// ResetContent response.
func ResetContent(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusResetContent)
	fmt.Fprintln(w, message(msg, http.StatusResetContent))
}

// PartialContent response.
func PartialContent(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPartialContent)
	fmt.Fprintln(w, message(msg, http.StatusPartialContent))
}

// MultipleChoices response.
func MultipleChoices(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMultipleChoices)
	fmt.Fprintln(w, message(msg, http.StatusMultipleChoices))
}

// MovedPermanently response.
func MovedPermanently(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMovedPermanently)
	fmt.Fprintln(w, message(msg, http.StatusMovedPermanently))
}

// Found response.
func Found(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusFound)
	fmt.Fprintln(w, message(msg, http.StatusFound))
}

// SeeOther response.
func SeeOther(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusSeeOther)
	fmt.Fprintln(w, message(msg, http.StatusSeeOther))
}

// NotModified response.
func NotModified(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotModified)
	fmt.Fprintln(w, message(msg, http.StatusNotModified))
}

// UseProxy response.
func UseProxy(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUseProxy)
	fmt.Fprintln(w, message(msg, http.StatusUseProxy))
}

// TemporaryRedirect response.
func TemporaryRedirect(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprintln(w, message(msg, http.StatusTemporaryRedirect))
}

// BadRequest response.
func BadRequest(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, message(msg, http.StatusBadRequest))
}

// Unauthorized response.
func Unauthorized(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, message(msg, http.StatusUnauthorized))
}

// PaymentRequired response.
func PaymentRequired(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPaymentRequired)
	fmt.Fprintln(w, message(msg, http.StatusPaymentRequired))
}

// Forbidden response.
func Forbidden(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintln(w, message(msg, http.StatusForbidden))
}

// NotFound response.
func NotFound(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, message(msg, http.StatusNotFound))
}

// MethodNotAllowed response.
func MethodNotAllowed(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, message(msg, http.StatusMethodNotAllowed))
}

// NotAcceptable response.
func NotAcceptable(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprintln(w, message(msg, http.StatusNotAcceptable))
}

// ProxyAuthRequired response.
func ProxyAuthRequired(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusProxyAuthRequired)
	fmt.Fprintln(w, message(msg, http.StatusProxyAuthRequired))
}

// RequestTimeout response.
func RequestTimeout(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestTimeout)
	fmt.Fprintln(w, message(msg, http.StatusRequestTimeout))
}

// Conflict response.
func Conflict(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintln(w, message(msg, http.StatusConflict))
}

// Gone response.
func Gone(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusGone)
	fmt.Fprintln(w, message(msg, http.StatusGone))
}

// LengthRequired response.
func LengthRequired(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusLengthRequired)
	fmt.Fprintln(w, message(msg, http.StatusLengthRequired))
}

// PreconditionFailed response.
func PreconditionFailed(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPreconditionFailed)
	fmt.Fprintln(w, message(msg, http.StatusPreconditionFailed))
}

// RequestEntityTooLarge response.
func RequestEntityTooLarge(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	fmt.Fprintln(w, message(msg, http.StatusRequestEntityTooLarge))
}

// RequestURITooLong response.
func RequestURITooLong(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestURITooLong)
	fmt.Fprintln(w, message(msg, http.StatusRequestURITooLong))
}

// UnsupportedMediaType response.
func UnsupportedMediaType(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	fmt.Fprintln(w, message(msg, http.StatusUnsupportedMediaType))
}

// RequestedRangeNotSatisfiable response.
func RequestedRangeNotSatisfiable(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	fmt.Fprintln(w, message(msg, http.StatusRequestedRangeNotSatisfiable))
}

// ExpectationFailed response.
func ExpectationFailed(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusExpectationFailed)
	fmt.Fprintln(w, message(msg, http.StatusExpectationFailed))
}

// Teapot response.
func Teapot(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusTeapot)
	fmt.Fprintln(w, message(msg, http.StatusTeapot))
}

// InternalServerError response.
func InternalServerError(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, message(msg, http.StatusInternalServerError))
}

// NotImplemented response.
func NotImplemented(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, message(msg, http.StatusNotImplemented))
}

// BadGateway response.
func BadGateway(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintln(w, message(msg, http.StatusBadGateway))
}

// ServiceUnavailable response.
func ServiceUnavailable(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintln(w, message(msg, http.StatusServiceUnavailable))
}

// GatewayTimeout response.
func GatewayTimeout(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusGatewayTimeout)
	fmt.Fprintln(w, message(msg, http.StatusGatewayTimeout))
}

// HTTPVersionNotSupported response.
func HTTPVersionNotSupported(w http.ResponseWriter, msg ...string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	fmt.Fprintln(w, message(msg, http.StatusHTTPVersionNotSupported))
}
