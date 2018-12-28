# ICFG
A simple and easily use config manager.
> Note: You can use `icfg` freely in your project. It’s always great to mark the `icfg` github address in your project. HaHa~
### Usage
Key search priority: flag > config > env

#### Set Default key
```
func TestConfigDefault(t *testing.T) {
	project := "icfg"
	icfg.SetDefaultKey("project", &project)
	assert.Equal(t, project, cfg.String("project"))
}
```

#### Flag
```
func TestConfigFlag(t *testing.T) {
	var name = "zt"
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
```

#### JSON Config file
```
func TestConfigJson(t *testing.T) {
	// default config path is "./config.json"
	cfg.LoadCfg("./config_demo.json")
	assert.Equal(t, "./config_demo.json", cfg.String("config"))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, cfg.IntArray("int_arr"))
	assert.Equal(t, []float64{0.1, 0.2, 0.3}, cfg.FloatArray("float_arr"))
	assert.Equal(t, "INetwork", cfg.String("network.name"))
	assert.Equal(t, "Boom", cfg.String("A.B.C"))
}
```

#### Load Env
```
func TestConfigEnv(t *testing.T) {
	keys := []string{"GOROOT", "GOPATH", "PATH"}
	cfg.LoadEnv(keys)
	assert.Equal(t, "/opt/soft/go", cfg.String("GOROOT"))
}
```

#### RegExp Match
```
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
		"int_arr.10":10,
		"int_arr.11":11,
	}
	assert.Equal(t, intExcept, cfg.Match(pattern).IntMap())
}
```

#### Dump Config Keys
Use the original configuration path plus the second-level timestamp as the file name backup. 
For example: `config_2018-12-28_11:15:48.json`. 
```
func TestConfig_Dump(t *testing.T) {
	if testDump {
		cfg.LoadCfg("./config_demo.json")
		peerName := "IPeer"
		cfg.Set("peer_name", &peerName)
		assert.Equal(t, nil, cfg.Dump())
		// Check new config
		cfg.LoadCfg()
        assert.Equal(t, peerName, cfg.String("peer_name"))
	}
}
```