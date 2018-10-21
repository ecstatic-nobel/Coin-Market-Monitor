# [coin_market_monitor]  
##### Simple terminal-based monitor of Coin Market Cap written in Go.  

#### Description  
This project is used to monitor coins listed on https://coinmarketcap.com/.  

#### Prerequisites  
- Git   
- Go  

#### Setup  
Open a terminal and run the following commands:  
```bash
git clone https://github.com/leunammejii/simple_market_monitor.git
cd simple_market_monitor
```

then run:  
```bash
chmod u+x cmm
./cmm
```

or:  
```bash
go get github.com/olekukonko/tablewriter
go run cmm.go
```

You will get the following output:  

![alt text](https://github.com/leunammejii/coin_market_monitor/blob/master/monitor.png)  

To monitor different coins, change the contents of `coins.txt` (or any file containing a list of coin names) and run the following command:  
```bash
./cmm -f coins.txt
```

#### Destroy  
To remove the project completely,  run the following commands:  
```bash
rm -rf simple_market_monitor
```

#### Limitations  
CoinMarketCap currently asks that requests to the API are limited to 10 per minute. If there are more that 10 coins in the specified file, the script will only print 10 coins to the screen every 65 seconds. The site updates every 5 minutes.  

Also, make sure that the coin names listed in the file are spelled the same as the `id`.  The top 100 coins and their `id` can be found at `https://api.coinmarketcap.com/v1/ticker/`. You may be able to find other names at `https://coinmarketcap.com/currencies/COIN_NAME/#tools` under `API Ticker`.  
