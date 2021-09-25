package main

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"time"
)

// PriceCalculator estimates the VAT-inclusive retail prices of items.
type PriceCalculator struct {
	priceSrv Caller
	vatSrv   Caller
}

// NewPriceCalculator creates a PriceCalculator instance that queries the provided endpoints for item price and VAT information.
func NewPriceCalculator(priceEndpoint, vatEndpoint string) *PriceCalculator {
	return &PriceCalculator{
		priceSrv: restCaller(priceEndpoint),
		vatSrv:   restCaller(vatEndpoint),
	}
}

// PriceForItem calculates the VAT-inclusive retail price of itemUUID with the currently applicable VAT rates.
func (pc *PriceCalculator) PriceForItem(itemUUID string) (float64, error) {
	return pc.PriceForItemAtDate(itemUUID, time.Now())
}

// PriceForItemAtDate calculates the VAT-inclusive retail price of itemUUID with the VAT rates that applied at a particular date.
func (pc *PriceCalculator) PriceForItemAtDate(itemUUID string, date time.Time) (float64, error) {
	priceRes := struct {
		Price float64 `json:"price"`
	}{}
	// performs an RPC and decodes the response into priceRes.
	if err := pc.callService(
		pc.priceSrv,
		map[string]interface{}{
			"item":   itemUUID,
			"period": date,
		},
		&priceRes,
	); err != nil {
		return 0, xerrors.Errorf("unable to retrieve item price: %w", err)
	}

	vatRes := struct {
		Rate float64 `json:"vat_rate"`
	}{}
	// performs an RPC and decodes the response into vatRes.
	if err := pc.callService(
		pc.vatSrv,
		map[string]interface{}{"period": date},
		&vatRes,
	); err != nil {
		return 0, xerrors.Errorf("unable to retrieve vat percent: %w", err)
	}

	// applies a vat rate to a price and returns the result.
	return priceRes.Price * (1.0 + vatRes.Rate), nil
}

// callService performs an RPC and decodes the response into res.
func (pc *PriceCalculator) callService(svc Caller, req map[string]interface{}, res interface{}) error {
	svcRes, err := svc.Call(req)
	if err != nil {
		return xerrors.Errorf("call to remote service failed: %w", err)
	}
	defer drainAndClose(svcRes)

	if err = json.NewDecoder(svcRes).Decode(res); err != nil {
		return xerrors.Errorf("unable to decode remote service response: %w", err)
	}

	return nil
}
