package leaf

import(
    "log"
    "runtime"

    "_lua"
    "_leaf/nsleaf"
    "_leaf/nsruntime"
)

type Runtime struct {
    ls lua.State
}

func NewRuntime(ls lua.State) *Runtime {
    return &Runtime{ls: ls,}
}

func (this *Runtime) Start(apppath string) {    
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    this.ls.Openlibs()
    
    // Push our pcall function
    this.ls.Pushfunction(func(ls lua.State) int {
        log.Println("ERROR: "+ls.Tostring(-1))
        ls.Dostring(`print(debug.traceback())`)
        log.Fatal("Exiting")
        return 0 
    })
    
    // Load modules
    this.ls.Pushmodule("leaf", nsleaf.NewModule().Loader)
    this.ls.Pushmodule("leaf.runtime", nsruntime.NewModule().Loader)
    
    //Load in a lua chunk
    if loadfileerr := this.ls.Loadfile(apppath+"/main.lua"); loadfileerr != nil {
        log.Fatal("LOAD MAIN:",loadfileerr)
    }
    
    //Call the lua chunk
    pcallerr := this.ls.Pcall(0,0,1); if pcallerr != nil {
        log.Fatal(pcallerr)        
    }
}