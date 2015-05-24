// rest
package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"net/http"
)

type Server struct {
	addr   string
	engine *gin.Engine
	api    Api
}

type Api interface {
	GetVoiceCfg() map[string]string
	GetCmdCfg() map[string]string
	ExecuteVoice(r model.Request)
	ExecuteCommand(r model.Request)
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (t *Server) Init(api Api) error {
	t.api = api

	e := gin.Default()

	e.POST("/say", func(c *gin.Context) {

		var res model.Response
		var body model.Request
		if !c.Bind(&body) {
			//jbg
			return
		}

		res.Success = true
		res.Message = "Executed"
		c.JSON(http.StatusOK, &res)

		go func() {
			log.Println(body)
			t.api.ExecuteCommand(body)
			t.api.ExecuteVoice(body)
		}()

	})

	e.GET("/cfg", func(c *gin.Context) {

		cmd := t.api.GetCmdCfg()
		cmdText := ""
		for key, val := range cmd {
			cmdText += key + ":" + val + "\n"
		}

		voice := t.api.GetVoiceCfg()
		voiceText := ""
		for key, val := range voice {
			voiceText += key + ":" + val + "\n"
		}

		c.String(
			http.StatusOK,
			"Command config\n%s\nVoice config\n %s\n", cmdText, voiceText,
		)
	})

	t.engine = e

	return nil
}

func (t *Server) ListenAndServe() error {
	return t.engine.Run(t.addr)
}
