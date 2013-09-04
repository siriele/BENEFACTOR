package channels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//settings that can't be changed on the fly
type Options struct {
	MaxDuration    int64 `json:"max_duration"`
	MaxParticpants int64 `json:"max_participants"`
	Synchronous    bool  `json:"synchronous"`
}

func NewOptions() *Options {
	o := &Options{}
	o.MaxDuration = 9999999999
	o.MaxParticpants = 9999999999
	o.Synchronous = false
	return o
}

//add lock
//add players

//probably make something that can be changed on the fly

const (
	BEGINNING = 0
	BEFORE    = -1
	END       = -2
	AFTER     = -3
)

type Action string

const (
	REMOVE Action = "remove"
	ADD    Action = "add"
	SET    Action = "set"
)

type Delta struct {
	Path   []interface{} `json:"path"`
	Value  interface{}   `json:"value"`
	Action Action        `json:"action"`
}

type QueryType string

const (
	PARTICIPANT QueryType = "participant"
	CLOSER      QueryType = "closer"
	UPDATER     QueryType = "updater"
	TYPE        QueryType = "type"
	CREATOR     QueryType = "creator"
	IDS         QueryType = "ids"
)

type Status string

const (
	should pass in config
	DB = "http://btuser:XXXXX@<host>5985/"
)
const (
	OPEN = iota - 1
	CLOSED
	ALL
)

func NewAddDelta(path []interface{}, value interface{}) *Delta {
	return newDelta(ADD, path, value)
}

func NewSetDelta(path []interface{}, value interface{}) *Delta {
	return newDelta(SET, path, value)
}

func NewRemoveDelta(path []interface{}, value interface{}) *Delta {
	return newDelta(REMOVE, path, value)
}

func newDelta(a Action, p []interface{}, v interface{}) *Delta {
	return &Delta{Action: a, Path: p, Value: v}
}

func MakeChannel(g, u, t string, c *http.Client, o *Options, s map[string]interface{}) interface{} {

	params := map[string]string{
		"gameId":       g,
		"breaktimeId":  u,
		"channel_type": t,
	}

	query := make([]string, 0, len(params))
	for name, value := range params {
		query = append(query, name+"="+url.QueryEscape(value))
	}

	channel := map[string]interface{}{
		"options": *o,
		"state":   s,
	}
	mb, _ := json.Marshal(&channel)
	req, _ := http.NewRequest("PUT", DB+g+"_channels/_design/channels/_update/new_channel?"+strings.Join(query, "&"), bytes.NewReader(mb))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(mb)))
	r, _ := c.Do(req)
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println("status:", r.Status)
	var v interface{}
	json.Unmarshal(b, &v)
	return v
}

func UpdateChannel(g, u, t, id string, c *http.Client, ds []Delta) interface{} {
	params := map[string]string{
		"gameId":       g,
		"breaktimeId":  u,
		"channel_type": t,
	}

	query := make([]string, 0, len(params))
	for name, value := range params {
		query = append(query, name+"="+url.QueryEscape(value))
	}
	mb, _ := json.Marshal(&ds)
	req, _ := http.NewRequest("PUT", DB+g+"_channels/_design/channels/_update/update_channel/"+id+"?"+strings.Join(query, "&"), bytes.NewReader(mb))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(mb)))
	r, _ := c.Do(req)
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println("status:", r.Status)
	var v interface{}
	json.Unmarshal(b, &v)
	return v
}

func JoinChannel(g, u, t, id string, c *http.Client) interface{} {
	params := map[string]string{
		"gameId":       g,
		"breaktimeId":  u,
		"channel_type": t,
	}

	query := make([]string, 0, len(params))
	for name, value := range params {
		query = append(query, name+"="+url.QueryEscape(value))
	}
	req, _ := http.NewRequest("PUT", DB+g+"_channels/_design/channels/_update/join_channel/"+id+"?"+strings.Join(query, "&"), nil)
	req.Header.Add("Content-Type", "application/json")
	r, _ := c.Do(req)
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println("status:", r.Status)
	var v interface{}
	json.Unmarshal(b, &v)
	return v
}

func CloseChannel(g, u, t, id string, c *http.Client) interface{} {
	params := map[string]string{
		"gameId":       g,
		"breaktimeId":  u,
		"channel_type": t,
	}

	query := make([]string, 0, len(params))
	for name, value := range params {
		query = append(query, name+"="+url.QueryEscape(value))
	}
	req, _ := http.NewRequest("PUT", DB+g+"_channels/_design/channels/_update/close_channel/"+id+"?"+strings.Join(query, "&"), nil)
	req.Header.Add("Content-Type", "application/json")
	r, _ := c.Do(req)
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println("status:", r.Status)
	var v interface{}
	json.Unmarshal(b, &v)
	return v
}

func LeaveChannel(g, u, t, id string, c *http.Client) interface{} {
	params := map[string]string{
		"gameId":       g,
		"breaktimeId":  u,
		"channel_type": t,
	}

	query := make([]string, 0, len(params))
	for name, value := range params {
		query = append(query, name+"="+url.QueryEscape(value))
	}
	req, _ := http.NewRequest("PUT", DB+g+"_channels/_design/channels/_update/leave_channel/"+id+"?"+strings.Join(query, "&"), nil)
	req.Header.Add("Content-Type", "application/json")
	r, _ := c.Do(req)
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println("status:", r.Status)
	var v interface{}
	json.Unmarshal(b, &v)
	return v
}

func EditChannel() {

}

func GetChannels() {

}

func GetChannelState(g, u, t, id string, c *http.Client) interface{} {

}

func GetChannelsBy(q QueryType, g, u, t, id string, c *http.Client) {

}
