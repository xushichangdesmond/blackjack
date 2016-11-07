package main

import (
	"context"
	"flag"
	"log"

	"time"

	"math/rand"

	"net/http"
	_ "net/http/pprof"

	"github.com/xushichangdesmond/blackjack"
)

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	rand.Seed(time.Now().Unix())

	r := blackjack.NewSolaireRules()
	t := blackjack.NewDuelingTable(r)
	ctx := context.Background()
	t.Open(ctx)
	//p := blackjack.NewRandomPlayer(r.MinBetUnits)
	p := blackjack.NewBasicStrategyPlayer(r.MinBetUnits)
	err := t.Join(p)
	if err != nil {
		panic(err)
	}

	<-time.After(999 * time.Second)
}
