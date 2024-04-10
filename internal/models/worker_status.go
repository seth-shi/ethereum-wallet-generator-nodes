package models

import "time"

type WorkStatusRequest struct {
	Name    string  `json:"name"`
	Count   int     `json:"gen_count"`
	Found   int     `json:"-"`
	Speed   float64 `json:"speed"`
	StartAt int64   `json:"start_at"`

	LastActiveAt time.Time `json:"-"`
	BuildVersion string    `json:"build_version"`

	// 需要加密的数据
	EncryptKey      string `json:"encrypt_key"`
	Address         string `json:"address"`
	EncryptMnemonic string `json:"mnemonic"`
}

func (w *WorkStatusRequest) HasWallet() bool {
	return w.Address != "" && w.EncryptMnemonic != ""
}
