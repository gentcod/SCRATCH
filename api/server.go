package api

import (
	"net/http"
	"time"

	db "github.com/gentcod/RSSAggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
	store *db.Store
	router *chi.Mux
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	go startScraping(store.Queries, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", server.handlerCreateUser)
	v1Router.Get("/users", server.middlewareAuth(server.handlerGetUser))

	v1Router.Post("/feeds", server.middlewareAuth(server.handlerCreateFeed))
	v1Router.Get("/feeds", server.handlerGetFeeds)

	v1Router.Post("/feed_follows", server.middlewareAuth(server.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", server.middlewareAuth(server.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", server.middlewareAuth(server.handlerDeleteFeedFollow))

	v1Router.Get("/posts", server.middlewareAuth(server.handlerGetPostsForUser))

	router.Mount("/v1", v1Router)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	srv := &http.Server{
		Handler: server.router,
		Addr: address,
	}
	return srv.ListenAndServe()
}