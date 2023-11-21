package conf

type (
	Me struct {
		Ip string `json:"ip"`
		Port string `json:"port"`
		UploadIpDomain string `json:"upload_ip_domain"`
		LogInfo string `json:"log_info"`
		LogError string `json:"log_error"`
		LogWarn string `json:"log_warn"`
	}

	Elasticsearch struct {
		Ip string `json:"ip"`
		Port string `json:"port"`
		Ssl string `json:"ssl"`
	}

	Mongodb struct {
		Ip string `json:"ip"`
		Port string `json:"port"`
		DbName string `json:"db_name"`
		DbUser string `json:"db_user"`
		DbPaaaword string `json:"db_password"`
	}

	Config struct {
		Me Me `json:"me"`
		Elasticsearch Elasticsearch `json:"elasticsearch"`
		Mongodb Mongodb `json:"mongodb"`
	}
)
