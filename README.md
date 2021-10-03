# pokcli
## What is this about
A poker CLI application that pick a winner from 2 player.
```
pokcli 222KQ AAKKQ
```

## Assumptions and limitations
- Cards:
    - Each player gets a hand of 5 cards
    - Available symbols are 23456789TJQKA,in order of value (A being the more valuable)
- Combinations in order of value:
    - Four of a kind, like 77377
    - Fullhouse, means 3 of a kind, and 2 of a kind, in the same hand,like KK2K2
    - Triple, like 32666
    - Two pairs, like 77332
    - A pair, like 43K9K
    - High card, when there’s none of the above, like 297QJ

## How to use
### Manual
1. Clone this repository
2. From the project root directory, run below commands to setup and build the application
```
make setup
make build
``` 
3. After successfully run previous commands, the output binary would be available inside `bin/pokcli`
4. Run the binary
```
bin/pokcli 222KQ AAKKQ
```

### Docker
1. Clone this repository
2. From the project root directory, run below command to build the docker image
```
docker image build -t <image_name> .
```
3. Execute the command from inside docker image container
```
docker run <image_name> 77332 77441
```

### Test scenarios
To run all test scenarios of different kind of input, run below command
```
make test
```
