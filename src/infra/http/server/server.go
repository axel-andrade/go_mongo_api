package server

import (
	"crypto/tls"
	"go_mongo_api/src/infra/bootstrap"
	"go_mongo_api/src/infra/http/middlewares"
	"go_mongo_api/src/infra/http/routes"
	"go_mongo_api/src/shared/constants"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *http.Server
}

// NewServer cria e retorna uma nova instância do servidor Gin com configurações padrão de middleware.
func NewServer(port string) Server {
	r := gin.New()

	// Adiciona middlewares padrão
	// gzip: Comprime as respostas HTTP com GZIP.
	// cors: Define as configurações padrão de CORS para permitir todas as solicitações de origem cruzada.
	// requestid: Gera um ID de solicitação exclusivo para cada solicitação.
	// definições de cabeçalho de segurança padrão para proteger contra ataques comuns.
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.Default())
	r.Use(requestid.New())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	})
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
	})
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Frame-Options", "DENY")
	})

	// Adiciona o middleware de controle de acesso
	r.Use(middlewares.Cors())

	// Configura o cache das respostas
	r.Use(middlewares.Cache(time.Minute))

	// Configuração do certificado SSL e arquivos de chave
	certFile := "certs/api.crt"
	keyFile := "certs/api.key"

	// Cria a configuração do TLS usando os arquivos de certificado e chave
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{},
	}

	// Carrega o arquivo de certificado e chave
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	// Adiciona a chave ao TLSConfig
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

	if os.Getenv("ENV") == constants.PROD_ENV {
		gin.SetMode(gin.ReleaseMode)
	}

	// Cria um novo servidor HTTP com a configuração TLS e o roteador Gin como handler.
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		TLSConfig:    tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return Server{
		port:   port,
		server: srv,
	}
}

func (s *Server) AddRoutes(dependencies *bootstrap.Dependencies) {
	router := routes.ConfigRoutes(s.server.Handler.(*gin.Engine), dependencies)
	router.SetTrustedProxies([]string{"127.0.0.1"})
}

func (s *Server) Run() {
	log.Printf("Server starting on port %s", s.port)

	if err := s.server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func (s *Server) Shutdown() {
	if s.server == nil {
		return
	}

	if err := s.server.Shutdown(nil); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	}

	log.Println("Server shutdown completed")
}
