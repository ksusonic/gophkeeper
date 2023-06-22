package proto2json

import (
	"encoding/json"
	"testing"
	"unicode"

	datapb "github.com/ksusonic/gophkeeper/proto/data"
	"github.com/stretchr/testify/assert"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestConvertToJson(t *testing.T) {
	secret := &datapb.Secret{
		Name:    "test",
		Version: 1,
		Data:    &datapb.SecretValue{},
	}

	marshalled, err := protojson.Marshal(secret)
	assert.NoError(t, err)
	t.Logf("marshalled: %s", marshalled)

	jsonMap := map[string]string{}
	err = json.Unmarshal(marshalled, &jsonMap)
	for k := range jsonMap {
		assert.True(t, len(k) > 0)
		assert.True(t, unicode.IsLower(rune(k[0])))
	}
}
