local limit_req = require "resty.limit.req"

// START OMIT
local lim = limit_req.new("stream_limit_req_store", 1, 3)
local key = ngx.var.remote_addr
local delay, err = lim:incoming(key, true)
if not delay then
    return ngx.exit(1)
end
if delay >= 0.001 then
    ngx.sleep(delay)
end
// END OMIT