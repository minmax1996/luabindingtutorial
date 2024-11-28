# HasherBinder

This project is a binder between the Golang implementation of a hasher and Lua. It allows you to use the hashing functionality provided by the Golang code within Lua scripts.
## Usage

To use the hasher binder in your Lua scripts, follow these steps:

1. Load the binder module in your Lua script:
    ```lua
    local hasher = require("hasherbinder")
    ```

2. Use the provided functions to hash data:
    ```lua
    local hash256 = hasher.sha256_raw("https://example.com/some/link?q1=1&q2=2&g=3"))
    print("Hash:", hash256)
    local hashmd5 = hasher.double_md5_sorted("https://example.com/some/link?q1=1&q2=2&g=3")
    print("Hash:", hashmd5)
    ```