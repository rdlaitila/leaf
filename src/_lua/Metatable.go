package lua

type Metatable struct {
    IndexFunc Function
    NewindexFunc Function
    TostringFunc Function    
    GCFunc Function
}

func (this *Metatable) Index() Function {
    return this.IndexFunc
}

func (this *Metatable) Newindex() Function {
    return this.NewindexFunc
}

func (this *Metatable) Tostring() Function {
    return this.TostringFunc
}

func (this *Metatable) GC() Function {
    return this.GCFunc
}