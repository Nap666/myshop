// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalCartInfoDao is internal type for wrapping internal DAO implements.
type internalCartInfoDao = *internal.CartInfoDao

// cartInfoDao is the data access object for table cart_info.
// You can define custom methods on it to extend its functionality as you wish.
type cartInfoDao struct {
	internalCartInfoDao
}

var (
	// CartInfo is globally public accessible object for table cart_info operations.
	CartInfo = cartInfoDao{
		internal.NewCartInfoDao(),
	}
)

// Fill with you ideas below.
