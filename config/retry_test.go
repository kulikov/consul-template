package config

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRetryConfig_Copy(t *testing.T) {
	cases := []struct {
		name string
		a    *RetryConfig
	}{
		{
			"nil",
			nil,
		},
		{
			"empty",
			&RetryConfig{},
		},
		{
			"same_enabled",
			&RetryConfig{
				Attempts: Int(25),
				Backoff:  TimeDuration(20 * time.Second),
				Enabled:  Bool(true),
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			r := tc.a.Copy()
			if !reflect.DeepEqual(tc.a, r) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.a, r)
			}
		})
	}
}

func TestRetryConfig_Merge(t *testing.T) {
	cases := []struct {
		name string
		a    *RetryConfig
		b    *RetryConfig
		r    *RetryConfig
	}{
		{
			"nil_a",
			nil,
			&RetryConfig{},
			&RetryConfig{},
		},
		{
			"nil_b",
			&RetryConfig{},
			nil,
			&RetryConfig{},
		},
		{
			"nil_both",
			nil,
			nil,
			nil,
		},
		{
			"empty",
			&RetryConfig{},
			&RetryConfig{},
			&RetryConfig{},
		},
		{
			"attempts_overrides",
			&RetryConfig{Attempts: Int(10)},
			&RetryConfig{Attempts: Int(20)},
			&RetryConfig{Attempts: Int(20)},
		},
		{
			"attempts_empty_one",
			&RetryConfig{Attempts: Int(10)},
			&RetryConfig{},
			&RetryConfig{Attempts: Int(10)},
		},
		{
			"attempts_empty_two",
			&RetryConfig{},
			&RetryConfig{Attempts: Int(10)},
			&RetryConfig{Attempts: Int(10)},
		},
		{
			"attempts_same",
			&RetryConfig{Attempts: Int(10)},
			&RetryConfig{Attempts: Int(10)},
			&RetryConfig{Attempts: Int(10)},
		},

		{
			"backoff_overrides",
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
			&RetryConfig{Backoff: TimeDuration(20 * time.Second)},
			&RetryConfig{Backoff: TimeDuration(20 * time.Second)},
		},
		{
			"backoff_empty_one",
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
			&RetryConfig{},
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
		},
		{
			"backoff_empty_two",
			&RetryConfig{},
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
		},
		{
			"backoff_same",
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
			&RetryConfig{Backoff: TimeDuration(10 * time.Second)},
		},
		{
			"enabled_overrides",
			&RetryConfig{Enabled: Bool(true)},
			&RetryConfig{Enabled: Bool(false)},
			&RetryConfig{Enabled: Bool(false)},
		},
		{
			"enabled_empty_one",
			&RetryConfig{Enabled: Bool(true)},
			&RetryConfig{},
			&RetryConfig{Enabled: Bool(true)},
		},
		{
			"enabled_empty_two",
			&RetryConfig{},
			&RetryConfig{Enabled: Bool(true)},
			&RetryConfig{Enabled: Bool(true)},
		},
		{
			"enabled_same",
			&RetryConfig{Enabled: Bool(true)},
			&RetryConfig{Enabled: Bool(true)},
			&RetryConfig{Enabled: Bool(true)},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			r := tc.a.Merge(tc.b)
			if !reflect.DeepEqual(tc.r, r) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.r, r)
			}
		})
	}
}

func TestRetryConfig_Finalize(t *testing.T) {
	cases := []struct {
		name string
		i    *RetryConfig
		r    *RetryConfig
	}{
		{
			"empty",
			&RetryConfig{},
			&RetryConfig{
				Attempts: Int(DefaultRetryAttempts),
				Backoff:  TimeDuration(DefaultRetryBackoff),
				Enabled:  Bool(true),
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			tc.i.Finalize()
			if !reflect.DeepEqual(tc.r, tc.i) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.r, tc.i)
			}
		})
	}
}
