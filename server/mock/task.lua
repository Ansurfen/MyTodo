---@module "utils"
local utils = import("./utils")

local task = function(format, ...)
    return utils.baseURL .. "/task" .. string.format(format, ...)
end

local jwt =
[[eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQzMzAyNzksImp0aSI6IjQiLCJpYXQiOjE2OTM3MjU0NzksImlzcyI6Im9yZy5teV90b2RvIiwic3ViIjoidXNlciB0b2tlbiJ9.CkL_X8BbsHtCAl-KkO4290EEXuDowWY6_xVqwSs5hZw]]

job("get", function(ctx)
    local data, err = curl({
        header = {
            ["x-token"] = jwt,
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            page = { tostring(1) },
            limit = { tostring(10) }
        })
    }, task("/get"))
    yassert(err)
    print(data)
end)
