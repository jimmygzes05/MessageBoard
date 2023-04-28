# 簡易留言板

## 環境建置

### Step.1

git clone下來後，docker使用以下指令(需先安裝docker)

```bash
# 建立mysql container，包含phpmyadmin
# container建好，連到127.0.0.1:80，應該可以進入phpmyadmin頁面
# 帳號root，密碼admin
docker-compose -f docker/db.yaml up
```

### Step.2

再來建立主要程式的Container，使用以下指令

```bash
# 建立主要程式的container，可以不用安裝Golang，但如果想要修改程式還是建議安裝
# container建好，連到127.0.0.1:1234/web/login，應該可以看到首頁
ddocker-compose -f board.yaml up
```
