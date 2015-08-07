package lua

type State interface {
    Dostring(string) int
    Error()
    Getfield(int, string)
    Getglobal(string)
    Gettop() int
    Insert(int)
    Isnil(int) bool
    Loadfile(string) error
    Newtable()
    Newthread() State
    Openlibs()
    Pcall(int,int,int) error
    Pop(int)
    Pushclosure(Function, int)
    Pushfunction(Function)
    Pushmetatable(*Metatable)
    Pushmodule(string, Function)
    Pushstring(string)
    Resume(int) (bool, error)
    Setfield(int, string)
    Setglobal(string)
    Setmetatable(int) int
    Tostring(int) string
    Xmove(State, int)
}