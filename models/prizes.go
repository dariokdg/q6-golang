package models

import "github.com/shopspring/decimal"

type Prizes struct {
	TotalTradicionalSales  decimal.Decimal `json:"totalTradicionalSales"`
	TotalRevanchaSales     decimal.Decimal `json:"totalRevanchaSales"`
	TotalSiempreSaleSales  decimal.Decimal `json:"totalSiempreSaleSales"`
	TradicionalFirstPrize  decimal.Decimal `json:"tradicionalFirstPrize"`
	TradicionalSecondPrize decimal.Decimal `json:"tradicionalSecondPrize"`
	TradicionalThirdPrize  decimal.Decimal `json:"tradicionalThirdPrize"`
	SegundaFirstPrize      decimal.Decimal `json:"segundaFirstPrize"`
	SegundaSecondPrize     decimal.Decimal `json:"segundaSecondPrize"`
	SegundaThirdPrize      decimal.Decimal `json:"segundaThirdPrize"`
	RevanchaPrize          decimal.Decimal `json:"revanchaPrize"`
	SiempreSalePrize       decimal.Decimal `json:"siempreSalePrize"`
	PozoExtraPrize         decimal.Decimal `json:"pozoExtraPrize"`
}

func GetPrizes(totalTradicionalSales decimal.Decimal, totalRevanchaSales decimal.Decimal, totalSiempreSaleSales decimal.Decimal) Prizes {
	pOne := decimal.NewFromFloat(0.1)
	pFour := decimal.NewFromFloat(0.4)
	pFive := decimal.NewFromFloat(0.5)
	pSix := decimal.NewFromFloat(0.6)
	pSeven := decimal.NewFromFloat(0.7)
	pEight := decimal.NewFromFloat(0.8)
	pNine := decimal.NewFromFloat(0.9)
	tFirstPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pSeven).Round(2)
	tSecondPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pOne).Round(2)
	tThirdPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.03)).Round(2)

	sFirstPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pSeven).Round(2)
	sSecondPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pOne).Round(2)
	sThirdPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.03)).Round(2)

	rPrize := totalRevanchaSales.Mul(pNine).Mul(pSix).Mul(pEight).Round(2)

	ssPrize := totalSiempreSaleSales.Mul(pNine).Mul(pSix).Mul(pSix).Round(2)

	tempOne := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.163)).Round(2)
	tempTwo := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.163)).Round(2)
	tempThree := totalRevanchaSales.Mul(pNine).Mul(pSix).Mul(decimal.NewFromFloat(0.192)).Round(2)
	tempFour := totalSiempreSaleSales.Mul(pNine).Mul(pSix).Mul(pFour).Round(2)
	pePrize := decimal.Sum(tempOne, tempTwo, tempThree, tempFour).Round(2)
	return Prizes{TotalTradicionalSales: totalTradicionalSales, TotalRevanchaSales: totalRevanchaSales, TotalSiempreSaleSales: totalSiempreSaleSales, TradicionalFirstPrize: tFirstPrize, TradicionalSecondPrize: tSecondPrize, TradicionalThirdPrize: tThirdPrize, SegundaFirstPrize: sFirstPrize, SegundaSecondPrize: sSecondPrize, SegundaThirdPrize: sThirdPrize, RevanchaPrize: rPrize, SiempreSalePrize: ssPrize, PozoExtraPrize: pePrize}
}
