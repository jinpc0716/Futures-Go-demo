﻿package config

// API KEY
const (
	// todo: replace with your own AccessKey and Secret Key
	ACCESS_KEY string = ""  // huobi申请的apiKey
	SECRET_KEY string = ""  // huobi申请的secretKey



	// API请求地址, 不要带最后的/
	MARKET_URL string = "https://api.hbdm.com"
	TRADE_URL  string = "https://api.hbdm.com"
	WS_URL  string = "wss://www.hbdm.com/ws"

	//replace with real URLs and HostName
	HOST_NAME  string = "api.hbdm.com"



	ENABLE_PRIVATE_SIGNATURE bool = false

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// replace with your own PrivateKey from privatekey.pem
	PRIVATE_KEY_PRIME_256 string = ``


)

