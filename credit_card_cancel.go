package payuni

//
//import (
//	"encoding/json"
//	"strconv"
//	"time"
//)
//
//const (
//	TestCreditCardCancelUrl = "https://ccore.newebpay.com/API/CreditCard/Cancel"
//	CreditCardCancelUrl     = "https://core.newebpay.com/API/CreditCard/Cancel"
//	CreditCardCancelVersion = "1.0"
//)
//
//type CreditCardCancelRequestCall struct {
//	//商店代號
//	MerchantID_ string
//	//加密資料
//	PostData_ string
//}
//type CreditCardCancelRequest struct {
//	//回傳格式
//	RespondType RespondType
//	//串接程式版本
//	Version string
//	//取消授權金額
//	Amt int
//	//商店訂單編號
//	MerchantOrderNo string
//	//藍新金流交易序號
//	TradeNo string
//	//單號類別
//	IndexType IndexType
//	//時間戳記
//	TimeStamp string
//}
//type CreditCardCancelResponse struct {
//	//回傳狀態
//	Status string
//	//回傳訊息
//	Message string
//	//回傳資料
//	Result CreditCardCancelResult
//}
//type CreditCardCancelResult struct {
//	//商店代號
//	MerchantID string
//	//藍新金流交易序號
//	TradeNo string
//	//交易金額
//	Amt int
//	//商店訂單編號
//	MerchantOrderNo string
//	//檢核碼
//	CheckCode string
//}
//
//func NewCreditCardCancelRequest() *CreditCardCancelRequest {
//	return &CreditCardCancelRequest{}
//}
//
//func (c *CreditCardCancelRequest) SetIndexType(indexType IndexType) {
//	c.IndexType = indexType
//}
//
//func (c *CreditCardCancelRequest) SetCreditCardCancelRequest(Amt int, MerchantOrderNo, TradeNo string) {
//	c.Amt = Amt
//	c.MerchantOrderNo = MerchantOrderNo
//	c.TradeNo = TradeNo
//}
//
//func (c *Client) CreditCardCancel(Data *CreditCardCancelRequest) *CreditCardCancelRequestCall {
//	Data.RespondType = RespondType_JSON
//	Data.Version = CreditCardCancelVersion
//	Data.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
//	params := StructToParamsMap(Data)
//	paramStr := NewValuesFromMap(params).Encode()
//	PostData := Aes256(paramStr, c.HashKey, c.HashIV)
//	return &CreditCardCancelRequestCall{
//		MerchantID_: c.MerID,
//		PostData_:   PostData,
//	}
//}
//
//func (c CreditCardCancelRequestCall) Do() (*CreditCardCancelResponse, error) {
//	response := new(CreditCardCancelResponse)
//	PostData := make(map[string]string)
//	PostData["MerchantID_"] = c.MerchantID_
//	PostData["PostData_"] = c.PostData_
//	body, err := SendNewebPayRequest(&PostData, CreditCardCancelUrl)
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(body, response)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}
//
//func (c CreditCardCancelRequestCall) DoTest() (*CreditCardCancelResponse, error) {
//	response := new(CreditCardCancelResponse)
//	PostData := make(map[string]string)
//	PostData["MerchantID_"] = c.MerchantID_
//	PostData["PostData_"] = c.PostData_
//	body, err := SendNewebPayRequest(&PostData, TestCreditCardCancelUrl)
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(body, response)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}
