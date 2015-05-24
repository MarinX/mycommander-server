// decoder
package decoder

import (
	"model"
	"strings"
)

type Decoder struct {
	cfg Configuration
}

type Configuration interface {
	GetVoiceCfg() map[string]string
	GetCmdCfg() map[string]string
}

func New(cfg Configuration) *Decoder {
	return &Decoder{
		cfg: cfg,
	}
}

func (t *Decoder) DecodeCommand(r model.Request) string {
	r.Text = strings.ToLower(r.Text)

	var ret string

	for key, val := range t.cfg.GetCmdCfg() {
		if strings.Contains(r.Text, key) {
			if len(ret) == 0 {
				ret += val + " "
			} else {
				ret += "&& " + val + " "
			}
		}
	}

	return ret
}

func (t *Decoder) DecodeVoice(r model.Request) (string, bool) {
	var ret string

	r.Text = strings.ToLower(r.Text)

	split := strings.Split(r.Text, " ")
	for key, val := range t.cfg.GetVoiceCfg() {
		for i := range split {
			if strings.Contains(key, split[i]) {
				if len(split) > (i + 1) {
					ret += val + " " + split[i+1] + ". "
				} else {
					ret += val + ". "
				}

			}
		}
	}

	return ret, t.isIntroduction(r.Text)
}

func (t *Decoder) isIntroduction(text string) bool {

	for _, val := range model.GetIntroductionWords() {
		if strings.Contains(text, val) {
			return true
		}
	}

	return false
}
