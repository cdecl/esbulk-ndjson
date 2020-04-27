package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func esConnect(host string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{Addresses: []string{host}}
	es, err := elasticsearch.NewClient(cfg)
	return es, err
}

func esBulk(es *elasticsearch.Client, index string, docs string) (*esapi.Response, error) {
	res, err := es.Bulk(strings.NewReader(docs), es.Bulk.WithIndex(index))
	return res, err
}

func esGetIDValue(js string, fid string) string {
	idval := ""
	if len(fid) > 0 {
		dic := make(map[string]interface{})
		json.Unmarshal([]byte(js), &dic)

		if v, ok := dic[fid].(string); ok {
			idval = string(v)
		}
		if v, ok := dic[fid].(int); ok {
			idval = strconv.Itoa(v)
		}
		if v, ok := dic[fid].(float64); ok {
			idval = strconv.Itoa(int(v))
		}
	}
	return idval
}

func esDoc(js string, fid string) string {
	type ID struct {
		Id string `json:"_id"`
	}
	type NoID struct{}

	type Index struct {
		Index interface{} `json:"index"`
	}

	index := Index{NoID{}}
	idvalue := esGetIDValue(js, fid)
	if len(idvalue) > 0 {
		index = Index{ID{idvalue}}
	}

	meta, _ := json.Marshal(index)
	metastr := string(meta)
	docs := fmt.Sprintf("%s\n%s\n", metastr, js)

	return docs
}

func esInvokeBulk(wg *sync.WaitGroup, es *elasticsearch.Client, index string, docs string, count int) {
	defer wg.Done()
	fmt.Printf("bulk -> %s : %d \n", index, count)

	_, err := esBulk(es, index, docs)
	assertPanic(err)
}

func assertPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type Args struct {
	File  *string
	Host  *string
	Index *string
	Id    *string
	Size  *int
}

func getArgs() (Args, bool) {
	args := Args{}

	args.File = flag.String("f", "", "input json file; nd(newline delimeter) json format (require)")
	args.Host = flag.String("h", "", "elasticsearch host : http://es-host:9200 (require) ")
	args.Index = flag.String("i", "", "index name (require)")
	args.Id = flag.String("id", "", "_id field match json key name (default not set)")
	args.Size = flag.Int("s", 1000, "bulk size")
	flag.Bool("", false, "ver. 200427.0")

	flag.Parse()

	isFlagPassed := func(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		return found
	}

	found := isFlagPassed("f")
	found = found && isFlagPassed("h")
	found = found && isFlagPassed("i")

	if !found {
		flag.Usage()
	}

	return args, found
}

func Run(args Args) {
	buff := ""
	fin, err := os.Open(*args.File)
	assertPanic(err)

	defer fin.Close()
	scanner := bufio.NewScanner(fin)
	// reader := bufio.NewReader(fin)

	es, err := esConnect(*args.Host)
	assertPanic(err)

	wg := sync.WaitGroup{}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		// line, err := reader.ReadString('\n')
		// if err != nil {
		// 	break
		// }

		if len(strings.Trim(line, " ")) == 0 {
			fmt.Println(line)
			continue
		}

		js := esDoc(line, *args.Id)
		buff += js

		if count%*args.Size == 0 && count != 0 {
			wg.Add(1)
			go esInvokeBulk(&wg, es, *args.Index, buff, count)
			buff = ""
		}
		count++
	}

	if len(buff) > 0 {
		wg.Add(1)
		go esInvokeBulk(&wg, es, *args.Index, buff, count)
		buff = ""
	}

	wg.Wait()
	fmt.Printf("bulk insert done -> %s : %d \n", *args.Index, count)
}

func main() {
	args, ok := getArgs()
	if !ok {
		return
	}

	Run(args)
}
