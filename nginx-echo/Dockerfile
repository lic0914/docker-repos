FROM debian:bullseye-slim AS base
RUN apt update && apt install -y ca-certificates curl telnet && rm -rf /var/lib/apt/lists/*
LABEL maintainer="leechao lic0914@163.com"
ENV NGINX_VERSION   1.28.0
WORKDIR /etc/nginx
EXPOSE 80
EXPOSE 443

FROM ubuntu:20.04 AS nginx-build
WORKDIR /src
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y build-essential libpcre3 libpcre3-dev zlib1g zlib1g-dev libssl-dev libgd-dev libxml2 libxml2-dev uuid-dev git
ADD nginx-1.28.0.tar.gz ./
RUN ls -lh

WORKDIR "/src/nginx-1.28.0"

RUN git clone https://github.com/chobits/ngx_http_proxy_connect_module.git
RUN patch -p1 < ngx_http_proxy_connect_module/patch/proxy_connect_rewrite_102101.patch

RUN git clone https://github.com/openresty/echo-nginx-module.git
RUN ./configure \
--add-dynamic-module=echo-nginx-module \
--add-dynamic-module=ngx_http_proxy_connect_module \
--with-compat \
--prefix=/etc/nginx \
--sbin-path=/usr/sbin/nginx \
--modules-path=/usr/lib/nginx/modules \
--conf-path=/etc/nginx/nginx.conf \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--pid-path=/var/run/nginx.pid \
--lock-path=/var/run/nginx.lock \
--with-debug \
--with-file-aio \
--with-threads \
--with-http_addition_module \
--with-http_auth_request_module \
--with-http_dav_module \
--with-http_flv_module \
--with-http_gunzip_module \
--with-http_gzip_static_module \
--with-http_mp4_module \
--with-http_random_index_module \
--with-http_realip_module \
--with-http_secure_link_module \
--with-http_slice_module \
--with-http_ssl_module \
--with-http_stub_status_module \
--with-http_sub_module \
--with-http_v2_module \
--with-mail \
--with-mail_ssl_module \
--with-stream \
--with-stream_realip_module \
--with-stream_ssl_module \
--with-stream_ssl_preread_module 

RUN make -j2 && make install 
RUN ls -lh objs/

FROM base AS final
COPY --from=nginx-build /etc/nginx /etc/nginx
COPY --from=nginx-build /usr/sbin/nginx /usr/sbin/nginx
COPY --from=nginx-build /usr/lib/nginx/modules /etc/nginx/modules
COPY --from=nginx-build /src/nginx-1.28.0/objs/*.so /etc/nginx/modules/
COPY --from=nginx-build /var/log/nginx /var/log/nginx
COPY config/ ./

RUN ls -lh /etc/nginx/modules/

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log 
STOPSIGNAL SIGQUIT
CMD ["nginx", "-g", "daemon off;"]
