package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Index     int          // 保存数据在链中的位置
	Data      BookCheckout // 区块中需要保存的信息(借阅信息)
	TimeStamp string       // 区块创建时时间
	Hash      string       // 区块生成的哈希值
	PrevHash  string       // 前一个块的哈希值
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1] //?

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func (b *Block) GenerateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(rune(b.Index)) + b.TimeStamp + string(bytes) + b.PrevHash
	// data := string(b.Index) + b.TimeStamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (b *Block) validateHash(hash string) bool {
	b.GenerateHash()
	return b.Hash == hash
}

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	// todo:检查传入参数
	block := &Block{}
	block.Index = prevBlock.Index + 1
	block.TimeStamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.GenerateHash()

	return block
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true}) //?
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}} //?
}

func validBlock(block, prevBlock *Block) bool {
	// 确认哈希值
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	// 确认哈希值合法
	if !block.validateHash(block.Hash) {
		return false
	}

	// 检查位置是否自增
	if prevBlock.Index+1 != block.Index {
		return false
	}

	return true
}
