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
	t := blackjack.NewDuelingTable(r, blackjack.ShuffleMaster126)

	ctx := context.Background()
	t.Open(ctx)

	//p := blackjack.NewRandomPlayer(r.MinBetUnits)
	//p := blackjack.NewBasicStrategyPlayer(r.MinBetUnits)

	csmC, cl, sl := blackjack.NewCSMCounter()
	p := blackjack.NewBetSizeOnRunningCountPlayer(
		csmC,
		blackjack.NewBasicStrategyPlayer(r.MinBetUnits),
		[]uint{1, 1, 1, 1, 1, 40, 50},
	)

	t.SubscribeCardListener(cl)
	t.SubscribeShuffleListener(sl)

	err := t.Join(p)
	if err != nil {
		panic(err)
	}

	<-time.After(999 * time.Second)
}
