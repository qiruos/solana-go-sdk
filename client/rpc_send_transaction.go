package client

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

type SendTransactionConfig struct {
	SkipPreflight       bool
	PreflightCommitment rpc.Commitment
	MaxRetries          uint64
	MinContextSlot      uint64
}

func (c SendTransactionConfig) toRpc() rpc.SendTransactionConfig {
	return rpc.SendTransactionConfig{
		Encoding:            rpc.SendTransactionConfigEncodingBase64,
		PreflightCommitment: c.PreflightCommitment,
		MaxRetries:          c.MaxRetries,
		SkipPreflight:       c.SkipPreflight,
		MinContextSlot:      c.MinContextSlot,
	}
}

// SendTransaction send transaction struct directly
func (c *Client) SendTransaction(ctx context.Context, tx types.Transaction) (string, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.SendTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				SendTransactionConfig{}.toRpc(),
			)
		},
		forward[string],
	)
}

// SendTransaction send transaction struct directly
func (c *Client) SendTransactionWithConfig(ctx context.Context, tx types.Transaction, cfg SendTransactionConfig) (string, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return "", fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.SendTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				cfg.toRpc(),
			)
		},
		forward[string],
	)
}
