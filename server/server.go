package server

import (
	"github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
	"../database"
	"runtime"
	"log"
	c "../controllers"
)


type Server struct {
	server *fasthttp.Server
	router *routing.Router
}


func NewServer() *Server {
	return &Server{server: &fasthttp.Server{}, router: routing.New()}
}


func (myServer *Server) Run() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	myServer.setHandlers()
	//log.SetFlags(log.Llongfile | log.Ltime)
	//log.SetOutput(ioutil.Discard)
	err := database.InitDataBase()
	if err != nil {
		log.Fatal(err)
	}
	myServer.server.Handler = myServer.router.HandleRequest
	//svc.server.MaxConnsPerIP = 10000
	//svc.server.Concurrency = 10000
	return myServer.server.ListenAndServe(":5000")
}

func (myServer *Server) setHandlers() {

	//:5000/api/forum/...
	myServer.router.Post("/api/forum/create",  c.CreateForum) //createForum
	//myServer.router.Post("/api/forum/<slug>/create", forumHandler) //createThread
	//myServer.router.Get("/api/forum/<slug>/details", forumHandler) //GetForumDetails
	//myServer.router.Get("/api/forum/<slug>/threads", forumHandler) //getThread
	//myServer.router.Get("/api/forum/<slug>/users", forumHandler) //GetForumUsers
	//
	////:5000/api/post/...
	//myServer.router.Post("/api/post/<id>/details", forumHandler) //updatePosts
	//myServer.router.Get("/api/post/<id>/details", forumHandler) //getPost
	//
	////:5000/api/thread/...
	//myServer.router.Post("/api/thread/<slug_or_id>/create", forumHandler) //createPosts
	//myServer.router.Post("/api/thread/<slug_or_id>/details", forumHandler) //updateThread
	//myServer.router.Post("//api/thread/<slug_or_id>/vote", forumHandler) //voteForThread
	//myServer.router.Get("/api/thread/<slug_or_id>/details", forumHandler) //getThread
	//myServer.router.Get("/api/thread/<slug_or_id>/posts", forumHandler) //getPosts
	//
	////:5000/api/user/...
	myServer.router.Post("/api/user/<nickname>/create", c.CreateUser) //createUser
	//myServer.router.Post("/api/user/<nickname>/profile", forumHandler) //updateUser
	myServer.router.Get("/api/user/<nickname>/profile", c.GetUser) //getUser
	//
	////:5000/api/service/...
	//myServer.router.Post("/api/service/clear", forumHandler) //clear db
	//myServer.router.Get("/api/service/status", forumHandler) //getCount of tables

}