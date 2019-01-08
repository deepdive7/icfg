package icfg_test

import (
	"fmt"
	"testing"

	"github.com/deepdive7/icfg"
	"github.com/stretchr/testify/assert"
)

var cfg = icfg.NewConfig()

func TestICFG(t *testing.T) {
	icfg.StringVar("system", "icfg", "system name")
	icfg.StringVar("config", "./config_demo.json", "json config path")
	icfg.LoadEnv([]string{"GOPATH", "GOROOT"})
	icfg.Parse()
	cfg := icfg.String("config")
	icfg.LoadCfg(cfg)

	fmt.Println(icfg.String("system")) // icfg
	fmt.Println(icfg.String("GOROOT")) // /opt/soft/go
	fmt.Println(icfg.IntMap("maps.int")) // map[i:1]
	fmt.Println(icfg.FloatMap("maps.float")) // map[f:0.01]
	fmt.Println(icfg.StrMap("maps.str")) //map[a:a]
	fmt.Println(icfg.IntArray("int_arr")) // [0 1 2 3 4 5 6 7 8 9 10 11]
	fmt.Println(icfg.Int("int_arr.1")) // 1

	pattern := "int_arr.[^1234567]"
	fmt.Println(icfg.Match(pattern).IntMap()) // map[int_arr.0:0 int_arr.9:9 int_arr.8:8]
	m := icfg.Map("maps.int")
	fmt.Println(m["i"].Int())
}

func TestConfigDefault(t *testing.T) {
	project := "icfg"
	icfg.SetDefaultKey("project", &project)
	assert.Equal(t, project, cfg.String("project"))
}

func TestConfigFlag(t *testing.T) {
	var name = "haha"
	var age = 21
	var phone = int64(17877652365)
	var id uint64 = 192839
	var money = 0.01

	icfg.StringVar("name", name, "my name")
	icfg.IntVar("age", age, "Age")
	icfg.Int64Var("phone", phone, "Phone number")
	icfg.Uint64Var("id", id, "ID number")
	icfg.Float64Var("money", money, "Left money")
	icfg.Parse()

	assert.Equal(t, name, cfg.String("name"))
	assert.Equal(t, age, cfg.Int("age"))
	assert.Equal(t, phone, cfg.Int64("phone"))
	assert.Equal(t, id, cfg.Uint64("id"))
	assert.Equal(t, money, cfg.Float("money"))
}

func TestConfigJson(t *testing.T) {
	// default config path is "./config.json"
	cfg.LoadCfg("./config_demo.json")
	assert.Equal(t, "./config_demo.json", cfg.String("config"))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, cfg.IntArray("int_arr"))
	assert.Equal(t, []float64{0.1, 0.2, 0.3}, cfg.FloatArray("float_arr"))
	assert.Equal(t, "INetwork", cfg.String("network.name"))
	assert.Equal(t, "Boom", cfg.String("A.B.C"))
}

func TestConfigEnv(t *testing.T) {
	keys := []string{"GOROOT", "GOPATH", "PATH"}
	icfg.LoadEnv(keys)
	assert.Equal(t, "/opt/soft/go", cfg.String("GOROOT"))
}

func TestConfig_Match(t *testing.T) {
	pattern := `match.sub_[a-z]*.[a-z]a[a-z]`
	assert.Equal(t, map[string]int{
		"match.sub_match.cab": 3,
		"match.sub_match.bac": 2,
	}, cfg.Match(pattern).IntMap())
	pattern = `network.listeners.[0-9].protocol`
	except := map[string]string{
		"network.listeners.0.protocol": "udp",
		"network.listeners.1.protocol": "tcp",
		"network.listeners.2.protocol": "kcp",
	}
	assert.Equal(t, except, cfg.Match(pattern).StrMap())

	pattern = `int_arr.[0-9]{2}`
	intExcept := map[string]int{
		"int_arr.10": 10,
		"int_arr.11": 11,
	}
	assert.Equal(t, intExcept, cfg.Match(pattern).IntMap())
}

func TestConfig_Dump(t *testing.T) {
	if false {
		cfg.LoadCfg("./config_demo.json")
		peerName := "NewPeer"
		cfg.Set("peer_name", &peerName)
		assert.Equal(t, nil, cfg.Dump())
		// Check new config
		cfg.LoadCfg()
		assert.Equal(t, peerName, cfg.String("peer_name"))
	}
}
