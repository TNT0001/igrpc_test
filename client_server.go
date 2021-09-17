package main

import (
	"context"
	"crypto/x509"
	"flag"
	"fmt"
	"git.ghtk.vn/gmicro/gmicro/config"
	"git.ghtk.vn/gmicro/gmicro/service/grpcclient"
	"git.ghtk.vn/gmicro/ig/igrpc-proto/generated/igdata-service"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
	interceptor "tungnguyen/client_interceptor"
)

var IgdataClient igrpcproto.IgdataClient

var (
	appName        = "iGHTK gateway service"
	configFileFlag = flag.String("config.file", "", "Path to configuration file.")
)

func init(){
	flag.Parse()
	config.Init(
		config.New(
			config.WithDefaultEnvVars(""),
			config.WithDefaultConfigFile(appName, *configFileFlag),
		),
	)

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certCalocal) {
		log.Println("/home/tungnguyenthanh/Downloads/ghtk-cert.pem")
	}
	cred := credentials.NewClientTLSFromCert(certPool, "")
	//igdata-grpc.ghtklab.com:8443
	//cred, err := LoadTLSCredentialsOld()
	//if err != nil {
	//	log.Fatalln("error when load cert")
	//}
	log.Println("load cert ok")
	client := grpcclient.NewClient("localhost:8087",
		grpcclient.WithTransportCredentials(cred),
		grpcclient.WithInterceptor(
			interceptor.NewMiddleware(),
			),
		grpcclient.WithBlock(),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	conn, err := client.Connect(ctx)
	if err != nil {
		fmt.Printf("can't establish connect err: %s\n", err)
		os.Exit(1)
	}
	IgdataClient = igrpcproto.NewIgdataClient(conn)
	result, err := IgdataClient.Ping(context.Background(), &igrpcproto.PingRequest{Message: "tungnguyen"})
	fmt.Println(result)
}
