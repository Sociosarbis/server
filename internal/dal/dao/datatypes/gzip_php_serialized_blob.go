package datatypes

import (
	"bytes"
	"compress/gzip"
	"database/sql/driver"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/elliotchance/phpserialize"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GzipPhpSerializedBlob []interface{}

func decompress(data []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		defer gr.Close()
		return nil, err
	}
	defer gr.Close()
	bytes, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (b *GzipPhpSerializedBlob) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal MEDIUMBLOB value:", value))
	}
	bytes, err := decompress(data)
	if err != nil {
		return err
	}
	result, err := phpserialize.UnmarshalIndexedArray(bytes)
	*b = result
	return err
}

func (b GzipPhpSerializedBlob) Value() (driver.Value, error) {
	return nil, errors.New("write to db is not supported now")
}

func (b GzipPhpSerializedBlob) GormDataType() string {
	return "GzipPhpSerializedBlob"
}

func (b GzipPhpSerializedBlob) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "MEDIUMBLOB"
}
