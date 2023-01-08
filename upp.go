package payuni

import "time"

const (
	Version    = "1.0"
	TestUPPUrl = "https://sandbox-api.payuni.com.tw/api/upp"
	UPPUrl     = "https://api.payuni.com.tw/api/upp"
)

type UPPRequestCall struct {
	HashKey    string
	HashIV     string
	UPPRequest *UPPRequest
}

type UPPRequest struct {
	//商店代號
	MerID string `json:"MerchantID"`
	//版本
	Version string
	//交易資料 AES加密字串
	EncryptInfo string `json:"EncryptInfo"`
	//交易資料 SHA256加密字串
	HashInfo string `json:"HashInfo"`
}

type UPPTradeInfo struct {
	//商店代號
	MerID string `json:"MerID"`
	//商店訂單編號
	MerTradeNo string `json:"MerTradeNo"`
	//訂單金額
	TradeAmt int `json:"TradeAmt,omitempty"`
	//時間戳記
	Timestamp int `json:"Timestamp,omitempty"`
	//繳費有效期限
	ExpireDate string `json:"ExpireDate,omitempty"`
	//支付完成 返回商店網址
	ReturnURL string `json:"ReturnURL,omitempty"`
	//支付通知網址
	NotifyURL string `json:"NotifyURL,omitempty"`
	//付款頁面顯示返回按鈕點擊返回後指定網址
	BackURL string `json:"BackURL,omitempty"`
	//付款人電子信箱
	UsrMail string `json:"UsrMail,omitempty"`
	//信用卡Token 類型
	UseTokenType int `json:"UseTokenType,omitempty"`
	//信用卡(記憶卡號) 顯示類型
	CreditShowType int `json:"CreditShowType,omitempty"`
	//信用卡Token
	CreditToken string `json:"CreditToken,omitempty"`
	//商品說明
	ProdDesc string `json:"ProdDesc,omitempty"`
	//信用卡 Token 紀錄類型
	CreditTokenType int `json:"CreditTokenType,omitempty"`
	//信用卡 Token 有效期間
	CreditTokenExpired string `json:"CreditTokenExpired,omitempty"`
	//信用卡 一次付清啟用
	Credit int `json:"CREDIT,omitempty"`
	//icash Pay支付
	ICash int `json:"ICash,omitempty"`
	//AFTEE先享後付
	Aftee int `json:"Aftee,omitempty"`
	//信用卡分期支付 同時開啟多期別時，將此參數用”，”(半形) 分隔，例如:3,6,12，代表開啟 分 3、6、12 期的功能
	CreditInst string `json:"CreditInst,omitempty"`
	//信用卡紅利啟用
	CreditRed int `json:"CreditRed,omitempty"`
	//信用卡 銀聯卡啟用
	CreditUnionPay int `json:"CreditUnionPay,omitempty"`
	//虛擬帳號支付
	ATM int `json:"ATM,omitempty"`
	//金融機構
	BankType string `json:"BankType,omitempty"`
	//超商代碼繳費 啟用
	CVS int `json:"CVS,omitempty"`
	//ApplePay
	ApplePay int `json:"ApplePay,omitempty"`
}

type AtmPayment struct {
	//繳費有效期限
	ExpireDate string `json:"ExpireDate,omitempty"`
	//ATM 啟用
	ATM int `json:"ATM,omitempty"`
}

func NewUPPTradeInfo() *UPPTradeInfo {
	return &UPPTradeInfo{}
}

func (m *UPPTradeInfo) SetExpireDate(ExpireDate string) *UPPTradeInfo {
	m.ExpireDate = ExpireDate
	return m
}

func (m *UPPTradeInfo) SetCREDIT(InstFlag string, CreditRed int, CreditUnionPay int) *UPPTradeInfo {
	m.Credit = 1
	m.CreditInst = InstFlag
	m.CreditRed = CreditRed
	m.CreditUnionPay = CreditUnionPay
	return m
}

func (m *UPPTradeInfo) SetATM() *UPPTradeInfo {
	m.ATM = 1
	return m
}

func (m *UPPTradeInfo) SetICash() *UPPTradeInfo {
	m.ICash = 1
	return m
}

func (m *UPPTradeInfo) SetAFTEE() *UPPTradeInfo {
	m.Aftee = 1
	return m
}

func (m *UPPTradeInfo) SetCVS() *UPPTradeInfo {
	m.CVS = 1
	return m
}

func (m *UPPTradeInfo) SetApplePay() *UPPTradeInfo {
	m.ApplePay = 1
	return m
}

func (m *UPPTradeInfo) CreateOrder(MerTradeNo string, TradeAmt int, ProdDesc string) *UPPTradeInfo {
	m.MerTradeNo = MerTradeNo
	m.TradeAmt = TradeAmt
	m.ProdDesc = ProdDesc
	return m
}

func (c *Client) UPP(Data *UPPTradeInfo) *UPPRequestCall {
	Data.MerID = c.MerID
	Data.Timestamp = int(time.Now().Unix())
	params := StructToParamsMap(Data)
	paramStr := ParamsMapToURLEncode(params)
	req := new(UPPRequest)
	req.MerID = c.MerID
	req.EncryptInfo = Aes256GCMEncrypt(paramStr, c.HashKey, c.HashIV)
	req.HashInfo = SHA256(c.HashKey + req.EncryptInfo + c.HashIV)
	req.Version = Version
	return &UPPRequestCall{
		HashIV:     c.HashIV,
		HashKey:    c.HashKey,
		UPPRequest: req,
	}
}

func (m *UPPRequestCall) Do() string {
	params := StructToParamsMap(m.UPPRequest)
	html := GenerateAutoSubmitHtmlForm(params, UPPUrl)
	return html
}

func (m *UPPRequestCall) DoTest() string {
	params := StructToParamsMap(m.UPPRequest)
	html := GenerateAutoSubmitHtmlForm(params, TestUPPUrl)
	return html
}
