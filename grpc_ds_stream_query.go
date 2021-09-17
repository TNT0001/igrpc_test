package main

import (
	"context"
	"fmt"
	"git.ghtk.vn/gmicro/ig/igrpc-proto/generated/igdata-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

func dsStreamQueryTest(numIterator int, printResult bool){
	var reqs []*igrpcproto.SQLQueryRequest
	for i := 0; i < numIterator; i++ {
		request, err := createDSRequest(1000)
		if err != nil {
			log.Fatalf("error when create req err : %s", err.Error())
		}
		reqs = append(reqs, request)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()
	fmt.Printf("ds stream query start %s\n", strings.Repeat("-", 40))
	start := time.Now()
	stream, err := IgdataClient.SQLQueryStream(context.Background(), grpc.WaitForReady(true))
	if err != nil {
		fmt.Printf("can't get stream err : %s", err)
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
				errStatus, _ := status.FromError(err)
				fmt.Printf("can't get result, err : %s code : %d\n", errStatus.Message(), errStatus.Code())
			}
			if err != nil {
				fmt.Printf("error when recieve %s\n", err.Error())
				return
			}
			if printResult{
				fmt.Println(string(result.GetData()))
			}
		}
	}()
	wg.Wait()
	fmt.Printf("ds stream query complete\n total time : %v\n", time.Since(start))
}
