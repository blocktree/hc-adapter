/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package hypercash

import (
	"testing"
)

func TestWalletManager_GetOmniBalance(t *testing.T) {
	propertyID := uint64(2)
	address := "Tsa6c6ivPG3fjRnJwKjGPUyMauPYDvK6Ckk"
	balance, err := tw.GetOmniBalance(propertyID, address)
	if err != nil {
		t.Errorf("GetOmniBalance failed unexpected error: %v\n", err)
		return
	}
	t.Logf("balance: %v\n", balance)
}

func TestWalletManager_IsHaveOmniAssets(t *testing.T) {
	address := "HsC2XDBCvxu4Z8nmWSVtLoZrCGiznRK2mpr"
	//address := "mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn"
	bool := tw.IsHaveOmniAssets(address)
	t.Logf("IsHaveOmniAssets: %v\n", bool)
}

func TestWalletManager_GetOmniTransaction(t *testing.T) {
	txid := "0244db052f01a788e9f582b5c1ae8ecc1f0dbdbaba6f3e03d21d92f3db243553"
	transaction, err := tw.GetOmniTransaction(txid)
	if err != nil {
		t.Errorf("GetOmniBalance failed unexpected error: %v\n", err)
		return
	}
	t.Logf("transaction: %+v", transaction)
}

func TestWalletManager_GetOmniInfo(t *testing.T) {
	result, err := tw.GetOmniInfo()
	if err != nil {
		t.Errorf("TestWalletManager_GetOmniInfo failed unexpected error: %v\n", err)
		return
	}
	t.Logf("OmniInfo: %+v", result)
}

func TestWalletManager_GetOmniProperty(t *testing.T) {
	propertyID := uint64(2)
	result, err := tw.GetOmniProperty(propertyID)
	if err != nil {
		t.Errorf("GetOmniProperty failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetOmniProperty: %+v", result)
}

func TestWalletManager_GetOmniBestBlockHash(t *testing.T) {
	blockhash, err := tw.GetOmniBestBlockHash()
	if err != nil {
		t.Errorf("GetOmniBestBlockHash failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockhash: %+v", blockhash)
}

func TestWalletManager_GetOmniBlockHeight(t *testing.T) {
	blockheight, err := tw.GetOmniBlockHeight()
	if err != nil {
		t.Errorf("GetOmniBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockheight: %+v", blockheight)
}

func TestWalletManager_GetOmniBlockHash(t *testing.T) {
	blockheight, err := tw.GetOmniBlockHash(279992)
	if err != nil {
		t.Errorf("GetOmniBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockheight: %+v", blockheight)
}