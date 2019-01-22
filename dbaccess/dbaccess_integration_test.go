package dbaccess_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/twistingmercury/go-tsqlsh/dbaccess"
)

func Test_Integration_Connect_Success(t *testing.T) {
	usr := "sa"
	pwd := "6yhnMJU&"

	result, err := dbaccess.Connect(usr, pwd, nil, nil, nil, nil)

	assert.NoError(t, err)

	assert.NotNil(t, result)

	result.Close()
}
