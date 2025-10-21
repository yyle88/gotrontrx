package gotrongrpc

import (
	"math/big"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/pkg/errors"
)

const (
	trc20TransferFromMethodSignature = "0x23b872dd" // TRC20 transferFrom method signature // TRC20 transferFrom 方法签名
	trc20AllowanceMethodSignature    = "0xdd62ed3e" // TRC20 allowance method signature // TRC20 allowance 方法签名
)

// TRC20TransferFrom creates a TRC20 transferFrom transaction on TRON blockchain
// In the `transferFrom(address from, address to, uint256 value)` function, the `from` parameter represents the token holder's address, not the caller's address
// Before using `transferFrom`, ensure the caller has been authorized by the `from` address through the `approve` function
// Returns transaction extension that can be signed and broadcast to the network
//
// TRC20TransferFrom 在 TRON 区块链上创建 TRC20 transferFrom 交易
// 在 `transferFrom(address from, address to, uint256 value)` 函数中，`from` 参数表示代币持有者的地址，而非调用者的地址
// 使用 `transferFrom` 之前，需确保调用者已通过 `approve` 函数获得 `from` 地址的授权
// 返回可签名并广播到网络的交易扩展
func TRC20TransferFrom(grpcClient *client.GrpcClient, operatorAddress, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error) {
	addrA, err := address.Base58ToAddress(from)
	if err != nil {
		return nil, err
	}
	addrB, err := address.Base58ToAddress(to)
	if err != nil {
		return nil, err
	}
	req := trc20TransferFromMethodSignature
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrA.Hex())-4:] + addrA.Hex()[4:]
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-4:] + addrB.Hex()[4:]
	req += common.Bytes2Hex(common.LeftPadBytes(amount.Bytes(), 32))
	return grpcClient.TRC20Call(operatorAddress, contract, req, false, feeLimit)
}

// TRC20Allowance queries the allowance amount that spenderAddress is authorized to spend from ownerAddress
// Checks the TRC20 token allowance for the specified owner-spender pair on the specified contract
// Returns the allowance amount as *big.Int representing the authorized spending limit
//
// TRC20Allowance 查询 spenderAddress 被授权从 ownerAddress 花费的额度
// 检查指定合约上指定所有者-花费者对的 TRC20 代币授权额度
// 返回授权额度 *big.Int，代表授权的花费限额
func TRC20Allowance(grpcClient *client.GrpcClient, ownerAddress string, spenderAddress string, contractAddress string) (*big.Int, error) {
	addrA, err := address.Base58ToAddress(ownerAddress)
	if err != nil {
		return nil, errors.Errorf("invalid address %s: %v", ownerAddress, ownerAddress)
	}
	addrB, err := address.Base58ToAddress(spenderAddress)
	if err != nil {
		return nil, errors.Errorf("invalid address %s: %v", spenderAddress, spenderAddress)
	}
	req := trc20AllowanceMethodSignature
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrA.Hex())-4:] + addrA.Hex()[4:]
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-4:] + addrB.Hex()[4:]
	result, err := grpcClient.TRC20Call("", contractAddress, req, true, 0)
	if err != nil {
		return nil, err
	}
	data := common.BytesToHexString(result.GetConstantResult()[0])
	allowance, err := grpcClient.ParseTRC20NumericProperty(data)
	if err != nil {
		return nil, errors.Errorf("contract address %s: %v", contractAddress, err)
	}
	if allowance == nil {
		return nil, errors.Errorf("contract address %s: invalid allowance of %s %s", contractAddress, ownerAddress, spenderAddress)
	}
	return allowance, nil
}
