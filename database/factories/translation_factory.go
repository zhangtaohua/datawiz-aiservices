package factories

import "datawiz-aiservices/app/models/translation"

func MakeTranslations() []translation.Translation {

	var objs []translation.Translation

	translationModel := translation.Translation{
		TranslationId:  "c_test",
		Language:       "zh-CN",
		TranslatedText: "测试",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_test",
		Language:       "zh-TW",
		TranslatedText: "測試",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_test",
		Language:       "en",
		TranslatedText: "test",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_name",
		Language:       "zh-CN",
		TranslatedText: "名称",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_name",
		Language:       "zh-TW",
		TranslatedText: "名稱",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_name",
		Language:       "en",
		TranslatedText: "name",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_language",
		Language:       "zh-CN",
		TranslatedText: "语言",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_language",
		Language:       "zh-TW",
		TranslatedText: "語言",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "c_language",
		Language:       "en",
		TranslatedText: "language",
	}
	objs = append(objs, translationModel)

	return objs
}
