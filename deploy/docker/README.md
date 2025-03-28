# Docker 部署

## SQLite
- [sqlite.compose.yml](sqlite.compose.yml)
```yaml
services:
  shortener:
    image: idevsig/shortener:latest
    container_name: shortener
    restart: unless-stopped        
    environment:
      - TZ=Asia/Shanghai
      - GIN_MODE=release
      - DATABASE_TYPE=sqlite
    ports:
      - "8080:8080"
    volumes:
      - ./config.toml:/app/config.toml      
```      

## PostgreSQL
- [postgres.compose.yml](postgres.compose.yml)
```yaml
services:
  shortener:
    image: idevsig/shortener:latest
    container_name: shortener
    restart: unless-stopped        
    environment:
      - TZ=Asia/Shanghai
      - GIN_MODE=release
      - DATABASE_TYPE=postgres
    ports:
      - "8080:8080"
    volumes:
      - ./config.toml:/app/config.toml      
    depends_on:
      - postgres      

  postgres:
    image: postgres:latest
    container_name: postgres
    restart: unless-stopped
    environment:
      - TZ=Asia/Shanghai
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=shortener      
```

## MySQL
- [mysql.compose.yml](mysql.compose.yml)
```yaml
services:
  shortener:
    image: idevsig/shortener:latest
    container_name: shortener
    restart: unless-stopped
    environment:
      - TZ=Asia/Shanghai
      - GIN_MODE=release
      - DATABASE_TYPE=mysql
    ports:
      - "8080:8080"
    volumes:
      - ./config.toml:/app/config.toml      
    depends_on:
      - mysql      
      
  mysql:
    image: mysql:latest
    container_name: mysql
    restart: unless-stopped
    environment:
      - TZ=Asia/Shanghai
      - MYSQL_USER=root
      - MYSQL_PASSWORD=root
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=shortener
```
