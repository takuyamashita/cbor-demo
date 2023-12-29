package main

func main() {

	/*
		0	0x00
		1	0x01
		10	0x0a
		23	0x17
		24	0x1818
		25	0x1819
		100	0x1864
		1000	0x1903e8
		1000000	0x1a000f4240
		1000000000000	0x1b000000e8d4a51000

		-1	0x20
		-10	0x29
		-100	0x3863

		"hello"	0x4568656c6c6f

		"IETF"	0x6449455446
	*/

	/*
		intZero := []byte{0x00}
		int100 := []byte{0x18, 0x64}
		int1000 := []byte{0x19, 0x03, 0xe8}
		int1000000 := []byte{0x1a, 0x00, 0x0f, 0x42, 0x40}
		int1000000000000 := []byte{0x1b, 0x00, 0x00, 0x00, 0xe8, 0xd4, 0xa5, 0x10, 0x00}

		cbordemo.Decode(bytes.NewReader(intZero))
		cbordemo.Decode(bytes.NewReader(int100))
		cbordemo.Decode(bytes.NewReader(int1000))
		cbordemo.Decode(bytes.NewReader(int1000000))
		cbordemo.Decode(bytes.NewReader(int1000000000000))

		intZeroint100 := []byte{0x00, 0x18, 0x64}

		cbordemo.Decode(bytes.NewReader(intZeroint100))

		nInt1 := []byte{0x20}
		nInt10 := []byte{0x29}
		nInt100 := []byte{0x38, 0x63}

		cbordemo.Decode(bytes.NewReader(nInt1))
		cbordemo.Decode(bytes.NewReader(nInt10))
		cbordemo.Decode(bytes.NewReader(nInt100))

		strHello := []byte{0x45, 0x68, 0x65, 0x6c, 0x6c, 0x6f}
		cbordemo.Decode(bytes.NewReader(strHello))

		strIETF := []byte{0x64, 0x49, 0x45, 0x54, 0x46}
		cbordemo.Decode(bytes.NewReader(strIETF))
	*/

}
