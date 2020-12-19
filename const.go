package decimal

import "math/big"

// Exported const for some common number
var (
	Zero = BigDecimal{
		new(big.Int).SetInt64(0),
		0, 0, 0, "",
	}
	One = BigDecimal{
		big.NewInt(1),
		0, 0, 0, "",
	}
	Two = BigDecimal{
		big.NewInt(2),
		0, 0, 0, "",
	}
	Ten = BigDecimal{
		big.NewInt(10),
		0, 0, 0, "",
	}
)
