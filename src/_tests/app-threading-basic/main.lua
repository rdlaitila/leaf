local leaf = require('leaf')
threads = {}

local wg = leaf.WaitGroup()
local mu = leaf.Mutex()

wg:add()
leaf.Thread(function()
    while true do
        print('thread1')
    end
end):run()

wg:add()
leaf.Thread(function()
    while true do
        print('thread2')
    end
end):run()

wg:wait()

print(count)