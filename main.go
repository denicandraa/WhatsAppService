package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"

	"./controllers"
	"./handlers"

	cors "github.com/rs/cors/wrapper/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file

	f, err := os.OpenFile("logs/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	//gin.SetMode(gin.ReleaseMode)
	m.router = gin.Default()
	m.router.Use(cors.AllowAll())

	return nil
}

func main() {

	m := Main{}

	if m.initServer() != nil {
		return
	}

	w, err := os.OpenFile("logs/whatsapp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer w.Close()

	log.SetOutput(w)

	waInit := handlers.WaInit{}

	go waInit.WhatsAppInit()

	ctrl := controllers.GinCtrl{}

	was := m.router.Group("/whatsapp/")

	was.Static("login/", "./")

	//was.POST("reset", ctrl.ResetSession)

	send := was.Group("send/")
	send.POST("text", ctrl.SendText)
	send.POST("text_group", ctrl.SendTextGroup)

	//<-time.After(60 * time.Minute)

	fmt.Println("Copyright - SiMotor")

	m.router.Run(":8824")
}
