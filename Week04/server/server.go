package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	newServer http.Server
}
func serverWork(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "The server is working")
}

func New(addr string) (*Server, error){
	return &Server{http.Server{Addr:addr}}, nil
}

func (s *Server)OnStart(ctx context.Context) error{

	http.HandleFunc("/", serverWork)
	//服务正常会阻塞
	err := s.newServer.ListenAndServe()
	if err !=nil{
		return err
	}
	return nil
}


func (s *Server) OnStop(ctx context.Context) error{
	err := s.newServer.Shutdown(ctx)
	if err != nil{
		return err
	}
	return nil
}