package app

import (
	"net/http"
	"wsw/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func newRouter() *chi.Mux {
	gqlHandler := createGQLHandler()

	router := chi.NewRouter()
	router.Use(middleware.Logger,
		middleware.Recoverer,
		middleware.RealIP,
		cors.New(cors.Options{
			AllowedOrigins:     []string{"*"},
			AllowCredentials:   true,
			AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
			OptionsPassthrough: true,
			Debug:              true,
		}).
			Handler)

	router.Options("/graphql", gqlHandler)
	router.Post("/graphql", gqlHandler)
	return router
}

func createGQLHandler() http.HandlerFunc {
	scheme := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	handler := handler.NewDefaultServer(scheme)
	return func(w http.ResponseWriter, r *http.Request) { handler.ServeHTTP(w, r) }
}
