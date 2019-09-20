package configs

import (
	"context"
	"github.com/google/wire"
	"reflect"

	"errors"

	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	DefStructTagName         = "config"
	defStructTagDefaultValue = "config_def"
	defTimeLayout            = "200601021504"
)

var (
	//WireSet wire set
	WireSet = wire.NewSet(
		ProvideDecodeOption,
		ProvideStore,
	)
	mutex    = &sync.Mutex{}
	instance Store
)

func init() {
	viper.AutomaticEnv()
}

//StoreImpl config store
type StoreImpl struct {
	decodeOption viper.DecoderConfigOption
}

//Store config store interface
type Store interface {
	// GetConfig get config struct from environment variables
	GetConfig(ctx context.Context, val interface{}) error
}

//ProvideDecodeOption  decode option provider
func ProvideDecodeOption(ctx context.Context) viper.DecoderConfigOption {
	return func(option *mapstructure.DecoderConfig) {
		option.TagName = DefStructTagName
		option.WeaklyTypedInput = true
		option.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			// adding string to time decoder for which isn't enabled by default in ciper
			// ref: https://github.com/spf13/viper/blob/master/viper.go:838
			// YYYYmmDDHHMM
			mapstructure.StringToTimeHookFunc(defTimeLayout),
		)
	}
}

//ProvideStore config store provider
func ProvideStore(ctx context.Context, option viper.DecoderConfigOption) Store {
	mutex.Lock()
	defer mutex.Unlock()

	if instance != nil {
		return instance
	}

	instance = &StoreImpl{
		decodeOption: option,
	}

	return instance
}

//GetConfig get config from env, default val
func (c *StoreImpl) GetConfig(ctx context.Context, val interface{}) error {
	if err := c.loadTagConfig(val); err != nil {
		return err
	}

	return viper.Unmarshal(val, c.decodeOption)
}

//setDefault sets the default value for this key.
func (c *StoreImpl) setDefault(key string, val interface{}) error {
	viper.SetDefault(key, val)
	return viper.BindEnv(key)
}

//loadTagConfig load custom tag config
func (c *StoreImpl) loadTagConfig(val interface{}) error {

	dataValues := reflect.TypeOf(val)
	if dataValues.Kind() != reflect.Ptr {
		return errors.New("only accept struct pointer")
	}
	dataFields := dataValues.Elem()

	for i := 0; i < dataFields.NumField(); i++ {
		field := dataFields.Field(i)

		configKeyTag := field.Tag.Get(DefStructTagName)
		defaultValTag := field.Tag.Get(defStructTagDefaultValue)
		if err := c.setDefault(configKeyTag, defaultValTag); err != nil {
			return err
		}
	}
	return nil
}
