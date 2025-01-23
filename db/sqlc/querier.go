// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	ChangePassword(ctx context.Context, arg ChangePasswordParams) (Users, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	CheckWalletExists(ctx context.Context, arg CheckWalletExistsParams) (bool, error)
	CreateCampaignType(ctx context.Context, arg CreateCampaignTypeParams) (Campaigns, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (UserSession, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	CreateUserWallet(ctx context.Context, arg CreateUserWalletParams) (UserWalletAddresses, error)
	DeleteSession(ctx context.Context, id uuid.UUID) (UserSession, error)
	DeleteUser(ctx context.Context, username string) (Users, error)
	GetAllActiveDonations(ctx context.Context) ([]Donations, error)
	GetAllCampaignType(ctx context.Context) ([]Campaigns, error)
	GetSession(ctx context.Context, id uuid.UUID) (UserSession, error)
	GetUser(ctx context.Context, username string) (Users, error)
	GetUserByAddress(ctx context.Context, address string) (Users, error)
	GetUserWallets(ctx context.Context, arg GetUserWalletsParams) ([]UserWalletAddresses, error)
	GetWalletByAddress(ctx context.Context, arg GetWalletByAddressParams) (UserWalletAddresses, error)
	GetWalletById(ctx context.Context, arg GetWalletByIdParams) (UserWalletAddresses, error)
	HardDeleteUserWallet(ctx context.Context, arg HardDeleteUserWalletParams) (UserWalletAddresses, error)
	SoftDeleteUserWallet(ctx context.Context, arg SoftDeleteUserWalletParams) (UserWalletAddresses, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
	UpdateUserWalletStatus(ctx context.Context, arg UpdateUserWalletStatusParams) (UserWalletAddresses, error)
}

var _ Querier = (*Queries)(nil)
