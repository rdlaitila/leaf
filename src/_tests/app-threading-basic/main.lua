local apprun = true
local mu = leap.Mutex()
local wg = leap.WaitGroup()
local ct = 0

wg:add()
leap.Thread(function()
    while true do end
end):run()

wg:add()
leap.Thread(function()
    while true do end
end):run()

wg:add()
leap.Thread(function()
    for a=1, 1000 do
        print(a)
    end
end):run()

wg:wait()
print(ct)