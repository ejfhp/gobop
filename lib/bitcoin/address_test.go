package bitcoin_test

import (
	"fmt"
	"testing"

	"github.com/ejfhp/gobop/lib/bitcoin"
)

func TestAddress_WIF(t *testing.T) {
	for i := 0; i < 100; i++ {
		wif, err := bitcoin.WIF()
		if err != nil {
			t.Logf("cannot generate WIF: %v", err)
			t.Fail()
		}
		if len(wif) < 10 {
			t.Fatalf("WIF too short: %s", wif)
		}
		if wif[0] != 'K' && wif[0] != 'L' {
			t.Fatalf("WIF starts with a wrong char: %s", wif)

		}
		t.Logf("WIF: %s\n", wif)
	}
}

func TestAddress_DecodeWIF(t *testing.T) {
	wifs := []string{
		"L48cWSssxbFnRuuJCVes9NEYP1W987kfpSgWG2RKSaZtcs6iCHpT",
		"KzxKMJoJ13Ug2E8mBb9npbqavs9hbX3rZ3XPq3jBNUriQNk5rMUc",
		"Kyypokm7KphGVa6QdqpWM4bkdTocQmD6f2waMBREcVq9UJKHow3o",
		"KzLEePeTR2utHtBLfoPRjf7hJeDzBodnfApN8WFb4gaEkneCP7KP",
		"L5n7n4ntJD3YyqUtcaekyqHZiv5nB71yhZE5SRzwWwtQocqEgwiv",
		"L4JnikU8C8z8nJgipUEAbwQfqCRW19FhpXs8cWnw25mYjjVu32jC",
		"KwGkpXM5Chd7J2CDFKTgU9YQm1M4gkN2njic9MmsqQFhmCbYykXK",
		"L2nuX1XJbbvUUyAVbBzjZzLLiniw1sJf6jQsMgkoZm15ajG8NuLP",
		"KyU8g3jfuFiDAJmTqFEyYbfFUY58783K4qBbo7QBKNjpo5VGqG1V",
		"KwSLmi6PiyngRxrX5QVX6QL5L9eSSYM2Ehw5q59vKsbgA6qnQLgp",
		"KymHqwJasnhAcJgyw6UA9SWeCcZS1h2jg5qPc58BUnQbvUFXePyq",
		"KzdbgKhTGeFLNwj2c2g4cMLHHUyhtHLn4YVvqGF8YxDW7sa9cWQ5",
		"L18KvKE3q31ohiFjQiDQiVM1ocE3e7BdHS3Fz8KRAx44xmm1ovaD",
	}
	var err error
	for _, wif := range wifs {
		_, err = bitcoin.DecodeWIF(wif)
		if err != nil {
			t.Fatalf("cannot decode WIF %s: %v", wif, err)
		}
		add, err := bitcoin.AddressOf(wif)
		if err != nil {
			t.Fatalf("cannot get address of WIF %s: %v", wif, err)
		}
		if len(add) < 10 {
			t.Fatalf("address too short for WIF %s: '%s'", wif, add)
		}
		if add[0] != '1' {
			t.Fatalf("address starts with the wrong char: '%s'", add)
		}
		fmt.Printf("Key: %s  Address: %s\n", wif, add)
	}
}
