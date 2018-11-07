package migrations

import (
	"github.com/ProtocolONE/p1pay.api/database/model"
	"github.com/ProtocolONE/p1pay.api/manager"
	"github.com/globalsign/mgo"
	"github.com/xakep666/mongo-migrate"
	"time"
)

var countries = []*model.Country{
	{
		CodeInt:   36,
		CodeA2:    "AU",
		CodeA3:    "AUS",
		Name:      &model.Name{RU: "Австралия", EN: "Australia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   40,
		CodeA2:    "AT",
		CodeA3:    "AUT",
		Name:      &model.Name{RU: "Австрия", EN: "Austria"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   31,
		CodeA2:    "AZ",
		CodeA3:    "AZE",
		Name:      &model.Name{RU: "Азербайджан", EN: "Azerbaijan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   8,
		CodeA2:    "AL",
		CodeA3:    "ALB",
		Name:      &model.Name{RU: "Албания", EN: "Albania"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   12,
		CodeA2:    "DZ",
		CodeA3:    "DZA",
		Name:      &model.Name{RU: "Алжир", EN: "Algeria"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   24,
		CodeA2:    "AO",
		CodeA3:    "AGO",
		Name:      &model.Name{RU: "Ангола", EN: "Angola"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   20,
		CodeA2:    "AD",
		CodeA3:    "AND",
		Name:      &model.Name{RU: "Андорра", EN: "Andorra"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   10,
		CodeA2:    "AQ",
		CodeA3:    "ATA",
		Name:      &model.Name{RU: "Антарктика", EN: "Antarctic"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   28,
		CodeA2:    "AG",
		CodeA3:    "ATG",
		Name:      &model.Name{RU: "Антигуа и Барбуда", EN: "Antigua and Barbuda"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   32,
		CodeA2:    "AR",
		CodeA3:    "ARG",
		Name:      &model.Name{RU: "Аргентина", EN: "Argentina"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   51,
		CodeA2:    "AM",
		CodeA3:    "ARM",
		Name:      &model.Name{RU: "Армения", EN: "Armenia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   533,
		CodeA2:    "AW",
		CodeA3:    "ABW",
		Name:      &model.Name{RU: "Аруба", EN: "Aruba"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   4,
		CodeA2:    "AF",
		CodeA3:    "AFG",
		Name:      &model.Name{RU: "Афганистан", EN: "Afghanistan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   44,
		CodeA2:    "BS",
		CodeA3:    "BHS",
		Name:      &model.Name{RU: "Багамы", EN: "Bahamas"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   50,
		CodeA2:    "BD",
		CodeA3:    "BGD",
		Name:      &model.Name{RU: "Бангладеш", EN: "Bangladesh"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   52,
		CodeA2:    "BB",
		CodeA3:    "BB",
		Name:      &model.Name{RU: "Барбадос", EN: "Barbados"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   48,
		CodeA2:    "BH",
		CodeA3:    "BHR",
		Name:      &model.Name{RU: "Бахрейн", EN: "Bahrain"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   112,
		CodeA2:    "BY",
		CodeA3:    "BLR",
		Name:      &model.Name{RU: "Беларусь", EN: "Belarus"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   84,
		CodeA2:    "BZ",
		CodeA3:    "BLZ",
		Name:      &model.Name{RU: "Белиз", EN: "Belize"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   56,
		CodeA2:    "BE",
		CodeA3:    "BEL",
		Name:      &model.Name{RU: "Бельгия", EN: "Belgium"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   204,
		CodeA2:    "BJ",
		CodeA3:    "BEN",
		Name:      &model.Name{RU: "Бенин", EN: "Benin"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   60,
		CodeA2:    "BM",
		CodeA3:    "BMU",
		Name:      &model.Name{RU: "Бермуды", EN: "Bermuda"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   100,
		CodeA2:    "BG",
		CodeA3:    "BGR",
		Name:      &model.Name{RU: "Болгария", EN: "Bulgaria"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   68,
		CodeA2:    "BO",
		CodeA3:    "BOL",
		Name:      &model.Name{RU: "Боливия", EN: "Bolivia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   70,
		CodeA2:    "BA",
		CodeA3:    "BIH",
		Name:      &model.Name{RU: "Босния и Герцеговина", EN: "Bosnia & Herzegovina"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   72,
		CodeA2:    "BW",
		CodeA3:    "BWA",
		Name:      &model.Name{RU: "Ботсвана", EN: "Botswana"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   76,
		CodeA2:    "BR",
		CodeA3:    "BRA",
		Name:      &model.Name{RU: "Бразилия", EN: "Brazil"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   96,
		CodeA2:    "BN",
		CodeA3:    "BRN",
		Name:      &model.Name{RU: "Бруней Дарассалам", EN: "Brunei Darussalam"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   854,
		CodeA2:    "BF",
		CodeA3:    "BFA",
		Name:      &model.Name{RU: "Буркина‐Фасо", EN: "Burkina‐Faso"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   108,
		CodeA2:    "BI",
		CodeA3:    "BDI",
		Name:      &model.Name{RU: "Бурунди", EN: "Burundi"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   64,
		CodeA2:    "BT",
		CodeA3:    "BTN",
		Name:      &model.Name{RU: "Бутан", EN: "Bhutan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   548,
		CodeA2:    "VU",
		CodeA3:    "VUT",
		Name:      &model.Name{RU: "Вануату", EN: "Vanuatu"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   336,
		CodeA2:    "VA",
		CodeA3:    "VAT",
		Name:      &model.Name{RU: "Ватикан", EN: "Vatican (Holy See)"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   826,
		CodeA2:    "GB",
		CodeA3:    "GBR",
		Name:      &model.Name{RU: "Великобритания", EN: "Great Britain (United Kingdom)"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   348,
		CodeA2:    "HU",
		CodeA3:    "HUN",
		Name:      &model.Name{RU: "Венгрия", EN: "Hungary"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   862,
		CodeA2:    "VE",
		CodeA3:    "VEN",
		Name:      &model.Name{RU: "Венесуэла", EN: "Venezuela"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   626,
		CodeA2:    "TP",
		CodeA3:    "TMP",
		Name:      &model.Name{RU: "Восточный Тимор", EN: "East Timor"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   704,
		CodeA2:    "VN",
		CodeA3:    "VNM",
		Name:      &model.Name{RU: "Вьетнам", EN: "Vietnam"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   266,
		CodeA2:    "GA",
		CodeA3:    "GAB",
		Name:      &model.Name{RU: "Габон", EN: "Gabon"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   332,
		CodeA2:    "HT",
		CodeA3:    "HTI",
		Name:      &model.Name{RU: "Гаити", EN: "Haiti"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   328,
		CodeA2:    "GY",
		CodeA3:    "GUY",
		Name:      &model.Name{RU: "Гайана", EN: "Guyana"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   270,
		CodeA2:    "GM",
		CodeA3:    "GMB",
		Name:      &model.Name{RU: "Гамбия", EN: "Gambia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   288,
		CodeA2:    "GH",
		CodeA3:    "GHA",
		Name:      &model.Name{RU: "Гана", EN: "Ghana"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   312,
		CodeA2:    "GP",
		CodeA3:    "GLP",
		Name:      &model.Name{RU: "Гваделупа", EN: "Guadeloupe"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   320,
		CodeA2:    "GT",
		CodeA3:    "GTM",
		Name:      &model.Name{RU: "Гватемала", EN: "Guatemala"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   324,
		CodeA2:    "GN",
		CodeA3:    "GIN",
		Name:      &model.Name{RU: "Гвинея", EN: "Guinea"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   624,
		CodeA2:    "GW",
		CodeA3:    "GNB",
		Name:      &model.Name{RU: "Гвинея‐Бисау", EN: "Guinea‐Bissau"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   276,
		CodeA2:    "DE",
		CodeA3:    "DEU",
		Name:      &model.Name{RU: "Германия", EN: "Germany"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   292,
		CodeA2:    "GI",
		CodeA3:    "GIB",
		Name:      &model.Name{RU: "Гибралтар", EN: "Gibraltar"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   340,
		CodeA2:    "HN",
		CodeA3:    "HND",
		Name:      &model.Name{RU: "Гондурас", EN: "Honduras"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   344,
		CodeA2:    "HK",
		CodeA3:    "HKG",
		Name:      &model.Name{RU: "Гонконг", EN: "Hong Kong"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   308,
		CodeA2:    "GD",
		CodeA3:    "GRD",
		Name:      &model.Name{RU: "Гренада", EN: "Grenada"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   304,
		CodeA2:    "GL",
		CodeA3:    "GRL",
		Name:      &model.Name{RU: "Гренландия", EN: "Greenland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   300,
		CodeA2:    "GR",
		CodeA3:    "GRC",
		Name:      &model.Name{RU: "Греция", EN: "Greece"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   268,
		CodeA2:    "GE",
		CodeA3:    "GEO",
		Name:      &model.Name{RU: "Грузия", EN: "Georgia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   316,
		CodeA2:    "GU",
		CodeA3:    "GUM",
		Name:      &model.Name{RU: "Гуам", EN: "Guam"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   208,
		CodeA2:    "DK",
		CodeA3:    "DNK",
		Name:      &model.Name{RU: "Дания", EN: "Denmark"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   180,
		CodeA2:    "CD",
		CodeA3:    "COD",
		Name:      &model.Name{RU: "Конго", EN: "Congo"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   262,
		CodeA2:    "DJ",
		CodeA3:    "DJI",
		Name:      &model.Name{RU: "Джибути", EN: "Djibouti"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   212,
		CodeA2:    "DM",
		CodeA3:    "DMA",
		Name:      &model.Name{RU: "Доминика", EN: "Dominica"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   214,
		CodeA2:    "DO",
		CodeA3:    "DOM",
		Name:      &model.Name{RU: "Доминиканская Республика", EN: "Dominican Republic"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   818,
		CodeA2:    "EG",
		CodeA3:    "EGY",
		Name:      &model.Name{RU: "Египет", EN: "Egypt"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   894,
		CodeA2:    "ZM",
		CodeA3:    "ZMB",
		Name:      &model.Name{RU: "Замбия", EN: "Zambia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   732,
		CodeA2:    "EH",
		CodeA3:    "ESH",
		Name:      &model.Name{RU: "Западная Сахара", EN: "Western Sahara"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   716,
		CodeA2:    "ZW",
		CodeA3:    "ZWE",
		Name:      &model.Name{RU: "Зимбабве", EN: "Zimbabwe"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   376,
		CodeA2:    "IL",
		CodeA3:    "ISR",
		Name:      &model.Name{RU: "Израиль", EN: "Israel"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   356,
		CodeA2:    "IN",
		CodeA3:    "IND",
		Name:      &model.Name{RU: "Индия", EN: "India"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   360,
		CodeA2:    "ID",
		CodeA3:    "IDN",
		Name:      &model.Name{RU: "Индонезия", EN: "Indonesia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   400,
		CodeA2:    "JO",
		CodeA3:    "JOR",
		Name:      &model.Name{RU: "Иордания", EN: "Jordan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   368,
		CodeA2:    "IQ",
		CodeA3:    "IRQ",
		Name:      &model.Name{RU: "Ирак", EN: "Iraq"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   364,
		CodeA2:    "IR",
		CodeA3:    "IRN",
		Name:      &model.Name{RU: "Иран", EN: "Iran"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   372,
		CodeA2:    "IE",
		CodeA3:    "IRL",
		Name:      &model.Name{RU: "Ирландия", EN: "Ireland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   352,
		CodeA2:    "IS",
		CodeA3:    "ISL",
		Name:      &model.Name{RU: "Исландия", EN: "Iceland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   724,
		CodeA2:    "ES",
		CodeA3:    "ESP",
		Name:      &model.Name{RU: "Испания", EN: "Spain"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   380,
		CodeA2:    "IT",
		CodeA3:    "ITA",
		Name:      &model.Name{RU: "Италия", EN: "Italy"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   887,
		CodeA2:    "YE",
		CodeA3:    "YEM",
		Name:      &model.Name{RU: "Йемен", EN: "Yemen"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   132,
		CodeA2:    "CV",
		CodeA3:    "CPV",
		Name:      &model.Name{RU: "Кабо‐Верде", EN: "Cape Verde"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   398,
		CodeA2:    "KZ",
		CodeA3:    "KAZ",
		Name:      &model.Name{RU: "Казахстан", EN: "Kazakhstan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   116,
		CodeA2:    "KH",
		CodeA3:    "KHM",
		Name:      &model.Name{RU: "Камбоджа", EN: "Cambodia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   120,
		CodeA2:    "CM",
		CodeA3:    "CMR",
		Name:      &model.Name{RU: "Камерун", EN: "Cameroon"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   124,
		CodeA2:    "CA",
		CodeA3:    "CAN",
		Name:      &model.Name{RU: "Канада", EN: "Canada"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   634,
		CodeA2:    "QA",
		CodeA3:    "QAT",
		Name:      &model.Name{RU: "Катар", EN: "Qatar"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   404,
		CodeA2:    "KE",
		CodeA3:    "KEN",
		Name:      &model.Name{RU: "Кения", EN: "Kenya"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   196,
		CodeA2:    "CY",
		CodeA3:    "CYP",
		Name:      &model.Name{RU: "Кипр", EN: "Cyprus"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   417,
		CodeA2:    "KG",
		CodeA3:    "KGZ",
		Name:      &model.Name{RU: "Киргизстан", EN: "Kirghizstan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   296,
		CodeA2:    "KI",
		CodeA3:    "KIR",
		Name:      &model.Name{RU: "Кирибати", EN: "Kiribati"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   156,
		CodeA2:    "CN",
		CodeA3:    "CHN",
		Name:      &model.Name{RU: "Китай", EN: "China"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   170,
		CodeA2:    "CO",
		CodeA3:    "COL",
		Name:      &model.Name{RU: "Колумбия", EN: "Colombia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   178,
		CodeA2:    "CG",
		CodeA3:    "COG",
		Name:      &model.Name{RU: "Конго", EN: "Congo"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   188,
		CodeA2:    "CR",
		CodeA3:    "CRI",
		Name:      &model.Name{RU: "Коста‐Рика", EN: "Costa Rica"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   384,
		CodeA2:    "CI",
		CodeA3:    "CIV",
		Name:      &model.Name{RU: "Кот де Ивуар", EN: "Cote d'Ivoire"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   192,
		CodeA2:    "CU",
		CodeA3:    "CUB",
		Name:      &model.Name{RU: "Куба", EN: "Cuba"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   414,
		CodeA2:    "KW",
		CodeA3:    "KWT",
		Name:      &model.Name{RU: "Кувейт", EN: "Kuwait"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   428,
		CodeA2:    "LV",
		CodeA3:    "LVA",
		Name:      &model.Name{RU: "Латвия", EN: "Latvia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   426,
		CodeA2:    "LS",
		CodeA3:    "LSO",
		Name:      &model.Name{RU: "Лесото", EN: "Lesotho"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   430,
		CodeA2:    "LR",
		CodeA3:    "LBR",
		Name:      &model.Name{RU: "Либерия", EN: "Liberia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   422,
		CodeA2:    "LB",
		CodeA3:    "LBN",
		Name:      &model.Name{RU: "Ливан", EN: "Lebanon"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   434,
		CodeA2:    "LY",
		CodeA3:    "LBY",
		Name:      &model.Name{RU: "Ливия", EN: "Libyan Arab Jamahiriya"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   440,
		CodeA2:    "LT",
		CodeA3:    "LTU",
		Name:      &model.Name{RU: "Литва", EN: "Lithuania"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   438,
		CodeA2:    "LI",
		CodeA3:    "LIE",
		Name:      &model.Name{RU: "Лихтенштейн", EN: "Liechtenstein"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   442,
		CodeA2:    "LU",
		CodeA3:    "LUX",
		Name:      &model.Name{RU: "Люксембург", EN: "Luxembourg"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   480,
		CodeA2:    "MU",
		CodeA3:    "MUS",
		Name:      &model.Name{RU: "Маврикий", EN: "Mauritius"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   478,
		CodeA2:    "MR",
		CodeA3:    "MRT",
		Name:      &model.Name{RU: "Мавритания", EN: "Mauritania"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   450,
		CodeA2:    "MG",
		CodeA3:    "MDG",
		Name:      &model.Name{RU: "Мадагаскар", EN: "Madagascar"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   446,
		CodeA2:    "MO",
		CodeA3:    "MAC",
		Name:      &model.Name{RU: "Макао", EN: "Macau"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   807,
		CodeA2:    "MK",
		CodeA3:    "MKD",
		Name:      &model.Name{RU: "Македония", EN: "Macedonia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   454,
		CodeA2:    "MW",
		CodeA3:    "MWI",
		Name:      &model.Name{RU: "Малави", EN: "Malawi"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   458,
		CodeA2:    "MY",
		CodeA3:    "MYS",
		Name:      &model.Name{RU: "Малайзия", EN: "Malaysia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   466,
		CodeA2:    "ML",
		CodeA3:    "MLI",
		Name:      &model.Name{RU: "Мали", EN: "Mali"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   462,
		CodeA2:    "MV",
		CodeA3:    "MDV",
		Name:      &model.Name{RU: "Мальдивы", EN: "Maldives"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   470,
		CodeA2:    "MT",
		CodeA3:    "MLT",
		Name:      &model.Name{RU: "Мальта", EN: "Malta"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   504,
		CodeA2:    "MA",
		CodeA3:    "MAR",
		Name:      &model.Name{RU: "Марокко", EN: "Marocco"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   474,
		CodeA2:    "MQ",
		CodeA3:    "MTQ",
		Name:      &model.Name{RU: "Мартиника", EN: "Martinique"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   484,
		CodeA2:    "MX",
		CodeA3:    "MEX",
		Name:      &model.Name{RU: "Мексика", EN: "Mexico"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   508,
		CodeA2:    "MZ",
		CodeA3:    "MOZ",
		Name:      &model.Name{RU: "Мозамбик", EN: "Mozambique"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   498,
		CodeA2:    "MD",
		CodeA3:    "MDA",
		Name:      &model.Name{RU: "Молдова", EN: "Moldova"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   492,
		CodeA2:    "MC",
		CodeA3:    "MCO",
		Name:      &model.Name{RU: "Монако", EN: "Monaco"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   496,
		CodeA2:    "MN",
		CodeA3:    "MNG",
		Name:      &model.Name{RU: "Монголия", EN: "Mongolia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   104,
		CodeA2:    "MM",
		CodeA3:    "MMR",
		Name:      &model.Name{RU: "Мьянма", EN: "Myanmar"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   516,
		CodeA2:    "NA",
		CodeA3:    "NAM",
		Name:      &model.Name{RU: "Намибия", EN: "Namibia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   520,
		CodeA2:    "NR",
		CodeA3:    "NRU",
		Name:      &model.Name{RU: "Науру", EN: "Nauru"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   524,
		CodeA2:    "NP",
		CodeA3:    "NPL",
		Name:      &model.Name{RU: "Непал", EN: "Nepal"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   562,
		CodeA2:    "NE",
		CodeA3:    "NER",
		Name:      &model.Name{RU: "Нигер", EN: "Niger"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   566,
		CodeA2:    "NG",
		CodeA3:    "NGA",
		Name:      &model.Name{RU: "Нигерия", EN: "Nigeria"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   528,
		CodeA2:    "NL",
		CodeA3:    "NLD",
		Name:      &model.Name{RU: "Нидерланды", EN: "Netherlands"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   558,
		CodeA2:    "NI",
		CodeA3:    "NIC",
		Name:      &model.Name{RU: "Никарагуа", EN: "Nicaragua"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   554,
		CodeA2:    "NZ",
		CodeA3:    "NZL",
		Name:      &model.Name{RU: "Новая Зеландия", EN: "New Zealand"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   578,
		CodeA2:    "NO",
		CodeA3:    "NOR",
		Name:      &model.Name{RU: "Норвегия", EN: "Norway"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   784,
		CodeA2:    "AE",
		CodeA3:    "ARE",
		Name:      &model.Name{RU: "Объединенные Арабские Эмираты", EN: "United Arab Emirates"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   512,
		CodeA2:    "OM",
		CodeA3:    "OMN",
		Name:      &model.Name{RU: "Оман", EN: "Oman"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   586,
		CodeA2:    "PK",
		CodeA3:    "PAK",
		Name:      &model.Name{RU: "Пакистан", EN: "Pakistan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   591,
		CodeA2:    "PA",
		CodeA3:    "PAN",
		Name:      &model.Name{RU: "Панама", EN: "Panama"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   598,
		CodeA2:    "PG",
		CodeA3:    "PNG",
		Name:      &model.Name{RU: "Папуа‐Новая Гвинея", EN: "Papua New Guinea"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   600,
		CodeA2:    "PY",
		CodeA3:    "PRY",
		Name:      &model.Name{RU: "Парагвай", EN: "Paraguay"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   604,
		CodeA2:    "PE",
		CodeA3:    "PER",
		Name:      &model.Name{RU: "Перу", EN: "Peru"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   616,
		CodeA2:    "PL",
		CodeA3:    "POL",
		Name:      &model.Name{RU: "Польша", EN: "Poland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   620,
		CodeA2:    "PT",
		CodeA3:    "PRT",
		Name:      &model.Name{RU: "Португалия", EN: "Portugal"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   643,
		CodeA2:    "RU",
		CodeA3:    "RUS",
		Name:      &model.Name{RU: "Россия", EN: "Russia (Russian Federation)"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   646,
		CodeA2:    "RW",
		CodeA3:    "RWA",
		Name:      &model.Name{RU: "Руанда", EN: "Rwanda"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   642,
		CodeA2:    "RO",
		CodeA3:    "ROM",
		Name:      &model.Name{RU: "Румыния", EN: "Romania"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   222,
		CodeA2:    "SV",
		CodeA3:    "SLV",
		Name:      &model.Name{RU: "Сальвадор", EN: "El Salvador"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   882,
		CodeA2:    "WS",
		CodeA3:    "WSM",
		Name:      &model.Name{RU: "Самоа", EN: "Samoa"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   682,
		CodeA2:    "SA",
		CodeA3:    "SAU",
		Name:      &model.Name{RU: "Саудовская Аравия", EN: "Saudi Arabia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   748,
		CodeA2:    "SZ",
		CodeA3:    "SWZ",
		Name:      &model.Name{RU: "Свазиленд", EN: "Swaziland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   686,
		CodeA2:    "SN",
		CodeA3:    "SEN",
		Name:      &model.Name{RU: "Сенегал", EN: "Senegal"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   702,
		CodeA2:    "SG",
		CodeA3:    "SGP",
		Name:      &model.Name{RU: "Сингапур", EN: "Singapore"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   760,
		CodeA2:    "SY",
		CodeA3:    "SYR",
		Name:      &model.Name{RU: "Сирия", EN: "Syria"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   703,
		CodeA2:    "SK",
		CodeA3:    "SVK",
		Name:      &model.Name{RU: "Словакия", EN: "Slovak Republic"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   705,
		CodeA2:    "SI",
		CodeA3:    "SVN",
		Name:      &model.Name{RU: "Словения", EN: "Slovenia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   840,
		CodeA2:    "US",
		CodeA3:    "USA",
		Name:      &model.Name{RU: "Соединенные Штаты Америки", EN: "United States of America"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   706,
		CodeA2:    "SO",
		CodeA3:    "SOM",
		Name:      &model.Name{RU: "Сомали", EN: "Somali"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   736,
		CodeA2:    "SD",
		CodeA3:    "SDN",
		Name:      &model.Name{RU: "Судан", EN: "Sudan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   740,
		CodeA2:    "SR",
		CodeA3:    "SUR",
		Name:      &model.Name{RU: "Суринам", EN: "Surinam"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   762,
		CodeA2:    "TJ",
		CodeA3:    "TJK",
		Name:      &model.Name{RU: "Таджикистан", EN: "Tadjikistan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   764,
		CodeA2:    "TH",
		CodeA3:    "THA",
		Name:      &model.Name{RU: "Таиланд", EN: "Thailand"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   158,
		CodeA2:    "TW",
		CodeA3:    "TWN",
		Name:      &model.Name{RU: "Тайвань", EN: "Taiwan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   834,
		CodeA2:    "TZ",
		CodeA3:    "TZA",
		Name:      &model.Name{RU: "Танзания", EN: "Tanzania"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   768,
		CodeA2:    "TG",
		CodeA3:    "TGO",
		Name:      &model.Name{RU: "Того", EN: "Togo"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   776,
		CodeA2:    "TO",
		CodeA3:    "TON",
		Name:      &model.Name{RU: "Тонга", EN: "Tonga"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   780,
		CodeA2:    "TT",
		CodeA3:    "TTO",
		Name:      &model.Name{RU: "Тринидад и Тобаго", EN: "Trinidad and Tobago"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   798,
		CodeA2:    "TV",
		CodeA3:    "TUV",
		Name:      &model.Name{RU: "Тувалу", EN: "Tuvalu"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   788,
		CodeA2:    "TN",
		CodeA3:    "TUN",
		Name:      &model.Name{RU: "Тунис", EN: "Tunisia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   795,
		CodeA2:    "TM",
		CodeA3:    "TKM",
		Name:      &model.Name{RU: "Туркменистан", EN: "Turkmenistan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   792,
		CodeA2:    "TR",
		CodeA3:    "TUR",
		Name:      &model.Name{RU: "Турция", EN: "Turkey"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   800,
		CodeA2:    "UG",
		CodeA3:    "UGA",
		Name:      &model.Name{RU: "Уганда", EN: "Uganda"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   860,
		CodeA2:    "UZ",
		CodeA3:    "UZB",
		Name:      &model.Name{RU: "Узбекистан", EN: "Uzbekistan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   804,
		CodeA2:    "UA",
		CodeA3:    "UKR",
		Name:      &model.Name{RU: "Украина", EN: "Ukraine"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   858,
		CodeA2:    "UY",
		CodeA3:    "URY",
		Name:      &model.Name{RU: "Уругвай", EN: "Uruguay"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   242,
		CodeA2:    "FJ",
		CodeA3:    "FJI",
		Name:      &model.Name{RU: "Фиджи", EN: "Fiji"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   608,
		CodeA2:    "PH",
		CodeA3:    "PHL",
		Name:      &model.Name{RU: "Филиппины", EN: "Philippines"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   246,
		CodeA2:    "FI",
		CodeA3:    "FIN",
		Name:      &model.Name{RU: "Финляндия", EN: "Finland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   250,
		CodeA2:    "FR",
		CodeA3:    "FRA",
		Name:      &model.Name{RU: "Франция", EN: "France"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   191,
		CodeA2:    "HR",
		CodeA3:    "HRV",
		Name:      &model.Name{RU: "Хорватия", EN: "Croatia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   148,
		CodeA2:    "TD",
		CodeA3:    "TCD",
		Name:      &model.Name{RU: "Чад", EN: "Chad"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   203,
		CodeA2:    "CZ",
		CodeA3:    "CZE",
		Name:      &model.Name{RU: "Чехия", EN: "Czech Republic"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   152,
		CodeA2:    "CL",
		CodeA3:    "CHL",
		Name:      &model.Name{RU: "Чили", EN: "Chili"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   756,
		CodeA2:    "CH",
		CodeA3:    "CHE",
		Name:      &model.Name{RU: "Швейцария", EN: "Switzerland"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   752,
		CodeA2:    "SE",
		CodeA3:    "SWE",
		Name:      &model.Name{RU: "Швеция", EN: "Sweden"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   144,
		CodeA2:    "LK",
		CodeA3:    "LKA",
		Name:      &model.Name{RU: "Шри‐Ланка", EN: "Sri Lanka"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   218,
		CodeA2:    "EC",
		CodeA3:    "ECU",
		Name:      &model.Name{RU: "Эквадор", EN: "Ecuador"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   232,
		CodeA2:    "ER",
		CodeA3:    "ERI",
		Name:      &model.Name{RU: "Эритрия", EN: "Eritrea"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   233,
		CodeA2:    "EE",
		CodeA3:    "EST",
		Name:      &model.Name{RU: "Эстония", EN: "Estonia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   231,
		CodeA2:    "ET",
		CodeA3:    "ETH",
		Name:      &model.Name{RU: "Эфиопия", EN: "Ethiopia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   891,
		CodeA2:    "YU",
		CodeA3:    "YUG",
		Name:      &model.Name{RU: "Югославия", EN: "Yugoslavia"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   710,
		CodeA2:    "ZA",
		CodeA3:    "ZAF",
		Name:      &model.Name{RU: "Южная Африка", EN: "South Africa"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   410,
		CodeA2:    "KR",
		CodeA3:    "KOR",
		Name:      &model.Name{RU: "Южная Корея", EN: "South Korea"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   388,
		CodeA2:    "JM",
		CodeA3:    "JAM",
		Name:      &model.Name{RU: "Ямайка", EN: "Jamaica"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		CodeInt:   392,
		CodeA2:    "JP",
		CodeA3:    "JPN",
		Name:      &model.Name{RU: "Япония", EN: "Japan"},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func init() {
	err := migrate.Register(
		func(db *mgo.Database) error {
			var err error
			c := db.C(manager.TableCountry)

			err = c.EnsureIndex(mgo.Index{Name: "county_code_int_uniq", Key: []string{"code_int"}, Unique: true})

			if err != nil {
				return err
			}

			err = c.EnsureIndex(mgo.Index{Name: "county_is_active_idx", Key: []string{"is_active"}})

			if err != nil {
				return err
			}

			var iCountries []interface{}

			for _, t := range countries {
				iCountries = append(iCountries, t)
			}

			return c.Insert(iCountries...)
		},
		func(db *mgo.Database) error {
			return db.C(manager.TableCountry).DropCollection()
		},
	)

	if err != nil {
		return
	}
}