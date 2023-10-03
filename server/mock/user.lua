---@module "utils"
local utils = import("./utils")

local user = function(format, ...)
    return utils.baseURL .. "/user" .. string.format(format, ...)
end

job("sign", function(ctx)
    local data, err = curl({
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            email    = { "a@gmail.com" },
            password = { "root" }
        })
    }, user("/sign"))
    yassert(err)
    print(data)
end)
