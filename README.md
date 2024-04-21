### Reth-Go 脚本 
#### 项目概览
该脚本提供了能自动挖取 reth 的脚本，并且采用 golang 实现相比 js 版本的脚本性能更好

#### 功能摘要
该脚本负责执行以下任务：
1. **以太坊节点连接**：使用 `ethclient.Dial` 函数建立与以太坊节点的连接。
2. **私钥解析**：将十六进制私钥转换为 ECDSA 私钥对象。
3. **地址派生**：从 ECDSA 公钥派生出以太坊地址。
4. **挑战哈希计算**：计算给定挑战的哈希值，在本例中挑战为 "rETH"。
5. **寻找解决方案**：尝试找到一个满足特定条件的解决方案，通过生成随机值并将其与挑战哈希进行哈希运算。
6. **构建交易**：使用指定的参数（如 nonce、value、gas limit 和 gas price）构建交易。
7. **Gas 价格管理**：如果建议的 gas 价格高于预定义的最大值，则等待并重试。
8. **交易签名和发送**：使用私钥签名交易并将其发送到以太坊网络。

#### 使用说明
要使用 Go 脚本，请按照以下步骤操作：
1. **环境搭建**：确保您的系统上已安装 Go 编程语言。
2. **私钥和地址设置**：将占位符 `"地址"` 替换为您的以太坊账户地址，将 `"私钥"` 替换为您的实际十六进制私钥。
3. **连接字符串**：如果需要，更新以太坊节点 RPC URL。当前 URL 设置为 "https://rpc.flashbots.net"。
4. **运行脚本**：使用 Go 命令执行脚本：`go run <文件名>.go`。

#### 贡献指南
- Fork 该仓库并为您的特性创建一个新的分支（`Feat_xxx`）。
- 进行您的更改并提交，附带一个描述性的提交信息。
- 开启一个 Pull Request 将您的特性分支合并回主项目。

#### To-Do List
- [x] 修复 `ethclient.Dial` 中的 URL 字符串格式，确保正确连接到以太坊节点。
- [x] 实现 `findSolution` 函数的退出条件，以避免潜在的无限循环。
- [x] 验证 `jsonData` 结构体中的字段是否符合智能合约的ABI要求。
- [x] 对 `dataHex` 的生成逻辑进行测试，确保构造的交易数据正确无误。
- [x] 在交易发送失败时，实现错误处理和重试机制。
- [x] 添加日志记录功能，以便于监控脚本运行状态和问题调试。
- [x] 根据以太坊网络的实时 gas 价格调整脚本中的 gas 价格和 gas 限制设置。
- [x] 开发单元测试，验证 `findSolution` 和其他关键函数的正确性。
- [x] 实现交易发送后的确认机制，确保交易已成功上链。

请注意，必须将以太坊账户地址和私钥的占位符替换为实际值，脚本才能正确工作。此外，根据用户的需求或网络可用性，可能需要更改以太坊节点 RPC URL。