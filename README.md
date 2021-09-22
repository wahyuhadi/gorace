# Race condition tester

```sh
gorace -f http.req -stcode 200 -tr=400
```

```Usage of gorace:
  -f string
        http request file 
  -stcode int
        Expect status code (default 200)
  -tr int
        total  request   (default 100)
```

* http.req is http raw request from burp

```sh
GET /v1/buckets/main/collections/ms-language-packs/records/cfr-v1-en-US HTTP/1.1
Host: firefox.settings.services.mozilla.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:92.0) Gecko/20100101 Firefox/92.0
Accept: application/json
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Content-Type: application/json
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
Te: trailers
Connection: close


```
