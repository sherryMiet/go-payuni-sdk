package payuni

import (
	"bytes"
	"fmt"
	"github.com/sony/sonyflake"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

type (
	RespondType string
	IndexType   int
)

const (
	RespondType_JSON          RespondType = "JSON"
	RespondType_STRING        RespondType = "STRING"
	IndexType_MerchantOrderNo IndexType   = 1
	IndexType_TradeNo         IndexType   = 2
)

func PtrNilString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func GenSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(id, 16)
}

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" method="POST">
{{range $key,$element := .Values}}    <input type="hidden" name="{{$key}}" id="{{$key}}" value="{{$element}}" />
{{end -}}
</form>
<script>document.querySelector("#order_form").submit();</script>`

type OrderTmplArgs struct {
	Values map[string]string
	Action string
}

var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func GenerateAutoSubmitHtmlForm(params map[string]string, targetUrl string) string {

	var result bytes.Buffer
	err := OrderTmpl.Execute(&result, OrderTmplArgs{
		Values: params,
		Action: targetUrl,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}
func SendNewebPayRequest(postData *map[string]string, URL string) ([]byte, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, _ := http.NewRequest(http.MethodPost, URL, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := http.DefaultClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Printf("%s", data)
	return data, nil
}
