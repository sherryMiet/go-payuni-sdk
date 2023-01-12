package payuni

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const (
	TestQueryTradeInfoUrl = "https://ccore.newebpay.com/API/QueryTradeInfo"
	QueryTradeInfoUrl     = "https://core.newebpay.com/API/QueryTradeInfo"
	QueryTradeInfoVersion = "1.3"
)

type QueryTradeInfoRequestCall struct {
	//串接程式版本
	Version string
	//回傳格式
	RespondType RespondType
	//檢查碼
	CheckValue string
	//時間戳記
	TimeStamp string
	//資料來源
	Gateway               string
	QueryTradeInfoRequest *QueryTradeInfoRequest
}
type QueryTradeInfoRequest struct {
	//商店代號
	MerchantID string
	//商店訂單編號
	MerchantOrderNo string
	//訂單金額
	Amt int
}

func NewQueryTradeInfoRequest() *QueryTradeInfoRequest {
	return &QueryTradeInfoRequest{}
}
func (q *QueryTradeInfoRequest) SetQueryTradeInfoRequest(MerchantOrderNo string, Amt int) {
	q.MerchantOrderNo = MerchantOrderNo
	q.Amt = Amt
}

type QueryTradeInfoResponse struct {
	//回傳狀態
	Status string
	//回傳訊息
	Message string
	//回傳內容
	Result QueryTradeInfoResult
}

type QueryTradeInfoResult struct {
	//藍新金流商店代號
	MerchantID string
	//交易金額
	Amt int
	//藍新金流交易序號
	TradeNo string
	//支付狀態 以數字回應，其代表下列意含: 0=未付款 1=付款成功 2=付款失敗 3=取消付款
	MerchantOrderNo string
	//支付狀態
	TradeStatus string
	//支付方式
	PaymentType string
	//交易建立時間
	CreateTime string
	//支付完成時間
	PayTime string
	//檢核碼
	CheckCode string
	//預計撥款日
	FundTime string
	//實際交易藍新金流金流特店或閘道特店之商店代號
	ShopMerchantId string
	//信用卡專屬欄位:當該筆交易為信用卡交易時(包含:國外卡、國旅卡、ApplePay、 GooglePay、SamsungPay)，會回應下列欄位
	//金融機構回應碼
	RespondCode string
	//授權碼
	Auth string
	//ECI
	ECI string
	//請款金額
	CloseAmt string
	//請款狀態
	CloseStatus string
	//可退款餘額
	BackBalance string
	//退款狀態
	BackStatus string
	//授權結果訊息
	RespondMsg string
	//分期-期別
	Inst string
	//分期-首期金額
	InstFirst string
	//分期-每期金額
	InstEach string
	//交易類別
	PaymentMethod string
	//信用卡前 6 碼
	Card6No string
	//信用卡後 4 碼
	Card4No string
	//收單金融機構
	AuthBank string
	//超商代碼、超商條碼、ATM、WebATM 專屬欄位:當該筆交易為此四種支付方式時，會回應下列欄位
	//付款資訊
	PayInfo string
	//繳費有效期限
	ExpireDate string
	//交易狀態
	OrderStatus string
	//超商取貨付款專屬欄位:當該筆交易為超商取貨付款時，會回應下列欄位:
	//超商類別名稱
	StoreType string
	//超商門市編號
	StoreCode string
	//超商門市名稱
	StoreName string
	//物流訂單編號
	LgsNo string
	//物流型態
	LgsType string
}

func (c *Client) QueryTradeInfo(Request *QueryTradeInfoRequest) *QueryTradeInfoRequestCall {
	call := new(QueryTradeInfoRequestCall)
	Request.MerchantID = c.MerID
	params := StructToParamsMap(Request)
	paramStr := NewValuesFromMap(params).Encode()
	call.CheckValue = SHA256("IV=" + c.HashIV + "&" + paramStr + "&Key=" + c.HashKey)
	call.Version = QueryTradeInfoVersion
	call.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
	if strings.Contains(c.MerID, "MS5") {
		call.Gateway = "Composite"
	}
	call.QueryTradeInfoRequest = Request
	return call
}

func (q *QueryTradeInfoRequestCall) Do() (*QueryTradeInfoResponse, error) {
	response := new(QueryTradeInfoResponse)

	PostData := make(map[string]string)
	PostData["MerchantID"] = q.QueryTradeInfoRequest.MerchantID
	PostData["Version"] = q.Version
	PostData["RespondType"] = string(q.RespondType)
	PostData["CheckValue"] = q.CheckValue
	PostData["TimeStamp"] = q.TimeStamp
	PostData["MerchantOrderNo"] = q.QueryTradeInfoRequest.MerchantOrderNo
	PostData["Amt"] = strconv.Itoa(q.QueryTradeInfoRequest.Amt)
	PostData["Gateway"] = q.Gateway

	body, err := SendPayUniRequest(&PostData, TestQueryTradeInfoUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (q *QueryTradeInfoRequestCall) DoTest() (*QueryTradeInfoResponse, error) {

	response := new(QueryTradeInfoResponse)

	PostData := make(map[string]string)
	PostData["MerchantID"] = q.QueryTradeInfoRequest.MerchantID
	PostData["Version"] = q.Version
	PostData["RespondType"] = string(q.RespondType)
	PostData["CheckValue"] = q.CheckValue
	PostData["TimeStamp"] = q.TimeStamp
	PostData["MerchantOrderNo"] = q.QueryTradeInfoRequest.MerchantOrderNo
	PostData["Amt"] = strconv.Itoa(q.QueryTradeInfoRequest.Amt)
	PostData["Gateway"] = q.Gateway

	body, err := SendPayUniRequest(&PostData, TestQueryTradeInfoUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
