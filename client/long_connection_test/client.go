/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	pb "grpc-example/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	var wg sync.WaitGroup
	concurrencyNumber := 1000
	num := 1
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		if num <= concurrencyNumber {
			wg.Add(1)
			go func(num *int) {
				defer wg.Done()
				for {
					call(num, conn)
				}
			}(&num)
			if num%20 == 0 {
				time.Sleep(time.Second * 10)
			}
		} else {
			break
		}
		num++
	}
	wg.Wait()
}

func call(num *int, conn *grpc.ClientConn) {
	// Set up a connection to the server.
	var responseTime time.Time
	var responseDuration = int64(0)
	var logError error
	defer func() {
		fmt.Printf("%d	%d	%v\n", *num, responseDuration, logError)
	}()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	start := time.Now()
	_, logError = c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if logError != nil {
		return
	}
	responseTime = time.Now()
	responseDuration = (responseTime.UnixNano() - start.UnixNano()) / 1000000
}
