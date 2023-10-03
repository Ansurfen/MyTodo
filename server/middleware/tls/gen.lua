job("tmpl", function()
    sh([[
cfssl print-defaults config > ca-config.json
cfssl print-defaults csr > ca-csr.json
]])
end)

job("ca", function()
    sh([[
cfssl gencert -initca ca-csr.json | cfssljson -bare ca -
cfssl gencert -config ca-config.json -ca ca.pem -ca-key ca-key.pem -profile www ca-csr.json | cfssljson -bare server
]])
end)

job("clean", function()
    rm("ca-key.pem", "ca.csr", "ca.pem", "server-key.pem", "server.csr", "server.pem")
end)

jobs("all", "tmpl", "ca")
