package payuni

import "fmt"

type UPPResult struct {
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
	//商店訂單編號
	MerTradeNo string `db:"merchant_order_no"`
	//UNi序號
	TradeNo string `db:"trade_no"`
	//交易金額
	TradeAmt string `db:"amt"`
	//支付狀態
	TradeStatus string `db:"trade_status"`
	//支付方式
	PaymentType string `db:"payment_type"`
	//**信用卡支付回傳參數（一次付清、Google Pay、Samaung Pay、國民旅遊卡、銀聯）**
	//卡號前六碼
	Card6No string `db:"card_six_no"`
	//卡號末四碼
	Card4No string `db:"card_four_no"`
	//分期-期別
	CardInst string `db:"inst"`
	//分期-首期金額
	FirstAmt string `db:"inst_first"`
	//分期-每期金額
	EachAmt string `db:"inst_each"`
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
	//**虛擬帳號回傳參數**
	//	銀行(代碼)
	BankType string `db:"bank_code"`
	//繳費設定
	PaySet string `db:"pay_set"`
	//**超商代碼繳費回傳參數**
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
	ICPNo string `db:"icp_no"`

	//Gateway
	Gateway string `db:"gateway"`
}

func NewUPPResult() *UPPResult {
	return &UPPResult{}
}

func (m UPPResult) DecodeTradeInfo(HashKey string, HashIV string) (*EncryptInfo, error) {
	info := new(EncryptInfo)
	StrData := Aes256GCMDecrypt(m.EncryptInfo, HashKey, HashIV)
	params := StrToMap(StrData)
	err := info.FillStruct(params)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func (t *EncryptInfo) FillStruct(m map[string]string) error {
	for k, v := range m {
		err := SetField(t, k, v)
		if err != nil {
			fmt.Errorf(k, ":"+err.Error())
		}
	}
	return nil
}
