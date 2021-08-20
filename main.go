package main

import (
	"fmt"
	"exporter/collector"
	"log"
	"crypto/subtle"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func basicAuth(username string, passwordHash string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || CheckPasswordHash(pass, passwordHash) != true {
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised\n"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	IP := "127.0.0.1"
	PORT := "8080"
	basicAuthUser := ""
	basicAuthPassHash := ""
	enableAuth := false
	if len(os.Args) > 2 {
		IP = os.Args[1]
		PORT = os.Args[2]
	}
	if len(os.Args) > 3 {
		basicAuthUser = os.Args[3]
		basicAuthPassHash = os.Args[4]
		enableAuth = true
	}

	foo := collector.NewMemoryCollector()
	prometheus.MustRegister(foo)
	listen := IP + ":" + PORT

	fmt.Print("Exporter listening on: ", IP,":" , PORT ,"\n")

	if enableAuth {
		http.Handle("/metrics", basicAuth(basicAuthUser, basicAuthPassHash, promhttp.Handler()))
	} else {
		http.Handle("/metrics", promhttp.Handler())
	}

	log.Fatal(http.ListenAndServe(listen, nil))
}
