package vm

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm/lightclient"
	"github.com/stretchr/testify/require"
)

//const (
//	testHeight uint64 = 66848226
//)
/* 自定义INIT_CONSENSUS_STATE_BYTES,生成bytes->hexString*/
func TestTmHeaderValidateAndMerkleProofValidateTemp(t *testing.T) {
	// 初始共识状态
	// TendermintLightClient.sol默认的INIT_CONSENSUS_STATE_BYTES里的chainId为Binance-Chain-Nile
	consensusStateBytesStr := "3731350000000000000000000000000000000000000000000000000000000000000000000000000229eca254b3859bffefaf85f4c95da9fbd26527766b784272789c30ec56b380b6eb96442aaab207bc59978ba3dd477690f5c5872334fc39e627723daa97e441e88ba4515150ec3182bc82593df36f8abb25a619187fcfab7e552b94e64ed2deed000000e8d4a51000" // 十六进制字符串
	consensusStateBytes, err := hex.DecodeString(consensusStateBytesStr)
	require.NoError(t, err)

	cs, err := lightclient.DecodeConsensusState(consensusStateBytes)
	require.NoError(t, err)

	fmt.Printf("cs: %v \n", cs)
	// TODO:生成chainId可自定义的INIT_CONSENSUS_STATE_BYTES
	cs.ChainID = "715"	// 修改链id
	cs.Height = 2 // 修改初始高度


	// TODO:修改验证者集合ValidatorSet和Validators
	consensusStateBytes, err = cs.EncodeConsensusState()
	// 需要转成十六进制字符串
	newConsensusStateBytesStr := hex.EncodeToString(consensusStateBytes)
	fmt.Println("consensusStateBytesHexStr: ", newConsensusStateBytesStr)
	newConsensusStateBytes, err := hex.DecodeString(newConsensusStateBytesStr)
	newcs, err := lightclient.DecodeConsensusState(newConsensusStateBytes)
	require.NoError(t, err)
	fmt.Printf("newcs: %v \n", newcs)

}
