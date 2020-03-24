package monitor

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/mitchellh/mapstructure"

	"quorumengineering/quorum-report/client"
	"quorumengineering/quorum-report/database"
	"quorumengineering/quorum-report/graphql"
	"quorumengineering/quorum-report/types"
)

type BlockMonitor struct {
	db                 database.Database
	quorumClient       *client.QuorumClient
	syncStart          chan uint64
	transactionMonitor *TransactionMonitor
}

func NewBlockMonitor(db database.Database, quorumClient *client.QuorumClient) *BlockMonitor {
	return &BlockMonitor{
		db,
		quorumClient,
		make(chan uint64),
		NewTransactionMonitor(db, quorumClient),
	}
}

func (bm *BlockMonitor) Start() {
	// Pulling historical blocks since the last persisted while continuously listening to ChainHeadEvent.
	// For every block received, pull transactions/ events related to the registered contracts.

	fmt.Println("Start to sync blocks...")

	// 1. Fetch the current block height.
	currentBlockNumber, err := bm.currentBlockNumber()
	if err != nil {
		// TODO: should gracefully handle error (if quorum node is down, reconnect?)
		log.Fatalf("get current head error: %v.\n", err)
	}
	fmt.Printf("Current block head is: %v\n", currentBlockNumber)

	// 2. Sync from last persisted to current block height.
	bm.syncBlocks(bm.db.GetLastPersistedBlockNumber(), currentBlockNumber)

	// 3. Listen to ChainHeadEvent and sync.
	go bm.listenToChainHead()
	latestChainHead := <-bm.syncStart
	close(bm.syncStart)

	// 4. Sync from current block height + 1 to the first ChainHeadEvent if there is any gap.
	bm.syncBlocks(currentBlockNumber, latestChainHead-1)
}

func (bm *BlockMonitor) currentBlockNumber() (uint64, error) {
	var (
		resp         map[string]interface{}
		currentBlock graphql.CurrentBlock
	)
	resp, err := bm.quorumClient.ExecuteGraphQLQuery(context.Background(), graphql.CurrentBlockQuery())
	if err != nil {
		return 0, err
	}
	err = mapstructure.Decode(resp["block"].(map[string]interface{}), &currentBlock)
	if err != nil {
		return 0, err
	}
	return hexutil.DecodeUint64(currentBlock.Number)
}

func (bm *BlockMonitor) listenToChainHead() {
	headers := make(chan *ethTypes.Header)
	sub, err := bm.quorumClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		// TODO: should gracefully handle error (if quorum node is down, reconnect?)
		log.Fatalf("subscribe to chain head event failed: %v.\n", err)
	}
	for {
		select {
		case err := <-sub.Err():
			// TODO: should gracefully handle error (if quorum node is down, reconnect?)
			log.Fatalf("chain head event subscription error: %v.\n", err)
		case header := <-headers:
			if !isClosed(bm.syncStart) {
				bm.syncStart <- header.Number.Uint64()
			}
			blockOrigin, err := bm.quorumClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				// TODO: should gracefully handle error (if quorum node is down, reconnect?)
				log.Fatalf("get block %v error: %v.\n", header.Hash(), err)
			}
			bm.process(createBlock(blockOrigin))
		}
	}
}

func (bm *BlockMonitor) syncBlocks(start, end uint64) {
	fmt.Printf("Start to sync historic block from %v to %v. \n", start, end)
	for i := start + 1; i <= end; i++ {
		blockOrigin, err := bm.quorumClient.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			// TODO: should gracefully handle error (if quorum node is down, reconnect?)
			log.Fatalf("fetch block %v failed: %v.\n", i, err)
		}
		bm.process(createBlock(blockOrigin))
	}
}

func (bm *BlockMonitor) process(block *types.Block) {
	// Transaction monitor pulls all transactions for the given block.
	bm.transactionMonitor.PullTransactions(block)

	// Write block to DB.
	bm.db.WriteBlock(block)
}

func createBlock(block *ethTypes.Block) *types.Block {
	txs := []common.Hash{}
	for _, tx := range block.Transactions() {
		txs = append(txs, tx.Hash())
	}
	return &types.Block{
		Hash:         block.Hash(),
		ParentHash:   block.ParentHash(),
		StateRoot:    block.Root(),
		TxRoot:       block.TxHash(),
		ReceiptRoot:  block.ReceiptHash(),
		Number:       block.NumberU64(),
		GasLimit:     block.GasLimit(),
		GasUsed:      block.GasUsed(),
		Timestamp:    block.Time(),
		ExtraData:    block.Extra(),
		Transactions: txs,
	}
}