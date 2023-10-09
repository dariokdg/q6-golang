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
	tFirstPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pSeven)
	tSecondPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pOne)
	tThirdPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.03))

	sFirstPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pSeven)
	sSecondPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(pOne)
	sThirdPrize := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.03))

	rPrize := totalRevanchaSales.Mul(pNine).Mul(pSix).Mul(pEight)

	ssPrize := totalSiempreSaleSales.Mul(pNine).Mul(pSix).Mul(pSix)

	tempOne := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.163))
	tempTwo := totalTradicionalSales.Mul(pFive).Mul(pFour).Mul(decimal.NewFromFloat(0.163))
	tempThree := totalRevanchaSales.Mul(pNine).Mul(pSix).Mul(decimal.NewFromFloat(0.192))
	tempFour := totalSiempreSaleSales.Mul(pNine).Mul(pSix).Mul(pFour)
	pePrize := decimal.Sum(tempOne, tempTwo, tempThree, tempFour)
	return Prizes{TotalTradicionalSales: totalTradicionalSales, TotalRevanchaSales: totalRevanchaSales, TotalSiempreSaleSales: totalSiempreSaleSales, TradicionalFirstPrize: tFirstPrize, TradicionalSecondPrize: tSecondPrize, TradicionalThirdPrize: tThirdPrize, SegundaFirstPrize: sFirstPrize, SegundaSecondPrize: sSecondPrize, SegundaThirdPrize: sThirdPrize, RevanchaPrize: rPrize, SiempreSalePrize: ssPrize, PozoExtraPrize: pePrize}
}
