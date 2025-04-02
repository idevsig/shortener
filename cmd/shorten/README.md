# shorten

短网址 CLI 管理工具

## 使用帮助
### 安装
```bash
go install go.dsig.cn/shortener/cmd/shorten@latest
```

### 初始化
```bash
shorten init

# 初始化时同时设置 KEY 和 URL
shorten init --key KEY --url "http://127.0.0.1:8080"
```
创建配置文件 `$HOME/.config/shortener/config.toml`
```toml
# 密钥
key = ''
# 短网址服务器
url = 'http://127.0.0.1:8080'
```
亦可通过全局变量设置
```bash
SHORTENER_KEY=""
SHORTENER_URL="http://127.0.0.1:8080"
```
临时使用可通过传参方式
```bash
shorten --key KEY --url "http://127.0.0.1:8080"
```

### 更多帮助
```bash
shorten help

# 子命令帮助
shorten help list
```