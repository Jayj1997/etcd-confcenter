/*
 * @Author       : jayj
 * @Date         : 2021-11-15 14:08:34
 * @Description  :
 */
package main

import (
	"etcdgate/routes"
	"etcdgate/service"
	"etcdgate/utils"
	"flag"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	isAuth    = flag.Bool("auth", false, "enable authentication")
	isTLS     = flag.Bool("tls", false, "enable tls")
	ca        = flag.String("ca", "", "ca")
	cert      = flag.String("cert", "", "cert")
	keyfile   = flag.String("keyfile", "", "keyfile")
	timeout   = flag.Int("timeout", 5, "dial timeout, eg. 5")
	port      = flag.String("port", ":8080", "server listen port, eg. :8080")
	separator = flag.String("separator", "/", "key separator")
)

func main() {

	flag.Parse()

	v3 := &service.EtcdV3Service{
		IsAuth:      *isAuth,
		IsTls:       *isTLS,
		CaFile:      *ca,
		Cert:        *cert,
		Separator:   *separator,
		DialTimeout: time.Duration(*timeout) * time.Second,
		Mu:          sync.RWMutex{},
	}

	g := routes.LoadGin(v3)

	server := &http.Server{
		Addr:           *port,
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Errorln("server stop or failed, info: ", err)
		}
	}()

	utils.GracefullyShutdown(server)
}
