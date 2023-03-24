# GPT 聊天機器人

GPT 聊天機器人是一個基於 GPT-3 神經網路的聊天機器人應用程式，可以根據您的輸入，自動回應相關的訊息。

## 安裝

請確保您擁有以下軟件：

- Go (version >= 1.13)
- GPT-3 API KEY

接著，請按照以下步驟進行安裝：

1. 下載程式碼：

```
git clone https://github.com/your-repository.git
```

2. 填寫 `.env` 文件

根據您的 GPT-3 API KEY，請填寫 `.env` 文件，如下:

```
API_KEY=<your_API_KEY>
```

3. 安裝運行所需套件：

```
go get -u github.com/PullRequestInc/go-gpt3
go get -u github.com/gin-gonic/gin
```

## 運行

在完成安裝後，請按照以下步驟運行：

```
go run .
```

接著，您可以在瀏覽器中訪問 `http://localhost:8080/` 開始使用 GPT 聊天機器人。

## 功能介紹

GPT 聊天機器人提供以下功能：

- 傳送訊息：您可以輸入文字訊息，聊天機器人會根據您的訊息進行回應。
- 新增聊天：您可以新增一個新的聊天。
- 多執行緒聊天：聊天功能基於多執行緒編寫，可以同時支持多個聊天互動。

## 程式文件

GPT 聊天機器人包含以下程式文件：

- `main.go`：為主程式碼文件，負責應用程式初始化和運行。
- `chat.go`：提供聊天相關功能的程式碼文件。

## 貢獻指南

如有任何問題，歡迎提交 issue 或者 pull request 進行貢獻，感謝您的支持。

## 授權

GPT 聊天機器人是一個開源項目，採用 MIT 授權，有關詳細的授權信息，請參考 LICENSE 文件。