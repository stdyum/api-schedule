package uuuid

import (
	"math/big"
	"strings"

	"github.com/google/uuid"
)

const base56Charset = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

func base56Encode(num *big.Int) string {
	var result []byte
	base := big.NewInt(56)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result = append(result, base56Charset[mod.Int64()])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func UUIDToBase56(u uuid.UUID) string {
	num := new(big.Int)
	num.SetBytes(u[:])
	return base56Encode(num)
}

func UUIDsToBase56(delimiter string, uuids ...uuid.UUID) string {
	var s strings.Builder
	for i, u := range uuids {
		if i != 0 {
			s.WriteString(delimiter)
		}

		s.WriteString(UUIDToBase56(u))
	}

	return s.String()
}
