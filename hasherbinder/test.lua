-- test.lua
local hasherbinding = require("hasherbinding")
print(hasherbinding.sha256_raw("https://example.com/some/link?q1=1&q2=2&g=3")) -- Should print ca2a598def71c556327b531e2be48fbba879e2d596acdb60006b6ff4626dec2d
print(hasherbinding.double_md5_sorted("https://example.com/some/link?q1=1&q2=2&g=3")) -- Should print 72df8c4cc5ae9c072fc101de65124298