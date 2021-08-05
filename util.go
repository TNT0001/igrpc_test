package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"tungnguyen/grpc_proto"
)

func createDSRequest (maxSize int64) (*igrpcproto.SQLQueryRequest, error){
	jsFile, err := os.Open("dsReq.json")
	if err != nil {
		return nil, err
	}

	req := &igrpcproto.SQLQueryRequest{}
	err = json.NewDecoder(jsFile).Decode(req)
	if err != nil {
		return nil, err
	}

	limit := 1 + rand.Int63n(10)
	page := rand.Int63n(maxSize/limit)
	req.Paginate.Limit = limit
	req.Paginate.Page = page
	//b, err := json.Marshal([]int{900000,1999999,14324324})
	//if err != nil {
	//	log.Println("can't get json marshal")
	//}
	//req.Conjunctions[0].Conditions[0].Value = b

	return req, nil
}

func createESRequest (maxSize int64) (*igrpcproto.ESQuery, error){
	limit := 1 + rand.Int63n(20)
	page := rand.Int63n(maxSize/limit)
	body, err := json.Marshal(map[string]interface{}{
		"from": page,
		"size": limit,
	})
	if err != nil {
		log.Println("can't create request body")
		return nil, err
	}
	esRequest := &igrpcproto.ESQuery{
		ConnectionName: "report",
		IndexName:      "packages",
		Body:           body,
	}
	return esRequest, nil
}
