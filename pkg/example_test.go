package pkg_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_shipping-dx-dojo/pkg"
	"github.com/mercadolibre/fury_shipping-dx-dojo/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func (u userServiceMock) GetUsers() []pkg.User {
	args := u.Called()
	return args.Get(0).([]pkg.User)
}

func ExamplePrintln() {
	pkg.Println("hola mundo")
	pkg.Println("chau mundo")
	// Output:
	// hola mundo
	// chau mundo
}

func BenchmarkGetUserInUseCase(b *testing.B) {
	u := new(userServiceMock)
	u.On("GetUsers").Return([]pkg.User{
		{
			ID: 200,
		},
		{
			ID: 1500,
		},
	})
	c := pkg.NewGetUserUseCase(u)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		user, err := c.GetUser(200)
		b.StopTimer()

		if user.ID != 200 || err != nil {
			b.Fail()
		}
	}
}

func TestGetUserInUseCaseReturns200(t *testing.T) {
	u := new(userServiceMock)
	u.On("GetUsers").Return([]pkg.User{
		{
			ID: 200,
		},
		{
			ID: 1500,
		},
	}).Once()
	c := pkg.NewGetUserUseCase(u)

	user, err := c.GetUser(200)

	assert.Equal(t, 200, user.ID)
	assert.Equal(t, pkg.User{ID: 200}, user)
	assert.Nil(t, err)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users/create", pkg.PostUser)
	return r
}

func TestPostUserReturns200BodyCorrectly(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/create", nil)
	router.ServeHTTP(w, req)

	actual := w.Body.String()
	body := test.Get(t, []byte(actual), "test.post_user.200.golden")
	assert.Equal(t, []byte(actual), body)
}
