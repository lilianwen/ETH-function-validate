package eth

import "testing"

func TestNewAddress(t *testing.T) {
	var testcases = []struct {
		privk string
		address string
	} {
		{
			"E452FE5BB764BE2F3E030B6CE5C16E7A68BB341D0D2FDE13AAE5347690F12B52",
			"0x50352B904445576242444bc1924e93e61090738C",
		},
	}

	var (
		got string
		err error
	)
	for _, oneCase := range testcases {
		if got,err = NewAddress(oneCase.privk); err != nil {
			t.Error(err)
			return
		}
		if got != oneCase.address {
			t.Error("new address error")
			t.Error("want: ", oneCase.address)
			t.Error("got: ", got)
			return
		}
	}
}

func TestCheckEthAddress(t *testing.T) {
	var testcases = []struct{
		address string
		want bool
	} {
		{
			"0xD4a16aa11Bd0D3315698792F5E1F66770F9Cd78F",
			true,
		},
		{
			"0x40DAB7E81503AA1F8c1ef3574842017277755646",
			true,
		},
		{
			"0x50352B904445576242444bc1924e93e61090738c",
			false,
		},
	}

	for _, oneCase := range testcases {
		if oneCase.want != CheckEthAddress(oneCase.address) {
			t.Error("check eth address error")
			t.Error("address:", oneCase.address)
			return
		}
	}
}