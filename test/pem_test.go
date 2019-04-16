package test

import (
	"encoding/pem"
	"pem/key"
	"pem/rsa"
	"testing"
)

/*生成密钥对*/
func TestGenKeyPair(t *testing.T) {
	err := key.GenRsaKey(1024, "./key")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("密钥对生成成功")
	}
}

/*私钥加密公钥解密*/
func TestHelloWorld(t *testing.T) {
	t.Log("私钥加密公钥解密")
	txt := "测试文件内容"
	block := pem.Block{Type: "ENCRYPT DEMO",
		Headers: map[string]string{"PROC_TYPE": "Golang",}}
	cipherText, _ := rsa.Encrypt2PemByPriKey([]byte(txt), "", block)
	plainText, _ := rsa.Decrypt4PemByPubKey(cipherText, "")
	if txt != string(plainText) {
		t.Error("失败")
	} else {
		t.Log("成功")
	}
}

/*公钥加密私钥解密*/
func TestHelloWorld2(t *testing.T) {
	t.Log("公钥加密私钥解密")
	txt := "测试文件内容"
	block := pem.Block{Type: "ENCRYPT DEMO",
		Headers: map[string]string{"PROC_TYPE": "Golang",}}
	cipherText, err := rsa.Encrypt2PemByPubKey([]byte(txt), "", block)
	plainText, err := rsa.Decrypt4PemByPriKey(cipherText, "")
	if txt != string(plainText) {
		t.Error(err)
	} else {
		t.Log("成功")
	}
}
