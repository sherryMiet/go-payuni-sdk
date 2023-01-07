package payuni

import (
	"encoding/json"
	"fmt"
)

type MPGGatewayResult struct {
	//回傳狀態
	Status string
	//回傳訊息
	MerID string
	//加密字串
	EncryptInfo string
	//加密 Hash
	HashInfo string
	//串接程式版本
	Version string
}

type EncryptInfo struct {
	//回傳狀態
	Status string `db:"status"`
	//回傳訊息
	Message string `db:"message"`
	// **所有支付方式共同回傳參數**
	//商店代號
	MerID string `db:"merchant_id"`
	//交易金額
	TradeAmt int `db:"amt"`
	//UNi序號
	TradeNo string `db:"trade_no"`
	//商店訂單編號
	MerTradeNo string `db:"merchant_order_no"`
	//支付狀態
	TradeStatus int `db:"trade_status"`
	//支付方式
	PaymentType string `db:"payment_type"`
	//交易 IP
	IP string `db:"ip"`
	//款項保管銀行
	EscrowBank string `db:"escrow_bank"`
	//**信用卡支付回傳參數（一次付清、Google Pay、Samaung Pay、國民旅遊卡、銀聯）**
	//卡號前六碼
	Card6No string `db:"card_six_no"`
	//卡號末四碼
	Card4No string `db:"card_four_no"`
	//分期-期別
	CardInst int `db:"inst"`
	//分期-首期金額
	FirstAmt int `db:"inst_first"`
	//分期-每期金額
	EachAmt int `db:"inst_each"`
	//回應碼
	ResCode string `db:"res_code"`
	//回應碼敘述
	ResCodeMsg string `db:"res_code_msg"`
	//授權碼
	AuthCode string `db:"auth_code"`
	//收單金融機構
	AuthBank string `db:"auth_bank"`
	//授權銀行(名稱)
	AuthBankName string `db:"auth_bank_name"`
	//授權類型
	AuthType string
	//授權日期
	AuthDay string
	//授權時間
	AuthTime string
	//信用卡Token Hash
	CreditHash string
	//信用卡Token 有效日期
	CreditLife string
	//	發卡銀行(代碼)
	CardBank string
	//**WEBATM、ATM 繳費回傳參數**
	//	銀行(代碼)
	BankType string `db:"bank_code"`
	//繳費設定
	PaySet string `db:"pay_set"`
	//**超商代碼繳費回傳參數**
	//繳費代碼
	CodeNo string `db:"code_no"`
	//繳費門市類別 超商類別名稱
	StoreType string `db:"store_type"`
	//超商(代碼)
	Store string `db:"store_id"`
	//驗證碼
	VerifyCode string `db:"verify_code"`
	//**愛金卡（ＩＣＡＳＨ）**
	//支付完成時間
	PayTime string `db:"pay_time"`
	//**後支付 (Aftee)**
	//**ATM、超商代碼共同回傳參數**
	//繳費截止日期
	ExpireDate string `db:"expire_date"`
	//繳費虛擬帳號.繳費代碼
	PayNo string `db:"pay_no"`
	//**愛金卡（ＩＣＡＳＨ）、後支付 (Aftee)共同回傳參數**
	//愛金卡.Aftee交易序號
	ICPNo string
}

type Result struct {
	// **所有支付方式共同回傳參數**
	//商店代號
	MerchantID string `db:"merchant_id"`
	//交易金額
	Amt int `db:"amt"`
	//藍新金流交易序號
	TradeNo string `db:"trade_no"`
	//商店訂單編號
	MerchantOrderNo string `db:"merchant_order_no"`
	//支付狀態
	TradeStatus int `db:"trade_status"`
	//支付方式
	PaymentType string `db:"payment_type"`
	//回傳格式
	RespondType string `db:"response_type"`
	//支付完成時間
	PayTime string `db:"pay_time"`
	//交易 IP
	IP string `db:"ip"`
	//款項保管銀行
	EscrowBank string `db:"escrow_bank"`
	//**信用卡支付回傳參數（一次付清、Google Pay、Samaung Pay、國民旅遊卡、銀聯）**
	//收單金融機構
	AuthBank string `db:"auth_bank"`
	//金融機構回應碼
	RespondCode string `db:"response_code"`
	//授權碼
	Auth string `db:"auth"`
	//卡號前六碼
	Card6No string `db:"card_six_no"`
	//卡號末四碼
	Card4No string `db:"card_four_no"`
	//分期-期別
	Inst int `db:"inst"`
	//分期-首期金額
	InstFirst int `db:"inst_first"`
	//分期-每期金額
	InstEach int `db:"inst_each"`
	//ECI
	//1.3D 回傳值 eci=1,2,5,6，代表為 3D 交易。
	//2.若交易送至收單機構授權時已是失敗狀態，則本欄位的值會以空值回傳。
	ECI string `db:"eci"`
	//信用卡快速結帳使用狀態
	TokenUseStatus int `db:"token_use_status"`
	//紅利折抵後實際金額
	RedAmt string `db:"red_amt"`
	//交易類別
	PaymentMethod string `db:"payment_method"`
	//外幣金額
	DCC_Amt int `db:"dcc_amt"`
	//匯率
	DCC_Rate int `db:"dcc_rate"`
	//風險匯率
	DCC_Markup int `db:"dcc_markup"`
	//幣別
	DCC_Currency string `db:"dcc_currency"`
	//幣別代碼
	DCC_Currency_Code string `db:"dcc_currency_code"`
	//**WEBATM、ATM 繳費回傳參數**
	//付款人金融機構代碼
	PayBankCode string `db:"pay_bank_code"`
	//付款人金融機構帳號末五碼
	PayerAccount5Code string `db:"payer_account_five_code"`
	//**超商代碼繳費回傳參數**
	//繳費代碼
	CodeNo string `db:"code_no"`
	//繳費門市類別 超商類別名稱
	StoreType string `db:"store_type"`
	//繳費門市代號
	StoreID string `db:"store_id"`
	//**超商條碼繳費回傳參數**
	//第一段條碼
	Barcode_1 string `db:"barcode_1"`
	//第二段條碼
	Barcode_2 string `db:"barcode_2"`
	//第三段條碼
	Barcode_3 string `db:"barcode_3"`
	//付款次數
	RepayTimes string `db:"repay_times"`
	//繳費超商
	PayStore string `db:"pay_store"`
	//**超商物流回傳參數**
	//超商門市編號
	StoreCode string `db:"store_code"`
	//超商門市名稱
	StoreName string `db:"store_name"`
	//超商門市地址
	StoreAddr string `db:"store_addr"`
	//取件交易方式
	TradeType string `db:"trade_type"`
	//取貨人
	CVSCOMName string `db:"cvs_com_name"`
	//取貨人手機號碼
	CVSCOMPhone string `db:"cvs_com_phone"`
	//物流寄件單號
	LgsNo string `db:"lgs_no"`
	//物流型態
	LgsType string `db:"lgs_type"`
	//**跨境支付回傳參數(包含簡單付電子錢包、簡單付微信支付、簡單付支付寶)**
	//跨境通路類型
	ChannelID string `db:"channel_id"`
	//跨境通路 交易序號
	ChannelNo string `db:"channel_no"`
	//**玉山 Wallet 回傳參數** **台灣 Pay 回傳參數**
	//實際付款金額
	PayAmt string `db:"pay_amt"`
	//紅利折抵金額
	RedDisAmt string `db:"red_dis_amt"`
	//**ATM、超商代碼、超商條碼、超商取貨付款共同回傳參數**
	//繳費截止日期
	ExpireDate string `db:"expire_date"`
	//金融機構代碼
	BankCode string `db:"bank_code"`
}

func NewMPGGatewayResult() *MPGGatewayResult {
	return &MPGGatewayResult{}
}
func (m MPGGatewayResult) DecodeTradeInfo(HashKey string, HashIV string) (*EncryptInfo, error) {
	info := new(EncryptInfo)
	StrData := Aes256GCMDecrypt(m.EncryptInfo, HashKey, HashIV)
	err := json.Unmarshal([]byte(StrData), info)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	return info, nil
}
func (t *EncryptInfo) FillStruct(m map[string]string) error {
	for k, v := range m {
		err := SetField(t, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
