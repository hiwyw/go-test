package main

import (
	"encoding/xml"
	"fmt"
)

const testInput = `<ROOT>
<RESULT_TABLE>
  <Code>1</Code>
</RESULT_TABLE>
<RESULT_TABLE>
  <Code>2</Code>
</RESULT_TABLE>
</ROOT>`

type MDM_LOG_STAUS struct {
	XMLName xml.Name `xml:"ROOT"`

	RESULT_TABLE []RESULT_TABLE_TMP `xml:"RESULT_TABLE,omitempty" json:"RESULT_TABLE,omitempty"`
}

type RESULT_TABLE_TMP struct {
	// 编码
	Code string `xml:"Code,omitempty" json:"Code,omitempty"`

	// 远程系统代码
	RemoteSysCode string `xml:"RemoteSysCode,omitempty" json:"RemoteSysCode,omitempty"`

	// 分发状态
	Status string `xml:"Status,omitempty" json:"Status,omitempty"`

	// 反馈消息
	Message string `xml:"Message,omitempty" json:"Message,omitempty"`

	// 处理时间
	ProcessTime string `xml:"ProcessTime,omitempty" json:"ProcessTime,omitempty"`

	// 分发时间
	SendTime string `xml:"SendTime,omitempty" json:"SendTime,omitempty"`
}

func main() {
	s := &MDM_LOG_STAUS{
		RESULT_TABLE: []RESULT_TABLE_TMP{
			RESULT_TABLE_TMP{
				Code: "1",
			},
			RESULT_TABLE_TMP{
				Code: "2",
			},
		},
	}

	out, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	s1 := MDM_LOG_STAUS{}
	if err := xml.Unmarshal([]byte(testInput), &s1); err != nil {
		panic(err)
	}
	fmt.Println(s1)
}
