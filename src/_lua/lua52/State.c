#include <lua52/csrc/lua.h>
#include <lua52/csrc/lauxlib.h>
#include <lua52/csrc/lualib.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "_cgo_export.h"

static int lua52_panicf(lua_State *s);

void lua52_luainit(lua_State *s)
{
    printf("%s\n","lua52_luainit");
    lua_atpanic(s, lua52_panicf);
}

static int lua52_panicf(lua_State *s)
{
    printf("%s\n", "Oh Shit Happened!");    
    return 0;
}    

static int lua52_closurecallback(lua_State *s)
{
	float stateindex;
    float funcindex;
    
    // pull our GovalueRegistry indexes from the closure's upvalues
    stateindex = lua_tonumber(s, lua_upvalueindex(1));
    funcindex = lua_tonumber(s, lua_upvalueindex(2));
    
    // Call back into golang luajit.docallback
	return docallback(stateindex, funcindex);
}

void lua52_pushclosure(lua_State *s, int n)
{
	// pass a goluajit_closurecallback, +2 upvalues that should have been previously pushed:
    // 1: the gvindex of our golang State struct
    // 2: the gvindex of our golang Gofunction func
    lua_pushcclosure(s, lua52_closurecallback, n);
}