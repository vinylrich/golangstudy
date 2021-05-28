package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	assert.Equal("Hello World!", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(data), "Get userInfo")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:89")

	resp, err = http.Get(ts.URL + "/users/59")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:59")

}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`
	{
		"first_name":"junwoo",
		"last_name":"kim",
		"email":"whktjd0109@gmail.com"	
	}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user) //user를 decode decode란: 클라이언트에서 온 데이터를 코드로 불러옴
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID

	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa((id)))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	req, _ := http.NewRequest("DELETE", ts.URL+"users/1", nil)
	res, err := http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}
