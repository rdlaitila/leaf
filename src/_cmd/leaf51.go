package main

import(
    "os"
    "log"
    "path/filepath"
    "runtime"
    
    "_lua"
    "_lua/lua51"
    "_leaf/nsleaf"
)

func main() {
    // Attempt to resolve the app directory
    appdir, abserr := filepath.Abs(os.Args[1]); if abserr != nil {
        log.Fatal("APP PATH:", abserr)
    } 
    log.Println("APP PATH:",appdir)
    
    // Set GOMAXPROCS
    log.Println("MAX PROCS:", runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // Createstate and open libs
    state := lua51.Newstate()
    state.Openlibs()
    
    // Push our pcall function
    state.Pushfunction(func(ls lua.State) int {
        log.Fatal("OOB ERROR: "+ls.Tostring(-1))
        return 0 
    })
    
    // Load modules
    state.Pushmodule("leaf", nsleaf.NewModule().Loader)
    
    //Load in a lua chunk
    if loadfileerr := state.Loadfile(appdir+"/main.lua"); loadfileerr != nil {
        log.Fatal("LOAD MAIN:",loadfileerr)
    }
    
    //Call the lua chunk
    pcallerr := state.Pcall(0,0,1); if pcallerr != nil {
        log.Fatal(pcallerr)        
    }
}