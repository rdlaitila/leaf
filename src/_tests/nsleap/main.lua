local leaf = require('leaf')

-- Test leaf.Mutex
local mu = leaf.Mutex()
mu:lock()
mu:unlock()

-- Test leaf.WaitGroup
local wg = leaf.WaitGroup()
wg:add()
wg:done()
wg:wait()

-- Test leaf.Thread
local th = leaf.Thread(function() 
end)
th:run()