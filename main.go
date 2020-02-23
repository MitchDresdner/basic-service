package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v1.0.1"
	addr    = "0.0.0.0:8083"
	helpText = `
Basic service provides the ability to check continuity the to your Service.
In it's a basic form, you can determine whether you have a path to the
service, it's Hostname and IP Address'

   Commands
   ========
   http :8083/health
   http :8083/version

JSON messaging powered by: https://github.com/jakubroztocil/httpie
`
)

type Health struct {
	Hostname    string
	IpAddress 	string
}

func main() {
	msg := fmt.Sprintf("Basic service running: http://%s", addr)
	log.Println(msg)

	host, err := os.Hostname()
	if err != nil {
		log.Fatal("Unable to determine Hostname " + err.Error())
		return
	}

	ip, err := externalIP()
	if err != nil {
		log.Fatal("Unable to determine IP Address " + err.Error())
		return
	}

	// API handlers
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helpText)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Performing health check")

		health := Health{host, ip}

		js, err := json.Marshal(health)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Update any header fields
		w.Header().Set("Content-Type", "application/json")


		log.Println("Health is vigorous")

		w.Write(js)

	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

	// Listen for inbound API connections
	s := http.Server{Addr: addr}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	// Use signals to end processing
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Basic service received exit signal")

	s.Shutdown(context.Background())
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}