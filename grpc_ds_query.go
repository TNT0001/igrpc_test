package main

import (
	"context"
	"fmt"
	"git.ghtk.vn/gmicro/ig/igrpc-proto/generated/igdata-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

//test ds query
func dsQueryTest(numIterator int, printResult bool){
	fmt.Printf("ds query start %s\n", strings.Repeat("-", 40))
	var reqs []*igrpcproto.SQLQueryRequest
	for i := 0; i < numIterator; i++ {
		request, err := createDSRequest(1000)
		if err != nil {
			log.Fatalf("error when create req err : %s", err.Error())
		}
		reqs = append(reqs, request)
	}
	start := time.Now()
	for _, req := range reqs{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		result, err := IgdataClient.SQLQuery(ctx, req, grpc.WaitForReady(true))
		if err != nil {
			fmt.Printf("grpcsdfjasjkf lksdjkl  %s\n", err.Error())
			errStatus, _ := status.FromError(err)
			//fmt.Printf("can't get result, err : %s code : %d\n", errStatus.Message(), errStatus.Code())
			for _, detail := range errStatus.Details(){
				switch t := detail.(type) {
				case *igrpcproto.ErrorBaseResponse:
					fmt.Println(t)
				default:
					fmt.Println("some thing wrong")
				}
			}
			cancel()
			return
		}
		if printResult{
			if len(result.GetData()) > 100 && result.GetPaginate() != nil{
				fmt.Println(true)
			}
			fmt.Println(string(result.GetData()))
		}
		cancel()
	}
	fmt.Printf("ds query complte\ntotal time : %v\n", time.Since(start))
}
