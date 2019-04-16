package rsa

import (
	"bytes"
	"encoding/pem"
	"github.com/wenzhenxi/gorsa"
	"io/ioutil"
	"pem/key"
	"pem/util"
)

/*私钥加密*/
func Encrypt2PemByPriKey(input []byte, privateKey string, block pem.Block) ([]byte, error) {
	rsaSecurity := gorsa.RSASecurity{}
	if "" == privateKey {
		_ = rsaSecurity.SetPrivateKey(key.PriKey)
	} else {
		_ = rsaSecurity.SetPrivateKey(privateKey)
	}

	/*私钥加密生成密文*/
	cipherBytes, err := rsaSecurity.PriKeyENCTYPT(input)
	if err != nil {
		return nil, err
	}

	/*输出流*/
	output := bytes.NewBuffer(nil)
	_ = util.GeneratePemFile(block.Type, block.Headers, cipherBytes, output)

	return ioutil.ReadAll(output)
}

/*公钥解密*/
func Decrypt4PemByPubKey(input []byte, publicKey string) ([]byte, error) {
	rsaSecurity := gorsa.RSASecurity{}

	if "" == publicKey {
		_ = rsaSecurity.SetPublicKey(key.PubKey)
	} else {
		_ = rsaSecurity.SetPublicKey(publicKey)
	}

	_, plain := util.ParsePemFile(input)

	return rsaSecurity.PubKeyDECRYPT([]byte(plain))

}

/*公钥加密*/
func Encrypt2PemByPubKey(input []byte, publicKey string, block pem.Block) ([]byte, error) {
	rsaSecurity := gorsa.RSASecurity{}
	if "" == publicKey {
		_ = rsaSecurity.SetPublicKey(key.PubKey)
	} else {
		_ = rsaSecurity.SetPublicKey(publicKey)
	}

	cipherBytes, err := rsaSecurity.PubKeyENCTYPT(input)
	if err != nil {
		return nil, err
	}

	output := bytes.NewBuffer(nil)
	_ = util.GeneratePemFile(block.Type, block.Headers, cipherBytes, output)

	return ioutil.ReadAll(output)
}

/*私钥解密*/
func Decrypt4PemByPriKey(input []byte, privateKey string) ([]byte, error) {
	rsaSecurity := gorsa.RSASecurity{}
	if "" == privateKey {
		_ = rsaSecurity.SetPrivateKey(key.PriKey)
	} else {
		_ = rsaSecurity.SetPrivateKey(privateKey)
	}

	_, plain := util.ParsePemFile(input)

	return rsaSecurity.PriKeyDECRYPT([]byte(plain))

}
