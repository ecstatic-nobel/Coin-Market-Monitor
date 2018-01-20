# [coin_market_monitor]  
##### Simple terminal-based monitor of Coin Market Cap written in Go.  

#### Description  
This project is used to monitor coins from https://coinmarketcap.com/.  

#### Prerequisites  
- Git   
- Go  

#### Setup  
Open a terminal and run the following commands:  
```bash  
go get github.com/olekukonko/tablewriter  
mkdir ~/leunammejii  
cd ~/leunammejii  
git clone https://github.com/leunammejii/simple_market_monitor.git  
cd simple_market_monitor  
go run coin.go  
```  

You will get the following output:  

![alt text](https://github.com/leunammejii/coin_market_monitor/blob/master/monitor.png)  

To monitor different coins, change the contents of `coins.txt` (or any file containing a list of  
coin names) and run the following command:  
```bash
go run coin.go -f coins.txt  
```

#### Destroy  
To remove the project completely,  run the following commands:  
```bash  
rm -rf ~/leunammejii/simple_market_monitor  
```  

#### Limitations  
CoinMarketCap currently asks that requests to the API are limited to ten (10) per minute.  
If there are more that ten (10) coins in the specified file, the script will only print ten (10)  
coins to the screen every sixty (60) seconds.  

Also, make sure that the coin name is spelled the same as on `https://coinmarketcap.com/currencies/COIN_NAME/#tools`  
under `API Ticker`.  
