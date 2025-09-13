# web auth token
事前にdocker compose upしてシステムを起動したうえで3000番ポートを使用しているwebコンテナは停止する。
これはGoogle Auth Platformにlocalhost:3000/auth/callbackをリダイレクト先に設定しているため。
このディレクトリに移動してから以下を実行
```sh
npm install

export GOOGLE_CLIENT_ID=your_client_id
export GOOGLE_CLIENT_SECRET=your_client_secret

node index.js
```
