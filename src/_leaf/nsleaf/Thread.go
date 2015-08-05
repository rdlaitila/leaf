package nsleaf

import(
    "log"
    "sync"
    
    "_lua"
    //"_lua/lua51"
    //"_lua/lua52"
    //"_lua/lua53"
    
    "code.google.com/p/go-uuid/uuid"
)

type Thread struct {
    mu *sync.Mutex
    grindex int
}

func NewThread(ls lua.State) int {        
    if ls.Gettop() < 1 {
        ls.Pushstring("You must supply a function to leap.Thread() constructor")
        ls.Error()
    }
    
    grindex := lua.GlobalRegistry.ReserveValue()
    thread := &Thread{
        mu: &sync.Mutex{},
        grindex: grindex,
    }
    lua.GlobalRegistry.SetValue(grindex, thread)
    
    // Create new table. This will be returned
    ls.Newtable()
    
    // Swap arg with table, arg should be at tos
    ls.Insert(ls.Gettop()-1)    
    ls.Setfield(-2, "func")
    
    // Push func
    ls.Pushfunction(thread.run)
    ls.Setfield(-2, "run")
    
    // Create metatable for table
    ls.Pushmetatable(&lua.Metatable{
        IndexFunc: thread.index,
        GCFunc: thread.gc,
    })
    
    // Set metatable for table
    ls.Setmetatable(-2)
    
    return 1
}

func (this *Thread) index(ls lua.State) int {    
    ls.Getfield(-2, ls.Tostring(-1))
    
    return 1
}

func (this *Thread) run(ls lua.State) int {    
    if ls.Gettop() < 1 || ls.Gettop() > 1 {
        panic("Invalid Stack To Thread Run")
    }
    
    threadid := uuid.New()     
    ls.Getglobal("threads")    
    thread := ls.Newthread()
    ls.Setfield(ls.Gettop()-1, threadid)  
    ls.Pop(1)    
    ls.Getfield(-1, "func")
    thread.Xmove(ls, 1)    
    
    go func(thread lua.State) {
       if thread.Gettop() < 1 || thread.Gettop() > 1 {
            panic("Invalid Threadstate Stack")
        }
        //time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
        thread.Resume(0)                        
    }(thread)    
    
    return 0
}

func (this *Thread) gc(ls lua.State) int {
    log.Println("THREAD GC")
    return 0
}