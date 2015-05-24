// config
package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type Cfg struct {
	reader *bytes.Buffer
}

func New(path string) (*Cfg, error) {

	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &Cfg{
		reader: bytes.NewBuffer(buff),
	}, nil
}

func (t *Cfg) ReadConfig() (map[string]string, error) {

	cfg := make(map[string]string)

	var position int
	for {
		line, err := t.reader.ReadString('\n')
		if err != nil {
			break
		}

		if len(line) > 0 {
			split := strings.Split(line, "=")
			if len(split) != 2 {
				return cfg, fmt.Errorf("Config invalid on line %d near %s", position, line)
			}

			cfg[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		}

		position++
	}

	return cfg, nil
}
