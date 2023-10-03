---@module "utils"
local utils = import("./utils")

local chat = function(format, ...)
    return utils.baseURL .. "/chat" .. string.format(format, ...)
end


local jwt =
[[eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQzMzAyNzksImp0aSI6IjQiLCJpYXQiOjE2OTM3MjU0NzksImlzcyI6Im9yZy5teV90b2RvIiwic3ViIjoidXNlciB0b2tlbiJ9.CkL_X8BbsHtCAl-KkO4290EEXuDowWY6_xVqwSs5hZw]]

job("add", function(ctx)
    local data, err = curl({
        header = {
            ["x-token"] = jwt,
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = json.encode({

        })
    })
end)

job("del", function(ctx)

end)

job("get", function(ctx)

end)
