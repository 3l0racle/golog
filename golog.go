/*
  *Golog logs
  *   >  log an error into a file
  *   > all requests and responses (if enabled)
  *
*/

package golog

import (
  "io"
  "io/fs"
  "os"
  "log"
  "net/http"
  "path/filepath"
)


type ErrorLogger struct{
  Filename string
  Dir string
  Perm fs.FileMode
  Text interface{}
}

type RequestLogger struct{
  FileName string
  Dir string
  Perm fs.FileMode
  Handle http.Handler
}

//loggs to a perticular directory
func (el ErrorLogger) UniversalLog(){
  dir := filepath.Clean(el.Dir)
  if dir != "" || len(dir) <= 0{
    name := el.Filename + ".log"
    name = filepath.Join(dir,name)
    f, err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,el.Perm)
    if err != nil {
      logError(err)
    }
    defer f.Close()
    writer := io.MultiWriter(os.Stdout,f)
    log.SetOutput(writer)
    log.Println(el.Text)
  }
}

//logs the error to a file in the current diretory
func LogErrorToFile(name string,perm os.FileMode,text ...interface{}) {
  name = name + ".log"
  f,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,perm)
  if err != nil {
    logError(err)
  }
  defer f.Close()
  writer := io.MultiWriter(os.Stdout,f)
  log.SetOutput(writer)
  log.Println(text)
}


//Log error to file in a different directory
func LogErrorToFileInDir(name,dir string,text ...interface{}) {
  dir = filepath.Clean(dir)
  if dir != "" || len(dir) <= 0{
    name = name + ".log"
    name = filepath.Join(dir,name)
    f,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
      logError(err)
    }
    defer f.Close()
    writer := io.MultiWriter(os.Stdout,f)
    log.SetOutput(writer)
    log.Println(text)
  }
}

func logError(e error){
  if e != nil{
    panic(e)
  }
}


func Start(l RequestLogger) http.Handler{
  //open file for loging
  l.OpenLogFile()
  //set the flags
  log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
  //return the handler
  return l.LogRequest(l.Handle)
}

//opens or creates a file for logging
// To be addedfile permisions
func (l RequestLogger) OpenLogFile(){
  //clean the directory
  dir := filepath.Clean(l.Dir)
  if dir != "" || len(l.Dir) != 0{
    if l.FileName != "" && len(l.FileName) >= 0{
      name := l.FileName + ".log"
      name = filepath.Join(l.Dir,name)
      dataLog,err := os.OpenFile(name,os.O_WRONLY|os.O_CREATE|os.O_APPEND,l.Perm)
      if err != nil{
        log.Fatal("[-] Error logging to file: ",err)
      }
      log.SetOutput(dataLog)
    }
  }
}


func (l RequestLogger)LogRequest(handler http.Handler) http.Handler{
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
    log.Printf("%s %s %s\n",req.RemoteAddr,req.Method,req.URL)
    handler.ServeHTTP(res,req)
  })
}
