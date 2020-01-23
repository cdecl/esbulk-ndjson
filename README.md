

## esbulk-ndjson
ElasticSearch Bulk from ndjson file

```
$ esbulk-ndjson.exe
Usage of esbulk-ndjson.exe:
  -f string
        json file
  -i string
        index name
  -id string
        _id field in json (default not set)
  -n int
        bulk size (default 1000)
  -s string
        elasticsearch server : http://es-server

```


### Build 

```
## package get
$ make get 

## build
$ make 
```