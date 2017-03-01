package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

var mu sync.Mutex

// MainController is a controller about user
type MainController struct {
	beego.Controller
}

// UserController is a controller about user
type UserController struct {
	beego.Controller
}

// Get function is to get user info
func (c *MainController) Get() {
	c.Data["Website"] = "wentian"
	c.Data["Email"] = "jwentian@redhat.com"
	c.TplName = "websocket.html"
}

func (this *MainController) Join() {
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	this.TplName = "websocket.html"
	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		js, err := simplejson.NewJson(p)
		if err != nil {
			panic(err.Error())
		}
		//log.Println(js)
		mapdata, err := js.Map()
		if err != nil {
			panic(err.Error())
		}
		parallel := mapdata["parallel"].(string)
		jobname := mapdata["jobname"].(string)
		svcname := mapdata["svcname"].(string)
		intparallel, err := strconv.Atoi(parallel)
		runParallelHttpCurl(ws, jobname, svcname, intparallel)
	}
}

func runParallelHttpCurl(ws *websocket.Conn, jobname string, svcname string, parallel int) {
	parallel_num := 0
	begin_sum := time.Now()
	versionmap := make(map[string]int)
	chhttp := make(chan map[string]interface{}, 100)
	for i := 0; i < parallel; i++ {
		go curlHttp(ws, jobname, svcname, parallel, i, chhttp)
	}
	for {
		select {
		case data := <-chhttp:
			versionmap[data["version"].(string)]++
			parallel_num++
		default:
			if parallel_num == parallel {
				tr := make(map[string]interface{})
				dis := time.Since(begin_sum).Seconds()
				dis_str := strconv.FormatFloat(dis, 'g', 10, 64)
				tr["type"] = "summary"
				tr["dis"] = dis_str
				tr["versionmap"] = versionmap
				trm, _ := json.Marshal(tr)
				trmstr := strings.Replace(string(trm), " ", "", -1)
				trmstr = strings.Replace(trmstr, "\n", "", -1)
				mu.Lock()
				defer mu.Unlock()
				ws.WriteMessage(websocket.TextMessage, []byte(trmstr))
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func curlHttp(ws *websocket.Conn, jobname string, svcname string, parallel int, jobid int, chhttp chan map[string]interface{}) {
	begin := time.Now()
	result, err := httplib.Get(svcname).String()
	tr := make(map[string]interface{})

	if err != nil {
		panic(err.Error())
	}
	dis := time.Since(begin).Seconds()
	tr["data"] = result
	tr["dis"] = dis
	tr["jobid"] = jobid
	tr["type"] = "jobflow"

	userdatamap := make(map[string]interface{})
	json.Unmarshal([]byte(result), &userdatamap)

	for _, v := range userdatamap {
		userd := v.(map[string]interface{})
		tr["version"] = userd["Version"]
	}

	trm, _ := json.Marshal(tr)
	trmstr := strings.Replace(string(trm), " ", "", -1)
	trmstr = strings.Replace(trmstr, "\n", "", -1)
	mu.Lock()
	defer mu.Unlock()
	ws.WriteMessage(websocket.TextMessage, []byte(trmstr))
	chhttp <- tr
}

// Get function is to get user info
func (c *UserController) Get() {
	c.Data["Website"] = "user"
	c.Data["Email"] = "user@redhat.com"
	c.TplName = "index.tpl"
}
