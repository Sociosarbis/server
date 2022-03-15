package datatypes

import (
	"bytes"
	"compress/flate"
	"database/sql/driver"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/elliotchance/phpserialize"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GzipPhpSerializedBlob map[string]interface{}

func decompress(data []byte) ([]byte, error) {
	gr := flate.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	bytes, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func toValidJson(data interface{}) interface{} {
	t := reflect.TypeOf(data).Kind()
	if t == reflect.Array || t == reflect.Slice {
		arr := data.([]interface{})
		for i, val := range arr {
			arr[i] = toValidJson(val)
		}
		return arr
	} else if t == reflect.Map {
		m := data.(map[interface{}]interface{})
		ret := map[string]interface{}{}
		for k, v := range m {
			ret[fmt.Sprint(k)] = toValidJson(v)
		}
		return ret
	} else {
		return data
	}
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
	result, err := phpserialize.UnmarshalAssociativeArray(bytes)
	*b = toValidJson(result).(map[string]interface{})
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
