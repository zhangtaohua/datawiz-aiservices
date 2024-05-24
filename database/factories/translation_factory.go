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

	// --------------------------------------
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

	// --------------------------------------
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

	// --------------------------------------
	translationModel = translation.Translation{
		TranslationId:  "ai_classification",
		Language:       "zh-CN",
		TranslatedText: "分类",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_classification",
		Language:       "zh-TW",
		TranslatedText: "分類",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_classification",
		Language:       "en",
		TranslatedText: "Classification",
	}
	objs = append(objs, translationModel)

	// --------------------------------------
	translationModel = translation.Translation{
		TranslationId:  "ai_regression",
		Language:       "zh-CN",
		TranslatedText: "回归",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_regression",
		Language:       "zh-TW",
		TranslatedText: "迴歸",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_regression",
		Language:       "en",
		TranslatedText: "Regression",
	}
	objs = append(objs, translationModel)

	// --------------------------------------
	translationModel = translation.Translation{
		TranslationId:  "ai_clustering",
		Language:       "zh-CN",
		TranslatedText: "聚类",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_clustering",
		Language:       "zh-TW",
		TranslatedText: "聚類",
	}
	objs = append(objs, translationModel)

	translationModel = translation.Translation{
		TranslationId:  "ai_clustering",
		Language:       "en",
		TranslatedText: "Clustering",
	}
	objs = append(objs, translationModel)

	return objs
}
