# 短网址

一个超简单的短网址管理平台。

**配置前端：[shortener-frontend](https://git.jetsung.com/idev/shortener-frontend)**   
**命令行工具：[shortener](./cmd/shortener/README.md)**   

## 命令行
```bash
go install go.dsig.cn/shortener/cmd/shortener@latest
```

## [Docker](./deploy/docker/README.md)

> **版本：** `latest`, `main`, <`TAG`>

| Registry                                                                                   | Image                                                  |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------ |
| [**Docker Hub**](https://hub.docker.com/r/idevsig/shortener-server/)                                | `idevsig/shortener-server`                                    |
| [**GitHub Container Registry**](https://github.com/idevsig/shortener-server/pkgs/container/shortener-server) | `ghcr.io/idevsig/shortener-server`                            |
| **Tencent Cloud Container Registry**                                                       | `ccr.ccs.tencentyun.com/idevsig/shortener-server`             |
| **Aliyun Container Registry**                                                              | `registry.cn-guangzhou.aliyuncs.com/idevsig/shortener-server` |

## 开发

### 1. 拉取代码
```bash
git clone https://git.jetsung.com/idev/shortener-server.git
cd shortener-server
```

### 2. 修改配置
```bash
mkdir -p config/dev
cp config/config.toml config/dev/

# 修改开发环境的配置文件
vi config/dev/config.toml
```

### 3. 运行
```bash
go run .
```

### 4. 构建
```bash
go build

# 支持 GoReleaser 方式构建
goreleaser release --snapshot --clean
```

### 更多功能
```bash
just --list
```

## 部署
### Docker
```yaml
---
# https://github.com/idevsig/shortener-server

services:
  shortener:
    image: ghcr.io/idevsig/shortener-server:dev-amd64
    container_name: shortener
    restart: unless-stopped
    ports:
      - ${BACKEND_PORT:-8080}:8080
    volumes:
      - ./data:/app/data
      - ./config.toml:/app/config.toml
    depends_on:
      - valkey

  valkey:
    image: valkey/valkey:latest
    restart: unless-stopped
    environment:
      - TZ=Asia/Shanghai

  frontend:
    image: ghcr.io/idevsig/shortener-frontend:dev-amd64
    restart: unless-stopped
    ports:
      - ${FRONTEND_PORT:-8081}:80
```

## TODO

- [x] 实现全部功能接口
  - [x] `API` 权限校验
- [x] 支持数据库
  - [x] SQLite
  - [x] PostgreSQL
  - [x] MySQL
- [x] 支持缓存
  - [x] Redis
  - [x] Valkey
- [x] 制作 CLI 工具
  - [x] 添加 OpenAPI
- [x] 添加跳转请求日志记录
- [x] `CI/CD` 构建
  - [x] Docker 镜像构建与推送
- [x] 实现管理平台接口
- [x] 添加文档
- [ ] 添加测试

## 仓库镜像

- https://git.jetsung.com/idev/shortener-server
- https://framagit.org/idev/shortener-server
- https://gitcode.com/idev/shortener-server
- https://github.com/idevsig/shortener-server
