package main

import(
    "os"
    "log"
    "path/filepath"
    
    "_leaf"
    "_lua/lua52"
)

func main() {
    appdir, abserr := filepath.Abs(os.Args[1]); if abserr != nil {
        log.Fatal("APP PATH:", abserr)
    } 
   
    state := lua52.Newstate()
    
    leaf.NewRuntime(state).Start(appdir)
}