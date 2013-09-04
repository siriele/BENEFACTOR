package main

import (
	"fmt"
	"html"
	"log"
	//"log/syslog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	// "time"
	// "crypto/sha1"
	"encoding/base64"
	"errors"
	"flag"
	"strings"
)

var service string
var port int
var dbport int
var username string
var password string
var baseUrl string
var host string
var client *http.Client

const (
	USER           = "%susers/_design/users/_show/install_check/%s?gameId=%s"
	PING           = "%susers/_design/users/_upate/ping/%s?gameId=%s"
	SECURE_USER    = "%susers/_design/users/_show/secure_user/%s?gameId=%s&token=%s"
	REFRESH_TOKEN  = "%susers/_design/users/_update/refresh_token/%s?gameId=%s&refresh_token=%s"
	NEW_TOKEN      = "%susers/_design/users/_update/refresh_token/%s?gameId=%s"
	REJECT_ACTIVE  = "%susers/_design/users/_update/refresh_token/%s?gameId=%s"
	REJECT_REFRESH = "%susers/_design/users/_update/refresh_token/%s?gameId=%s"
)

/*
	This is where the actual running server logic will live.
	Shutdown hooks and so forth...It shouldn't be too expansive
	every server endpoint will decide if it needs its request to be authenticated or not

*/
func main() {
	flag.IntVar(&port, "port", 8000, "port number to bind to")
	flag.IntVar(&dbport, "dbport", 5985, "port number to bind to")
	flag.StringVar(&service, "service", "Super Service", "port number to bind to")
	flag.StringVar(&username, "username", "breaktime", "database username")
	flag.StringVar(&password, "password", "breaktime", "database password")
	flag.StringVar(&host, "host", "ec2-23-22-125-97.compute-1.amazonaws.com", "host for database")
	flag.Parse()
	// initialize take up all cores
	cores := runtime.NumCPU()
	client = &http.Client{}
	used := 1
	baseUrl = fmt.Sprintf("http://%s:%s@%s:%d/", username, password, host, dbport)
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/foo", SecureRequest(fooHandler))
	http.HandleFunc("/bar", SemiSecureRequest(barHandler))
	http.HandleFunc("/foo/bar", fooBarHandler)
	http.HandleFunc("/health", healthHandler)
	log.Printf("Starting %s service Server on port %d\n", service, port)
	log.Printf("Using %d of %d cores\n", used, cores)
	c := make(chan os.Signal)

	signal.Notify(c)

	// go func() {
	// 	timeout := time.After(time.Second * 15)
	// 	for {
	// 		select {
	// 		// case t := <-time.After(time.Second):
	// 		// 	//runtime.Gosched()
	// 		// 	log.Printf("ticking away %d\n", t.Second())

	// 		case t := <-timeout:
	// 			//runtime.Gosched()
	// 			log.Printf("timed out...jumpoing off a cliff at %d\n", t.Second())
	// 			os.Exit(0)
	// 		}
	// 	}
	// }()
	go func() {
		select {
		case s := <-c:
			log.Fatalf("Received Signal %s \n", s.String())
			log.Fatal("running shutdown hooks\n")
		}
	}()
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	//use later for more control
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	//log.Fatal(s.ListenAndServe())
}

func NewToken(rw http.ResponseWriter, req *http.Request) {

}

func RefreshToken(rw http.ResponseWriter, req *http.Request) {

}

func Ping(rw http.ResponseWriter, req *http.Request) {

}

func CheckUser(rw http.ResponseWriter, req *http.Request) {

}

func fooHandler(rw http.ResponseWriter, req *SecuredRequest) {
	var name, pass, ip string
	name = req.UserId
	pass = req.GameId
	ip = req.Ip
	//name := "Foo"
	defer log.Println("name", name)
	fmt.Fprint(rw, name, pass)
	fmt.Fprintf(rw, "Hello, %q\n listening on port %s , and you're Ip is %s", html.EscapeString(req.URL.Path), port, ip)
}

func barHandler(rw http.ResponseWriter, req *SemiSecuredRequest) {
	defer log.Println("name", req.UserId)
	fmt.Fprintf(rw, "Hello, %q\n listening on port %d you are %s and with game %s", html.EscapeString(req.URL.Path), port, req.UserId, req.GameId)
}

func fooBarHandler(rw http.ResponseWriter, req *http.Request) {
	//name := "FooBar"
	name, pass, ip, _, err := SafeRequest(rw, req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
	}
	defer log.Println("name", name, pass, ip)
	fmt.Fprint(rw, name)
	fmt.Fprintf(rw, "Hello, %q\n listening on port %d", html.EscapeString(req.URL.Path), port)
}

func healthHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, HEALTHY)
}

const (
	AUTH    = "Authorization"
	TO_AUTH = `WWW-Authenticate`
	REALM   = `Basic realm="Breaktime Realm"`
	HEALTHY = "HEALTHY"
)

func SafeRequest(rw http.ResponseWriter, req *http.Request) (name, pass, ip, ua string, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = errors.New("Failed Authentication")
		}
	}()
	ip = getIpAddress(req)
	//token := req.Header.Get("x-active-token")
	ua = req.UserAgent()
	s := req.Header.Get(AUTH)
	//b, _ := base64.StdEncoding.DecodeString(s)
	tokens := strings.Split(string(s), " ")
	credentials, _ := base64.StdEncoding.DecodeString(tokens[1])
	digest := strings.Split(string(credentials), ":")
	name = digest[0]
	pass = digest[1]
	//I'm thinking maybe just a url but whatever
	//url := fmt.Sprintf("http://%s:%s@%s:%s/", username, password) //host amd port

	//call some show fuction
	//in here launch some config
	return name, pass, ip, ua, err
}

//just used for extracting the IP address
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func getIpAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIp := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIp == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIp
}

func SecureRequest(f func(http.ResponseWriter, *SecuredRequest)) func(http.ResponseWriter, *http.Request) {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		http.Error(rw, errors.New("Authentication failed").Error(), http.StatusUnauthorized)
	// 	}
	// }()
	return func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				rw.Header().Set(TO_AUTH, REALM)
				http.Error(rw, errors.New("Authentication failed").Error(), http.StatusUnauthorized)
			}
		}()
		ip := getIpAddress(req)
		log.Println(req.Header)
		//token := req.Header.Get("x-active-token")
		//ua := req.UserAgent()
		s := req.Header.Get(AUTH)
		log.Println(s)
		//b, _ := base64.StdEncoding.DecodeString(s)
		tokens := strings.Split(s, " ")
		log.Println(tokens)
		credentials, _ := base64.StdEncoding.DecodeString(tokens[1])
		digest := strings.Split(string(credentials), ":")
		name := digest[0]
		pass := digest[1]
		//I'm thinking maybe just a url but whatever
		url := fmt.Sprintf("http://%s:%s@%s:%s/", username, password) //host amd port
		log.Println(url)
		//call some show fuction
		//in here launch some config
		nreq := &SecuredRequest{}
		nreq.Request = req
		nreq.UserId = name
		nreq.GameId = pass
		nreq.Ip = ip
		resp, _ := client.Get(fmt.Sprintf(USER, baseUrl, name, pass))
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			panic("User not found")
		}
		f(rw, nreq)
	}
}

func SemiSecureRequest(f func(http.ResponseWriter, *SemiSecuredRequest)) func(http.ResponseWriter, *http.Request) {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		http.Error(rw, errors.New("Authentication failed").Error(), http.StatusUnauthorized)
	// 	}
	// }()
	return func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				rw.Header().Set(TO_AUTH, REALM)
				http.Error(rw, errors.New("Authentication failed").Error(), http.StatusUnauthorized)
			}
		}()
		ip := getIpAddress(req)
		s := req.Header.Get(AUTH)
		tokens := strings.Split(s, " ")
		credentials, _ := base64.StdEncoding.DecodeString(tokens[1])
		digest := strings.Split(string(credentials), ":")
		name := digest[0]
		pass := digest[1]
		//I'm thinking maybe just a url but whatever
		//url := fmt.Sprintf("http://%s:%s@%s:%s/", username, password) //host amd port
		//call some show fuction
		//in here launch some config
		sreq := &SemiSecuredRequest{}
		sreq.Request = req
		sreq.UserId = name
		sreq.GameId = pass
		sreq.Ip = ip
		url := fmt.Sprintf(USER, baseUrl, name, pass)
		log.Println(url)
		resp, _ := client.Get(url)
		if resp.StatusCode != 200 {
			panic("User Not found")
		}
		f(rw, sreq)
	}
}

type SecuredRequest struct {
	SemiSecuredRequest
	Token        string
	Refresh      string
	Expires      int64
	RefreshAgent string
	ActiveAgent  string
}

type SemiSecuredRequest struct {
	*http.Request
	UserId string
	GameId string
	Ip     string
}
