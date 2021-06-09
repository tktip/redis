package main

import (
	"encoding/json"
	"fmt"

	"github.com/haraldfw/cfger"
	"github.com/tktip/redis/pkg/redis"
)

type S struct {
	X int `yaml:"debug"`
}

func (s *S) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}

func (s *S) UnmarshalBinary(data []byte) error {
	*s = S{}
	return json.Unmarshal(data, s)
}

func main() {
	//for test
	useDummy := true
	var h redis.Handler
	if useDummy {

		h = &redis.MockHandler{
			Cache: &map[string][]byte{},
		}
	} else {
		var h redis.Handler
		h = &redis.DefaultHandler{}

		_, err := cfger.ReadStructuredCfg("file::./dev_cfg/cfg.yml", &h)
		if err != nil {
			panic(err.Error())
		}
	}
	defer h.Close()

	var x *string
	f := "hey"
	x = &f

	err := h.Write("test", x)
	if err != nil {
		panic(err.Error())
	}

	var y string
	err = h.GetAndScanTo("test", &y)
	if err != nil {
		panic(err.Error())
	}

	s := S{X: 10}
	err = h.Write("test1", &s)
	if err != nil {
		panic(err.Error())
	}

	n := S{}
	err = h.GetAndUnmarshalBinary("test1", &n)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("didnt crash!")
}
