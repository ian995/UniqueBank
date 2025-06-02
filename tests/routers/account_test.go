package routers_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/internal/routers"
	"github.com/ian995/UniqueBank/pkg/utils"
	"github.com/ian995/UniqueBank/tests/mock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()

	testcase := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mock_test.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.IDAccount,
			buildStubs: func(store *mock_test.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.IDAccount)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder, account)
			},
		},{
			name:      "NotFound",
			accountID: account.IDAccount,
			buildStubs: func(store *mock_test.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.IDAccount)).
					Times(1).
					Return(repo.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},{
			name:      "InternalError",
			accountID: account.IDAccount,
			buildStubs: func(store *mock_test.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.IDAccount)).
					Times(1).
					Return(repo.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testcase {
		tc := testcase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mock_test.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := routers.NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", account.IDAccount)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})

	}
}

func randomAccount() repo.Account {
	return repo.Account{
		IDAccount: utils.RandomInt(1, 1000),
		Owner:     utils.RandomOwner(),
		Balance:   utils.RandomMoney(),
		Currency:  utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *httptest.ResponseRecorder, account repo.Account) {
	data, err := ioutil.ReadAll(body.Body)
	require.NoError(t, err)

	var gotAccount repo.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account.IDAccount, gotAccount.IDAccount)
}
