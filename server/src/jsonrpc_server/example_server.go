package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

type Args struct {
	Who string
}

type Reply struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *Args, reply *Reply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc", s)
	http.ListenAndServe("0.0.0.0:6061", nil)
}
