package bot

import (
	"bytes"
	"html/template"
)

const (
	HELP_TPL = `
    Заготовка бота

    Доступные команды:
     /start - регистрация пользователя. Необходимо указать прдопределенный пароль
     /help - справка
    `
)

func RenderTemplate(tpl string, data interface{}) (string, error) {
	parsedTpl, err := template.New("").Parse(tpl)

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := parsedTpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
