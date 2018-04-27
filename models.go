package main

//this is the block
type Block struct {
	Index     int `json:"index"`
	Timestamp int64 `json:"timestamp"`
	TheData   string  `json:"thedata"` // <- holding our "data"
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

//Data that gets added into the block
type Message struct {
	TheData string `json:"thedata"`
}
