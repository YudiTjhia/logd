package svc

import (
	"errors"
	"fmt"
)

func StringPointerRequired(fieldName string, val *string) error {
	if val == nil {
		return errors.New(fieldName + "_is_required")
	}
	return nil
}

func StringRequired(fieldName string, val string) error {
	if val == "" {
		return errors.New(fieldName + "_is_required")
	}
	return nil
}

func StringPointerMax(fieldName string, val *string, max int) error {
	if val != nil {
		if len(*val) > max {
			return errors.New(fmt.Sprintf("%s_max_length_is_%d_characters", fieldName, max))
		}
	}
	return nil
}

func StringMax(fieldName string, val string, max int) error {
	if val == "" {
		if len(val) > max {
			return errors.New(fmt.Sprintf("%s_max_length_is_%d_characters", fieldName, max))
		}
	}
	return nil
}

func StringMin(fieldName string, val string, min int) error {
	if val == "" {
		if len(val) < min {
			return errors.New(fmt.Sprintf("%s_min_length_is_%d_characters", fieldName, min))
		}
	}
	return nil
}

func IntGtz(fieldName string, val int) error {
	if val <= 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_zero", fieldName))
	}
	return nil
}

func IntMax(fieldName string, val int, max int) error {
	if val > max {
		return errors.New(fmt.Sprintf("%s_maximum_value_is_%d", fieldName, max))
	}
	return nil
}

func IntMin(fieldName string, val int, min int) error {
	if val < min {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_%d", fieldName, min))
	}
	return nil
}

func IntGtez(fieldName string, val int) error {
	if val < 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_or_equals_to_zero", fieldName))
	}
	return nil
}

func Int64Gtz(fieldName string, val int64) error {
	if val <= 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_zero", fieldName))
	}
	return nil
}

func Int64Max(fieldName string, val int64, max int64) error {
	if val > max {
		return errors.New(fmt.Sprintf("%s_must_be_less_than_%d", fieldName, max))
	}
	return nil
}

func Int64Min(fieldName string, val int64, min int64) error {
	if val < min {
		return errors.New(fmt.Sprintf("%s_minimum_value_is_%d", fieldName, min))
	}
	return nil
}

func Int64Gtez(fieldName string, val int64) error {
	if val < 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_zero", fieldName))
	}
	return nil
}

func Float32Gtz(fieldName string, val float32) error {
	if val <= 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_or_equals_to_zero", fieldName))
	}
	return nil
}

func Float32Max(fieldName string, val float32, max float32) error {
	if val > max {
		return errors.New(fmt.Sprintf("%s_maximum_value_is_%f", fieldName, max))
	}
	return nil
}

func Float32Min(fieldName string, val float32, min float32) error {
	if val < min {
		return errors.New(fmt.Sprintf("%s_minimum_value_is_%f", fieldName, min))
	}
	return nil
}

func Float32Gtez(fieldName string, val float32) error {
	if val < 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_or_equals_to_zero", fieldName))
	}
	return nil
}

func Float64Gtz(fieldName string, val float64) error {
	if val <= 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_zero", fieldName))
	}
	return nil
}

func Float64Max(fieldName string, val float64, max float64) error {
	if val > max {
		return errors.New(fmt.Sprintf("%s_must_be_less_than_%f", fieldName, val))
	}
	return nil
}

func Float64Min(fieldName string, val float64, min float64, errMsg string) error {
	if val < min {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_%f", fieldName, val))
	}
	return nil
}

func Float64Gtez(fieldName string, val float64) error {
	if val < 0 {
		return errors.New(fmt.Sprintf("%s_must_be_greater_than_or_equals_to_zero", fieldName))
	}
	return nil
}
