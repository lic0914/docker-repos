FROM python:3.14.0a7-slim-bookworm AS base 
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    rm -rf /var/lib/apt/lists/*

RUN pip install uvx
RUN uvx --version

# 添加 NodeSource 源（Node.js 18）
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -

# 安装 Node.js 和 npm
RUN apt-get update && \
    apt-get install -y --no-install-recommends nodejs && \
    rm -rf /var/lib/apt/lists/*

# 验证安装
RUN node --version && npm --version
RUN corepack enable # 开启yarn

RUN yarn --verison

# 1. 安装必要依赖（仅临时使用）
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        gnupg

# 2. 添加 Docker 官方 GPG 密钥和 APT 源（Bookworm）
RUN mkdir -p /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
    chmod a+r /etc/apt/keyrings/docker.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian bookworm stable" > /etc/apt/sources.list.d/docker.list

# 3. 安装 docker-ce-cli 并清理临时依赖
RUN apt-get update && \
    apt-get install -y --no-install-recommends docker-ce-cli && \
    apt-get purge -y curl gnupg && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*

# 4. 验证
RUN docker --version