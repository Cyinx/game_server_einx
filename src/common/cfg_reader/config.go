package cfg_reader

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Segment struct {
	values map[string]string
	name   string
}

type Config struct {
	cfgs map[string]*Segment
	path string
}

func NewConfig() *Config {
	return &Config{
		cfgs: make(map[string]*Segment),
	}
}

var name_match = regexp.MustCompile(`\[.+\]`)

func (this *Config) ReadConfig(path string) {
	this.path = path

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var curr_seg *Segment = nil
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if len(b) == 0 {
			continue
		}

		s := strings.TrimSpace(string(b))

		if strings.Index(s, "#") == 0 {
			continue
		}

		name := name_match.FindString(s)

		if len(name) > 0 {
			curr_seg = &Segment{
				values: make(map[string]string),
				name:   name[1 : len(name)-1],
			}
			this.cfgs[curr_seg.name] = curr_seg
			continue
		}

		if curr_seg == nil {
			continue
		}

		values := curr_seg.values

		e := strings.Index(s, "=")
		if e == -1 {
			panic("read wrong config file!")
		}

		object_name := s[0:e]
		value := s[e+1:]

		if i := strings.Index(value, "#"); i != -1 {
			value = value[0:i]
		}

		values[object_name] = value
	}

}

func (this *Config) GetNamedInt(seg_name string, name string) int {
	if seg_cfg, ok := this.cfgs[seg_name]; ok == true {
		if val_str, ok := seg_cfg.values[name]; ok == true {
			if val, err := strconv.Atoi(val_str); err == nil {
				return val
			}
		}
	}
	return 0
}

func (this *Config) GetNamedString(seg_name string, name string) string {
	if seg_cfg, ok := this.cfgs[seg_name]; ok == true {
		if val_str, ok := seg_cfg.values[name]; ok == true {
			return val_str
		}
	}
	return ""
}
