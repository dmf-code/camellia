# https://registry.hub.docker.com/_/php?tab=tags&page=1&ordering=last_updated&name=7
ARG PHP_VERSION
FROM php:${PHP_VERSION}

ARG TIME_ZONE

# 设置时区
ENV TZ=${TIME_ZONE}

ARG PHP_EXT

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN ${PHP_EXT}

ARG COMPOSER_PATH
# 安装 Composer
ENV COMPOSER_HOME ${COMPOSER_PATH}

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
ENV PATH $COMPOSER_HOME/vendor/bin:$PATH
