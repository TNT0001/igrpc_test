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
	"os"
	"tungnguyen/client_interceptor"
)

var IgdataClient igrpcproto.IgdataClient

var (
	appName        = "iGHTK gateway service"
	configFileFlag = flag.String("config.file", "", "Path to configuration file.")
	versionFlag    = flag.Bool("version", false, "Show version information.")
	rsaCertPEM = [] byte("-----BEGIN CERTIFICATE-----\nMIIFHzCCAwegAwIBAgIUG0NpL8v2J5PEEeWBltdAfNAzwPwwDQYJKoZIhvcNAQEL\nBQAwFDESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIxMDgwMzA1NDIzN1oXDTIyMDgw\nMzA1NDIzN1owFDESMBAGA1UEAwwJbG9jYWxob3N0MIICIjANBgkqhkiG9w0BAQEF\nAAOCAg8AMIICCgKCAgEA1V+e++aO0popPMqd937QaGRRlwgB+C6SlnTiQY+gGY+E\nU16ToywgNj4s6fg6VKhhia2urUEPwX91KlKweiJ+PdDJBNKe0oufhmLrs7cO4OZp\nQr9Z6XZeh39zNjYDsWK1Tp6T8zJ4E+nfqffjRiAJ9rszr1waT6ElAf9QIHSrGaZN\nDVaCM95jgN/nVx46exVqDcIjuWAXIFkO6NVBaC+ev76DqNC/raZpkxbUGnFP7PGV\nrVbEILVmw4s2czcv9hVHhtOZWb7eXv0VFMy8FoF8HiLjW+eGwWqI0r2yJ6YPRbsk\nHrW4cviZ7cLV5nOdfMyWfCIGa8Pj2+eOuDFANf2PRkzUf0itIp6kG9CCVviC917b\nBGHKpYofwLe5BXI7+Qqx+M0IuDouiRE6TOrQxZFiPXA+miIWUXd+0g2YKhXRb6V/\njbPm7VPNPqpOqePZ97PySRMQDIjtg4HhpK20hXV090uFX/+KDhW0Mnxd6eDpQilM\nJvUmzakyuUNvNUP3ep7/I+/EZd9bUuNmZUR3bjrOqlyW46RfgtuyODhoiZ5gQYI/\niWrU4IPMftyXM5FRqHKdJFmAD8YrxAt7HmO8rJZ4Bk42HgYX3yQpmKOAFGMuD7kv\nRRwCon8bUjMR7KiIjJkAwtBqmO4OEFlz7zn/jgi15TB0BGWFAtdhgVRyu4waYlsC\nAwEAAaNpMGcwHQYDVR0OBBYEFHIh9FI3Cn8Pu4bDup6pmTmKadRzMB8GA1UdIwQY\nMBaAFHIh9FI3Cn8Pu4bDup6pmTmKadRzMA8GA1UdEwEB/wQFMAMBAf8wFAYDVR0R\nBA0wC4IJbG9jYWxob3N0MA0GCSqGSIb3DQEBCwUAA4ICAQCk7zi5ZtG80NozZnT6\n2b1lkseryKJ3yMbhGE+saexSahbhg9r6He1pyl61MCU0vFhVGJoLDX11nEkF+QNW\nNTFpGvp+rOWix2SP8lo3Lb8gUU5A84lPZUvPYyaq1LGTsMuO+HydqpzLPHX+nUKl\nPdvc9COeJDkTkKBTFUvi0Skd+iaQ43Sfpjz4lvujG2TPQlHHkmKY2DX8QAHg18yJ\nUJ0iNp7tmc1OFqPLQw6K9J7TNsw06Oa2akrgaKatFCEwDcO5JF5iJ18GwIUlpWxv\njcTbDqRKbcF4s3c/1m992TVgZATNHddt8nRIr1bTMR8e9fLRfk/DfsQFpkhCGsDe\nW5Q9XokDNBAp/ZBEE3zW2Zpb3yVC9Js7iYrs4QZo5pR3HR7skgL0HlHxvpTmlLC/\n30kImNMv8onK3TjI42TyrH6qXqvQSiJgodJ7WrvpeDNeu20swm/tanHh4ysMVaUd\ngjG9e/lyoNEcyqstyRMvgnotBAcdp3ZIndX1N+wxWkrVHVWJmPenKDHAPTm7dsqR\nZfUDh+mSo89BpZOXM+uRsbGBf/t6KGM4fOM0apnJ6/p7ym2v4JCjatirXXMeP7YH\nIqZlcUeZrLnJYjP6NFDBRuxYDc2HRNBsTSRdxV9CL9A8+PAKfJxvSakPYEU9t3Oi\nBPBd+1MJVBKJURtEIpZD499bug==\n-----END CERTIFICATE-----\n")
)

func init(){
	flag.Parse()

	config.Init(
		config.New(
			config.WithDefaultEnvVars(""),
			config.WithDefaultConfigFile(appName, *configFileFlag),
		),
	)
	a := x509.NewCertPool()
	ok := a.AppendCertsFromPEM(rsaCertPEM)
	if !ok {
		fmt.Println("can't not create cert pool")
		os.Exit(1)
	}
	cred:= credentials.NewClientTLSFromCert(a, "")

	client := grpcclient.NewClient("localhost:50443",
		grpcclient.WithTransportCredentials(cred),
		grpcclient.WithInterceptor(
			interceptor.NewMiddleware(),
			),
		grpcclient.WithBlock(),
	)
	conn, err := client.Connect(context.Background())
	if err != nil {
		fmt.Printf("can't establish connect err: %s\n", err)
		os.Exit(1)
	}
	IgdataClient = igrpcproto.NewIgdataClient(conn)
	result, err := IgdataClient.Ping(context.Background(), &igrpcproto.PingRequest{Message: "tungnguyen"})
	if err != nil {
		fmt.Printf("cann't get response, err : %v\n", err)
		return
	}
	fmt.Println(result.Message)
}
