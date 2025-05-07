<p align="center">
  <img 
    alt="wojack-cartoon logo" 
    src="assets/wojack-cartoon.jpeg" 
    style="max-height: 500px; width: auto; max-width: 100%;" 
  />
</p>
<h3 align="center">golang-tron</h3>
<p align="center">create/sign <code>tron transaction</code> with golang</p>

---

# gotrontrx

`gotrontrx` is a Go toolkit exploring TRON blockchain tech without participating in crypto coins.

`gotrontrx` package interacts with the TRON network via a gRPC client, enabling developers to seamlessly execute the full workflow of transaction creation, signing, and broadcasting.

---

## CHINESE README

[中文说明](README.zh.md)

---

## Features

- **gRPC Client Support**: Establish connections with TRON nodes via gRPC, supporting both mainnet and testnet nodes.
- **Account and Transfer Operations**: Facilitate transactions, including specifying sender, recipient, and amount.
- **Transaction Signing**: Enable transaction signing using private keys to ensure security.
- **Transaction Broadcasting**: Broadcast signed transactions to the blockchain network.
- **Transaction Hash Calculation**: Provide tools for calculating transaction hashes.
- **Response Handling**: Offer structured processing of TRON gRPC API responses.

## Dependencies

- `github.com/fbsobreira/gotron-sdk`: Basic client of TRON gRPC APIs.
- `github.com/yyle88/gotrontrx`: Simple TRON operations.
- `neatjson`: Neat and structured output of information.
- `must`: Simple assertions in conditions.
- `rese`: Reduce boilerplate code for error handling.

## Installation

```bash
go get github.com/yyle88/gotrontrx
```

## Quick Start

Here are the primary functions:

- **`gotrongrpc.NewClient`**: Initialize a gRPC client to connect to TRON nodes.
- **`client.GetGrpc().Transfer`**: Create a transfer transaction.
- **`gotronsign.Sign`**: Sign a transaction using a private key.
- **`client.GetGrpc().Broadcast`**: Broadcast the signed transaction to the network.

`gotrontrx` allow developers to efficiently build TRON blockchain-based applications.

## Important Notes

1. **Security Precautions**: Never input your private key directly into the terminal in production environments to prevent leaks.
2. **Test Environment**: Use the testnet for debugging to avoid financial losses from unintended operations.
3. **Data Validation**: Ensure that input addresses and amounts are valid and compliant with TRON blockchain requirements.

## TRON Guide

`gotrontrx` package provides a straightforward introduction to TRON, along with essential knowledge for working with its blockchain.

### Step 1: Create a Wallet

You can generate a wallet offline using the following code:  
https://gist.github.com/motopig/c680f53897429fd15f5b3ca9aa6f6ed2

Copy the code and run it on your local machine. Alternatively, you can use other offline tools for wallet creation.

**Important:** Blockchain wallets should always be created offline. Never use online services to generate private keys as they may not safe.

### Step 2: Check Wallet Information

Once the wallet is created, you can check its balance and information on the blockchain. For example, use the TRON testnet explorer website:  
https://shasta.tronscan.org/#/address/TBYHGsFkshasvB3R6Zys4627h98owvUNFn

### Step 3: Obtain Test TRX

Developers can join the official TRON Telegram groups to request 5,000 TRX test tokens:
- Chinese Support: [TRON 官方中文客服群](https://t.me/TronOfficialTechSupport)
- English Support: [TRON Official Developers Group](https://t.me/TronOfficialDevelopersGroupEn)

Send the following message in either group to receive instructions:
```
!help
```

### Step 4: Use This SDK to Perform Transactions

Use a testnet wallet to try out the SDK functionalities. see: [Basic-TRX-Transfer-DemoCode](internal/demos/sendtrx/main.go).

---

## DISCLAIMER

Crypto coin, at its core, is nothing but a scam. It thrives on the concept of "air coins"—valueless digital assets—to exploit the hard-earned wealth of ordinary people, all under the guise of innovation and progress. This ecosystem is inherently devoid of fairness or justice.

For the elderly, cryptocurrencies present significant challenges and risks. The so-called "high-tech" façade often excludes them from understanding or engaging with these tools. Instead, they become easy targets for financial exploitation, stripped of the resources they worked a lifetime to accumulate.

The younger generation faces a different but equally insidious issue. By the time they have the opportunity to engage, the early adopters have already hoarded the lion’s share of resources. The system is inherently tilted, offering little chance for new entrants to gain a fair footing.

The idea that cryptocurrencies like BTC, ETH, or TRX could replace global fiat currencies is nothing more than a pipe dream. This notion serves only as the shameless fantasy of early adopters, particularly those from the 1980s generation, who hoarded significant amounts of crypto coin before the general public even had an opportunity to participate.

Ask yourself this: would someone holding thousands, or even tens of thousands, of Bitcoin ever genuinely believe the system is fair? The answer is unequivocally no. These systems were never designed with fairness in mind but rather to entrench the advantages of a select few.

The rise of cryptocurrencies is not the endgame. It is inevitable that new innovations will emerge, replacing these deeply flawed systems. At this moment, my interest lies purely in understanding the underlying technology—nothing more, nothing less.

This project exists solely for the purpose of technical learning and exploration. The author of this project maintains a firm and unequivocal stance of *staunch resistance to cryptocurrencies*.

--- 

## License

MIT License. See [LICENSE](LICENSE).

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with `gotrontrx`!** 🎉

Give me stars. Thank you!!!

## GitHub Stars

[![starring](https://starchart.cc/yyle88/gotrontrx.svg?variant=adaptive)](https://starchart.cc/yyle88/gotrontrx)
