# go-cipher

Simple SSL cipher check for an input URL.

## Use

Build the executable:

```
go build go-cipher.go
```

Run and input a URL:

```
./go-cipher
Enter a URL: https://www.google.com
ECDHE-RSA-AES256-GCM-SHA384
ECDHE-RSA-AES128-GCM-SHA256
ECDHE-RSA-AES256-SHA384
ECDHE-RSA-AES128-SHA256
AES256-GCM-SHA384
AES128-GCM-SHA256
AES256-SHA256
AES128-SHA256
RC4-SHA
```
