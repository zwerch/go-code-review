package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"coupon_service/internal/service/entity"

	"github.com/gin-gonic/gin"
)

type Service interface { // TODO: should be declared by service package
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) (*entity.Coupon, error)
	GetCoupons([]string) ([]entity.Coupon, error)
}

type Config struct {
	Host string
	Port int
}

type API struct {
	srv *http.Server
	mux *gin.Engine
	svc Service
	cfg Config
}

// why generics here?
func New(cfg Config, svc Service) API {
	gin.SetMode(gin.ReleaseMode) // could make use of different environments
	r := new(gin.Engine)         // unnecessary line
	r = gin.New()                // variable naming, also should probably use gin.Default()
	r.Use(gin.Recovery())

	return API{
		mux: r,
		cfg: cfg,
		svc: svc,
	}.withServer().withRoutes()
}

func (a API) withServer() API {

	// why?
	// ch := make(chan API)
	// go func() {
	a.srv = &http.Server{
		// Addr:    fmt.Sprintf(":%d", a.cfg.Port), // missing host
		Addr:    fmt.Sprintf(":%d", 8080), // missing host
		Handler: a.mux,
	}
	// ch <- a
	// }()

	return a
}

func (a API) withRoutes() API { // function was unused before
	apiGroup := a.mux.Group("/api") // should be versioned (v1)
	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)

	return a
}

func (a API) Start() { // should return an error and simply wrap
	if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (a API) Close() {
	// why not use time.Sleep? also this doubles the "sleep" time.
	<-time.After(5 * time.Second)

	// timeout could be a constant
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
