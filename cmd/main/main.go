package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest_api/internal/config"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	/*cfgMongo := cfg.MongoDb
	mongoDBClient, err := mongodb.NewClient(
		context.Background(),
		cfgMongo.Host,
		cfgMongo.Port,
		cfgMongo.Username,
		cfgMongo.Password,
		cfgMongo.Database,
		cfgMongo.AuthDB,
	)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfgMongo.Collection, logger)*/

	/*user1 := user.User{
		ID:           "",
		Email:        "theartdevel@gmail.com",
		Username:     "theartofdevel",
		PasswordHash: "12345",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)

	user2 := user.User{
		ID:           "",
		Email:        "more@mail.here",
		Username:     "moremail",
		PasswordHash: "54321",
	}
	user2ID, err := storage.Create(context.Background(), user2)
	if err != nil {
		panic(err)
	}
	logger.Info(user2ID)

	user2Found, err := storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	logger.Info(user2Found)

	user2Found.Email = "newEmail@here.ok"
	err = storage.Update(context.Background(), user2Found)
	if err != nil {
		panic(err)
	}

	err = storage.Delete(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}

	_, err = storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}

	users, err := storage.FindAll(context.Background())
	fmt.Println(users)*/

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix soket %s", socketPath)
	} else {
		logger.Info("listen tcp ")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}
