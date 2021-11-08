package secp256k1

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

// Creates and checks a KeccakSecp256k1 instruction against the following rust code:
//
//  let sk_bytes = base64::decode("bNyQVhCtQ86p9CCtzVkrg3Fm6WJqiYb+dMO4HDtbl6o=").unwrap();
//  let sk = libsecp256k1::SecretKey::parse_slice(&sk_bytes).unwrap();
//  let instr = solana_sdk::secp256k1_instruction::new_secp256k1_instruction(&sk, "message".as_bytes());
//  println!("{}", base64::encode(instr.data))
//
// As all that's in a KeccakSecp256k1 instruction is the program id and the data (no accounts are passed),
// we just check that the instruction data matches that generated by the rust code

func TestNewSecp256k1InstructionMulti(t *testing.T) {
	skBytes, err := base64.StdEncoding.DecodeString("bNyQVhCtQ86p9CCtzVkrg3Fm6WJqiYb+dMO4HDtbl6o=")
	if err != nil {
		t.Error(err)
	}
	sk := crypto.ToECDSAUnsafe(skBytes)

	instr, err := NewSecp256k1InstructionMultipleSigs([]*ecdsa.PrivateKey{sk}, [][]byte{[]byte("message")}, 0)
	if err != nil {
		t.Error(err)
	}

	checkDataStr := "ASAAAAwAAGEABwAArx8O5L8N25rze03Dr4YXi9E+/YsraZi2z1/W/WElzaSnacJNnqFn12Ggd8AMdb3NQIF+0VN43WVhRkSChRmuSV+J+dl5ZQlS6NKFkqBi3OgoxpEKAW1lc3NhZ2U="
	instrDataStr := base64.StdEncoding.EncodeToString(instr.Data)
	if checkDataStr != instrDataStr {
		t.Error(fmt.Sprintf("\ncheckDataStr: %s \ninstrDataStr: %s", checkDataStr, instrDataStr))
	}
}

func TestNewSecp256k1InstructionSingle(t *testing.T) {
	skBytes, err := base64.StdEncoding.DecodeString("bNyQVhCtQ86p9CCtzVkrg3Fm6WJqiYb+dMO4HDtbl6o=")
	if err != nil {
		t.Error(err)
	}
	sk := crypto.ToECDSAUnsafe(skBytes)

	instr, err := NewSecp256k1Instruction(sk, []byte("message"), 0)
	if err != nil {
		t.Error(err)
	}

	checkDataStr := "ASAAAAwAAGEABwAArx8O5L8N25rze03Dr4YXi9E+/YsraZi2z1/W/WElzaSnacJNnqFn12Ggd8AMdb3NQIF+0VN43WVhRkSChRmuSV+J+dl5ZQlS6NKFkqBi3OgoxpEKAW1lc3NhZ2U="
	instrDataStr := base64.StdEncoding.EncodeToString(instr.Data)
	if checkDataStr != instrDataStr {
		t.Error(fmt.Sprintf("\ncheckDataStr: %s \ninstrDataStr: %s", checkDataStr, instrDataStr))
	}
}
