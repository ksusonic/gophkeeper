package crypta

import (
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/models"
	datapb "github.com/ksusonic/gophkeeper/proto/data"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func DecryptSecret(service *Service, s *models.Secret) (*datapb.Secret, error) {
	decryptedData, err := service.Decrypt(s.Data)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt %s secret secretData: %v", s.Name, err)
	}
	secretData := &datapb.Secret_Data{}
	if err := proto.Unmarshal(decryptedData, secretData); err != nil {
		return nil, fmt.Errorf("could not unmarshall %s secretData: %v", s.Name, err)
	}
	meta, err := structpb.NewStruct(s.Meta)
	if err != nil {
		return nil, fmt.Errorf("could not create struct from map: %v", err)
	}
	return &datapb.Secret{
		Name:       s.Name,
		Meta:       meta,
		SecretData: secretData,
	}, nil
}
