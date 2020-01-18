expirationd = require('expirationd')
clock = require('clock')

box.cfg{
    listen = 3301
}

-- need to remove session from here
local storageinfo = {
    {name = 'session', lifetime = 3600 * 24 * 10},

    {name = 'register', lifetime = 3600},
    {name = 'steam', lifetime = 3600},
    {name = 'change_password', lifetime = 3600},

    {name = 'pubg', lifetime = 3600}
}

function generate_tokener_storage(name, lifetime)
    if not box.space[name] then
        box.schema.space.create(name)
        box.space[name]:create_index('primary', {
            type = "hash",
            parts = {1, 'string'}
        })

        print("tokener storage", name, "was created")
    else
        print("tokener storage", name, "has already created")
    end

    local tokener = {
        space = box.space[name],
        lifetime = lifetime
    }

    function tokener:add(token, ...)
        return self.space:insert{token, os.time(), ...}
    end

    function tokener:select(token)
        return self.space:select{token}
    end

    function tokener:remove(token)
        return self.space:delete{token}
    end

    function tokener:is_expired(args, tuple)
        if tuple[2] + self.lifetime > os.time() then
            return false
        end

        return true
    end

    function tokener:delete_tuple(space_id, args, tuple)
        print("delete by", tuple[1])
        self.space:delete{tuple[1]}
    end

    return tokener
end

storage = {}
for _, info in ipairs(storageinfo) do
    local tokener = generate_tokener_storage(info.name, info.lifetime)

    expirationd.start(info.name, tokener.space.id, function (...) return tokener:is_expired(...) end, {
        process_expired_tuple = function (...) return tokener:delete_tuple(...) end,
        args = nil,
        tuples_per_iteration = 50,
        full_scan_time = 3600
    })

    storage[info.name] = tokener
end
