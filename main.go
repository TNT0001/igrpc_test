package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

var (
	service = flag.String("service", "", "provide service number to run")
	num = flag.Int("num", 100, "num of iterator for each service")
	printAble = flag.Bool("print", false, "if true, print response of each rpc")

	serviceList = []func(int, bool){dsQueryTest, dsStreamQueryTest, esQueryTest, esStreamQueryTest}
)

func main() {
	flag.Parse()

	var serviceNums []int
	if len(*service) == 0 {
		serviceNums = []int{0, 1, 2, 3}
	} else{
		services := strings.Split(*service, ",")
		for _, i := range services {
			j, err := strconv.Atoi(i)
			if err != nil || j < 0 || j > len(serviceList) {
				log.Fatalf("service number must be int, gt than -1 and lt than %d", len(serviceList))
			}
			serviceNums = append(serviceNums, j)
		}
	}

	if *num < 0 {
		*num = 1
	}

	if *num > 1000 {
		*num = 1000
	}

	for _, i := range serviceNums {
		serviceList[i](*num, *printAble)
	}
}

