package apiserver

// ============================================================================
/*

curl -XGET http://localhost:6666/v1/user/userID-roo6vn

curl -XPOST -H'Content-Type: application/json' http://localhost:6666/api/v1/user -d'{"username": "Tom", "password": "123456789", "nickname": "shabi", "email": "123456@qq.com", "phone": "18012345678"}'

- token

install jq
sudo apt-get update && sudo apt-get install jq

token=$(curl -s -XPOST -H'Content-Type: application/json' http://localhost:6666/login -d'{"username": "Tom", "password": "123456789"}' | jq -r '.token')



curl -XPOST -H'Content-Type: application/json' -H"Authorization: Bearer ${token}" http://localhost:6666/api/v1/post -d'{"title": "文章Title", "content": "这是一个文章"}'

*/
// ============================================================================
