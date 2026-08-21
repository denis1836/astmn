package validator

import (
	"fmt"
	"reflect"
	"strings"

	"astmn/internal/manifest"
)

func Validate(m *manifest.Manifest) error {
	if m == nil {
		return fmt.Errorf("manifest structure is nil")
	}

	val := reflect.ValueOf(*m)
	typ := val.Type()

	var missing []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		if fieldValue.Kind() == reflect.String && strings.TrimSpace(fieldValue.String()) == "" {
			yamlTag := field.Tag.Get("yaml")
			if yamlTag != "" {
				yamlTag = strings.Split(yamlTag, ",")[0]
			} else {
				yamlTag = field.Name
			}
			missing = append(missing, yamlTag)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("manifest validation failed, missing required fields: %s", strings.Join(missing, ", "))
	}

	return nil
}
