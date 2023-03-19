package auth_test

import (
	"fmt"
	"testing"

	fake "github.com/brianvoe/gofakeit/v5"
	authservice "github.com/icaroribeiro/go-code-challenge-template-2/internal/application/service/auth"
	domainentity "github.com/icaroribeiro/go-code-challenge-template-2/internal/core/domain/entity"
	authdatastoremockrepository "github.com/icaroribeiro/go-code-challenge-template-2/internal/core/ports/infrastructure/datastore/mockrepository/auth"
	logindatastoremockrepository "github.com/icaroribeiro/go-code-challenge-template-2/internal/core/ports/infrastructure/datastore/mockrepository/login"
	userdatastoremockrepository "github.com/icaroribeiro/go-code-challenge-template-2/internal/core/ports/infrastructure/datastore/mockrepository/user"
	"github.com/icaroribeiro/go-code-challenge-template-2/pkg/customerror"
	"github.com/icaroribeiro/go-code-challenge-template-2/pkg/security"
	securitypkg "github.com/icaroribeiro/go-code-challenge-template-2/pkg/security"
	mockauth "github.com/icaroribeiro/go-code-challenge-template-2/tests/mocks/pkg/mockauth"
	mocksecurity "github.com/icaroribeiro/go-code-challenge-template-2/tests/mocks/pkg/mocksecurity"
	mockvalidator "github.com/icaroribeiro/go-code-challenge-template-2/tests/mocks/pkg/mockvalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func (ts *TestSuite) TestLogIn() {
	credentials := security.Credentials{}

	login := domainentity.Login{}

	auth := domainentity.Auth{}

	AuthFactory := domainentity.Auth{}

	tokenExpTimeInSec := fake.Number(2, 10)

	token := ""

	errorType := customerror.NoType

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInLoggingIn",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				auth = domainentity.Auth{
					UserID: login.UserID,
				}

				id = uuid.NewV4()

				AuthFactory = domainentity.Auth{
					ID:     id,
					UserID: userID,
				}

				token = fake.Word()

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{nil},
					{domainentity.Auth{}, nil},
					{AuthFactory, nil},
					{token, nil},
				}
			},
			WantError: false,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfTheCredentialsAreNotValid",
			SetUp: func(t *testing.T) {
				credentials = security.Credentials{}

				returnArgs = ReturnArgs{
					{customerror.New("failed")},
					{domainentity.Login{}, nil},
					{nil},
					{domainentity.Auth{}, nil},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.BadRequest
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenGettingALoginByUsername",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				returnArgs = ReturnArgs{
					{nil},
					{domainentity.Login{}, customerror.New("failed")},
					{nil},
					{domainentity.Auth{}, nil},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfTheUsernameIsNotRegistered",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				returnArgs = ReturnArgs{
					{nil},
					{domainentity.Login{}, nil},
					{nil},
					{domainentity.Auth{}, nil},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.NotFound
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenVerifyingThePasswords",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{customerror.New("failed")},
					{domainentity.Auth{}, nil},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenCreatingAnAuth",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				auth = domainentity.Auth{
					UserID: login.UserID,
				}

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{nil},
					{domainentity.Auth{}, customerror.New("failed")},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfTheUserIDIsAlreadyRegistered",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				auth = domainentity.Auth{
					UserID: login.UserID,
				}

				id = uuid.NewV4()

				AuthFactory = domainentity.Auth{
					ID:     id,
					UserID: login.UserID,
				}

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{nil},
					{auth, nil},
					{domainentity.Auth{}, nil},
					{"", nil},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenCreatingAAuthFactory",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				auth = domainentity.Auth{
					UserID: login.UserID,
				}

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{nil},
					{domainentity.Auth{}, nil},
					{domainentity.Auth{}, customerror.New("failed")},
					{"", nil},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenCreatingAToken",
			SetUp: func(t *testing.T) {
				credentials = securitypkg.CredentialsFactory(nil)

				id := uuid.NewV4()
				userID := uuid.NewV4()

				login = domainentity.Login{
					ID:       id,
					UserID:   userID,
					Username: credentials.Username,
					Password: credentials.Password,
				}

				auth = domainentity.Auth{
					UserID: login.UserID,
				}

				id = uuid.NewV4()

				AuthFactory = domainentity.Auth{
					ID:     id,
					UserID: login.UserID,
				}

				returnArgs = ReturnArgs{
					{nil},
					{login, nil},
					{nil},
					{domainentity.Auth{}, nil},
					{AuthFactory, nil},
					{"", customerror.New("failed")},
				}

				errorType = customerror.NoType
			},
			WantError: true,
			TearDown:  func(t *testing.T) {},
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			validator := new(mockvalidator.Validator)
			validator.On("Validate", credentials).Return(returnArgs[0]...)

			loginDatastoreRepository := new(logindatastoremockrepository.Repository)
			loginDatastoreRepository.On("GetByUsername", credentials.Username).Return(returnArgs[1]...)

			security := new(mocksecurity.Security)
			security.On("VerifyPasswords", login.Password, credentials.Password).Return(returnArgs[2]...)

			authDatastoreRepository := new(authdatastoremockrepository.Repository)
			authDatastoreRepository.On("GetByUserID", login.UserID.String()).Return(returnArgs[3]...)
			authDatastoreRepository.On("Create", auth).Return(returnArgs[4]...)

			authN := new(mockauth.Auth)
			authN.On("CreateToken", AuthFactory, tokenExpTimeInSec).Return(returnArgs[5]...)

			userDatastoreRepository := new(userdatastoremockrepository.Repository)

			authService := authservice.New(authDatastoreRepository, loginDatastoreRepository, userDatastoreRepository,
				authN, security, validator, tokenExpTimeInSec)

			returnedToken, err := authService.LogIn(credentials)

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.Equal(t, token, returnedToken)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
				assert.Equal(t, errorType, customerror.GetType(err))
				assert.Empty(t, returnedToken)
			}

			tc.TearDown(t)
		})
	}
}
