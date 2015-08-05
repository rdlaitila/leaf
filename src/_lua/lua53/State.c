#include <lua53/csrc/lua.h>
#include <lua53/csrc/lauxlib.h>
#include <lua53/csrc/lualib.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "_cgo_export.h"

static int lua53_panicf(lua_State *s);

void lua53_luainit(lua_State *s)
{
    printf("%s\n","lua53_luainit");
    lua_atpanic(s, lua53_panicf);
}

static int lua53_panicf(lua_State *s)
{
    printf("%s\n", "Oh Shit Happened!");    
    return 0;
}    

static int lua53_closurecallback(lua_State *s)
{
	float stateindex;
    float funcindex;
    
    // pull our GovalueRegistry indexes from the closure's upvalues
    stateindex = lua_tonumber(s, lua_upvalueindex(1));
    funcindex = lua_tonumber(s, lua_upvalueindex(2));
    
    // Call back into golang luajit.docallback
	return docallback(stateindex, funcindex);
}

void lua53_pushclosure(lua_State *s, int n)
{
	// pass a goluajit_closurecallback, +2 upvalues that should have been previously pushed:
    // 1: the gvindex of our golang State struct
    // 2: the gvindex of our golang Gofunction func
    lua_pushcclosure(s, lua53_closurecallback, n);
}

int lua53_yield(lua_State *s, int n)
{
    return lua_yield(s, n);
}

void lua53_rawseti(lua_State *s, int i, int n)
{
    lua_rawgeti(s, i, n);
}

void lua53_rawgeti(lua_State *s, int i, int n)
{
    lua_rawgeti(s, i, n);
}

int lua53_pcall(lua_State *s, int nargs, int nresults, int errfunc)
{
    return lua_pcall(s, nargs, nresults, errfunc);
}

void lua53_call(lua_State *s, int nargs, int nresults)
{
    lua_call(s, nargs, nresults);
}