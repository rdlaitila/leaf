package lua

type State interface {
    Error()
    Getfield(int, string)
    Getglobal(string)
    Gettop() int
    Insert(int)
    Newtable()
    Newthread() State
    Pop(int)
    Pushclosure(Function, int)
    Pushfunction(Function)
    Pushmetatable(*Metatable)
    Pushstring(string)
    Resume(int) (bool, error)
    Setfield(int, string)
    Setmetatable(int) int
    Tostring(int) string
    Xmove(State, int)
}