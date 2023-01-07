package payuni

//import (
//	"encoding/json"
//	"strconv"
//	"time"
//)
//
//const (
//	TestCreditCardCloseUrl           = "https://ccore.newebpay.com/API/CreditCard/Close"
//	CreditCardCloseUrl               = "https://core.newebpay.com/API/CreditCard/Close"
//	CreditCardCloseVersion           = "1.1"
//	CloseType_Ask          CloseType = 1
//	CloseType_Cancel       CloseType = 2
//	Cancel_Y               Cancel    = 1
//)
//
//type (
//	CloseType int
//	Cancel    int
//)
//type CreditCardCloseRequestCall struct {
//	//商店代號
//	MerchantID_ string
//	//加密資料
//	PostData_ string
//}
//type CreditCardCloseRequest struct {
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
//	//請款或退款
//	CloseType CloseType
//	//取消請款或退款
//	Cancel Cancel
//}
//type CreditCardCloseResponse struct {
//	//回傳狀態
//	Status string
//	//回傳訊息
//	Message string
//	//回傳資料
//	Result CreditCardCloseResult
//}
//type CreditCardCloseResult struct {
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
//func NewCreditCardCloseRequest() *CreditCardCloseRequest {
//	return &CreditCardCloseRequest{}
//}
//func (c *CreditCardCloseRequest) SetCloseType(closeType CloseType) {
//	c.CloseType = closeType
//}
//func (c *CreditCardCloseRequest) SetIndexType(indexType IndexType) {
//	c.IndexType = indexType
//}
//func (c *CreditCardCloseRequest) SetCancel() {
//	c.Cancel = Cancel_Y
//}
//func (c *CreditCardCloseRequest) SetCreditCardCloseRequest(Amt int, MerchantOrderNo, TradeNo string) {
//	c.Amt = Amt
//	c.MerchantOrderNo = MerchantOrderNo
//	c.TradeNo = TradeNo
//}
//func (c *Client) CreditCardClose(Data *CreditCardCloseRequest) *CreditCardCloseRequestCall {
//	Data.RespondType = RespondType_JSON
//	Data.Version = CreditCardCloseVersion
//	Data.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
//	params := StructToParamsMap(Data)
//	paramStr := NewValuesFromMap(params).Encode()
//	PostData := Aes256(paramStr, c.HashKey, c.HashIV)
//	return &CreditCardCloseRequestCall{
//		MerchantID_: c.MerID,
//		PostData_:   PostData,
//	}
//}
//
//func (c CreditCardCloseRequestCall) Do() (*CreditCardCloseResponse, error) {
//	response := new(CreditCardCloseResponse)
//	PostData := make(map[string]string)
//	PostData["MerchantID_"] = c.MerchantID_
//	PostData["PostData_"] = c.PostData_
//	body, err := SendNewebPayRequest(&PostData, CreditCardCloseUrl)
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
//
//func (c CreditCardCloseRequestCall) DoTest() (*CreditCardCloseResponse, error) {
//	response := new(CreditCardCloseResponse)
//	PostData := make(map[string]string)
//	PostData["MerchantID_"] = c.MerchantID_
//	PostData["PostData_"] = c.PostData_
//	body, err := SendNewebPayRequest(&PostData, TestCreditCardCloseUrl)
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
