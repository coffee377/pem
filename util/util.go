package util

import (
	"encoding/pem"
	"io"
)

/* 根据 Block 生成 PEM 格式文件*/
func GeneratePemFile4Block(block *pem.Block, out io.Writer) error {
	return pem.Encode(out, block)
}

/* 生成 PEM 格式文件*/
func GeneratePemFile(typeName string, headers map[string]string, plainDatas []byte, out io.Writer) error {
	block := &pem.Block{
		Bytes: plainDatas,
	}

	/* PEM 类型不为空 */
	if "" != typeName {
		block.Type = typeName
	} else {
		block.Type = "ENCRYPT DATA"
	}

	/* PEM 头信息存相关信息 */
	if nil != headers && len(headers) > 0 {
		block.Headers = headers
	}

	return GeneratePemFile4Block(block, out)
}

/* 解析 PEM 文件 */
func ParsePemFile(input []byte) (p *pem.Block, output []byte) {
	block, _ := pem.Decode([]byte(input))
	return block, block.Bytes
}

/* 生成帆软授权文件 */
func GenerateFRLicense(datas []byte, out io.Writer) error {
	var (
		LicenceType = "LICENCE DATA"
		LicenceHead = map[string]string{
			"Proc-Type":           "FineSoft Licence",
			"License-Restriction": "Only for test! Please support genuine!!!"}
	)
	return GeneratePemFile(LicenceType, LicenceHead, datas, out)
}

/* 生成私钥文件 */
func GenPriKey(datas []byte, out io.Writer) error {
	var (
		LicenceType = "PRIVATE KEY"
		LicenceHead = map[string]string{
			"Proc-Type": "Golang",
		}
	)
	return GeneratePemFile(LicenceType, LicenceHead, datas, out)
}

/* 生成公钥文件 */
func GenPubKey(datas []byte, out io.Writer) error {
	var (
		LicenceType = "PUBLIC KEY"
		LicenceHead = map[string]string{
			"Proc-Type": "Golang",
		}
	)
	return GeneratePemFile(LicenceType, LicenceHead, datas, out)
}
