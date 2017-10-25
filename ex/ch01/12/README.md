```bash
$ apt-get install unzip
```

```
$ curl http://s3.amazonaws.com/alexa-static/top-1m.csv.zip | funzip | head -n 50 | awk -F "," '{print $2}' | go run gopl/ex/01.12.go

0.19s    3154 http://t.co
0.40s   11776 http://google.es
0.43s    1064 http://reddit.com
0.45s   11515 http://google.fr
0.45s   11588 http://google.it
0.45s   11078 http://google.ca
0.47s   12246 http://google.ru
0.48s   14313 http://google.co.in
0.54s   14552 http://twitch.tv
0.57s   12196 http://google.com.hk
0.59s   86413 http://wikipedia.org
0.73s  220944 http://facebook.com
0.75s  245356 http://qq.com
0.96s  406829 http://yahoo.com
1.20s  132770 http://imgur.com
1.31s   70400 http://yandex.ru
1.32s   55957 http://netflix.com
1.34s  328111 http://pornhub.com
1.71s   20115 http://yahoo.co.jp
2.18s   21451 http://alipay.com
2.56s  231017 http://taobao.com
Get http://googleusercontent.com: dial tcp: lookup googleusercontent.com on 127.0.0.11:53: no such host
5.22s   36480 http://diply.com
5.32s  157368 http://cnn.com
5.36s  109028 http://bing.com
5.37s   11178 http://google.com
5.40s   10735 http://google.co.jp
5.42s   11536 http://google.de
5.43s  101598 http://adobe.com
5.47s   11435 http://google.co.uk
5.47s   10945 http://google.com.br
5.48s   38263 http://msn.com
5.48s   11588 http://google.com.mx
5.48s   57801 http://wordpress.com
5.50s   11446 http://google.com.au
5.49s   11362 http://google.com.ua
5.49s   11498 http://google.com.pk
5.53s   13689 http://instagram.com
5.54s   10424 http://google.com.tw
5.55s   46024 http://apple.com
5.54s   11879 http://google.pl
5.56s   10873 http://google.com.tr
5.57s  248966 http://stackoverflow.com
5.56s   11538 http://google.com.ar
5.58s   36530 http://whatsapp.com
5.61s   72469 http://tumblr.com
5.61s  211773 http://wikia.com
5.61s   55452 http://aliexpress.com
5.62s   14038 http://popads.net
5.64s   50692 http://github.com
5.66s  137493 http://uptodown.com
5.67s  451598 http://espn.com
5.71s  115617 http://savefrom.net
5.73s   12854 http://google.co.id
5.76s    3295 http://onclkds.com
5.76s  165498 http://imdb.com
5.79s   43877 http://linkedin.com
5.86s   42041 http://twitter.com
5.89s   15712 http://live.com
5.90s   51640 http://office.com
5.94s  202203 http://sohu.com
5.94s   94591 http://blogspot.com
5.98s   46494 http://coccoc.com
6.00s  361712 http://amazon.com
6.05s   48762 http://microsoftonline.com
6.07s  232066 http://ebay.com
6.05s   62498 http://xnxx.com
6.16s  600898 http://sina.com.cn
6.15s   71669 http://xhamster.com
6.18s  371224 http://amazon.co.jp
6.20s  198651 http://livejasmin.com
6.20s  112100 http://nicovideo.jp
6.22s   96426 http://xvideos.com
6.22s   83981 http://pinterest.com
6.25s  191250 http://amazon.in
6.28s  303629 http://txxx.com
6.31s  227683 http://bongacams.com
6.33s  205459 http://amazon.de
6.36s  198840 http://amazon.co.uk
6.46s   72969 http://youth.cn
6.48s  161329 http://microsoft.com
6.47s    7865 http://tianya.cn
6.55s  119320 http://dropbox.com
6.66s   32670 http://paypal.com
6.67s  293688 http://mail.ru
6.72s  132616 http://ok.ru
6.75s  149739 http://jd.com
6.78s   45032 http://weibo.com
6.84s    6494 http://vk.com
6.84s  491273 http://youtube.com
6.93s   98923 http://csdn.net
7.33s  149779 http://360.cn
7.56s  232304 http://tmall.com
7.78s  120579 http://pixnet.net
8.22s  306875 http://so.com
8.37s  232304 http://list.tmall.com
10.04s  463528 http://hao123.com
10.86s      81 http://baidu.com
Get https://detail.tmall.com/?tbpm=3: stopped after 10 redirects
65.22s      92 http://gmw.cn
65.25s elapsed
