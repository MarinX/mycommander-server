// service
package service

import (
	"decoder"
	"log"
	"model"
	"os/exec"
	"rest"
	"speech"
)

type Service struct {
	useVoice bool
	voiceCfg map[string]string
	cmdCfg   map[string]string
	dec      *decoder.Decoder
	spch     *speech.Speech
}

func New(cmdCfg map[string]string) *Service {
	return &Service{
		cmdCfg: cmdCfg,
	}
}

func (t *Service) UseVoice(cfg map[string]string) {
	t.useVoice = true
	t.voiceCfg = cfg
	t.spch = speech.New(cfg[model.CFG_VOICE])
}

func (t *Service) Run() error {

	r := rest.New(model.HTTP_ADDR)

	if err := r.Init(t); err != nil {
		return err
	}

	t.dec = decoder.New(t)

	return r.ListenAndServe()
}

func (t *Service) ExecuteVoice(r model.Request) {
	if !t.useVoice {
		return
	}

	voice, intro := t.dec.DecodeVoice(r)
	log.Println("VOICE", voice)

	var introUsed bool
	if intro {
		t.spch.Introduction(r.DeviceName, r.User)
		introUsed = true
	}

	if len(voice) == 0 && !introUsed {
		t.spch.Unknow()
		return
	}

	t.spch.Voice(voice)
}

func (t *Service) ExecuteCommand(r model.Request) {
	cmd := t.dec.DecodeCommand(r)
	log.Println("EXEC", cmd)

	if len(cmd) == 0 {
		return
	}

	err := exec.Command(model.DEFAULT_SHELL, "-c", cmd).Start()
	if err != nil {
		log.Println(err)
		if t.useVoice {
			t.spch.Error()
		}
	}
}

func (t *Service) GetVoiceCfg() map[string]string {
	return t.voiceCfg
}

func (t *Service) GetCmdCfg() map[string]string {
	return t.cmdCfg
}
