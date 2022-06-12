package main

import (
	"errors"
	"fmt"
	"html/template"
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

func main() {
	server := Server{
		router: mux.NewRouter(),
		logger: initLogger(),
		db:     nil,
	}

	server.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Witaj na stronie")
	})

	tmpl := getFormTemplate()
	server.router.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
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

		server.logger.Info(fmt.Sprintf("%v", user))
		// logger.Info(user.Firstname)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	err := http.ListenAndServe(":8080", server.router)
	if err != nil {
		server.logger.Fatal(err.Error())
	}
}
