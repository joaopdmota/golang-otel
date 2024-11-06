package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPService struct {
	server *http.Server
	echo   *echo.Echo
}

func NewHTTPService(port string) *HTTPService {
	e := echo.New()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	return &HTTPService{
		server: server,
		echo:   e,
	}
}

func (s *HTTPService) AddRoute(method, pattern string, handler echo.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.echo.GET(pattern, handler)
	case http.MethodPost:
		s.echo.POST(pattern, handler)
	case http.MethodPut:
		s.echo.PUT(pattern, handler)
	case http.MethodDelete:
		s.echo.DELETE(pattern, handler)
	default:
		fmt.Printf("Método %s não é suportado\n", method)
	}
}

func (s *HTTPService) Start() error {
	fmt.Printf("Servidor rodando na porta %s...\n", s.server.Addr)
	return s.echo.StartServer(s.server)
}

func (s *HTTPService) Stop() error {
	fmt.Println("Parando o servidor...")
	return s.echo.Close()
}
