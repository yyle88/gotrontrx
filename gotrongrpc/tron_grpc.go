package gotrongrpc

import (
	"math/big"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/pkg/errors"
)

const trc20TransferFromMethodSignature = "0x23b872dd"
const trc20AllowanceMethodSignature = "0xdd62ed3e"

type Client struct {
	grpcClient *client.GrpcClient
}

func NewClient(grpcClient *client.GrpcClient) *Client {
	return &Client{grpcClient: grpcClient}
}

func NewClientWithAddress(address string) (*Client, error) {
	grpcClient, err := NewGrpcClient(address)
	if err != nil {
		return nil, errors.WithMessage(err, "new_grpc_client is wrong")
	}
	return NewClient(grpcClient), nil
}

func (g *Client) GetGrpcClient() *client.GrpcClient {
	return g.grpcClient
}

/*
TRC20TransferFrom
在函数`transferFrom(address from, address to, uint256 value) public returns (bool)`中，参数`from`表示代币的拥有者或持有者的地址，而不是发起人的地址。
该函数的目的是从一个地址（`from`）向另一个地址（`to`）转移一定数量的代币（`value`）。在使用`transferFrom`函数之前，需要确保调用者已经获得了`from`地址的授权，这通常是通过调用`approve`函数进行授权的。
所以，`from`参数应该填入代币的持有者或拥有者的地址，而不是发起人的地址。通过这种方式，代币的持有者可以授权其他地址代表自己进行代币转移操作。

智能合约中，`transferFrom` 函数的发起人（也称为消息发送者）就是对该函数进行调用的账户或地址。
*/
func (g *Client) TRC20TransferFrom(operator, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error) {
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
	return g.grpcClient.TRC20Call(operator, contract, req, false, feeLimit)
}

/*
TRC20Allowance
Owner      string //ownerAddress是代币的持有者地址，它具有授权权限，并可以设置授权给其他地址。
Spender    string //spenderAddress是被授权执行转账操作的地址，它可以从ownerAddress处获得授权并执行转账。
Contract   string //代币合约地址
*/
func (g *Client) TRC20Allowance(owner string, spender string, contractAddress string) (*big.Int, error) {
	addrA, err := address.Base58ToAddress(owner)
	if err != nil {
		return nil, errors.Errorf("invalid address %s: %v", owner, owner)
	}
	addrB, err := address.Base58ToAddress(spender)
	if err != nil {
		return nil, errors.Errorf("invalid address %s: %v", spender, spender)
	}
	req := trc20AllowanceMethodSignature
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrA.Hex())-4:] + addrA.Hex()[4:]
	req += "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-4:] + addrB.Hex()[4:]
	result, err := g.grpcClient.TRC20Call("", contractAddress, req, true, 0)
	if err != nil {
		return nil, err
	}
	data := common.BytesToHexString(result.GetConstantResult()[0])
	r, err := g.grpcClient.ParseTRC20NumericProperty(data)
	if err != nil {
		return nil, errors.Errorf("contract address %s: %v", contractAddress, err)
	}
	if r == nil {
		return nil, errors.Errorf("contract address %s: invalid allowance of %s %s", contractAddress, owner, spender)
	}
	return r, nil
}
