function main()
    local svr = http.Server()
    
    svr:handle('/foo/bar', function(req, resp)    
    end)
    
    svr:listen('0.0.0.0:443')
end