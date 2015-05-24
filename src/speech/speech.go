// speech
package speech

import (
	"fmt"
	"os/exec"
)

const (
	DEFAULT_ERROR  = "Error occured, please check log for details"
	DEFAULT_UKNOWN = "I am not programmed for saying this command"
	INTRODUCTION   = "My name is %s, nice to meet you %s!.With your %s you can control me. Try it out!"
)

type Speech struct {
	name string
}

func New(name string) *Speech {
	return &Speech{
		name: name,
	}
}

func (t *Speech) Voice(text string) {
	t.execute(text)
}

func (t *Speech) Error() {
	t.execute(DEFAULT_ERROR)
}

func (t *Speech) Unknow() {
	t.execute(DEFAULT_UKNOWN)
}

func (t *Speech) Introduction(device string, user string) {
	text := fmt.Sprintf(INTRODUCTION, t.name, user, device)
	t.execute(text)
}

func (t *Speech) execute(text string) {
	exec.Command("/usr/bin/pico2wave", "-wtemp.wav", text).Run()
	exec.Command("aplay", "temp.wav").Run()
	exec.Command("rm", "temp.wav").Run()

}
