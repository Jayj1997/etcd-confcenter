/*
 * @Author       : jayj
 * @Date         : 2021-11-15 14:08:34
 * @Description  :
 */
package main

import (
	"confcenter/routes"
	"confcenter/service"
	"confcenter/utils"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	isAuth  = flag.Bool("auth", false, "enable authentication")
	isTLS   = flag.Bool("tls", false, "enable tls")
	ca      = flag.String("ca", "", "")
	cert    = flag.String("cert", "", "")
	keyfile = flag.String("keyfile", "", "")
	timeout = flag.Int("timeout", 5, "dial timeout")
)

func main() {

	flag.Parse()

	v3 := &service.EtcdV3Service{
		IsAuth:      *isAuth,
		IsTls:       *isTLS,
		CaFile:      *ca,
		Cert:        *cert,
		DialTimeout: time.Duration(*timeout) * time.Second,
		Mu:          sync.RWMutex{},
	}

	fmt.Println(v3)

	g := routes.LoadGin(v3)

	server := &http.Server{
		Addr:           ":8080",
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
