#include <lua.h>
#include <lauxlib.h>
#include "libhasher.h"

int lua_sha256_raw(lua_State *L) {
    const char* str = luaL_checkstring(L, 1);
    char* result = Sha256Raw((char*)str);
    lua_pushstring(L, result);
    free(result);
    return 1;
}

int lua_double_md5_sorted(lua_State *L) {
    const char* str = luaL_checkstring(L, 1);
    char* result = DoubleMd5Sorted((char*)str);
    lua_pushstring(L, result);
    free(result);
    return 1;
}

// Register your functions
int luaopen_hasherbinding(lua_State *L) {
    static const struct luaL_Reg mylib[] = {
        {"sha256_raw", lua_sha256_raw},
        {"double_md5_sorted", lua_double_md5_sorted},
        {NULL, NULL}
    };
    luaL_newlib(L, mylib);
    return 1;
}