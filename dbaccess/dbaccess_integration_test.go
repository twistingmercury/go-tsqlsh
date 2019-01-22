package dbaccess_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/twistingmercury/go-tsqlsh/dbaccess"
)

func Test_Integration_Connect_Success(t *testing.T) {
	usr := "sa"
	pwd := "6yhnMJU&"

	result, err := dbaccess.Connect(usr, pwd, "localhost", "master", 1433, true)

	assert.NoError(t, err)

	assert.NotNil(t, result)

	result.Close()
}

func Test_New_all_values_provided(t *testing.T) {
	obj, err := dbaccess.New("sa", "6yhnMJU&", &host, &db, &port, &mutable)

	assert.NoError(t, err)
	assert.NotNil(t, obj)
	assert.Equal(t, "sa", obj.User())
	assert.Equal(t, host, obj.Host())
	assert.Equal(t, db, obj.Database())
	assert.Equal(t, port, obj.Port())
	assert.Equal(t, mutable, obj.Mutable())
	assert.Equal(t, "open", obj.State())

	err = obj.Close()
	assert.NoError(t, err)
	assert.Equal(t, "closed", obj.State())
}

func Test_New_All_Defaults(t *testing.T) {
	obj, err := dbaccess.New("sa", "6yhnMJU&", nil, nil, nil, nil)

	assert.NoError(t, err)
	assert.NotNil(t, obj)
	assert.Equal(t, "sa", obj.User())
	assert.Equal(t, "localhost", obj.Host())
	assert.Equal(t, "master", obj.Database())
	assert.Equal(t, 1433, obj.Port())
	assert.Equal(t, false, obj.Mutable())
	assert.Equal(t, "open", obj.State())

	err = obj.Close()
	assert.NoError(t, err)
	assert.Equal(t, "closed", obj.State())
}
