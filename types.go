// Code generated by generate_types.go; DO NOT EDIT.

package greenpass

import (
	"fmt"
)

type DiseaseAgentTargeted string

func (dat DiseaseAgentTargeted) String() string {
	name, ok := diseaseAgentTargetedNames[dat]
	if !ok {
		return fmt.Sprintf("unknown('%s')", dat)
	}
	return name
}

var diseaseAgentTargetedNames = map[DiseaseAgentTargeted]string{
	"840539006": "COVID-19",
}

type Country2Codes string

func (cc Country2Codes) String() string {
	name, ok := country2CodesNames[cc]
	if !ok {
		return fmt.Sprintf("unknown('%s')", cc)
	}
	return name
}

var country2CodesNames = map[Country2Codes]string{
	"DZ": "Algeria",
	"FJ": "Fiji",
	"HT": "Haiti",
	"NG": "Nigeria",
	"CN": "China",
	"JE": "Jersey",
	"LK": "Sri Lanka",
	"ZM": "Zambia",
	"GY": "Guyana",
	"BM": "Bermuda",
	"BO": "Bolivia, Plurinational State of",
	"MU": "Mauritius",
	"BH": "Bahrain",
	"RS": "Serbia",
	"CF": "Central African Republic",
	"BG": "Bulgaria",
	"LT": "Lithuania",
	"FM": "Micronesia, Federated States of",
	"GA": "Gabon",
	"GB": "United Kingdom of Great Britain and Northern Ireland",
	"GW": "Guinea-Bissau",
	"MP": "Northern Mariana Islands",
	"MV": "Maldives",
	"PK": "Pakistan",
	"BR": "Brazil",
	"WF": "Wallis and Futuna",
	"KG": "Kyrgyzstan",
	"PM": "Saint Pierre and Miquelon",
	"SA": "Saudi Arabia",
	"ST": "Sao Tome and Principe",
	"BF": "Burkina Faso",
	"IT": "Italy",
	"TT": "Trinidad and Tobago",
	"IO": "British Indian Ocean Territory",
	"GF": "French Guiana",
	"PE": "Peru",
	"SM": "San Marino",
	"UZ": "Uzbekistan",
	"CV": "Cabo Verde",
	"BW": "Botswana",
	"FK": "Falkland Islands (Malvinas)",
	"FO": "Faroe Islands",
	"IR": "Iran, Islamic Republic of",
	"NF": "Norfolk Island",
	"AQ": "Antarctica",
	"DM": "Dominica",
	"HR": "Croatia",
	"RW": "Rwanda",
	"AM": "Armenia",
	"BN": "Brunei Darussalam",
	"BL": "Saint Barthélemy",
	"DK": "Denmark",
	"RE": "Réunion",
	"CD": "Congo, the Democratic Republic of the",
	"SI": "Slovenia",
	"VC": "Saint Vincent and the Grenadines",
	"NO": "Norway",
	"LR": "Liberia",
	"SD": "Sudan",
	"IQ": "Iraq",
	"GM": "Gambia",
	"MM": "Myanmar",
	"SE": "Sweden",
	"TW": "Taiwan, Province of China",
	"GI": "Gibraltar",
	"MA": "Morocco",
	"QA": "Qatar",
	"HN": "Honduras",
	"GE": "Georgia",
	"TN": "Tunisia",
	"CW": "Curaçao",
	"MY": "Malaysia",
	"EH": "Western Sahara",
	"CU": "Cuba",
	"EC": "Ecuador",
	"GD": "Grenada",
	"TF": "French Southern Territories",
	"TL": "Timor-Leste",
	"VA": "Holy See",
	"BY": "Belarus",
	"UG": "Uganda",
	"MW": "Malawi",
	"CR": "Costa Rica",
	"CY": "Cyprus",
	"LI": "Liechtenstein",
	"BD": "Bangladesh",
	"JO": "Jordan",
	"SO": "Somalia",
	"UM": "United States Minor Outlying Islands",
	"CA": "Canada",
	"ER": "Eritrea",
	"ET": "Ethiopia",
	"AD": "Andorra",
	"OM": "Oman",
	"LY": "Libya",
	"VI": "Virgin Islands,",
	"PG": "Papua New Guinea",
	"UA": "Ukraine",
	"MF": "Saint Martin (French part)",
	"NA": "Namibia",
	"PR": "Puerto Rico",
	"LC": "Saint Lucia",
	"GT": "Guatemala",
	"MT": "Malta",
	"SX": "Sint Maarten (Dutch part)",
	"ZA": "South Africa",
	"AE": "United Arab Emirates",
	"DE": "Germany",
	"YT": "Mayotte",
	"CH": "Switzerland",
	"GH": "Ghana",
	"IE": "Ireland",
	"NI": "Nicaragua",
	"PL": "Poland",
	"SS": "South Sudan",
	"CI": "Côte d''Ivoire",
	"CC": "Cocos (Keeling) Islands",
	"IM": "Isle of Man",
	"NL": "Netherlands",
	"PA": "Panama",
	"AZ": "Azerbaijan",
	"LS": "Lesotho",
	"CX": "Christmas Island",
	"EE": "Estonia",
	"ID": "Indonesia",
	"BV": "Bouvet Island",
	"MD": "Moldova, Republic of",
	"MK": "Macedonia, the former Yugoslav Republic of",
	"ML": "Mali",
	"MX": "Mexico",
	"SG": "Singapore",
	"SL": "Sierra Leone",
	"HM": "Heard Island and McDonald Islands",
	"DO": "Dominican Republic",
	"GN": "Guinea",
	"KP": "Korea, Democratic People''s Republic of",
	"WS": "Samoa",
	"AU": "Australia",
	"JM": "Jamaica",
	"NP": "Nepal",
	"SK": "Slovakia",
	"TK": "Tokelau",
	"CL": "Chile",
	"SB": "Solomon Islands",
	"VU": "Vanuatu",
	"PW": "Palau",
	"PN": "Pitcairn",
	"SC": "Seychelles",
	"AL": "Albania",
	"PH": "Philippines",
	"AT": "Austria",
	"US": "United States of America",
	"ZW": "Zimbabwe",
	"SZ": "Swaziland",
	"MH": "Marshall Islands",
	"MO": "Macao",
	"MR": "Mauritania",
	"KW": "Kuwait",
	"MS": "Montserrat",
	"NE": "Niger",
	"UY": "Uruguay",
	"ES": "Spain",
	"KY": "Cayman Islands",
	"LU": "Luxembourg",
	"IN": "India",
	"GS": "South Georgia and the South Sandwich Islands",
	"LB": "Lebanon",
	"MN": "Mongolia",
	"AG": "Antigua and Barbuda",
	"FR": "France",
	"GP": "Guadeloupe",
	"KI": "Kiribati",
	"SR": "Suriname",
	"TR": "Turkey",
	"TV": "Tuvalu",
	"TZ": "Tanzania, United Republic of",
	"BA": "Bosnia and Herzegovina",
	"BQ": "Bonaire, Sint Eustatius and Saba",
	"CG": "Congo",
	"MG": "Madagascar",
	"PY": "Paraguay",
	"RO": "Romania",
	"SV": "El Salvador",
	"TM": "Turkmenistan",
	"AW": "Aruba",
	"CO": "Colombia",
	"JP": "Japan",
	"CM": "Cameroon",
	"ME": "Montenegro",
	"SN": "Senegal",
	"TD": "Chad",
	"KE": "Kenya",
	"BI": "Burundi",
	"IL": "Israel",
	"MC": "Monaco",
	"PS": "Palestine, State of",
	"TC": "Turks and Caicos Islands",
	"AF": "Afghanistan",
	"MQ": "Martinique",
	"RU": "Russian Federation",
	"TH": "Thailand",
	"GQ": "Equatorial Guinea",
	"KR": "Korea, Republic of",
	"LV": "Latvia",
	"NR": "Nauru",
	"TJ": "Tajikistan",
	"YE": "Yemen",
	"CK": "Cook Islands",
	"BZ": "Belize",
	"HK": "Hong Kong",
	"TG": "Togo",
	"AR": "Argentina",
	"IS": "Iceland",
	"BJ": "Benin",
	"GU": "Guam",
	"NC": "New Caledonia",
	"VE": "Venezuela, Bolivarian Republic of",
	"EG": "Egypt",
	"GL": "Greenland",
	"AO": "Angola",
	"AS": "American Samoa",
	"BB": "Barbados",
	"CZ": "Czechia",
	"FI": "Finland",
	"KN": "Saint Kitts and Nevis",
	"LA": "Lao People''s Democratic Republic",
	"NU": "Niue",
	"SY": "Syrian Arab Republic",
	"AI": "Anguilla",
	"NZ": "New Zealand",
	"VN": "Viet Nam",
	"KZ": "Kazakhstan",
	"TO": "Tonga",
	"AX": "Åland Islands",
	"SH": "Saint Helena, Ascension and Tristan da Cunha",
	"SJ": "Svalbard and Jan Mayen",
	"HU": "Hungary",
	"MZ": "Mozambique",
	"PF": "French Polynesia",
	"PT": "Portugal",
	"VG": "Virgin Islands, British",
	"BT": "Bhutan",
	"BS": "Bahamas",
	"DJ": "Djibouti",
	"GG": "Guernsey",
	"GR": "Greece",
	"KH": "Cambodia",
	"KM": "Comoros",
	"BE": "Belgium",
}
