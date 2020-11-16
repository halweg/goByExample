package main

import "context"

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {

	}()

}