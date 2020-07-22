package eth

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"paper/common"
	"strings"
)

func NewAddress(strPrivKey string) (string, error) {
	var (
		uncompressPubKey []byte
		pubKeyHash256 []byte
		err error
		buf []byte
	)
	if buf, err = hex.DecodeString(strPrivKey); err != nil {
		return "", err
	}
	privKey := secp256k1.PrivKeyFromBytes(buf)

	//拿公钥（非压缩公钥）来hash，计算公钥的 Keccak-256 哈希值（32bytes）
	uncompressPubKey = append(privKey.PubKey().X().Bytes(), privKey.PubKey().Y().Bytes()...)

	//计算公钥的Keccak256哈希值
	pubKeyHash256 = common.Keccak256Hash(uncompressPubKey)

	//取上一步结果取后20bytes即以太坊地址
	address := string("0x")+hex.EncodeToString(pubKeyHash256[len(pubKeyHash256)-20:])

	return ToValidateAddress(string(address)), nil
}
func ToValidateAddress(address string) string {
	addrLowerStr := strings.ToLower(address)
	if strings.HasPrefix(addrLowerStr, "0x") {
		addrLowerStr = addrLowerStr[2:]
		address = address[2:]
	}
	var binaryStr string
	addrBytes := []byte(addrLowerStr)
	hash256 := common.Keccak256Hash([]byte(addrLowerStr))//注意，这里是直接对字符串转换成byte切片然后哈希

	for i, e := range addrLowerStr {
		//如果是数字则跳过
		if e>='0' && e<='9' {
			continue
		} else {
			binaryStr = fmt.Sprintf("%08b", hash256[i/2])//注意，这里一定要填充0
			if binaryStr[4*(i % 2)] == '1'{
				addrBytes[i] -= 32
			}
		}
	}

	return "0x"+string(addrBytes)
}

//检查有大小写区别的以太坊地址是否合法
func CheckEthAddress(address string) bool {
	return ToValidateAddress(address) == address
}
