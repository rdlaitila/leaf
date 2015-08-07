-- Test to ensure that the module is preloaded
if package.preload['leaf.runtime'] == nil then
    error("leaf.runtime not properly preloaded")
end

-- Test module require
local runtime = require('leaf.runtime')

-- Test numcpu 
local numcpu = runtime:numcpu()
if type(numcpu) ~= 'number' then
    error("runtime:numcpu() returns a non-number")
elseif numcpu == 0 then
    error("runtime:numcpu() is not returning non-zero")
end

-- Test runtime:schedule()
runtime:schedule()

-- Test runtime:numthread()
local numthreads = runtime:numthread()
if type(numthreads) ~= 'number' then
    error("runtime:numthread() returns a non-number")
end

-- Test runtime.arch
local arch = runtime.arch
if type(arch) ~= 'string' then
    error('runtime.arch did not return a string')
end

-- Test runtime.os
local rtos = runtime.os
if type(rtos) ~= 'string' then
    error('runtime.os did not return a string')
end