package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	// Метод POST используется для отправки данных на сервер
	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)

	// с помощью GET модно получить информацию
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)

	// если есть GET значит что кто то хочеть получить имеющие задачи
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTasks)

	// Handler для всех невыолненных задач.        ключ      значение
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetAllUncompletedTasks)

	// кто то хочет изменить ресурс а значитьизменить какую-то задачу с помощью PATCH
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompleteTask)

	// кто то хочет удалить задачу
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	// использовал return чтобы не использовать проверку if err != nil { fmt.Println("HTTP server error", err)}
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return nil
	}

	return nil
}