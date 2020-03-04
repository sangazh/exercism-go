package diffiehellman

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	if p.Int64() < 0 {
		p = big.NewInt(math.MaxInt64)
	}

	rand.Seed(time.Now().UnixNano())
	min := int64(2)      // >=1
	max := p.Int64() - 1 // <=p
	r := rand.Int63n(max - min + 1)

	return big.NewInt(r + min)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	//A = g ** a(private) mod p
	return big.NewInt(1).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	a := PrivateKey(p)
	return a, PublicKey(a, p, g)
}

func SecretKey(a, B, p *big.Int) *big.Int {
	//s = B**a mod p
	return big.NewInt(1).Exp(B, a, p)
}
