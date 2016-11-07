# A casino-style blackjack simulator

A personal project of the authors' to simulate blackjack in the style that casinos offer them.

## Supported features
* Customizable rules
    - Double down on (any, 9-11, 10-11, not allowed)
    - Surrender on dealer card (any, not Ace, not allowed)
    - Max splits(non-ace) per hand
    - Max Ace-splits per hand
    - Can hit after ace-split
    - Double After splits (or not)
    - Dealer Hits on Soft 17 (or not)
    - Blackjack Payout (default 1.5)
    - Early Surrender (Late surrender to be implemented)
    - Lose only original bet on dealer BJ (or lose all bets)
    - Minimum betting units
    - Maximum betting units
    - Bet unit size
    - Number of decks in shoe
    - Shuffling point (in number of cards)
    - Out-of-box we will provide some rulesets of various real-world casinos
* Dueling Table (1 player box vs 1 dealer)
    - other table types to be implemented
* Customizable Player (playing strategy)
    - out-of-box we provide
        * random player
        * basic strategy player
        * player that sizes bets according to running count
* Player balance (simple money account)
* HiLo running counter
* CSM running counter (keeps counts of previous 2 rounds)
* Customizable shufflers
    - out-of-box we provide
        * perfect shufflers
        * SM126 shuffler (next 16 cards in the shoe/shuffler will not be shuffled)

## Features that the author is working on now
* Different counting strategies
* Player that vary decisions based on count
* Make it faster (99.5% of the cpu spent is actually because of logging)

## Features the author would work on if he has time/money
* A player that makes decisions from human input
* A UI to make a human-playable game out of this simulation
* Record historical results for starting player hand vs dealer upcard combinations (eg 12 v 2, 15 v 9)
* Some 'intelligent stuff' like coming out with strategies and evaluating/simming them
* More casino rules
* Pontoon table and rules

## Running the sim

There is (yet) to be much customizations for the sim thru command line options.
You will need to get your hands wet and edit sim/main.go to customize the sim for now.

Out-of-box, you can run the following command
```shell
go run sim/main.go -logtostderr -v=50
```

It is preconfigured with below
- rules from a certain partciular real-world casino
- CSM Shuffling (emulate some props from a popular CSM model) 
- player that uses a simple hilo counting strategy adapted for counting CSM
- player also sizes bets according to the running count

Note that logging makes the sim much much much slower. If you want to the sim to run faster, then 
```shell
go run sim/main.go -logtostderr -v=40
```

[Donate to the author](https://www.paypal.me/powerDancer)