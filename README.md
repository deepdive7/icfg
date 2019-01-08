# icfg
简单的配置读取和管理, flag和env变量保存在全局变量flagKV和envKV中，所有\*Config都可
以访问，LoadCfg读取文件的配置只属于该\*Config. 配置获取时的搜索顺序依次为
flag>config>env，最先搜索flag, SetDefaultKey设置的变量保存在envKV中.

## 支持接口
### 全局函数
- NewConfig() *Config
- TimeStamp() string
- SetDefaultKey(k string, v interface{})
- LoadEnv(keys \[\]string) // Load env vars
- Parse() //Parse flag vars
- BoolVar(name string, value bool, usage string)
- StringVar(name string, value string, usage string)
- IntVar(name string, value int, usage string)
- Int64Var(name string, value int64, usage string)
- Uint64Var(name string, value uint64, usage string)
- Float64Var(name string, value float64, usage string)

### 以下函数属于\*Config, 但是包级别也有一样的函数集，封装了一个默认\*Config
- LoadCfg(path ...string) error
- Ele(path string) (v *Element, ok bool)
- Set(path string, value interface{})
- Bool(path string) bool
- String(path string) string
- Int(path string) int
- IntArray(path string) []int
- IntMap(path string) map[string]int
- Int64(path string) int64
- Uint64(path string) uint64
- Float(path string) float64
- FloatArray(path string) []float64
- FloatMap(path string) map[string]float64
- Array(path string) []*Element
- StrArray(path string) []string
- Map(path string) map[string]*Element
- StrMap(path string) map[string]string
- Match(pattern string) *Element

### Test
- config.json

```
{
  "host": "127.0.0.1",
  "peer_name": "IPeer",
  "network": {
    "name": "INetwork",
    "listeners": [
      {
        "protocol":"udp",
        "port": "1008",
        "name": "udp_listener"
      },
      {
        "protocol":"tcp",
        "port": "1009",
        "name": "tcp_listener"

      },
      {
        "protocol":"kcp",
        "port": "1010",
        "name": "kcp_listener"
      }
    ]
  },
  "A": {
    "B": {
      "C": "Boom"
    }
  },
  "int_arr": [0,1,2,3,4,5,6,7,8,9,10,11],
  "float_arr": [0.1, 0.2, 0.3],
  "match": {
    "sub_match": {
      "acb": 1,
      "bac": 2,
      "cab": 3
    },
    "submatch": {
      "a": 1
    }
  },
  "maps" : {
    "str": {
      "a":"a"
    },
    "float": {
      "f":0.01
    },
    "int": {
      "i":1
    }
  }
}
```

- Simple Test

```
unc TestICFG(t *testing.T) {
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
```