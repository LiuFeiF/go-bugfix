#开启vpn后github提交代码一九会出现fatal: unable to access 'https://github.com/xxx/xxx.git/': Failed to connect to github.com port 443 after 21077 ms: Couldn't connect to server


解决方案：
#设置git本地代理，相关命令
### git config --global http.proxy "127.0.0.1:xxx"
### git config --global https.proxy "127.0.0.1:xxx"

输入命令：git config --global -l 会出现以下参数：http.proxy=127.0.0.1:xxx 即可

