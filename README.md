# Findit

Findit is a tool used to parse a canary value from a web request.  
This is useful in bugbounties when you want to find a return value from a list of URL's.  
Think xss, ssti, sqli etc. 


## Install 

```
go install github.com/AssassinUKG/findit@latest
```

## Usage

```
cat URLFile.txt | findit -c "alert(0)"
```
