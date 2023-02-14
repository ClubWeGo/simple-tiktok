package main

import (
	"context"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (resp *favorite.FavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteListMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	// TODO: Your code here...
	return
}
