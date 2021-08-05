package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
	"tungnguyen/grpc_proto"
)

func dsStreamQueryTest(numIterator int, printResult bool){
	var reqs []*igrpcproto.SQLQueryRequest
	for i := 0; i < numIterator; i++ {
		request, err := createDSRequest(20)
		if err != nil {
			log.Fatalf("error when create req err : %s", err.Error())
		}
		reqs = append(reqs, request)
	}

	fmt.Printf("ds stream query start %s\n", strings.Repeat("-", 40))
	start := time.Now()
	stream, err := IgdataClient.SQLQueryStream(context.Background())
	if err != nil {
		fmt.Println("can't get stream")
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, item := range reqs{
			err := stream.Send(item)
			if err != nil {
				fmt.Println("send fail")
				return
			}
		}
		err := stream.CloseSend()
		if err != nil {
			fmt.Println("close send error")
			return
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("get eof")
				return
			}
			if err != nil {
				fmt.Printf("error when recieve %s", err.Error())
				return
			}
			if printResult{
				fmt.Println(string(result.GetData()))
			}
		}
	}()
	wg.Wait()
	fmt.Printf("ds stream query complete\n total time : %v\n", time.Now().Sub(start))
}
