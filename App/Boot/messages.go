package Boot

func Messages(value string) string {
	var messages = map[string]string{
		"login success":          "با موفقیت وارد شدید",
		"login faild":            "اطلاعات ورود اشتباه است",
		"user made success":      "کاربر با موفقیت افزوده شد",
		"user made faild":        "افزودن کاربر با مشکل مواجه شد",
		"product made success":   "محصول با موفقیت افزوده شد",
		"product made faild":     "افزودن محصول با مشکل مواجه شد",
		"Export removed success": "فاکتور با موفقیت حذف شد ",
	}
	return messages[value]
}
