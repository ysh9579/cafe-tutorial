package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hello-cafe/middleware"

	"hello-cafe/handler"
	"hello-cafe/repository"
	"hello-cafe/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"hello-cafe/internal/api"
	"hello-cafe/internal/db"
)

type server struct {
	ginEngine *gin.Engine

	adminHandler handler.AdminHandler
	itemHandler  handler.ItemHandler

	adminService service.AdminService
	itemService  service.ItemService

	repo repository.Repository
}

func newServer() (*server, error) {
	s := &server{ginEngine: gin.Default()}

	cfg, err := api.Config()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get api configuration")
	}

	if err = db.Connect(cfg.DB); err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	if err := s.initRepository(); err != nil {
		return nil, errors.Wrap(err, "failed to init repository")
	}

	if err := s.initService(); err != nil {
		return nil, errors.Wrap(err, "failed to init service")
	}

	if err := s.initHandler(); err != nil {
		return nil, errors.Wrap(err, "failed to init handler")
	}

	s.initRoutes()

	return s, nil
}

func (s *server) initRepository() (err error) {
	if s.repo, err = repository.NewRepository(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *server) initService() (err error) {
	if s.adminService, err = service.NewAdminService(s.repo); err != nil {
		return errors.WithStack(err)
	}

	if s.itemService, err = service.NewItemService(s.repo); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *server) initHandler() (err error) {
	if s.adminHandler, err = handler.NewAdminHandler(s.adminService); err != nil {
		return errors.Wrap(err, "failed to create admin handler")
	}

	if s.itemHandler, err = handler.NewItemHandler(s.itemService); err != nil {
		return errors.Wrap(err, "failed to create item handler")
	}

	return nil
}

func (s *server) initRoutes() {
	v1 := s.ginEngine.Group("v1")

	{
		user := v1.Group("/admin")
		user.POST("/sign-in", s.adminHandler.SignIn)   // 로그인
		user.POST("/sign-up", s.adminHandler.SignUp)   // 회원가입
		user.POST("/sign-out", s.adminHandler.SignOut) // 로그아웃
	}

	{
		item := v1.Group("/items", middleware.TokenAuthMiddleware)
		item.POST("", s.itemHandler.Create)             // 상품 등록
		item.PUT("/:item_seq", s.itemHandler.Update)    // 상품 수정
		item.DELETE("/:item_seq", s.itemHandler.Delete) // 상품 삭제
		item.GET("", s.itemHandler.Find)                // 상품 리스트 조회
		item.GET("/:item_seq", s.itemHandler.Get)       // 상품 상세 조회
		item.GET("/search", s.itemHandler.Search)       // 상품 이름 검색
	}
}

func (s *server) start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.ginEngine,
	}

	go func() {
		// 서비스 접속
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 5초의 타임아웃으로 인해 인터럽트 신호가 서버를 정상종료 할 때까지 기다립니다.
	quit := make(chan os.Signal)
	// kill (파라미터 없음) 기본값으로 syscanll.SIGTERM를 보냅니다
	// kill -2 는 syscall.SIGINT를 보냅니다
	// kill -9 는 syscall.SIGKILL를 보내지만 캐치할수 없으므로, 추가할 필요가 없습니다.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 5초의 타임아웃으로 ctx.Done()을 캐치합니다.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
