# Go-React-Docker

GolangとReactとDockerを学びたかったので作ったもの

## Feature
- アカウント作成，認証
- リアルタイムチャット
- ルーム作成（public,private）


## setup (dev)
1. `git clone {this repository} `
2. `docker-compose build`
3. `docker-compose up -d`
4. `docker-compose exec backend go run main/main.go`
5. `docker-compose exec frontend npm i`
6. `docker-compose exec frontend npm run dev`
7. Access to http://localhost:8000/

## setup (production)
### toto


## やりたいことリスト
- [x] 基本的な機能の実装
- [ ] リファクタリング
- [ ] セキュリティ関連(csrfとか)
- [ ] テスト
- [ ] 認証サーバーを別で立てたい
- [ ] CI/DI
- [ ] 別のAPIをつなぐ
- [ ] Nginx
- [ ] READMEを書く


