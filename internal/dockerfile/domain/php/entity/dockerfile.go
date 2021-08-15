package entity

type Dockerfile struct {
	Args map[string]*string
}

func New() *Dockerfile {
	phpVersion := "7-fpm-alpine"
	timeZone := "Asia/Shanghai"
	composerPath := "/root/composer"
	phpExt := `apk update && apk add --no-cache vim git \
			&& docker-php-ext-install pdo_mysql \
			&& docker-php-ext-install mysqli \
			&& docker-php-ext-install bcmath `
	return &Dockerfile{
		Args: map[string]*string{
			"PHP_VERSION": &phpVersion,
			"TIME_ZONE": &timeZone,
			"COMPOSER_PATH": &composerPath,
			"PHP_EXT": &phpExt,
		},
	}
}
