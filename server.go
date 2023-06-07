package main

import (
	"fmt"
	"log"
	"os"

	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"strings"

	"github.com/niuee/hrGraphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt/v5"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("SV_PORT")
	if port == "" {
		port = defaultPort
	}
	router := gin.Default()
	// router.Use(VerifyJWT)
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/hrgraphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func VerifyJWT(c *gin.Context) {
	if authHeader, found := c.Request.Header["Authorization"]; !found {
		c.Error(fmt.Errorf("Authorization header not found"))
		c.String(401, "Unauthorized")
		c.Abort()
		return
	} else {

		jwt_token := strings.Split(authHeader[0], " ")
		log.Println(jwt_token[1])
		token, err := jwt.Parse(jwt_token[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("your_jwt_secret"), nil
		})
		if err != nil {
			c.Error(err)
		}
		log.Println(token.Claims)
		c.Next()
	}

}
