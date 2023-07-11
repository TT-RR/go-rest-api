# go-rest-api
VSCodeのoriginはこのリポジトリのURL

# 起動手順 
1.docker-composeの起動

```
docker compose up -d
```

2.起動しているか確認

```
docker ps
```

3.main.goの実行

```
go run main.go -GO_ENV=dev
```

# Postmanでの使い方

```: URL
http://localhost:8080/メソッド名
```

### User
1.signup
```: 例
{
    "email":"user1@test.com",
    "password":"pass1234"
}
```
2.csrf(GET)

-> CSRFトークンが付与される

3.login(POST)

    HeadersのKeyにX-CSRF-TOKENと入力し、Valueに付与されたCSRFトークンを貼り付ける
    
    signupで使ったemailとパスワードでログイン
    -> jwtのトークンが付与される

4.logout(POST)

### Task
1.GetAllTasks tasks(GET)
    []が返ってくる

2.GetTaskByID tasks/調べたいuserId(GET)

3.CreateTask tasks(POST)
    作りたいタスクのタイトルを入力

    ```
    {
        "title":"test"
    }
    ```

    -> id,title,create_at,update_atが返ってくる

4.UpdateTask tasks/変えたいタスクのID(PUT)
    変えたいタスクのタイトルを入力

    ```
    {
        "title":"test5"
    }
    ```

5.DeleteTask tasks/消したいタスクのID(DELETE)
