-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---yock inherits golang's goroutine, allowing
---you use it in lua script with easy, fast and
---free. Less like lua native coroutine, the goroutine
---supports nested call.
---
---### Example:
---```lua
---go(function()
---    local idx = 0
---    while idx ~= 5 do
---        print("task 1")
---        time.Sleep(1 * time.Second)
---        idx = idx + 1
---    end
---    notify("x")
---    print("task1 fine")
---end)
---
---go(function()
---    print("task 2")
---    wait("x")
---    print("task2 fine")
---end)
---
---go(function()
---    time.Sleep(8 * time.Second)
---    notify("y")
---end)
---
---waits("x", "y")
---```
---view [document](https://ansurfen.github.io/YockNav/guide/concurrency.html#goroutine-coroutines-with-stack)
---@async
---@param callback fun()
function go(callback) end

---wait blocks self and waits for sig.
---If lack signal, it'll block for ever.
---If not, you can set deadline to unblock
---when timeout.
---
---### Example:
---```lua
---# waiting for three second
---go(function()
---     time.Sleep(3 * time.Second)
---     notify("x")
---end)
---wait("x")
---# do something
---
---# the following code will block for ever if don't wait `blocked` signal.
---wait("blocked")
---
---# sets deadline to unblock, and it does not mean signal arrives and is received.
---wait("blocked", time.Second * 20)
---```
---view [document](https://ansurfen.github.io/YockNav/guide/concurrency.html#semaphores)
---@param sig string
---@param timeout? time
function wait(sig, timeout) end

---waits just like wait, blocks self and waits for sig
---supporting for setting deadline to unblock when timeout.
---The difference is waits can collect a variable number of
---signals than wait.
---
---### Example:
---```lua
---# just like wait when uses
---waits("x", "y", time.Second * 20)
---```
---view [document](https://ansurfen.github.io/YockNav/guide/concurrency.html#semaphores)
---@param ... string|time
function waits(...) end

---notify sends sig to signal stream, which can unblock
---wait or waits.
---
---### Example:
---```lua
---# waiting for three second util signal `x` was sent and received
---go(function()
---     time.Sleep(3 * time.Second)
---     notify("x")
---end)
---wait("x")
---```
---view [document](https://ansurfen.github.io/YockNav/guide/concurrency.html#semaphores)
---@param sig string
function notify(sig) end