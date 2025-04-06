# 短网址

一个超简单的短网址管理平台。

**配置前端：[shortener-frontend](https://git.jetsung.com/idev/shortener-frontend)**

## [Docker](./deploy/docker/README.md)

> **版本：** `latest`, `main`, <`TAG`>

| Registry                                                                                   | Image                                                  |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------ |
| [**Docker Hub**](https://hub.docker.com/r/idevsig/filetas/)                                | `idevsig/shortener`                                    |
| [**GitHub Container Registry**](https://github.com/idevsig/filetas/pkgs/container/filetas) | `ghcr.io/idevsig/shortener`                            |
| **Tencent Cloud Container Registry**                                                       | `ccr.ccs.tencentyun.com/idevsig/shortener`             |
| **Aliyun Container Registry**                                                              | `registry.cn-guangzhou.aliyuncs.com/idevsig/shortener` |

## 开发

### 1. 拉取代码
```bash
git clone https://git.jetsung.com/idev/shortener.git
cd shortener
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

- https://git.jetsung.com/idev/shortener
- https://framagit.org/idev/shortener
- https://gitcode.com/idev/shortener
- https://github.com/idevsig/shortener
