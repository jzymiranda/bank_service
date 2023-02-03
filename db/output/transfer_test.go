package db

import (
	"context"
	"testing"
	"time"

	"github.com/jzymiranda/bank_service/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountF Account, accountT Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: accountF.ID,
		ToAccountID:   accountT.ID,
		Amount:        util.RandomMoney(),
	}

	Transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Transfer)

	require.Equal(t, arg.Amount, Transfer.Amount)
	require.Equal(t, arg.FromAccountID, Transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, Transfer.ToAccountID)

	require.NotZero(t, Transfer.ID)
	require.NotZero(t, Transfer.CreatedAt)

	return Transfer

}

func TestCreateTransfer(t *testing.T) {
	accountF := createRandomAccount(t)
	accountT := createRandomAccount(t)
	createRandomTransfer(t, accountF, accountT)
}

func TestGetTransfer(t *testing.T) {
	accountF := createRandomAccount(t)
	accountT := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, accountF, accountT)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
	}

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account2.ID && transfer.ToAccountID == account1.ID)
	}
}
