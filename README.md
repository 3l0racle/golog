# GOLOG

golog is a lightweight package that allows logging of requests and errors into a file
The http Logger takes in the filename permision and a handler
initialize a logger and call start


## HTTP request loging
> Log all your requests

 _NOTE_ *This has not beeen tested on https*

```golang
package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/3l0racle/golog"
)

func main(){
  ReqtLog := golog.RequestLogger{
    FileName:"requests",
    Dir:"./.data/logs/",
    Perm:0666,
    Handle:http.DefaultServeMux,
  }
  http.HandleFunc("/",RootHandler)
  err := http.ListenAndServe(fmt.Sprintf(":%d", 3000),golog.Start(ReqtLog))
  if err != nil{
    log.Fatal(err)
  }
}


func RootHandler(res http.ResponseWriter, req *http.Request){
  fmt.Fprintf(res, "<h1>Hello World</h1></br><div>Welcome to whereever you are</div>")
}

```

## Error Loging mechanisms

*Later on will update log.Fatal and panic plus panic recovery*

> Log error to file in this specified directory

```golang
package main

import (
  "fmt"
  "errors"
  "github.com/3l0racle/golog"
)

func main(){
  err := errors.New("Test error logger final")
  // LogErrorToFileInDir("NameOfTheErroFile","Directory","Error or fmt.Sprintf(Error of randomparameter %s %s\n,randPrameter,err)")
  golog.LogErrorToFileInDir("test","./.data/logs/",err)
  randPrameter := "This is that random thing I was saying"
  golog.LogErrorToFileInDir("test","./.data/logs/",fmt.Sprintf("Error of %s plus error %s",randPrameter,err))

  //log error to file in the current directory
  golog.LogErrorToFile("test",0666,errors.New("this is an error"))
}
```

> Log error universally
> *PS :) I honestly don know why i created this it flipped my mind*

``` golang
package main

import (
  "errors"
  "github.com/3l0racle/golog"
)

func main(){
  err := errors.New("Test error logger")
  logggerr := golog.ErrorLogger{
    Filename: "test",
    Dir: "./.data/logs/",
    Perm:0666,
    Text:err,
  }
  logggerr.UniversalLog()
}

```
