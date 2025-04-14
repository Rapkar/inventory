package Boot

func Messages(value string) string {
	var messages = map[string]string{
		"login success":            "با موفقیت وارد شدید",
		"login faild":              "اطلاعات ورود اشتباه است",
		"user made success":        "کاربر با موفقیت افزوده شد",
		"user made faild":          "افزودن کاربر با مشکل مواجه شد",
		"user remove success":      "کاربر با موفقیت حذف شد",
		"user remove faild":        "حذف کاربر با مشکل مواجه شد",
		"product remove success":   "محصول با موفقیت حذف شد",
		"product remove faild":     "حذف محصول با مشکل مواجه شد",
		"product made success":     "محصول با موفقیت افزوده شد",
		"product made faild":       "افزودن محصول با مشکل مواجه شد",
		"Export removed success":   "فاکتور با موفقیت حذف شد ",
		"Export removed faild":     "حذف فاکتور با خطا مواچه شد ",
		"payments removed success": "پرداختی با موفقیت حذف شد",
		"payments removed faild":   "حذف پرداختی با مشکل مواجه شد",
	}
	return messages[value]
}
