package nsleaf

import(
    "log"
    "sync"
    
    "_lua"
    //"_lua/lua51"
    //"_lua/lua52"
    //"_lua/lua53"
)

type Mutex struct {
    Ticket chan int
    grindex int
    mu *sync.Mutex
}

func NewMutex(ls lua.State) int {  
    grindex := lua.GlobalRegistry.ReserveValue()  
    mu := &Mutex{
        Ticket: make(chan int, 1),
        mu: &sync.Mutex{},
        grindex: grindex,
    }
    mu.Ticket <- 1    
    lua.GlobalRegistry.SetValue(grindex, mu)
    
    // Create new userdata. This will be returned
    ls.Newtable()
    
    // Push mutex.Lock
    ls.Pushfunction(mu.lock)
    ls.Setfield(-2, "lock")
    
    // Push mutex.Unlock
    ls.Pushfunction(mu.unlock)
    ls.Setfield(-2, "unlock")
    
    // Create metatable for userdat
    ls.Pushmetatable(&lua.Metatable{
        GCFunc: mu.gc,
    })    
    
    // Set metatable for userdata
    ls.Setmetatable(-2)
    
    return 1
}

func (this *Mutex) gc(ls lua.State) int {
    log.Println("MUTEX GC")
    return 0
}

func (this *Mutex) lock(ls lua.State) int {    
    <- this.Ticket    
    
    return 0
}

func (this *Mutex) unlock(ls lua.State) int {    
    this.Ticket <- 1
    
    return 0
}