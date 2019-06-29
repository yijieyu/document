FROM alpine:3.4

MAINTAINER tvblack <github@tvblack.com>

ENV PATH /usr/local/bin:$PATH

ENV GPG_KEYS A917B1ECDA84AEC2B568FED6F50ABC807BD5DCD0 528995BFEDFBA7191D46839EF9BA0ADA31CBD89E
ENV PHP_VERSION 7.1.8
ENV PHP_URL="https://secure.php.net/get/php-7.1.8.tar.xz/from/this/mirror"
ENV PHP_ASC_URL="https://secure.php.net/get/php-7.1.8.tar.xz.asc/from/this/mirror"
ENV PHP_SHA256="8943858738604acb33ecedb865d6c4051eeffe4e2d06f3a3c8f794daccaa2aab"

ENV PHP_INI_DIR /usr/local/etc

ENV PHPREDIS_VERSION 3.1.3
ENV PHPREDIS_SRC_URL = "https://github.com/phpredis/phpredis/archive/3.1.3.tar.gz"

# Apply stack smash protection to functions using local buffers and alloca()
# Make PHP's main executable position-independent (improves ASLR security mechanism, and has no performance impact on x86_64)
# Enable optimization (-O2)
# Enable linker optimization (this sorts the hash buckets to improve cache locality, and is non-default)
# Adds GNU HASH segments to generated executables (this is used if present, and is much faster than sysv hash; in this configuration, sysv hash is also generated)
# https://github.com/docker-library/php/issues/272
ENV PHP_CFLAGS="-fstack-protector-strong -fpic -fpie -O2"
ENV PHP_CPPFLAGS="$PHP_CFLAGS"
ENV PHP_LDFLAGS="-Wl,-O1 -Wl,--hash-style=both -pie"

# 创建用户
RUN set -xe \
    && addgroup -S www \
    && adduser -D -S -G www www

RUN set -xe \
    && apk add --no-cache --virtual .persistent-deps \
        ca-certificates \
        curl \
        tar \
        xz

# 下载php源码
RUN set -xe \
    && apk add --no-cache --virtual .fetch-deps \
        gnupg \
        openssl \
        \
    && mkdir -p /usr/src \
    && cd /usr/src \
           \
    && wget -O php.tar.xz "$PHP_URL" \
    && if [ -n "$PHP_SHA256" ]; then \
            echo "$PHP_SHA256 *php.tar.xz" | sha256sum -c -; \
        fi \
    && if [ -n "$PHP_MD5" ]; then \
               echo "$PHP_MD5 *php.tar.xz" | md5sum -c -; \
           fi \
    && if [ -n "$PHP_ASC_URL" ]; then \
               wget -O php.tar.xz.asc "$PHP_ASC_URL"; \
               export GNUPGHOME="$(mktemp -d)"; \
               for key in $GPG_KEYS; do \
                   gpg --keyserver ha.pool.sks-keyservers.net --recv-keys "$key"; \
               done; \
               gpg --batch --verify php.tar.xz.asc php.tar.xz; \
               rm -rf "$GNUPGHOME"; \
           fi \
    && apk del .fetch-deps

RUN set -xe \
    && apk add --no-cache --virtual .build-deps \
       autoconf \
       dpkg-dev dpkg \
       file \
       g++ \
       gcc \
       libc-dev \
       make \
       pcre-dev \
       pkgconf \
       re2c \
       coreutils \
       curl-dev \
       libedit-dev \
       libxml2-dev \
       libmcrypt-dev \
       openssl-dev \
       sqlite-dev \
       curl-dev \
       jpeg-dev \
       libpng-dev \
       libxpm-dev \
       libwebp-dev \
       freetype-dev \
       gettext-dev \
       imap-dev \
       bzip2-dev \
       libxslt-dev \
    \
    && mkdir -p $PHP_INI_DIR/conf.d  \
    && mkdir -p /usr/src/php \
    && tar -Jxf /usr/src/php.tar.xz -C /usr/src/php --strip-components=1\
    && export CFLAGS="$PHP_CFLAGS" \
        CPPFLAGS="$PHP_CPPFLAGS" \
        LDFLAGS="$PHP_LDFLAGS" \
    && cd /usr/src/php \
    && gnuArch="$(dpkg-architecture --query DEB_BUILD_GNU_TYPE)" \
    && ./configure \
        --build="$gnuArch" \
        --with-config-file-path="$PHP_INI_DIR" \
        --with-config-file-scan-dir="$PHP_INI_DIR/conf.d" \
        --with-libxml-dir=/usr/local/ \
        --with-zlib-dir=/usr/local/ \
        --with-jpeg-dir=/usr/local/ \
        --with-png-dir=/usr/local/ \
        --with-freetype-dir=/usr/local/ \
        --with-webp-dir \
        --with-mhash=/usr/local/ \
        --with-mcrypt=/usr/local/ \
        --with-iconv-dir=/usr/local/ \
        --with-curl \
        --with-gd  \
        --enable-gd-native-ttf \
        --with-openssl \
        --with-zlib \
        --with-bz2 \
        --with-gettext \
        --with-xmlrpc  \
        --with-xsl \
        --disable-cgi \
        --enable-bcmath \
        --enable-shmop \
        --enable-sysvsem \
        --enable-inline-optimization \
        --enable-mbregex \
        --enable-mbstring \
        --enable-xml  \
        --enable-sockets \
        --enable-ftp  \
        --enable-soap  \
        --enable-zip \
        --with-xpm-dir \
        --enable-exif \
        --enable-opcache \
        --enable-mysqlnd \
        --with-pdo-mysql=mysqlnd \
        --with-mysqli=mysqlnd \
        --with-openssl \
        --with-libedit \
        --with-pcre-regex \
        --enable-fpm \
        --with-fpm-user=www \
        --with-fpm-group=www \
    \
    && make -j "$(nproc)" \
    && make install \
    && { find /usr/local/bin /usr/local/sbin -type f -perm +0111 -exec strip --strip-all '{}' + || true; } \
    && make clean \
    && cd / \
    && rm -rf /usr/src/php \
    \
    && runDeps="$( \
        scanelf --needed --nobanner --recursive /usr/local \
            | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
            | sort -u \
            | xargs -r apk info --installed \
            | sort -u \
    )" \
    && apk add --no-cache --virtual .php-rundeps $runDeps \
    \
    && apk del .build-deps \
    \
    && pecl update-channels \
    && rm -rf /tmp/pear ~/.pearrc

RUN set -ex \
    && apk add --no-cache --virtual .fetch-deps \
        gnupg \
        openssl \
        \
    && wget -c https://github.com/phpredis/phpredis/archive/4.0.0.tar.gz \
    && tar -zxf phpredis-4.0.0.tar.gz \
    && rm -rf phpredis-4.0.0.tar.gz \
    && cd phpredis-4.0.0 \
    && /usr/local/bin/phpize \
    && ./configure --with-php-config=/usr/local/php7/bin/php-config \
    && make \
    && make install \
    && rm -rf phpredis-4.0.0 \
    && apk del .fetch-deps

RUN set -ex \
    && cd /usr/local/etc \
    && if [ -d php-fpm.d ]; then \
        # for some reason, upstream's php-fpm.conf.default has "include=NONE/etc/php-fpm.d/*.conf"
        sed 's!=NONE/!=!g' php-fpm.conf.default | tee php-fpm.conf > /dev/null; \
        cp php-fpm.d/www.conf.default php-fpm.d/www.conf; \
    else \
        # PHP 5.x doesn't use "include=" by default, so we'll create our own simple config that mimics PHP 7+ for consistency
        mkdir php-fpm.d; \
        cp php-fpm.conf.default php-fpm.d/www.conf; \
        { \
            echo '[global]'; \
            echo 'include=etc/php-fpm.d/*.conf'; \
        } | tee php-fpm.conf; \
    fi \
    && { \
        echo '[global]'; \
        echo 'error_log = /proc/self/fd/2'; \
        echo; \
        echo '[www]'; \
        echo '; if we send this to /proc/self/fd/1, it never appears'; \
        echo 'access.log = /proc/self/fd/2'; \
        echo; \
        echo 'clear_env = no'; \
        echo; \
        echo '; Ensure worker stdout and stderr are sent to the main error log.'; \
        echo 'catch_workers_output = yes'; \
    } | tee php-fpm.d/docker.conf \
    && { \
        echo '[global]'; \
        echo 'daemonize = no'; \
        echo; \
        echo '[www]'; \
        echo 'listen = [::]:9000'; \
    } | tee php-fpm.d/zz-docker.conf


WORKDIR /var/www

EXPOSE 9000
CMD ["php-fpm"]
