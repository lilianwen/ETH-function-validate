package eth

import "testing"

func TestCreateContractAddr(t *testing.T) {
	var testcases = []struct {
		sender string
		nonce int
		want string
	} {
		{
			"D4a16aa11Bd0D3315698792F5E1F66770F9Cd78F",
			2,
			"a79fa249cad974b1f40124fd11452f8dc325440c",
		},
		{
			"d4a16aa11bd0d3315698792f5e1f66770f9cd78f",
			0,
			"7e4ca94147bea90fe22575e92f89b186af3ea523",
		},
		{
			"6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0",
			0,
			"cd234a471b72ba2f1ccf0a70fcaba648a5eecd8d",
		},
		{
			"6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0",
			1,
			"343c43a37d37dff08ae8c4a11544c718abb4fcf8",
		},
		{
			"6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0",
			2,
			"f778b86fa74e846c4f0a1fbd1335fe81c00a0c91",
		},
		{
			"6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0",
			3,
			"fffd933a0bc612844eaf0c6fe3e5b8e9b6c1d19c",
		},
		{
			"B20e2D38128E976b6B4852293C4e32f3A8D75C22",
			1,
			"53c4554e4e3dc3bf299a46f51dd3dbccaa5ea47a",
		},
		{
			"B20e2D38128E976b6B4852293C4e32f3A8D75C22",
			2,
			"1a125fd865e0f21ade5a37a54f22b08a8e154273",
		},
		//{//这个例子不是用CREATE操作码实现的，而是用CREATE2操作码实现的
		//	"40c84310ef15b0c0e5c69d25138e0e16e8000fe9",
		//	6816,
		//	"8b55c928602896a1e078e23a3fee33393821eec7",
		//},
	}
	for _, oneCase := range testcases {
		var got string
		var err error
		if got, err = CreateContractAddr(oneCase.sender, oneCase.nonce); err != nil {
			t.Error(err)
			return
		}

		if got != oneCase.want {
			t.Error("generate contract address error")
			t.Error("want:", oneCase.want)
			t.Error("got: ", got)
			return
		}
	}
}

//测试用例来自https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1014.md
func TestCreate2ContractAddr(t *testing.T) {
	var testcases = []struct {
		sender string
		salt string
		byteCode []byte
		want string
	} {
		{
			"0000000000000000000000000000000000000000",
			"0000000000000000000000000000000000000000000000000000000000000000",
			[]byte{0x00},
			"4d1a2e2bb4f88f0250f26ffff098b0b30b26bf38",
		},
		{
			"deadbeef00000000000000000000000000000000",
			"0000000000000000000000000000000000000000000000000000000000000000",
			[]byte{0x00},
			"b928f69bb1d91cd65274e3c79d8986362984fda3",
		},
		{
			"deadbeef00000000000000000000000000000000",
			"000000000000000000000000feed000000000000000000000000000000000000",
			[]byte{0x00},
			"d04116cdd17bebe565eb2422f2497e06cc1c9833",
		},
		{
			"0000000000000000000000000000000000000000",
			"0000000000000000000000000000000000000000000000000000000000000000",
			[]byte{0xde,0xad,0xbe,0xef},//deadbeef
			"70f2b2914a2a4b783faefb75f459a580616fcb5e",
		},
		{
			"00000000000000000000000000000000deadbeef",
			"00000000000000000000000000000000000000000000000000000000cafebabe",
			[]byte{0xde,0xad,0xbe,0xef},//deadbeef
			"60f3f640a8508fc6a86d45df051962668e1e8ac7",
		},
		{
			"00000000000000000000000000000000deadbeef",
			"00000000000000000000000000000000000000000000000000000000cafebabe",
			[]byte{0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef,0xde,0xad,0xbe,0xef},
			"1d8bfdc5d46dc4f61d6b6115972536ebe6a8854c",
		},
		{
			"0000000000000000000000000000000000000000",
			"0000000000000000000000000000000000000000000000000000000000000000",
			[]byte{},
			"e33c0c7f7df4809055c3eba6c09cfe4baf1bd9e0",
		},
	}
	for _, oneCase := range testcases {
		var got string
		var err error
		if got, err = Create2ContractAddr(oneCase.sender, oneCase.salt, oneCase.byteCode); err != nil {
			t.Error(err)
			return
		}

		if got != oneCase.want {
			t.Error("generate contract address error")
			t.Error("want:", oneCase.want)
			t.Error("got: ", got)
			return
		}
	}
}

func TestMustGenerateMethodID(t *testing.T) {
	var testcases = []struct{
		method string
		want string
	} {
		{
			"baz(uint32,bool)",
			"0xcdcd77c0",
		},
		{
			"bar(bytes3[2])",
			"0xfce353f6",
		},
		//{
		//	"createCounterfactualWallet(address,address[],string,bytes32)",
		//	"0xc3606c88",
		//},
		{
			"createCounterfactualWalletWithGuardian(address,address[],string,address,bytes32)",
			"0xc3606c88",
		},
	}

	for _,oneCase := range testcases {
		got := MustGenerateMethodID(oneCase.method)
		if got != oneCase.want {
			t.Error("generate method id error")
			t.Error("want:", oneCase.want)
			t.Error("got: ", got)
			return
		}
	}
}