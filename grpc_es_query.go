package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"git.ghtk.vn/gmicro/ig/igrpc-proto/generated/igdata-service"
)

//test es query
func esQueryTest (numIterator int, printResult bool){
	fmt.Printf("es query start %s\n", strings.Repeat("-", 40))
	var reqs []*igrpcproto.ESQuery
	for i := 0; i < numIterator; i++ {
		request, err := createESRequest(100)
		if err != nil {
			log.Fatalf("error when create req err : %s", err.Error())
		}
		reqs = append(reqs, request)
	}

	start := time.Now()
	ctx := context.Background()
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for i := 0; i < numIterator; i++ {
		esResult, err := IgdataClient.EsQuery(newCtx, reqs[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if printResult{
			fmt.Println(string(esResult.GetResponse()))
		}
	}
	fmt.Printf("es query complte\n total time : %v\n", time.Now().Sub(start))
}
