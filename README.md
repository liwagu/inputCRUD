# InputCRUD

InputCRUD是一个简单的前后端分离项目，允许用户通过 Web 页面输入文本并将其存储到数据库中。项目使用 Go 语言编写后端，并使用原生 HTML 和 JavaScript 编写前端。

## 主要功能

- 用户可以在 Web 页面上输入文本。
- 用户可以提交文本，将其保存到数据库中。
- 页面会显示数据库中已保存的所有文本。

## 技术栈

- 后端：Go 语言
- 数据库：MySQL
- 前端：HTML, JavaScript

## 安装与运行

### 环境要求

- Go 1.15 或更高版本
- MySQL 5.7 或更高版本

### 步骤

1. 克隆此项目：

2. 使用 MySQL 创建数据库和表结构。可以参考项目中的 schema.sql 文件。
3. 修改 main.go 文件中的数据库连接配置（例如，用户名、密码和数据库名称）。
4. 在项目根目录下运行 Go 程序：

    过程见B站视频 https://www.bilibili.com/video/BV1nX4y1d7Z1/