package nsleaf

import(
    "_lua"
)

type Module struct {
}

func NewModule() *Module {
    return &Module{}
}

func (this *Module) Loader(ls lua.State) int { 
    // push module table to stack, this will be returned
    ls.Newtable() 
    
    // Push nsleap.Mutex
    ls.Pushfunction(NewMutex)
    ls.Setfield(-2, "Mutex")
    
    // Push nsleap.WaitGroup
    ls.Pushfunction(NewWaitGroup)
    ls.Setfield(-2, "WaitGroup")
    
    // Push nsleap.Thread
    ls.Pushfunction(NewThread)
    ls.Setfield(-2, "Thread")
    
    return 1
}
