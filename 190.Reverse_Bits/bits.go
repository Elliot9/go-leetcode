package main

func reverseBits(num uint32) uint32 {
	var reversed uint32
	reversed = 0

	for i := 0; i < 32; i++ {
		reversed = reversed<<1 + num&1
		num = num >> 1
	}

	return reversed
}
