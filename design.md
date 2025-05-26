# WebSocket 抽獎伺服器 - 系統設計文件

## 架構設計 (Clean Architecture)

### 目錄結構
```
websocket-lottery/
├── main.go                 # 應用程式入口
├── domain/                 # 領域模型層
│   └── message.go          # 定義核心資料結構
├── usecase/                # 商業邏輯層
│   └── lottery_service.go  # 抽獎邏輯處理
└── infrastructure/         # 基礎設施層
    └── websocket_server.go # WebSocket 伺服器實作
```

## 領域模型 (Domain)
- `LotteryMessage` 結構定義抽獎訊息
  - `RoomID`: 房間標識
  - `Number`: 抽獎號碼
  - `Result`: 中獎結果
  - `Prize`: 獎金金額

## 商業邏輯 (UseCase)
- `ProcessLottery` 方法實現抽獎規則
  - 特殊號碼 7 → 特獎 10000 元
  - 偶數 → 中獎 50 元
  - 其他數字 → 未中獎 0 元

## 基礎設施 (Infrastructure)
- WebSocket 伺服器實作
  - 使用 gorilla/websocket 套件
  - 支持跨域連線
  - 即時處理和回傳抽獎結果

## 連線與互動
- 伺服器監聽 `:8080` 端口
- WebSocket 路徑 `/lottery`
- 支持 JSON 格式訊息交換

## 範例訊息
### 輸入
```json
{
    "room_id": "room1",
    "number": 7
}
```

### 輸出
```json
{
    "room_id": "room1", 
    "number": 7,
    "result": "特獎",
    "prize": 10000
}
```

## 擴充性考量
1. 可輕鬆新增更複雜的抽獎邏輯
2. 領域模型可彈性擴充
3. 遵循 Clean Architecture 原則，各層解耦

## 測試方式
1. 啟動伺服器：`go run main.go`
2. 使用 WebSocket 客戶端
3. 連線 `ws://localhost:8080/lottery`
4. 傳送不同數字測試中獎邏輯 