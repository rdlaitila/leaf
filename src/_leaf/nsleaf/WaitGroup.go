package nsleaf

import(
    "sync"
    "log"
    
    "_lua"
)

type WaitGroup struct{
    grindex int
    wg *sync.WaitGroup
    mu *sync.Mutex
}

func NewWaitGroup(ls lua.State) int {

    grindex := lua.GlobalRegistry.ReserveValue()
    wg := &WaitGroup{
        wg: &sync.WaitGroup{}, 
        mu: &sync.Mutex{},
        grindex: grindex,
    }
    lua.GlobalRegistry.SetValue(grindex, wg)    
    
    // Create new table. This will be returned
    ls.Newtable()
    
    // Push add
    ls.Pushfunction(wg.add)
    ls.Setfield(-2, "add")
    
    // Push done
    ls.Pushfunction(wg.done)
    ls.Setfield(-2, "done")
    
    // Push wait
    ls.Pushfunction(wg.wait)
    ls.Setfield(-2, "wait")
    
    // Create metatable for table
    ls.Pushmetatable(&lua.Metatable{
        GCFunc: wg.gc,
    })
    
    // Set metatable for userdata
    ls.Setmetatable(-2,)
    
    return 1
}

func (this *WaitGroup) add(ls lua.State) int {
    this.mu.Lock()
    this.wg.Add(1)
    this.mu.Unlock()
    
    return 0
}

func (this *WaitGroup) done(ls lua.State) int {
    this.mu.Lock()
    this.wg.Done()
    this.mu.Unlock()    
    return 0
}

func (this *WaitGroup) wait(ls lua.State) int {    
    this.wg.Wait()
    
    return 0
}

func (this *WaitGroup) gc(ls lua.State) int {
    log.Println("WAITGROUP GC")
    return 0
}