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
	txid := "300142e7f87bf669b09079cf25985f773ab4065498b15e51498533b949e054a0"
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