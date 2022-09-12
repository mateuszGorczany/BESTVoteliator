package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/services"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"go.uber.org/zap"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.status = code
	w.ResponseWriter.WriteHeader(code)
	w.wroteHeader = true

	return
}

func logging(next http.Handler) http.Handler {
	loggingMiddleware := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var rw *responseWriter = NewResponseWriter(w)
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				common.Logger.Error(
					fmt.Errorf("%v", err).Error(),
					zap.String("requestID", getRequestID(r).toString()),
					zap.String("host", r.Host),
					zap.String("endpoint", r.RequestURI),
					zap.String("method", r.Method),
					zap.Int("response", http.StatusInternalServerError),
					zap.Stack("trace"),
				)
			}
		}()

		next.ServeHTTP(rw, r)

		common.Logger.Info(
			"handled request",
			zap.String("requestID", getRequestID(r).toString()),
			zap.String("host", r.Host),
			zap.String("endpoint", r.RequestURI),
			zap.String("method", r.Method),
			zap.Int("response", rw.Status()),
			zap.Int64("time[ns]", time.Since(start).Nanoseconds()),
		)
	}

	return http.HandlerFunc(loggingMiddleware)
}

func allowCors(next http.Handler) http.Handler {
	corsAddingMiddleware := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(corsAddingMiddleware)
}

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		w.Header().Set("X-Request-Id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	auth := func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			respondWithError(w, http.StatusInternalServerError, errors.New("No Authorization header containing Google JWT Token"))
			return
		}
		jwtToken := strings.Replace(authorization, "Bearer ", "", 1)
		claims, err := services.ValidateGoogleJWT(jwtToken)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Errorf("Invalid google Auth, error: %s", err))
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(auth)
}

func getRequestID(r *http.Request) requestID {
	return requestID(r.Context().Value(requestIDKey).(string))
}

func getClaims(r *http.Request) datastruct.GoogleClaims {
	return r.Context().Value(claimsKey).(datastruct.GoogleClaims)
}
