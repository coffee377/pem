package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"github.com/coffee377/pem/util"
	"os"
)

var (
	PriKeyFile = "pri.key"
	PubKeyFile = "pub.key"
	PriKey     = `-----BEGIN PRIVATE KEY-----
Proc-Type: Golang

MIICWwIBAAKBgQClsRYOxb0QZO3yjEosgpu5dMjqOI35er9Gy5OU2o7T/ZRrRe7z
wVJZ85mbEMrhlHmFVeM/qBSLkNwFxULBkaboOpgkBJ1cSErtN1LpOAu81YzTK+T0
yumJNhkZFJCPbs3c021BxSYlsJXA9z8XgW8cdvmwtW3Hfq4fCb7aDtK0pQIDAQAB
AoGAPVHCNglcJPsVTJQ6xyo283WS/Yucj12r+ElS/t27yhyoluK0wiUjkwKiME8Q
pldKedwFVAOiRtx/cOtF/1kmCWdTciN9pJB+Sw88HtdzbPB3jdqrK6s0uVafsRW7
s8R1f9d23Tkp4fNbqXCVXNCh8fEQJtdH+BUzzssoPfScL8ECQQDUM1ws2H0PQvCk
hHhgNbOIqecTNiK7WsXCJDbp5tpP+AfUmUYknR00XW/QMl4iHJnrd5bN332bCIvQ
YIKK15iZAkEAx+Q0rPtFiocNTCkAZSSf23/0n1p2NSESErwJmyY9fSQau4BGWx07
8Cjl+ZIKI5K9yjQkrTMO6rhB5gQAMihH7QJADZETDZLxu+4NsJb/kzcbuVsTePj3
E39kMVtbX2zw/DfhWEhMYb7hxR1MLsVpm0i01ocYzyTAxQ6w0au57OKH0QJAcw0V
iI7jutT6wWBEGvMWk4c1bFbr/K55MZFLUiKTd6jFPjCZzi2oZxWTMK9u6IS4el7C
0XG230CIpXSeDECFDQJAfKdhnhW07c9lLTVyCTsPEcUj0gD91SApWGCmtyHo9pEw
SyVz4IOk++SRUcNWkj3YlFfb92rfksZXgahpqHDapw==
-----END PRIVATE KEY-----
`
	PubKey = `-----BEGIN PUBLIC KEY-----
Proc-Type: Golang

MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQClsRYOxb0QZO3yjEosgpu5dMjq
OI35er9Gy5OU2o7T/ZRrRe7zwVJZ85mbEMrhlHmFVeM/qBSLkNwFxULBkaboOpgk
BJ1cSErtN1LpOAu81YzTK+T0yumJNhkZFJCPbs3c021BxSYlsJXA9z8XgW8cdvmw
tW3Hfq4fCb7aDtK0pQIDAQAB
-----END PUBLIC KEY-----
`
)

/* 生成密钥对文件 */
func GenRsaKey(bits int, keyPairDir string) error {

	/*1.生成私钥*/
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	derPrivateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	if err != nil {
		return err
	}
	/*2.生成公钥*/
	publicKey := &privateKey.PublicKey
	derPublicStream, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	/*3.创建密钥存放目录*/
	if "" == keyPairDir {
		keyPairDir = "."
	}
	err = os.MkdirAll(keyPairDir, os.ModePerm)

	/*4.生成私钥文件*/
	priFile, err := os.Create(keyPairDir + "/" + PriKeyFile)
	if err != nil {
		return err
	}
	_ = util.GenPriKey(derPrivateStream, priFile)
	defer priFile.Close()

	/*5.生成公钥文件*/
	pubFile, err := os.Create(keyPairDir + "/" + PubKeyFile)
	if err != nil {
		return err
	}
	defer pubFile.Close()
	_ = util.GenPubKey(derPublicStream, pubFile)

	return nil
}
