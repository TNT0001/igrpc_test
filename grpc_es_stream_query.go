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

func esStreamQueryTest(numIterator int, printResult bool){
	fmt.Printf("es stream query start %s\n", strings.Repeat("-", 40))
	var reqs []*igrpcproto.ESQuery
	for i := 0; i < numIterator; i++ {
		request, err := createESRequest(100)
		if err != nil {
			log.Fatalf("error when create req err : %s\n", err.Error())
		}
		reqs = append(reqs, request)
	}


	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := IgdataClient.EsQueryStream(ctx)
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
			esResult, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("get eof")
				return
			}
			if err != nil {
				fmt.Printf("error when recieve %s\n", err.Error())
				return
			}
			if printResult{
				fmt.Println(string(esResult.GetResponse()))
			}
		}
	}()
	wg.Wait()
	fmt.Printf("es stream query complte\n total time : %v\n", time.Now().Sub(start))
}
