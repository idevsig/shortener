/*
Copyright © 2025 Jetsung Chan <i@jetsung.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"resty.dev/v3"
)

// Config 配置
type Config struct {
	APIURL string `mapstructure:"url"`
	APIKEY string `mapstructure:"key"`
}

// ResShorten 短链接响应
type ResShorten struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Describe    string `json:"describe"`
	Status      int8   `json:"status"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

// ResPage 分页响应
type ResPage struct {
	Page         int64 `json:"page"`          // 当前页码（从1开始）
	PageSize     int64 `json:"page_size"`     // 每页条数（可选返回，便于客户端验证）
	CurrentCount int64 `json:"current_count"` // 当前页实际条数
	TotalItems   int64 `json:"total_items"`   // 符合条件的总条数
	TotalPages   int64 `json:"total_pages"`   // 总页数
}

// ResSuccess 成功响应
type ResSuccess[T any] struct {
	Data T       `json:"data"` // 数据
	Meta ResPage `json:"meta"` // 元数据
}

// ResErr 错误响应
type ResErr struct {
	ErrCode int    `json:"errcode"`
	ErrInfo string `json:"errinfo"`
}

const (
	appName    = "shorten"
	configName = "config"
	configType = "toml"
	cfgDirName = "shortener"
)

var (
	cfg           Config
	configDir     string
	APIRequestURL = "/api/v1"
	APIShortenURL = "/shortens"
	rootCmd       = &cobra.Command{
		Use:           appName,
		Short:         "Short URL management CLI tool",
		SilenceUsage:  true,  // 隐藏错误时自动显示的用法帮助
		SilenceErrors: false, // 隐藏Cobra默认的错误前缀
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := initConfig(); err != nil {
				return err
			}

			if requiresAPIURL(cmd) {
				if cfg.APIURL == "" {
					return fmt.Errorf(`必须提供API地址，可用方式：
1. 命令行参数: --url
2. 环境变量: export SHORTENER_URL=your_url
3. 配置文件: 在 ~/.shortener/config.toml 添加 url`)
				}
			}

			return nil
		},
	}
)

func init() {
	viper.SetEnvPrefix("SHORTENER")
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringP("url", "u", "", "API URL")
	rootCmd.PersistentFlags().StringP("key", "k", "", "API KEY")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.AddCommand(newInitCmd())
	rootCmd.AddCommand(newShortenCreateCmd())
	rootCmd.AddCommand(newShortenDeleteCmd())
	rootCmd.AddCommand(newShortenUpdateCmd())
	rootCmd.AddCommand(newShortenGetCmd())
	rootCmd.AddCommand(newShortenListCmd())
}

func initConfig() error {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configDir = filepath.Join(userConfigDir, cfgDirName)
	log.Printf("configDir: %s", configDir)

	viper.SetConfigName(configName)
	viper.AddConfigPath(configDir)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	APIRequestURL = cfg.APIURL + APIRequestURL
	APIShortenURL = APIRequestURL + APIShortenURL

	log.Printf("cfg: %+v", cfg)
	return nil
}

// requiresAPIURL 判断命令是否需要API地址
func requiresAPIURL(cmd *cobra.Command) bool {
	requiredCmds := map[string]bool{
		"create": true,
		"delete": true,
		"update": true,
		"get":    true,
	}

	current := cmd
	for current != nil {
		if requiredCmds[current.Name()] {
			return true
		}
		current = current.Parent()
	}
	return false
}

// IsURL 判断是否为URL
func isURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := os.MkdirAll(configDir, 0700); err != nil {
				fmt.Printf("Create config directory failed: %s\n%v\n", configDir, err)
				return
			}

			apiURL := cmd.Flags().Lookup("url").Value.String()
			apiKey := cmd.Flags().Lookup("key").Value.String()

			if apiURL != "" {
				viper.Set("url", apiURL)
			}
			if apiKey != "" {
				viper.Set("key", apiKey)
			}

			configFile := filepath.Join(configDir, configName+"."+configType)
			if err := viper.WriteConfigAs(configFile); err != nil {
				fmt.Printf("Write config failed: %s\n%v\n", configFile, err)
				return
			}
		},
	}
}

func newShortenCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <origin_url>",
		Aliases: []string{"add"},
		Short:   "Create a short link",
		Args:    cobra.ExactArgs(1),
		Example: `  shorten create https://example.com/long/url
  shorten create https://example.com --code CUSTOM_CODE --desc "My special link"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("origin URL is required")
			}

			originURL := args[0]
			if !isURL(originURL) {
				return fmt.Errorf("invalid origin URL: %s", originURL)
			}

			customCode, _ := cmd.Flags().GetString("code")
			description, _ := cmd.Flags().GetString("desc")

			req := struct {
				Code        string `json:"code,omitempty"`
				OriginalURL string `json:"original_url" binding:"required"`
				Describe    string `json:"describe,omitempty"`
			}{
				Code:        customCode,
				OriginalURL: originURL,
				Describe:    description,
			}

			client := resty.New()
			defer client.Close()

			var response ResShorten
			var resErr ResErr

			res, err := client.R().
				SetContentType("application/json").
				SetBody(req).
				SetResult(&response).
				SetError(&resErr).
				Post(APIShortenURL)

			if err != nil {
				return fmt.Errorf("failed to create short URL: %w", err)
			}

			if res.StatusCode() != http.StatusCreated {
				return fmt.Errorf("failed to create short URL: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
					res.StatusCode(),
					resErr.ErrCode,
					resErr.ErrInfo)
			}

			if !isURL(response.OriginalURL) {
				return fmt.Errorf("invalid short url: %s", response.OriginalURL)
			}

			fmt.Printf("Created short Code: %s\n", response.Code)
			fmt.Printf("         Short URL: %s\n", response.ShortURL)
			fmt.Printf("      Original URL: %s\n", response.OriginalURL)
			fmt.Printf("       Description: %s\n", response.Describe)
			return nil
		},
	}

	cmd.Flags().StringP("code", "c", "", "Custom short code (optional)")
	cmd.Flags().StringP("desc", "d", "", "Link description (optional)")

	return cmd
}

func newShortenDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "delete <short_code>",
		Aliases: []string{"del"},
		Short:   "Delete a short code",
		Args:    cobra.ExactArgs(1),
		Example: `  shorten delete MySpecialCode`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("short code is required")
			}

			code := args[0]

			client := resty.New()
			defer client.Close()

			var resErr ResErr

			res, err := client.R().
				SetContentType("application/json").
				SetError(&resErr).
				Delete(APIShortenURL + "/" + code)

			if err != nil {
				return fmt.Errorf("failed to delete short URL: %w", err)
			}

			if res.StatusCode() != 204 {
				return fmt.Errorf("failed to delete short URL: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
					res.StatusCode(),
					resErr.ErrCode,
					resErr.ErrInfo)
			}

			fmt.Printf("Deleted short Code: %s\n", code)
			return nil
		},
	}
}

func newShortenUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <short_code>",
		Short: "Update a short code",
		Args:  cobra.ExactArgs(1),
		Example: `  shorten update MySpecialCode --ourl https://example.com
  shorten update MySpecialCode --ourl https://example.com --desc "My special link"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("short code is required")
			}

			code := args[0]

			originURL, _ := cmd.Flags().GetString("ourl")
			description, _ := cmd.Flags().GetString("desc")

			if originURL != "" && !isURL(originURL) {
				return fmt.Errorf("invalid origin URL: %s", originURL)
			}

			req := struct {
				OriginalURL string `json:"original_url,omitempty" binding:"omitempty,url"`
				Describe    string `json:"describe,omitempty"`
			}{
				OriginalURL: originURL,
				Describe:    description,
			}

			var response ResShorten
			var resErr ResErr

			client := resty.New()
			defer client.Close()

			res, err := client.R().
				SetContentType("application/json").
				SetBody(req).
				SetResult(&response).
				SetError(&resErr).
				Put(APIShortenURL + "/" + code)

			if err != nil {
				return fmt.Errorf("failed to update short code: %w", err)
			}

			if res.StatusCode() != http.StatusOK {
				return fmt.Errorf("failed to update short URL: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
					res.StatusCode(),
					resErr.ErrCode,
					resErr.ErrInfo)
			}

			if !isURL(response.OriginalURL) {
				return fmt.Errorf("invalid short code: %s", code)
			}

			fmt.Printf("Updated short Code: %s\n", response.Code)
			fmt.Printf("         Short URL: %s\n", response.ShortURL)
			fmt.Printf("      Original URL: %s\n", response.OriginalURL)
			fmt.Printf("       Description: %s\n", response.Describe)
			return nil
		},
	}

	cmd.Flags().StringP("ourl", "o", "", "Original URL (optional)")
	cmd.Flags().StringP("desc", "d", "", "Link description (optional)")

	return cmd
}

func newShortenGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "get <short_code>",
		Aliases: []string{"g"},
		Short:   "Get a short link",
		Args:    cobra.ExactArgs(1),
		Example: `  shorten get MySpecialCode`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("short code is required")
			}

			code := args[0]

			client := resty.New()
			defer client.Close()

			var response ResShorten
			var resErr ResErr

			res, err := client.R().
				SetContentType("application/json").
				SetResult(&response).
				SetError(&resErr).
				Get(APIShortenURL + "/" + code)

			if err != nil {
				return fmt.Errorf("failed to get short URL: %w", err)
			}

			if res.StatusCode() != http.StatusOK {
				return fmt.Errorf("failed to get short URL: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
					res.StatusCode(),
					resErr.ErrCode,
					resErr.ErrInfo)
			}

			if !isURL(response.OriginalURL) {
				return fmt.Errorf("invalid short code: %s", code)
			}

			fmt.Printf("Short Code: %s\n", response.Code)
			fmt.Printf("         Short URL: %s\n", response.ShortURL)
			fmt.Printf("      Original URL: %s\n", response.OriginalURL)
			fmt.Printf("       Description: %s\n", response.Describe)
			return nil
		},
	}
}

func newShortenListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List all short links",
		Example: `  shorten list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var isAll bool
			if len(args) > 0 && args[0] == "all" {
				isAll = true
			}

			client := resty.New()
			defer client.Close()

			var allData []ResShorten
			var response ResSuccess[[]ResShorten]
			var resErr ResErr

			// 处理获取所有数据的逻辑
			if isAll {
				page := int64(1)
				pageSize := int64(100) // 每页获取最大允许数量

				for {
					query := url.Values{}
					query.Set("page", strconv.FormatInt(page, 10))
					query.Set("page_size", strconv.FormatInt(pageSize, 10))
					query.Set("sort_by", "created_at")
					query.Set("order", "asc")

					res, err := client.R().
						SetContentType("application/json").
						SetResult(&response).
						SetError(&resErr).
						Get(APIShortenURL + "?" + query.Encode())

					if err != nil {
						return fmt.Errorf("failed to list short URLs: %w", err)
					}

					if res.StatusCode() != http.StatusOK {
						return fmt.Errorf("failed to list short URLs: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
							res.StatusCode(),
							resErr.ErrCode,
							resErr.ErrInfo)
					}

					// 合并数据
					allData = append(allData, response.Data...)

					// 检查是否还有更多数据
					if page >= response.Meta.TotalPages {
						break
					}
					page++
				}

				// 替换response为完整数据
				response.Data = allData
				response.Meta.Page = 1
				response.Meta.CurrentCount = int64(len(allData))
				response.Meta.TotalItems = int64(len(allData))
				response.Meta.TotalPages = 1
				response.Meta.PageSize = int64(len(allData))
			} else {
				// 原有分页查询逻辑
				page, _ := cmd.Flags().GetInt64("page")
				pageSize, _ := cmd.Flags().GetInt64("psize")
				sortBy, _ := cmd.Flags().GetString("sort")
				order, _ := cmd.Flags().GetString("order")

				// 设置默认值
				if page == 0 {
					page = 1
				}
				if pageSize == 0 {
					pageSize = 10
				}
				if sortBy == "" {
					sortBy = "created_at"
				}
				if order == "" {
					order = "desc"
				}

				query := url.Values{}
				query.Set("page", strconv.FormatInt(page, 10))
				query.Set("page_size", strconv.FormatInt(pageSize, 10))
				query.Set("sort_by", sortBy)
				query.Set("order", order)

				res, err := client.R().
					SetContentType("application/json").
					SetResult(&response).
					SetError(&resErr).
					Get(APIShortenURL + "?" + query.Encode())

				if err != nil {
					return fmt.Errorf("failed to list short URLs: %w", err)
				}

				if res.StatusCode() != http.StatusOK {
					return fmt.Errorf("failed to list short URLs: \n  status code: %d \n      errcode: %d \n      errinfo: %s",
						res.StatusCode(),
						resErr.ErrCode,
						resErr.ErrInfo)
				}
			}

			// 显示结果
			if len(response.Data) == 0 {
				fmt.Println("No short URLs found")
				return nil
			}

			for _, item := range response.Data {
				fmt.Printf("  Short Code: %s\n", item.Code)
				fmt.Printf("   Short URL: %s\n", item.ShortURL)
				fmt.Printf("Original URL: %s\n", item.OriginalURL)
				if item.Describe != "" {
					fmt.Printf(" Description: %s\n", item.Describe)
				}
				fmt.Println("--------------------------------")
			}

			fmt.Printf("  Total Items: %d\n", response.Meta.TotalItems)
			if !isAll {
				fmt.Printf("  Total Pages: %d\n", response.Meta.TotalPages)
				fmt.Printf("    Page Size: %d\n", response.Meta.PageSize)
				fmt.Printf(" Current Page: %d\n", response.Meta.Page)
				fmt.Printf("Current Count: %d\n", response.Meta.CurrentCount)
			}

			return nil
		},
	}

	cmd.Flags().Int64P("page", "p", 1, "Page number")
	cmd.Flags().Int64P("psize", "z", 10, "Page size")
	cmd.Flags().StringP("sort", "s", "created_at", "Sort by field")
	cmd.Flags().StringP("order", "o", "desc", "Sort order")

	return cmd
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
