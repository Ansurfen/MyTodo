---@module "utils"
local utils = import("./utils")

local captcha = function(format, ...)
    return utils.baseURL .. "/captcha" .. string.format(format, ...)
end

job("new", function(ctx)
    local data, err = curl(captcha("/new"))
    yassert(err)
    print(strings.TrimSpace(data))
end)

job("ls", function(ctx)
    local data, err = curl(captcha("/ls"))
    yassert(err)
    print(strings.TrimSpace(data))
end)

job("get", function(ctx)
    argsparse(ctx, {
        id = flag_type.str
    })
    local id = ctx.flags["id"]
    local data, err = curl(captcha("/get/%s", id))
    yassert(err)
    print(strings.TrimSpace(data))
end)
