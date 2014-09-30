package response

import "net/http"
import "fmt"

// Continue response.
func Continue(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusContinue)
	fmt.Fprintln(w, http.StatusText(http.StatusContinue))
}

// SwitchingProtocols response.
func SwitchingProtocols(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusSwitchingProtocols)
	fmt.Fprintln(w, http.StatusText(http.StatusSwitchingProtocols))
}

// OK response.
func OK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, http.StatusText(http.StatusOK))
}

// Created response.
func Created(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, http.StatusText(http.StatusCreated))
}

// Accepted response.
func Accepted(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, http.StatusText(http.StatusAccepted))
}

// NonAuthoritativeInfo response.
func NonAuthoritativeInfo(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNonAuthoritativeInfo)
	fmt.Fprintln(w, http.StatusText(http.StatusNonAuthoritativeInfo))
}

// NoContent response.
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, http.StatusText(http.StatusNoContent))
}

// ResetContent response.
func ResetContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusResetContent)
	fmt.Fprintln(w, http.StatusText(http.StatusResetContent))
}

// PartialContent response.
func PartialContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPartialContent)
	fmt.Fprintln(w, http.StatusText(http.StatusPartialContent))
}

// MultipleChoices response.
func MultipleChoices(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMultipleChoices)
	fmt.Fprintln(w, http.StatusText(http.StatusMultipleChoices))
}

// MovedPermanently response.
func MovedPermanently(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMovedPermanently)
	fmt.Fprintln(w, http.StatusText(http.StatusMovedPermanently))
}

// Found response.
func Found(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusFound)
	fmt.Fprintln(w, http.StatusText(http.StatusFound))
}

// SeeOther response.
func SeeOther(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusSeeOther)
	fmt.Fprintln(w, http.StatusText(http.StatusSeeOther))
}

// NotModified response.
func NotModified(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotModified)
	fmt.Fprintln(w, http.StatusText(http.StatusNotModified))
}

// UseProxy response.
func UseProxy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUseProxy)
	fmt.Fprintln(w, http.StatusText(http.StatusUseProxy))
}

// TemporaryRedirect response.
func TemporaryRedirect(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprintln(w, http.StatusText(http.StatusTemporaryRedirect))
}

// BadRequest response.
func BadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, http.StatusText(http.StatusBadRequest))
}

// Unauthorized response.
func Unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, http.StatusText(http.StatusUnauthorized))
}

// PaymentRequired response.
func PaymentRequired(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPaymentRequired)
	fmt.Fprintln(w, http.StatusText(http.StatusPaymentRequired))
}

// Forbidden response.
func Forbidden(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintln(w, http.StatusText(http.StatusForbidden))
}

// NotFound response.
func NotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
}

// MethodNotAllowed response.
func MethodNotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
}

// NotAcceptable response.
func NotAcceptable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprintln(w, http.StatusText(http.StatusNotAcceptable))
}

// ProxyAuthRequired response.
func ProxyAuthRequired(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusProxyAuthRequired)
	fmt.Fprintln(w, http.StatusText(http.StatusProxyAuthRequired))
}

// RequestTimeout response.
func RequestTimeout(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestTimeout)
	fmt.Fprintln(w, http.StatusText(http.StatusRequestTimeout))
}

// Conflict response.
func Conflict(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintln(w, http.StatusText(http.StatusConflict))
}

// Gone response.
func Gone(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusGone)
	fmt.Fprintln(w, http.StatusText(http.StatusGone))
}

// LengthRequired response.
func LengthRequired(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusLengthRequired)
	fmt.Fprintln(w, http.StatusText(http.StatusLengthRequired))
}

// PreconditionFailed response.
func PreconditionFailed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusPreconditionFailed)
	fmt.Fprintln(w, http.StatusText(http.StatusPreconditionFailed))
}

// RequestEntityTooLarge response.
func RequestEntityTooLarge(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	fmt.Fprintln(w, http.StatusText(http.StatusRequestEntityTooLarge))
}

// RequestURITooLong response.
func RequestURITooLong(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestURITooLong)
	fmt.Fprintln(w, http.StatusText(http.StatusRequestURITooLong))
}

// UnsupportedMediaType response.
func UnsupportedMediaType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	fmt.Fprintln(w, http.StatusText(http.StatusUnsupportedMediaType))
}

// RequestedRangeNotSatisfiable response.
func RequestedRangeNotSatisfiable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	fmt.Fprintln(w, http.StatusText(http.StatusRequestedRangeNotSatisfiable))
}

// ExpectationFailed response.
func ExpectationFailed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusExpectationFailed)
	fmt.Fprintln(w, http.StatusText(http.StatusExpectationFailed))
}

// Teapot response.
func Teapot(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusTeapot)
	fmt.Fprintln(w, http.StatusText(http.StatusTeapot))
}

// InternalServerError response.
func InternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
}

// NotImplemented response.
func NotImplemented(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusText(http.StatusNotImplemented))
}

// BadGateway response.
func BadGateway(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintln(w, http.StatusText(http.StatusBadGateway))
}

// ServiceUnavailable response.
func ServiceUnavailable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintln(w, http.StatusText(http.StatusServiceUnavailable))
}

// GatewayTimeout response.
func GatewayTimeout(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusGatewayTimeout)
	fmt.Fprintln(w, http.StatusText(http.StatusGatewayTimeout))
}

// HTTPVersionNotSupported response.
func HTTPVersionNotSupported(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	fmt.Fprintln(w, http.StatusText(http.StatusHTTPVersionNotSupported))
}
