

## esbulk-ndjson
ElasticSearch Bulk from ndjson file

```
$ bin\esbulk.exe
Usage of bin\esbulk.exe:
  -f string
        input json file; nd(newline delimeter) json format (require)
  -h string
        elasticsearch host : http://es-host:9200 (require)
  -i string
        index name (require)
  -id string
        _id field match json key name (default not set)
  -s int
        bulk size (default 1000)
```


### Build 

```
## package get
$ make get 

## build
$ make 
```