# ormsample

# localhostでMysql を立ち上げ（docker）
docker-compose up -d

# genのインストール
go install github.com/smallnest/gen@latest

# genでDBから構造体を作成（試行錯誤中）
gen --connstr="rescuenow:rescuenow119@tcp(localhost:3306)/anpidb" --out generated-local --sqltype mysql --database anpidb --gorm --no-json --overwrite --proto-fmt=snake --model_naming={{toSnakeCase}}
