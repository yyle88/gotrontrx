# gotrontrx
以前叫 gotron 但感觉这个名字会和别人的包重名和冲突，既然如此叫 gotrontrx 吧，这样相对不容易冲突

## 第一步创建钱包：

使用这段代码能创建钱包：
https://gist.github.com/motopig/c680f53897429fd15f5b3ca9aa6f6ed2
当然其它的也能创建钱包。

区块链的钱包创建是离线的，你可以使用任意你觉得趁手的离线工具生成你的钱包（任何通过网页在线创建私钥的行为都是耍流氓）

## 第二步去链上查钱包信息-主要是余额信息：

创建完即可在测试链查看资产情况：
https://shasta.tronscan.org/#/address/TBYHGsFkshasvB3R6Zys4627h98owvUNFn

只要你创建完钱包，在任意波场链（即主网，shasta测试网，nile测试网）查询都是能查到信息（但资产都是空的，是个空账号）

注意保存好你的私钥

你可以在任意波场链里给这个地址转账，结果都会让它有资金，也都可以向外转账。

但是请注意通常一个私钥只用在一个链，避免私钥泄露（只能在某种程度上避免）。

## 第三步领取测试币TRX
在波场官方中文客服群里即可领取测试币5000个TRX：
https://t.me/TronOfficialTechSupport

当然在英文Telegram群里也能领取5000个TRX测试币：
@TronOfficialDevelopersGroupEn

具体操作进群以后输入消息（在以上两个群里均可）:
```
!help
```
你自然就会啦。

## 第四步使用本SDK进行转账等操作吧
当然为避免私钥泄露，只建议使用测试链钱包验证以下的Demo:
[DEMO 最基本的trx转账](/demos/main_sendtrx/main.go)

## 其它的免责声明：
数字货币都是骗局

都是以空气币掠夺平民财富

没有公平正义可言

数字货币对中老年人是极不友好的，因为他们没有机会接触这类披着高科技外衣的割韭菜工具

数字货币对青少年也是极不友好的，因为当他们接触的时候，前面的人已经占据了大量的资源

因此妄图以数字货币，比如稍微主流的 BTC ETH TRX 代替世界货币的操作，都是不可能实现的

都不过是早先持有数字货币的八零后们的无耻幻想

扪心自问，持有几千甚至数万个比特币的人会觉得公平吗，其实不会的

因此未来还会有新事物来代替它们，而我现在也不过只是了解其中的技术，仅此而已。

该项目作者坚定持有“坚决抵制数字货币”的立场。
