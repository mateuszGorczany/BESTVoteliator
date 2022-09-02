package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ErrorLogInit = errors.New("cannot initalize logger")
)

type Server struct {
	router *mux.Router
	logger *zap.Logger
	db     *string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// func (s *Server) getDefaultHandler() http.Handler {
// 	return func (w http.Wr)
// }

// func (s *Server) getRestrictedHandler() http.Handler {

// }

func getFormTemplate() *template.Template {
	return template.Must(template.ParseFiles("./html/index.html"))
}

func initLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := loggerConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Errorf("%v, %v", ErrorLogInit, err).Error())
	}
	return logger
}

func jsonEncoder() *http.Handler {
	return nil
}

func logging(logger *zap.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("accepted incoming request")
		next.ServeHTTP(w, r)
		logger.Info("handled request")
	}
}

func authenticate(logger *zap.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Print(buf.String())
		if auth, ok := r.Header["Ok"]; !ok || auth[0] != "ok" {
			logger.Sugar().Warnf("unauthorized request from host: %v", r.Host)
			w.Write([]byte("Not authorized"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

type jsonI interface {
	Jsonify() []byte
}

type controller func(r *http.Request) (interface{}, error)

func toJsonResponse(service controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		toJson, _ := service(r)
		json.NewEncoder(w).Encode(toJson)
	}
}

type aa struct {
	B string
	C int
}

// func (aaa aa) Jsonify() []byte {
// 	return []byte("xxdd")
// }

func a(r *http.Request) (interface{}, error) {
	return aa{
		"ok11",
		32,
	}, nil
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Witaj na stronie")
}

func form() http.HandlerFunc {
	tmpl := getFormTemplate()
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		age, _ := strconv.Atoi(r.FormValue("age"))
		user := dto.User{
			Firstname: r.FormValue("firstname"),
			Lastname:  r.FormValue("lastname"),
			Age:       age,
		}

		log.Println("%v", user)
		// server.logger.Info(fmt.Sprintf("%v", user))
		// logger.Info(user.Firstname)

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("accepted incoming request")
		next.ServeHTTP(w, r)
		s.logger.Info("handled request")
	})
}

func (s *Server) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticate(s.logger, next.ServeHTTP)(w, r)
	})
}

type middleware func(http.HandlerFunc) http.HandlerFunc

func chain(f http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}

func authCallback(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, false)
}

func main() {
	server := Server{
		router: mux.NewRouter(),
		logger: initLogger(),
		db:     nil,
	}
	server.router.Use(server.loggingMiddleware)
	server.router.HandleFunc("/", authCallback)
	server.router.HandleFunc("/login", handleLogin)
	server.router.HandleFunc("/callback", authCallback)
	// server.router.HandleFunc("/", chain(greet, logging))
	// server.router.HandleFunc("/authenticate", jsonChain(a, logging))
	// server.router.HandleFunc("/form", chain(form(), authenticate, logging))
	// server.router.HandleFunc("/json", chain(toJsonResponse(a), authenticate, logging))
	api := server.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/subjects", greet)
	api.Use(server.authenticationMiddleware)
	err := http.ListenAndServe(":8081", server.router)
	if err != nil {
		server.logger.Fatal(err.Error())
	}
}
