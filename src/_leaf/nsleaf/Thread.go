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
    
    //Get or create _threads global
    ls.Getglobal("_threads")
    if ls.Isnil(-1) {
        ls.Pop(1)
        ls.Newtable()
        ls.Setglobal("_threads")
    }
    ls.Getglobal("_threads")
    
    //Create a new thread and push to global _threads var
    threadid := uuid.New()     
    thread := ls.Newthread()
    ls.Setfield(ls.Gettop()-1, threadid)  
    ls.Pop(1)    
    
    //Move the thread func to the new thread state
    ls.Getfield(-1, "func")
    thread.Xmove(ls, 1)    
    
    go func(thread lua.State) {
       if thread.Gettop() < 1 || thread.Gettop() > 1 {
            panic("Invalid Threadstate Stack")
        }        
        _, err := thread.Resume(0)        
        if err != nil {
            thread.Pushstring(err.Error())
            thread.Error()
        }        
    }(thread)    
    
    return 0
}

func (this *Thread) gc(ls lua.State) int {
    log.Println("THREAD GC")
    return 0
}