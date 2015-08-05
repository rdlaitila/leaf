package lua

import(
    "sync"
    "errors"
)

type Registry struct {
    mutex     *sync.Mutex    
    registry  map[int]interface{}
    currindex int
}

func NewRegistry() *Registry {
    return &Registry{
        mutex: &sync.Mutex{}, 
        registry: make(map[int]interface{}),
    }
}

func (this *Registry) AddValue(govalue interface{}) int {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    
    this.currindex++    
    this.registry[this.currindex] = govalue
    
    return this.currindex
}

func (this *Registry) GetValue(INDEX int) (interface{}, error) {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    
    val, ok := this.registry[INDEX]
    if !ok {
        return nil, errors.New("Invalid Index Supplied. Index does not exist")
    } else {
        return val, nil
    }    
}

func (this *Registry) RemoveValue(INDEX int) error {
    this.mutex.Lock()
    defer this.mutex.Unlock()

    _, ok := this.registry[INDEX]
    if !ok {
        return errors.New("Invalid Index Supplied. Index does not exist")
    } else {
        delete(this.registry, INDEX);
        return nil
    }
}

func (this *Registry) ReserveValue() int {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    
    this.currindex++
    
    return this.currindex
}

func (this *Registry) SetValue(index int, value interface{}) {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    
    this.registry[index] = value
}