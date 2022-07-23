package config

const DB_CONNECTION = "root:password@(blogs_api_db)/blogs_api?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s&parseTime=true"

const JWT_SECRET = "ibelieveinmiracles"
const TOKEN_TIME = 100000 // Token time validity in minutes
