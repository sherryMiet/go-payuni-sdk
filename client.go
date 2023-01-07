package payuni

type Client struct {
	//商店代號
	MerID string
	//交易資料 AES 加密
	HashKey string
	//交易資料 SHA256 加密
	HashIV string
}

func NewClient(MerID string, HashKey string, HashIV string) *Client {
	return &Client{
		MerID:   MerID,
		HashKey: HashKey,
		HashIV:  HashIV,
	}
}
