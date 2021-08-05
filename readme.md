# igdata test grpc

To test grpc service :
```bash
go run .
```

Flags options :

```bash
service int[,int...]
to specify what service to test, if don't provide, will run all service

options:
  1 - SQLQuery
  2 - SQLStreamQuery
  3 - ESQuery
  4 - ESStreamQuery
 
example usage :
to run service 1 and 2:
  go run . -service=1,2
```
```bash
num int
to set number of iterator for each run.
default 100, min = 1, max =10000
```
```bash
print bool
to print result of query or not
default False
```
