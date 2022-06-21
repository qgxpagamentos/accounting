package accounting

// Locale - Locale
type Locale struct {
	Name           string // currency name
	FractionLength int    // default decimal length
	ThouSep        string // thousands seperator
	DecSep         string // decimal seperator
	SpaceSep       string // space seperator
	UTFSymbol      string // UTF symbol
	HTMLSymbol     string // HTML symbol
	ComSymbol      string // Common symbol
	Pre            bool   // symbol before or after currency
}

const empty string = ""

// LocaleInfo - LocaleInfo
var LocaleInfo map[string]Locale = map[string]Locale{
	"AED": {"UAE Dirham", 2, ",", ".", " ", empty, empty, "Dhs.", true},
	"AFA": {"Afghani", 0, empty, empty, "060B", "&#x060B;", empty, "؋", true},
	"ALL": {"Lek", 2, empty, empty, "", empty, empty, "Lek", true},
	"AMD": {"Armenian Dram", 2, ",", ".", "", empty, empty, "֏", false},
	"ANG": {"Antillian Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "ƒ", true},
	"AOA": {"New Kwanza", 0, empty, empty, "", empty, empty, "Kz", true},
	"ARS": {"Argentine Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"ATS": {"Schilling", 2, ".", ",", " ", empty, empty, "öS", true},
	"AUD": {"Australian Dollar", 2, " ", ".", "", "0024", "&#x0024;", "$", true},
	"AWG": {"Aruban Guilder", 2, ",", ".", " ", "0192", "&#x0192;", "ƒ", true},
	"AZN": {"Azerbaijanian Manat", 2, empty, empty, "", empty, empty, "₼", true},
	"BAM": {"Convertible Marks", 2, ",", ".", "", empty, empty, "KM", false},
	"BBD": {"Barbados Dollar", 2, empty, empty, "", "0024", "&#x0024;", "Bds$", true},
	"BDT": {"Taka", 2, ",", ".", " ", empty, empty, "Tk", true},
	"BEF": {"Belgian Franc", 0, ".", "", " ", "20A3", "&#x20A3;", "BEF", true},
	"BGN": {"Lev", 2, " ", ",", " ", empty, empty, "лв", false},
	"BHD": {"Bahraini Dinar", 3, ",", ".", " ", empty, empty, "د.ب", true},
	"BIF": {"Burundi Franc", 0, empty, empty, "", empty, empty, "FBu", true},
	"BMD": {"Bermudian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BND": {"Brunei Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BOB": {"Bolivian Boliviano", 2, ",", ".", "", empty, empty, "$b", true},
	"BRL": {"Brazilian Real", 2, ".", ",", " ", "0052 0024", "R$", "R$", true},
	"BSD": {"Bahamian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BTN": {"Bhutan Ngultrum", 2, empty, empty, "", empty, empty, "BTN", true},
	"BWP": {"Pula", 2, ",", ".", "", empty, empty, "P", true},
	"BYR": {"Belarussian Ruble", 0, empty, empty, "", empty, empty, "p.", true},
	"BZD": {"Belize Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"CAD": {"Canadian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "CA$", true},
	"CDF": {"Franc Congolais", 2, empty, empty, "", empty, empty, "FC", true},
	"CHF": {"Swiss Franc", 2, "'", ".", " ", empty, empty, "CHF", true},
	"CLP": {"Chilean Peso", 0, ".", "", "", "20B1", "&#x20B1;", "$", true},
	"CNY": {"Yuan Renminbi", 2, ",", ".", "", "5713", "&#x5713;", "¥", true},
	"COP": {"Colombian Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"CRC": {"Costa Rican Colon", 2, ".", ",", " ", "20A1", "&#x20A1;", "₡", true},
	"CUP": {"Cuban Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"CVE": {"Cape Verde Escudo", 0, empty, empty, "", empty, empty, "$", true},
	"CYP": {"Cyprus Pound", 2, ".", ",", "", "00A3", "&#x00A3;", "£", true},
	"CZK": {"Czech Koruna", 2, ".", ",", " ", empty, empty, "Kč", false},
	"DEM": {"Deutsche Mark", 2, ".", ",", "", empty, empty, "DM", false},
	"DJF": {"Djibouti Franc", 0, empty, empty, "", empty, empty, "DJF", true},
	"DKK": {"Danish Krone", 2, ".", ",", "", empty, empty, "kr.", true},
	"DOP": {"Dominican Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"DZD": {"Algerian Dinar", 2, empty, empty, "", empty, empty, "DA", true},
	"ECS": {"Sucre", 0, empty, empty, "", empty, empty, "S.", true},
	"EEK": {"Kroon", 2, " ", ",", " ", empty, empty, "kr", false},
	"EGP": {"Egyptian Pound", 2, ",", ".", " ", "00A3", "&#x00A3;", "£", true},
	"ERN": {"Nakfa", 0, empty, empty, "", empty, empty, "NKf", true},
	"ESP": {"Spanish Peseta", 0, ".", "", " ", "20A7", "&#x20A7;", "Ptas", false},
	"ETB": {"Ethiopian Birr", 0, empty, empty, "", empty, empty, "BR", true},
	"EUR": {"Euro", 2, ".", ",", "", "20AC", "&#x20AC;", "€", true},
	"FIM": {"Markka", 2, " ", ",", " ", empty, empty, "mk", false},
	"FJD": {"Fiji Dollar", 0, empty, empty, "", "0024", "&#x0024;", "FJ$", true},
	"FKP": {"Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"FRF": {"French Franc", 2, " ", ",", " ", "20A3", "&#x20A3;", "Fr", false},
	"GBP": {"Pound Sterling", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GEL": {"Lari", 0, empty, empty, "", empty, empty, "GEL", true},
	"GHS": {"Cedi", 2, ",", ".", "", "20B5", "&#x20B5;", "₵", true},
	"GIP": {"Gibraltar Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GMD": {"Dalasi", 0, empty, empty, "", empty, empty, "GMD", true},
	"GNF": {"Guinea Franc", 0, empty, empty, empty, empty, empty, "FG", true},
	"GRD": {"Drachma", 2, ".", ",", " ", "20AF", "&#x20AF;", "GRD", false},
	"GTQ": {"Quetzal", 2, ",", ".", "", empty, empty, "Q.", true},
	"GWP": {"Guinea-Bissau Peso", 0, empty, empty, empty, empty, empty, "GWP", true},
	"GYD": {"Guyana Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"HKD": {"Hong Kong Dollar", 2, ",", ".", "", "0024", "&#x0024;", "HK$", true},
	"HNL": {"Lempira", 2, ",", ".", " ", empty, empty, "L", true},
	"HRK": {"Kuna", 2, ".", ",", " ", empty, empty, "kn", false},
	"HTG": {"Gourde", 0, empty, empty, "", empty, empty, "G", true},
	"HUF": {"Forint", 0, ".", "", " ", empty, empty, "Ft", false},
	"IDR": {"Rupiah", 0, ".", ",", "", empty, empty, "Rp", true},
	"IEP": {"Irish Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"ILS": {"New Israeli Sheqel", 2, ",", ".", " ", "20AA", "&#x20AA;", "₪", false},
	"INR": {"Indian Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "₹", true},
	"IQD": {"Iraqi Dinar", 3, empty, empty, "", empty, empty, "د.ع", true},
	"IRR": {"Iranian Rial", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "﷼", true},
	"ISK": {"Iceland Krona", 2, ".", ",", " ", empty, empty, "kr", false},
	"ITL": {"Italian Lira", 0, ".", "", " ", "20A4", "&#x20A4;", "L.", true},
	"JMD": {"Jamaican Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"JOD": {"Jordanian Dinar", 3, ",", ".", " ", empty, empty, "JD", true},
	"JPY": {"Yen", 0, ",", "", "", "00A5", "&#x00A5;", "¥", true},
	"KES": {"Kenyan Shilling", 2, ",", ".", "", empty, empty, "Ksh", true},
	"KGS": {"Som", 0, empty, empty, "", empty, empty, "лв", true},
	"KHR": {"Riel", 2, empty, empty, "", "17DB", "&#x17DB;", "៛", true},
	"KMF": {"Comoro Franc", 0, empty, empty, "", empty, empty, "KMF", true},
	"KPW": {"North Korean Won", 0, empty, empty, "", "20A9", "&#x20A9;", "₩", true},
	"KRW": {"Won", 0, ",", "", "", "20A9", "&#x20A9;", "₩", true},
	"KWD": {"Kuwaiti Dinar", 3, ",", ".", " ", empty, empty, "ك", true},
	"KYD": {"Cayman Islands Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"KZT": {"Tenge", 0, empty, empty, "", empty, empty, "₸", true},
	"LAK": {"Kip", 0, empty, empty, "", "20AD", "&#x20AD;", "₭", true},
	"LBP": {"Lebanese Pound", 0, " ", "", "", "00A3", "&#x00A3;", "ل.ل", false},
	"LKR": {"Sri Lanka Rupee", 0, empty, empty, "", "0BF9", "&#x0BF9;", "₨", true},
	"LRD": {"Liberian Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"LSL": {"Lesotho Maloti", 0, empty, empty, "", empty, empty, "LSL", true},
	"LTL": {"Lithuanian Litas", 2, " ", ",", " ", empty, empty, "Lt", false},
	"LUF": {"Luxembourg Franc", 0, "'", "", " ", "20A3", "&#x20A3;", "F", false},
	"LVL": {"Latvian Lats", 2, ",", ".", " ", empty, empty, "Ls", true},
	"LYD": {"Libyan Dinar", 0, empty, empty, "", empty, empty, "LD", true},
	"MAD": {"Moroccan Dirham", 0, empty, empty, "", empty, empty, "MAD", true},
	"MDL": {"Moldovan Leu", 0, empty, empty, "", empty, empty, "MDL", true},
	"MGF": {"Malagasy Franc", 0, empty, empty, "", empty, empty, "MF", true},
	"MKD": {"Denar", 2, ",", ".", " ", empty, empty, "ден", false},
	"MMK": {"Kyat", 0, empty, empty, "", empty, empty, "K", true},
	"MNT": {"Tugrik", 0, empty, empty, "", "20AE", "&#x20AE;", "₮", true},
	"MOP": {"Pataca", 0, empty, empty, "", empty, empty, "MOP$", true},
	"MRO": {"Ouguiya", 0, empty, empty, "", empty, empty, "MRO", true},
	"MTL": {"Maltese Lira", 2, ",", ".", "", "20A4", "&#x20A4;", "Lm", true},
	"MUR": {"Mauritius Rupee", 0, ",", "", "", "20A8", "&#x20A8;", "Rs", true},
	"MVR": {"Rufiyaa", 0, empty, empty, "", empty, empty, "MVR", true},
	"MWK": {"Kwacha", 2, ",", ".", "", empty, empty, "MK", true},
	"MXN": {"Mexican Peso", 2, ",", ".", " ", "0024", "&#x0024;", "$", true},
	"MYR": {"Malaysian Ringgit", 2, ",", ".", "", empty, empty, "RM", true},
	"MZN": {"Metical", 2, ".", ",", " ", empty, empty, "Mt", false},
	"NAD": {"Namibian Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"NGN": {"Naira", 2, ",", ".", ".", "20A6", "&#x20A6;", "₦", true},
	"NIO": {"Cordoba Oro", 0, empty, empty, "", empty, empty, "C$", true},
	"NLG": {"Netherlands Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "ƒ", true},
	"NOK": {"Norwegian Krone", 2, ".", ",", " ", "kr", "kr", "kr", true},
	"NPR": {"Nepalese Rupee", 2, ",", ".", " ", "20A8", "&#x20A8;", "Rs.", true},
	"NZD": {"New Zealand Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"OMR": {"Rial Omani", 3, ",", ".", " ", "FDFC", "&#xFDFC;", "RO", true},
	"PAB": {"Balboa", 0, empty, empty, "", empty, empty, "B/.", true},
	"PEN": {"Nuevo Sol", 2, ",", ".", " ", "S/.", "S/.", "S/.", true},
	"PGK": {"Kina", 0, empty, empty, "", empty, empty, "K", true},
	"PHP": {"Philippine Peso", 2, ",", ".", "", "20B1", "&#x20B1;", "₱", true},
	"PKR": {"Pakistan Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "Rs", true},
	"PLN": {"Zloty", 2, " ", ",", " ", empty, empty, "zł", false},
	"PTE": {"Portuguese Escudo", 0, ".", "", " ", empty, empty, "$", false},
	"PYG": {"Guarani", 0, empty, empty, "", "20B2", "&#x20B2;", "Gs", true},
	"QAR": {"Qatari Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", "﷼", true},
	"RON": {"Leu", 2, ".", ",", " ", empty, empty, "lei", false},
	"RSD": {"Serbian Dinar", 2, empty, empty, empty, empty, empty, "РСД", false},
	"RUB": {"Russian Ruble", 2, ".", ",", empty, "0440 0443 0431", "&#x0440;&#x0443;&#x0431;", "₽", true},
	"RWF": {"Rwanda Franc", 0, empty, empty, "", empty, empty, "RWF", true},
	"SAC": {"S. African Rand Commerc.", 0, empty, empty, "", empty, empty, "SAC", true},
	"SAR": {"Saudi Riyal", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "﷼", true},
	"SBD": {"Solomon Islands Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"SCR": {"Seychelles Rupee", 0, empty, empty, "", "20A8", "&#x20A8;", "₨", true},
	"SDG": {"Sudanese Dinar", 0, empty, empty, empty, empty, empty, "LSd", true},
	"SDP": {"Sudanese Pound", 0, empty, empty, "", empty, empty, "SDP", true},
	"SEK": {"Swedish Krona", 2, " ", ",", " ", empty, empty, "kr", false},
	"SGD": {"Singapore Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"SHP": {"St Helena Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"SIT": {"Tolar", 2, ".", ",", " ", empty, empty, "SIT", false},
	"SKK": {"Slovak Koruna", 2, " ", ",", " ", empty, empty, "Sk", false},
	"SLL": {"Leone", 0, empty, empty, "", empty, empty, "Le", true},
	"SOS": {"Somali Shilling", 0, empty, empty, "", empty, empty, "S", true},
	"SRG": {"Surinam Guilder", 0, empty, empty, "", empty, empty, "SRG", true},
	"STD": {"Dobra", 0, empty, empty, "", empty, empty, "DB", true},
	"SVC": {"El Salvador Colon", 2, ",", ".", "", "20A1", "&#x20A1;", "¢", true},
	"SYP": {"Syrian Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"SZL": {"Lilangeni", 2, "", ".", "", empty, empty, "E", true},
	"THB": {"Baht", 2, ",", ".", " ", "0E3F", "&#x0E3F;", "Bt", false},
	"TJR": {"Tajik Ruble", 0, empty, empty, "", empty, empty, "TJR", true},
	"TJS": {"Somoni", 0, empty, empty, empty, empty, empty, "TJS", true},
	"TMM": {"Manat", 0, empty, empty, "", empty, empty, "T", true},
	"TND": {"Tunisian Dinar", 3, empty, empty, "", empty, empty, "TND", true},
	"TOP": {"Pa'anga", 2, ",", ".", " ", empty, empty, "$", true},
	"TPE": {"Timor Escudo", 0, empty, empty, empty, empty, empty, "TPE", true},
	"TRY": {"Turkish Lira", 0, ",", "", "", "20A4", "&#x20A4;", "TL", false},
	"TTD": {"Trinidad and Tobago Dollar", 0, empty, empty, "", "0024", "&#x0024;", "TT$", true},
	"TWD": {"New Taiwan Dollar", 0, empty, empty, "", "0024", "&#x0024;", "NT$", true},
	"TZS": {"Tanzanian Shilling", 2, ",", ".", " ", empty, empty, "TZs", false},
	"UAH": {"Hryvnia", 2, " ", ",", "", "20B4", "&#x20B4", "UAH", false},
	"UGX": {"Uganda Shilling", 0, empty, empty, "", empty, empty, "UGX", true},
	"USD": {"US Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"UYU": {"Peso Uruguayo", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"UZS": {"Uzbekistan Sum", 0, empty, empty, "", empty, empty, "лв", true},
	"VEF": {"Bolivar", 2, ".", ",", " ", empty, empty, "Bs.", true},
	"VND": {"Dong", 2, ".", ",", " ", "20AB", "&#x20AB;", "₫", true},
	"VUV": {"Vatu", 0, ",", "", "", empty, empty, "VT", false},
	"WST": {"Tala", 0, empty, empty, "", empty, empty, "WST", true},
	"XAF": {"CFA Franc BEAC", 0, empty, empty, "", empty, empty, "$", true},
	"XCD": {"East Caribbean Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"XOF": {"CFA Franc BCEAO", 0, empty, empty, empty, empty, empty, "XOF", true},
	"XPF": {"CFP Franc", 0, empty, empty, "", empty, empty, "XPF", true},
	"YER": {"Yemeni Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", "﷼", true},
	"YUN": {"New Dinar", 0, empty, empty, "", empty, empty, "YUN", true},
	"ZAR": {"Rand", 2, " ", ".", " ", "0052", "&#x0052;", "R", true},
	"ZMK": {"Kwacha", 0, empty, empty, "", empty, empty, "ZMK", true},
	"ZRN": {"New Zaire", 0, empty, empty, empty, empty, empty, "ZRN", true},
	"ZWD": {"Zimbabwe Dollar ", 2, " ", ".", "", "0024", "&#x0024;", "Z$", true},
}