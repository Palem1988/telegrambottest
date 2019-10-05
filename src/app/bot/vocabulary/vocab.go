package vocab

var Translates = map[string]map[string]string{
	"Hello": {
		"en": "Hello, I'm an exchange bot for *BIP ⇆ BTC.* \n*Please choose your language.*",
		"ru": "Hello, I'm an exchange bot for *BIP ⇆ BTC.* \n*Please choose your language.*",
	},
	"Language": {
		"en": "Language",
		"ru": "Язык",
	},
	"english": {
		"en": "English",
		"ru": "английский",
	},
	"russian": {
		"en": "russian",
		"ru": "Русский",
	},
	"Installed": {
		"en": "Installed",
		"ru": "Установлен",
	},
	"Menu": {
		"en": "📕 Menu",
		"ru": "📕 Меню",
	},
	"Settings": {
		"en": "🔧 Settings",
		"ru": "🔧 Настройки",
	},
	"Yes": {
		"en": "Yes",
		"ru": "Да",
	},
	"No": {
		"en": "No",
		"ru": "Нет",
	},
	"Select": {
		"en": "Текущий курс: *$%.2f* (%s %%).\n\nЗдесь вы можете купить или продать *BIP*, а также отслеживать созданные заявки.",
		"ru": "Текущий курс: *$%.2f* (%s %%).\n\nЗдесь вы можете купить или продать *BIP*, а также отслеживать созданные заявки.",
	},
	// Купить
	// 1
	"New minter": {
		"en": "Введите ваш адрес в сети Minter.\n\n*Пример:* Mx00000000000000000000000000000000000001",
		"ru": "Введите ваш адрес в сети Minter.\n\n*Пример:* Mx00000000000000000000000000000000000001",
	},
	"Select minter": {
		"en": "Выберите ваш адрес в сети Minter или введите новый.\n\n*Пример:* Mx00000000000000000000000000000000000001",
		"ru": "Выберите ваш адрес в сети Minter или введите новый.\n\n*Пример:* Mx00000000000000000000000000000000000001",
	},
	"Wrong minter": {
		"en": "Проверьте правильность введённого адреса, он должен содержать *42 символа* и начинаться с *Mx*.",
		"ru": "Проверьте правильность введённого адреса, он должен содержать *42 символа* и начинаться с *Mx*.",
	},
	// 2
	"New Email": {
		"en": "Введите ваш почтовый адрес.\n\n*Пример:* mail@example.com",
		"ru": "Введите ваш почтовый адрес.\n\n*Пример:* mail@example.com",
	},
	"Select email": {
		"en": "Выберите ваш почтовый адрес или введите новый.\n\n*Пример:* mail@example.com",
		"ru": "Выберите ваш почтовый адрес или введите новый.\n\n*Пример:* mail@example.com",
	},
	"Wrong email": {
		"en": "Проверьте правильность введённого адреса.",
		"ru": "Проверьте правильность введённого адреса.",
	},
	// 3 Отправьте BTC ... 2 подтверждения...
	"Send deposit": {
		"en": "Отправьте BTC на следующий адрес, после *2* подтверждений сети, вы получите BIP на указанный вами адрес в сети Minter.\n\n*Текущий курс:* $%.2f (%s %%)\n\n" +
			"💡Сейчас за 1 BTC вы можете купить *250 631* BIP, это на *111%% больше* актуальной цены.",
		"ru": "Отправьте BTC на следующий адрес, после *2* подтверждений сети, вы получите BIP на указанный вами адрес в сети Minter.\n\n*Текущий курс:* $%.2f (%s %%)\n\n" +
			"💡Сейчас за 1 BTC вы можете купить *250 631* BIP, это на *111%% больше* актуальной цены.",
	},
	// 4
	"Check": {
		"en": "Проверить",
		"ru": "Проверить",
	},
	"Wait deposit": {
		"en": "Ожидание транзакции BTC…",
		"ru": "Ожидание транзакции BTC…",
	},
	"New deposit": {
		"en": "BTC уже в пути, вы получите как минимум %.2f BIP.",
		"ru": "BTC уже в пути, вы получите как минимум %.2f BIP.",
	},
	"No buy": {
		"en": "У вас нет заявок на покупку.",
		"ru": "У вас нет заявок на покупку.",
	},
	// 5
	"Exchange is successful": {
		"en": "🎉 *%.2f* BIP были отправлены на ваш адрес.",
		"ru": "🎉 *%.2f* BIP были отправлены на ваш адрес.",
	},
	// 1
	"Coin": {
		"en": "Введите название монеты, которую хотите продать.\n\n*Пример*: BIP",
		"ru": "Введите название монеты, которую хотите продать.\n\n*Пример*: BIP",
	},
	// 1
	"Wrong coin name": {
		"en": "⚠️Ошибка\n\nТакой монеты не существует.",
		"ru": "⚠️Ошибка\n\nТакой монеты не существует.",
	},
	// 2
	"Select price": {
		"en": "Укажите цену в *USD*, которую хотите установить для монет.\n\n*Пример*: 0.32",
		"ru": "Укажите цену в *USD*, которую хотите установить для монет.\n\n*Пример*: 0.32",
	},
	// 2
	"Wrong price": {
		"en": "⚠️Ошибка\n\nВозможный диапазон цены: от *$0.01* до *$0.32*, вводить цену нужно без символов обозначающих валюту и букв.",
		"ru": "⚠️Ошибка\n\nВозможный диапазон цены: от *$0.01* до *$0.32*, вводить цену нужно без символов обозначающих валюту и букв.",
	},
	// 3
	"Send bitcoin": {
		"en": "Введите ваш *Bitcoin* адрес.",
		"ru": "Введите ваш *Bitcoin* адрес.",
	},
	"Select bitcoin": {
		"en": "Select *Bitcoin* address or enter a new one.",
		"ru": "Выберите *Bitcoin* адрес или введите новый.",
	},
	// 3
	"Wrong bitcoin": {
		"en": "⚠️Ошибка\n\nПроверьте правильность введённого BTC адреса.",
		"ru": "⚠️Ошибка\n\nПроверьте правильность введённого BTC адреса.",
	},
	// 4
	"Save": {
		"en": "Сохранить введённый адрес для следующих продаж?",
		"ru": "Сохранить введённый адрес для следующих продаж?",
	},
	// 5
	"Send your coins": {
		"en": "Отправьте *%s* на указанный ниже адрес.\n\n⚠️ Не отправляйте меньше *1000 %s* в одной транзакции.\n\nВы можете отслеживать заявку по этой ссылке:\n%s",
		"ru": "Отправьте *%s* на указанный ниже адрес.\n\n⚠️ Не отправляйте меньше *1000 %s* в одной транзакции.\n\nВы можете отслеживать заявку по этой ссылке:\n%s",
	},
	//6
	"Share": {
		"en": "Поделиться",
		"ru": "Поделиться",
	},
	// 7
	"New deposit for sale": {
		"en": "Новая заявка на продажу: *%s* %s по *%.2f* $.",
		"ru": "Новая заявка на продажу: *%s* %s по *%.2f* $.",
	},
	// Заявки
	"Your loots": {
		"en": "📔*Заявки*\n\nВ этом разделе вы можете найти все активные заявки.",
		"ru": "📔*Заявки*\n\nВ этом разделе вы можете найти все активные заявки.",
	},
	"Loot": {
		"en": "Продажа %s %s по $%v",
		"ru": "Продажа %s %s по $%v",
	},
	"Empty loots": {
		"en": "You haven't got loots for sale.",
		"ru": "У вас нет лотов на продажу.",
	},
	// Кнопки
	"Sell": {
		"en": "Продать",
		"ru": "Продать",
	},
	"Buy": {
		"en": "Купить",
		"ru": "Купить",
	},
	"Loots": {
		"en": "Мои заявки",
		"ru": "Мои заявки",
	},
	"Cancel": {
		"en": "« Canlcel",
		"ru": "« Назад",
	},
	// ------------------------------------------------------
	"Now": {
		"en": "Текущий курс: $%.2f (+%.2f %)",
		"ru": "Текущий курс: $%.2f (+%.2f %)",
	},
	"Coin exchanged": {
		"en": "%.4f %s exchanged for %.4f BTC.",
		"ru": "%.4f %s обменяны на %.4f BTC.",
	},
	"Minter deposit and tag": {
		"en": "Your minter deposit address and tag:",
		"ru": "Твой minter адрес для депозита и tag:",
	},
	"Error": {
		"en": "⚠️*Ошибка*",
		"ru": "⚠️*Ошибка*",
	},
	"timeout": {
		"en": "Deposit timed out.",
		"ru": "Время ожидания депозита истекло.",
	},
}

func GetTranslate(key string, lang string) string {
	return Translates[key][lang]
}
