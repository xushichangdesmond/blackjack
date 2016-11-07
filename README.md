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
* red7 running counter
* Customizable shufflers
    - out-of-box we provide
        * perfect shufflers
        * SM126 shuffler (next 16 cards in the shoe/shuffler will not be shuffled)

## Features that the author is working on now
* Different counting strategies
* Player that vary decisions based on count

## Features the author would work on if he has time/money
* A player that makes decisions from human input
* A UI to make a human-playable game out of this simulation
* Record historical results for starting player hand vs dealer upcard combinations (eg 12 v 2, 15 v 9)
* Some 'intelligent stuff' like coming out with strategies and evaluating/simming them
* More casino rules
* Pontoon table and rules

## Running the sim

For a start, try the following to get a detailed view into what is being simmed
```shell
go run sim/main.go -logtostderr -v=50
```
The above uses the default config for the sim, but performs very details logging which slows down the sim by a factor of a few hundred times.

To let the sim run much faster and have it only report the player balance and shoe# and round# periodically, use 
```shell
go run sim/main.go -logtostderr
```

There are some other configuration parameters that you can optionally set to affect the parameters of the sim.
To see a full list of options (and the defaults used), use the -h option 
```shell
go run sim/main.go -h
```

The default config sets the sim up with the below
- rules from a certain partciular real-world casino that uses 4 deck shoe and shuffles when left with one deck
- perfect shuffling 
- player that uses a red7 count
- player sizes bets according to the running count
- uses basic strategy to make play decisions
- reports player balance and other stats every 10000 shoes

[Donate to the author](https://www.paypal.me/powerDancer)