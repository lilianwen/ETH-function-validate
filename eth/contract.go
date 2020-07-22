package eth

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

//根据RLP编码规则把int变量值转变成字节切片
func RlpInt2Bytes(i int)[]byte {
	var data [4]byte
	if i <= 255 {
		if i == 0 { //我靠，这个坑爹的玩意儿，害我好苦
			return nil
		}
		return []byte{byte(i)}
	} else  {
		binary.LittleEndian.PutUint32(data[:], uint32(i))
		if i<= 0xffff {
			return data[:2]
		} else if i <= 0xffffff {
			binary.LittleEndian.PutUint32(data[:], uint32(i))
			return data[:3]
		}
	}
	return data[:]
}

func Keccak256Hash(data []byte) []byte {
	keccak256Hash2 := sha3.NewLegacyKeccak256()
	keccak256Hash2.Write(data)
	return keccak256Hash2.Sum(nil)
}

//经测试，这种算法适合外部账号创建智能合约用
//同样是适用于简单的智能合约创建另一个智能合约
//但是不适用于用CREATE2 操作码创建新智能合约
func CreateContractAddr(senderAddr string, nonce int) (string, error) {
	var (
		data [][]byte
		buf []byte
		err error
	)
	if buf, err = hex.DecodeString(senderAddr); err != nil {
		return "",err
	}
	data = append(data, buf)
	buf = RlpInt2Bytes(nonce)
	data = append(data, buf)

	if buf, err = rlp.EncodeToBytes(data);err != nil {
		return "",nil
	}

	buf = Keccak256Hash(buf)
	return hex.EncodeToString(buf[12:]),nil
}

func Create2ContractAddr(sendAddr string, salt string, byteCode []byte) (string,error) {
	//keccak256( 0xff ++ sendAddr ++ salt ++ keccak256(byteCode))[12:]
	var (
		data []byte
		err error
		buf []byte
	)

	data = append(data, byte(0xff))

	if buf, err = hex.DecodeString(sendAddr); err != nil {
		return "",err
	}
	data = append(data, buf...)

	if buf, err = hex.DecodeString(salt); err != nil {
		return "",err
	}
	data = append(data, buf...)

	buf = Keccak256Hash(byteCode)
	data = append(data, buf...)

	buf = Keccak256Hash(data)
	return hex.EncodeToString(buf[12:]),nil
}

func MustGenerateMethodID(method string) string {
	buf := Keccak256Hash([]byte(method))
	return "0x"+hex.EncodeToString(buf[:4])
}
