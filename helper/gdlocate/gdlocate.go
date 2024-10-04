package gdlocate

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"shorten-v3-service/helper/gdhelper"
)

const (
	defaultLang = "en"
)

type IConfig interface {
	GetLocalize() string
	GetMessageFormat() string
	GetMessagePath() string
	GetMessageTag() string
}

type ILocalizeService interface {
	GetDefaultLanguage() string
}

type localizeService struct {
	defaultLang string
	languages   map[string]*i18n.Localizer
}

func NewLocalizeService(configs ...IConfig) ILocalizeService {
	var (
		languages = make(map[string]*i18n.Localizer)
	)

	for _, e := range configs {
		bundle := i18n.NewBundle(language.Make(e.GetMessageTag()))
		bundle.RegisterUnmarshalFunc(e.GetMessageFormat(), gdhelper.IIF(e.GetMessageFormat() == "json", json.Unmarshal, toml.Unmarshal))
		bundle.MustLoadMessageFile(e.GetMessagePath())
		languages[e.GetLocalize()] = i18n.NewLocalizer(bundle, e.GetLocalize())
	}
	return &localizeService{languages: languages, defaultLang: defaultLang}
}

func (s *localizeService) GetDefaultLanguage() string {
	return s.defaultLang
}
