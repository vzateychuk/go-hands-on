package main

import "errors"

type BackPressure struct {
	limitCh chan struct{}
}

func NewBackPressure(limit int) *BackPressure {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &BackPressure{
		limitCh: ch,
	}
}

func (bp *BackPressure) Process(fun func()) error {
	select {
	case <-bp.limitCh:
		fun()
		bp.limitCh <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}
