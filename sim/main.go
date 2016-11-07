package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"time"

	"math/rand"

	"net/http"
	_ "net/http/pprof"

	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/xushichangdesmond/blackjack"
)

var ruleSets = []string{
	"solaire", "cara",
}

var counters = []string{
	"red7Running", "hiloRunning", "csmRunning", "none",
}

var shufflers = []string{
	"perfect", "sm126",
}

var playingStrategies = []string{
	"basicStrategy",
}

var (
	ruleSet         = flag.String("ruleSets", "cara", fmt.Sprintf("one of %v", ruleSets))
	counter         = flag.String("counters", "red7Running", fmt.Sprintf("one of %v", counters))
	shuffler        = flag.String("shufflers", "perfect", fmt.Sprintf("one of %v", shufflers))
	playingStrategy = flag.String("playingStrategies", "basicStrategy", fmt.Sprintf("one of %v", playingStrategies))
	countToBetSize  = flag.String("countToBetSize", "3,3,3,3,3,3,3,3,20,30,40,50,60", "comma separated integer list. The nth number in the list is the number of unit bets to bet when the count is n. If the count is more than the length of the provided list, the last given number is used")
	report          = flag.Int("report", 10000, "Report player balance after each set of N shoes")
	httpHostPort    = flag.String("httpHostPort", "localhost:6060", "host:port to run http server on (for pprof)")
)

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe(*httpHostPort, nil))
	}()

	rand.Seed(time.Now().Unix())

	var r *blackjack.Rules
	switch *ruleSet {
	case "solaire":
		r = blackjack.NewSolaireRules()
	case "cara":
		r = blackjack.NewCaraRules()
	default:
		panic("Invalid ruleSet - " + *ruleSet)
	}

	if glog.V(0) {
		glog.Infof("Rules %##v\n", r)
	}

	var shf blackjack.Shuffler
	switch *shuffler {
	case "perfect":
		if glog.V(0) {
			glog.Infoln("Using perfect shuffler")
		}
		shf = blackjack.PerfectShuffler
	case "sm126":
		if glog.V(0) {
			glog.Infoln("Using ShuffleMaster 126")
		}
		shf = blackjack.ShuffleMaster126
	default:
		panic("Invalid shuffler - " + *shuffler)
	}

	t := blackjack.NewDuelingTable(r, shf)

	ctx := context.Background()
	t.Open(ctx)

	var ctr blackjack.RunningCounter
	switch *counter {
	case "red7Running":
		c, cl, sl := blackjack.NewRed7RunningCounter()
		ctr = c
		t.SubscribeCardListener(cl)
		t.SubscribeShuffleListener(sl)
	case "hiloRunning":
		c, cl := blackjack.NewHiLoRunningCounter()
		ctr = c
		t.SubscribeCardListener(cl)
	case "csmRunning":
		c, cl, sl := blackjack.NewCSMCounter()
		ctr = c
		t.SubscribeCardListener(cl)
		t.SubscribeShuffleListener(sl)
	case "none":
		ctr = blackjack.NewBlackHoleCounter()
	default:
		panic("Invalid counter - " + *counter)
	}

	var p blackjack.Player
	switch *playingStrategy {
	case "basicStrategy":
		p = blackjack.NewBasicStrategyPlayer(r.MinBetUnits)
	default:
		panic("Invalid playingStrategy - " + *playingStrategy)
	}

	c2bs_string := strings.Split(*countToBetSize, ",")
	c2bs_uint := make([]uint, len(c2bs_string))

	for i, c2bs := range c2bs_string {
		u, err := strconv.ParseUint(c2bs, 10, 0)
		if err != nil {
			panic("Cannot convert '" + c2bs + "' to a uint")
		}
		c2bs_uint[i] = uint(u)
	}

	p = blackjack.NewBetSizeOnRunningCountPlayer(
		ctr,
		p,
		c2bs_uint,
	)

	var shufflesSinceLastReport int
	t.SubscribeShuffleListener(func() {
		shufflesSinceLastReport++
		if shufflesSinceLastReport == *report {
			if glog.V(0) {
				shoesPlayed, roundsPlayed := t.Info()
				glog.Infoln("Shuffling after shoe #", shoesPlayed, "; rounds so far=", roundsPlayed, "; player balance=", p.Balance())
			}
			shufflesSinceLastReport = 0
		}
	})

	err := t.Join(p)
	if err != nil {
		panic(err)
	}

	<-time.After(99999 * time.Hour)
}
