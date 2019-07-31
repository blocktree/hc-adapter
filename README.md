# hypercash-adapter

hypercash-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。


## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建HC.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 1
# node api url, if RPC Server Type = 1, use bitbay insight-api
serverAPI = "http://{IP:PORT}/insight/api/"
# RPC Authentication Username
rpcUser = "hcd"
# RPC Authentication Password
rpcPassword = "hcd2019"
# Is network test?
isTestNet = false
# Omni Core RPC API
omniCoreAPI = "https://{IP:PORT}/"
# Omni Core RPC Authentication Username
omniRPCUser = "hcd"
# Omni Core RPC Authentication Password
omniRPCPassword = "hcd2019"
# Omni token transfer minimum cost
omniTransferCost = "0.01"
# support omnicore
omniSupport = true
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.0001"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 官方资料

HC浏览器

https://hc-explorer.h.cash/

HC OMNI浏览器

https://hcomni-explorer.h.cash/