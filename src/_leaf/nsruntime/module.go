package nsruntime

import(
    "runtime"
    
    "_lua"
)

type Module struct {
}

func NewModule() *Module {
    return &Module{}
}

func (this *Module) Loader(ls lua.State) int {
    //Module table
    ls.Newtable()
    
    //Schedule function
    ls.Pushfunction(this.schedule)
    ls.Setfield(-2, "schedule")
    
    //numcpu function
    ls.Pushfunction(this.numcpu)
    ls.Setfield(-2, "numcpu")
    
    //numthread function
    ls.Pushfunction(this.numthread)
    ls.Setfield(-2, "numthread")
    
    //arch property
    ls.Pushstring(runtime.GOARCH)
    ls.Setfield(-2, "arch")
    
    //os property
    ls.Pushstring(runtime.GOOS)
    ls.Setfield(-2, "os")
    
    return 1
}

func (this *Module) schedule(ls lua.State) int {
    runtime.Gosched()
    return 0
}

func (this *Module) numcpu(ls lua.State) int {
    ls.Pushnumber(float64(runtime.NumCPU()))
    return 1
}

func (this *Module) numthread(ls lua.State) int {
    ls.Pushnumber(float64(runtime.NumGoroutine()))
    return 1
}