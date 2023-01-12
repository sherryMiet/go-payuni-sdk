package payuni

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	TestCreditCardCloseUrl = "https://sandbox-api.payuni.com.tw/api/trade/close"
	CreditCardCloseUrl     = "https://api.payuni.com.tw/api/trade/close"
	CreditCardCloseVersion = "1.0"
)

// 關帳類型
type CloseTypeEnum int

// List of ChooseSubPaymentEnum
const (
	CloseType_Close        CloseTypeEnum = 1
	CloseType_Refund       CloseTypeEnum = 2
	CloseType_CancelClose  CloseTypeEnum = -1
	CloseType_CancelRefund CloseTypeEnum = -2
)

type CreditCardCloseRequestCall struct {
	HashKey                string
	HashIV                 string
	CreditCardCloseRequest *CreditCardCloseRequest
}

type CreditCardCloseRequest struct {
	//商店代號
	MerID string `json:"MerchantID"`
	//版本
	Version string `json:"Version"`
	//交易資料 AES加密字串
	EncryptInfo string `json:"EncryptInfo"`
	//交易資料 SHA256加密字串
	HashInfo string `json:"HashInfo"`
}

type CreditCardCloseEncryptInfoReq struct {
	MerID     string
	TradeNo   string
	Timestamp int
	CloseType CloseTypeEnum
	TradeAmt  int
}
type CreditCardCloseEncryptInfoRes struct {
	Status    string
	Message   string
	MerID     string
	TradeNo   string
	CloseType string
}
type CreditCardCloseResponse struct {
	//回傳狀態
	Status string
	//回傳訊息
	Message string
	//版本
	Version string `json:"Version"`
	//交易資料 AES加密字串
	EncryptInfo string `json:"EncryptInfo"`
	//交易資料 SHA256加密字串
	HashInfo string `json:"HashInfo"`
}
type CreditCardCloseResult struct {
	//商店代號
	MerchantID string
	//藍新金流交易序號
	TradeNo string
	//交易金額
	Amt int
	//商店訂單編號
	MerchantOrderNo string
	//檢核碼
	CheckCode string
}

func (c *CreditCardCloseEncryptInfoReq) SetTradeNo(TradeNo string) *CreditCardCloseEncryptInfoReq {
	c.TradeNo = TradeNo
	return c
}
func (c *CreditCardCloseEncryptInfoReq) SetTradeAmt(TradeAmt int) *CreditCardCloseEncryptInfoReq {
	c.TradeAmt = TradeAmt
	return c
}
func (c *CreditCardCloseEncryptInfoReq) SetCloseType(CloseType CloseTypeEnum) *CreditCardCloseEncryptInfoReq {
	c.CloseType = CloseType
	return c
}

func NewCreditCardCloseEncryptInfo() *CreditCardCloseEncryptInfoReq {
	return &CreditCardCloseEncryptInfoReq{}
}

func (c *Client) CreditCardClose(Data *CreditCardCloseEncryptInfoReq) *CreditCardCloseRequestCall {
	Data.MerID = c.MerID
	Data.Timestamp = int(time.Now().Unix())
	params := StructToParamsMap(Data)
	paramStr := ParamsMapToURLEncode(params)
	req := new(CreditCardCloseRequest)
	req.MerID = c.MerID
	req.EncryptInfo = Aes256GCMEncrypt(paramStr, c.HashKey, c.HashIV)
	req.HashInfo = SHA256(c.HashKey + req.EncryptInfo + c.HashIV)
	req.Version = CreditCardCloseVersion
	return &CreditCardCloseRequestCall{
		HashKey:                c.HashKey,
		HashIV:                 c.HashIV,
		CreditCardCloseRequest: req,
	}
}

func (c CreditCardCloseRequestCall) Do() (*CreditCardCloseEncryptInfoRes, error) {
	response := new(CreditCardCloseResponse)
	PostData := make(map[string]string)
	PostData["MerID"] = c.CreditCardCloseRequest.MerID
	PostData["Version"] = c.CreditCardCloseRequest.Version
	PostData["EncryptInfo"] = c.CreditCardCloseRequest.EncryptInfo
	PostData["HashInfo"] = c.CreditCardCloseRequest.HashInfo
	body, err := SendPayUniRequest(&PostData, CreditCardCloseUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	info := new(CreditCardCloseEncryptInfoRes)
	StrData := Aes256GCMDecrypt(response.EncryptInfo, c.HashKey, c.HashIV)
	params := StrToMap(StrData)
	err = info.FillStruct(params)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c CreditCardCloseRequestCall) DoTest() (*CreditCardCloseEncryptInfoRes, error) {
	response := new(CreditCardCloseResponse)
	PostData := make(map[string]string)
	PostData["MerID"] = c.CreditCardCloseRequest.MerID
	PostData["Version"] = c.CreditCardCloseRequest.Version
	PostData["EncryptInfo"] = c.CreditCardCloseRequest.EncryptInfo
	PostData["HashInfo"] = c.CreditCardCloseRequest.HashInfo
	body, err := SendPayUniRequest(&PostData, TestCreditCardCloseUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	info := new(CreditCardCloseEncryptInfoRes)
	StrData := Aes256GCMDecrypt(response.EncryptInfo, c.HashKey, c.HashIV)
	params := StrToMap(StrData)
	err = info.FillStruct(params)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func (t *CreditCardCloseEncryptInfoRes) FillStruct(m map[string]string) error {
	for k, v := range m {
		err := SetField(t, k, v)
		if err != nil {
			fmt.Errorf(k, ":"+err.Error())
		}
	}
	return nil
}
