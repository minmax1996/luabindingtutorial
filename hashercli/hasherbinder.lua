local hasherbinder = {}

local hasherformat = "hasher -algorithm %s -input '%s'"
function hasherbinder.hash(alg, input)
    local handle = io.popen(string.format(hasherformat, alg, input))
    local result = handle:read("*a")
    handle:close()
    return result
end

return hasherbinder