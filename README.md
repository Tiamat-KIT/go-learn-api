This is a minimal Go backend API service based on: https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service

Server should be run automatically when starting a workspace. Use `go run main.go` to run manually.

# 勤務シフト作成プログラムを作ってみたかった

## APIエンドポイントの構成
### 勤務者
#### データ形式
```typescript
{
  "name": string,　
  "grade": number, // 整数なのでint16
  "pref_work_times": string[],
  "pref_working_hours": number　// 小数点なのでfloat32
}
```
```go
package main

import (
	"fmt"
	"os"
)

type Employee struct {
	name string
	grade int16
	pref_work_times []string
	pref_working_hours float32
}

func (*Employee) new_employee (
	name string,
	grade int16,
	work_times []string,
	work_hours float32,
) Employee {
	if grade < 0 && 4 < grade {
		fmt.Printf("grade number input error")
		os.Exit(1)
	}
	return Employee{
		name: name,
		grade: grade,
		pref_work_times: work_times,
		pref_working_hours: work_hours,
	}	
}
```

#### 登録
勤務者データをデータベースに登録する
#### 削除
勤務者データをデータベースから削除する
#### 編集
データベース内の勤務者データを編集する
#### 読み込み
勤務者データをデータベースから読み込んで参照する

### 勤務シフト

#### 生成
勤務シフトデータを生成し、データベースに保存する

#### 削除
勤務シフトデータをデータベースから削除する

#### 読み込み
勤務シフトデータをデータベースから読み込む

#### 登録
勤務シフトデータを画像化し、正式なシフトとして送信する