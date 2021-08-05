package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"tungnguyen/grpc_proto"
)

//test ds query
func dsQueryTest(numIterator int, printResult bool){
	fmt.Printf("ds query start %s\n", strings.Repeat("-", 40))
	var reqs []*igrpcproto.SQLQueryRequest
	for i := 0; i < numIterator; i++ {
		request, err := createDSRequest(20)
		if err != nil {
			log.Fatalf("error when create req err : %s", err.Error())
		}
		reqs = append(reqs, request)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()
	start := time.Now()
	for _, req := range reqs{
		result, err := IgdataClient.SQLQuery(context.Background(), req)
		if err != nil {
			fmt.Printf("can't get result, err : %s\n", err.Error())
			return
		}
		if printResult{
			fmt.Println(string(result.GetData()))
		}
	}
	fmt.Printf("ds query complte\ntotal time : %v\n", time.Now().Sub(start))
}
