package rpc

import (
	"context"
)

type SendTransactionResponse JsonRpcResponse[string]

type SendTransactionConfigEncoding string

const (
	SendTransactionConfigEncodingBase58 SendTransactionConfigEncoding = "base58"
	SendTransactionConfigEncodingBase64 SendTransactionConfigEncoding = "base64"
)

type SendTransactionConfig struct {
	SkipPreflight       bool                          `json:"skipPreflight,omitempty"`       // default: false
	PreflightCommitment Commitment                    `json:"preflightCommitment,omitempty"` // default: finalized
	Encoding            SendTransactionConfigEncoding `json:"encoding,omitempty"`            // default: base58
	MaxRetries          uint64                        `json:"maxRetries,omitempty"`          // Maximum number of times for the RPC node to retry sending the transaction to the leader. If this parameter not provided, the RPC node will retry the transaction until it is finalized or until the blockhash expires.
	MinContextSlot      uint64                        `json:"minContextSlot"`                // set the minimum slot at which to perform preflight transaction checks
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransaction(ctx context.Context, tx string) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "sendTransaction", tx)
}

// SendTransaction submits a signed transaction to the cluster for processing
func (c *RpcClient) SendTransactionWithConfig(ctx context.Context, tx string, cfg SendTransactionConfig) (JsonRpcResponse[string], error) {
	return call[JsonRpcResponse[string]](c, ctx, "sendTransaction", tx, cfg)
}
