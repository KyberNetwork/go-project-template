package convert

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	expBase = 10
)

// Exp10 ...
func Exp10(n int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(expBase), big.NewInt(n), nil)
}

func FloatExp10(n int64) *big.Float {
	return big.NewFloat(math.Pow10(int(n)))
}

// WeiToFloat ..
func WeiToFloat(amount *big.Int, decimals int64) float64 {
	amountFloat := big.NewFloat(0).SetInt(amount)
	amountFloat.Quo(amountFloat, big.NewFloat(0).SetInt(Exp10(decimals)))
	output, _ := amountFloat.Float64()

	return output
}

// FloatToWei ...
func FloatToWei(amount float64, decimals int64) *big.Int {
	weiFloat := big.NewFloat(amount)
	decimalsBigFloat := big.NewFloat(0).SetInt(Exp10(decimals))
	amountBig := new(big.Float).Mul(weiFloat, decimalsBigFloat)
	r, _ := amountBig.Int(nil)

	return r
}

// IntToWei ...
func IntToWei(amount int64, decimals int64) *big.Int {
	weiFloat := big.NewInt(amount)
	decimalsBig := Exp10(decimals)
	amountBig := new(big.Int).Mul(weiFloat, decimalsBig)
	return amountBig
}

func RoundUp(value float64, tickSize float64) float64 {
	decs := int32(math.Abs(math.Round(math.Log10(tickSize))))
	v := decimal.NewFromFloat(value)
	rec := v.Round(decs)
	if rec.LessThan(v) {
		rec = rec.Add(decimal.NewFromFloat(tickSize))
	}
	r, _ := rec.Float64()
	return r
}

func RoundDown(value float64, tickSize float64) float64 {
	decs := int32(math.Abs(math.Round(math.Log10(tickSize))))
	v := decimal.NewFromFloat(value)
	rec := v.Round(decs)
	if rec.GreaterThan(v) {
		rec = rec.Sub(decimal.NewFromFloat(tickSize))
		if rec.IsNegative() {
			rec = decimal.NewFromInt(0)
		}
	}
	r, _ := rec.Float64()
	return r
}
