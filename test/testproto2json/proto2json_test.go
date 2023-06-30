package testproto2json

import (
	"encoding/json"
	"testing"
	"unicode"

	datapb "github.com/ksusonic/gophkeeper/proto/data"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestConvertToJson(t *testing.T) {
	meta, _ := structpb.NewStruct(map[string]interface{}{
		"version": 1,
		"owner":   "dandex",
	})

	secret := &datapb.Secret{
		Name:       "test",
		Meta:       meta,
		SecretData: &datapb.Secret_Data{Variant: &datapb.Secret_Data_Any{}},
	}

	marshalled, err := protojson.Marshal(secret)
	assert.NoError(t, err)
	t.Logf("marshalled: %s", marshalled)

	jsonMap := map[string]interface{}{}
	err = json.Unmarshal(marshalled, &jsonMap)
	assert.NoError(t, err)
	for k := range jsonMap {
		assert.True(t, len(k) > 0)
		assert.True(t, unicode.IsLower(rune(k[0])))
	}
}
