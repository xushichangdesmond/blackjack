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
* Dueling Table (1 player box vs 1 dealer)
    - other table types to be implemented
* Customizable Player (playing strategy)
    - out-of-box we provide
        * random player
        * basic strategy player
        * counting player is being implemented
* Player balance (simple money account)

## Features that the author is working on now
* Customizable counting strategies
* Counting strategy player
* Make it faster
* Use different loggers for different aspects - player performance, etc
    - different loggers to go to different output files (or blackhole or stderr)

## Features the author would work on if he has time/money
* A player that makes decisions from human input
* A UI to make a human-playable game out of this simulation
* Record historical results for starting player hand vs dealer upcard combinations (eg 12 v 2, 15 v 9)
* Make it faster
* More rules
* Pontoon table and rules

## Running the sim

There is (yet) to be much customizations for the sim thru command line options.
You will need to get your hands wet and edit sim/main.go to customize the sim for now.

Out-of-box, you can run a sim with Manila Solaire rules and a basic strategy player with
```shell
go run sim/main.go -logtostderr
```

[Donate to the author](https://www.paypal.me/powerDancer)