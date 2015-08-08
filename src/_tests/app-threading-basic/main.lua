local leaf = require('leaf')
local runtime = require('leaf.runtime')

local count = 0

local mu = leaf.Mutex()
local wg = leaf.WaitGroup()

wg:add()
leaf.Thread(function()
    for a=1, 1000 do
        print('1')
    end
    wg:done()
end):run()

wg:add()
leaf.Thread(function()
    for a=1, 1000 do
        print('2')
    end
    wg:done()
end):run()

wg:wait()

print(count)
