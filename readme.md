
# 1. Go 언어 개발의 통합 개발 환경(IDE) 구축

## 1-1. Go 언어 개발의 통합 개발 환경(IDE)의 종류

1. Visual Studio Code (VS Code): 가벼우면서도 강력한 편집기로, Go 언어 개발을 위한 확장 기능을 지원하며, 디버깅 기능과 통합된 터미널을 제공합니다. Go 개발 환경 설정이 비교적 쉽고, 다양한 플러그인을 통해 개인화할 수 있습니다.

2. GoLand: JetBrains에서 개발한 Go 전용 IDE로, Go 개발을 위한 전문적인 기능과 도구들을 제공합니다. 코드 완성, 리팩토링, 디버깅, 테스트 실행 등을 통합하여 제공하며, 대규모 프로젝트에 유용합니다.

3. Atom: GitHub에서 개발한 오픈 소스 편집기로, 다양한 언어를 지원하며 Go 언어 개발을 위한 패키지와 테마를 제공합니다. 특히 커스터마이징이 용이하다는 장점이 있습니다.

4. Sublime Text: 빠르고 가벼운 편집기로, Go 언어를 위한 패키지를 제공하며, 사용자가 원하는 대로 설정하여 사용할 수 있습니다.

5. Emacs: Emacs는 매우 강력한 편집기로, Go 개발을 위한 패키지들이 풍부하게 제공됩니다. Emacs 사용자들은 커스터마이징이 자유롭고, 다양한 통합 기능을 활용할 수 있습니다.

6. Vim: Vim은 CLI 환경에서도 뛰어난 성능을 보여주는 텍스트 에디터입니다. Go 언어 개발을 위한 플러그인들이 있어, 효율적으로 개발할 수 있습니다.

이 외에도 많은 개발자들이 자신이 편애하는 다른 편집기나 IDE를 사용하기도 합니다. Go 언어는 경량화와 성능 최적화가 강조되는 만큼, 다양한 개발 도구를 통해 개발자들이 편리하게 사용할 수 있습니다.

<br><br><br>

## 1-2. Go 언어 개발의 통합 개발 환경(IDE)의

### 1-2-1. Visual Studio Code 설치

먼저, Visual Studio Code 공식 웹사이트에서 OS에 맞는 설치 파일을 다운로드하고 설치합니다. 설치 후 VS Code를 실행합니다.

<br><br>

### 1-2-2. Go 확장 설치

Go 언어 개발을 위한 확장을 설치합니다. VS Code에서 Extensions (Ctrl + Shift + X)를 클릭하고 Go를 검색하여 설치합니다. 이 확장은 Go 언어 개발을 위한 여러 기능을 제공합니다.


<br><br>

### 1-2-3. Go 설치

만약, 시스템에 Go가 설치되어 있지 않다면, 공식 Go 웹사이트에서 OS에 맞는 설치 파일을 다운로드하고 설치합니다. 설치 후 터미널에서 go version 명령어를 실행하여 설치가 잘 되었는지 확인합니다.

<br><br>

### 1-2-4. Go 환경 설정

Go 언어 개발을 위해 GOPATH 환경 변수를 설정합니다. 일반적으로 $HOME/go를 사용하도록 설정합니다. 터미널에서 다음과 같이 환경 변수를 설정합니다:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

<br><br>

### 1-2-5. VS Code 설정

VS Code에서 Go 환경 설정을 변경합니다. File -> Preferences -> Settings를 선택하거나 Ctrl + ,를 눌러 설정 창을 엽니다. 검색창에 "go.toolsGopath"을 입력하고 Edit in settings.json을 클릭합니다. 다음 JSON 설정을 추가합니다:

```json
"go.toolsGopath": "/home/{your_username}/go"
```

위의 {your_username}를 자신의 사용자 이름으로 변경합니다

<br><br><br><br>

# 2. Go 언어의 개발 프로젝트 실습 

## 2-1. 프로젝트 구조 및 설정

### 2-1-1. 프로젝트 구조

```lua
go-rest-api/
├── config/
│   └── db.go
├── controllers/
│   ├── board_controller.go
│   ├── dataroom_controller.go
│   ├── member_controller.go
│   ├── product_controller.go
│   └── qna_controller.go
├── models/
│   ├── board.go
│   ├── dataroom.go
│   ├── member.go
│   ├── product.go
│   └── qna.go
├── services/
│   ├── board_service.go
│   ├── dataroom_service.go
│   ├── member_service.go
│   ├── product_service.go
│   └── qna_service.go
└── main.go
```

<br><br>

### 2-1-2. 의존성 관리 및 라이브러리 설치

아래 명령어로 필요한 Go 패키지를 설치합니다:

```bash
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm    # 선택적으로 사용 (ORM 라이브러리)
```

<br><br>

### 2-1-3. 데이터베이스 설정

MariaDB 설정은 config/db.go 파일에서 다룹니다:

```go
// config/db.go

package config

const (
    DBUser     = "root"
    DBPassword = "1234"
    DBName     = "company"
    DBHost     = "localhost"
    DBPort     = "3307"
)

func GetDBConnectionString() string {
    return DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"
}
```

<br><br><br>

## 2-2. Board 관련 기능 개발

### 2-2-1. board.go

먼저, models 디렉터리에 board.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// models/board.go

package models

import "time"

type Board struct {
    No      int       `json:"no"`
    Title   string    `json:"title"`
    Content string    `json:"content"`
    Author  string    `json:"author"`
    ResDate time.Time `json:"resdate"`
    Hits    int       `json:"hits"`
}
```

<br><br>

### 2-2-2. board_service.go

다음으로, services 디렉터리에 board_service.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// services/board_service.go

package services

import (
    "database/sql"
    "your-app/models"
)

type BoardService struct {
    DB *sql.DB
}

func (s *BoardService) GetBoardList() ([]models.Board, error) {
    var boards []models.Board
    rows, err := s.DB.Query("SELECT no, title, content, author, resdate, hits FROM board")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var board models.Board
        if err := rows.Scan(&board.No, &board.Title, &board.Content, &board.Author, &board.ResDate, &board.Hits); err != nil {
            return nil, err
        }
        boards = append(boards, board)
    }
    return boards, nil
}

func (s *BoardService) GetBoardDetail(no int) (*models.Board, error) {
    var board models.Board
    row := s.DB.QueryRow("SELECT no, title, content, author, resdate, hits FROM board WHERE no = ?", no)
    if err := row.Scan(&board.No, &board.Title, &board.Content, &board.Author, &board.ResDate, &board.Hits); err != nil {
        return nil, err
    }
    return &board, nil
}

func (s *BoardService) InsertBoard(board *models.Board) error {
    _, err := s.DB.Exec("INSERT INTO board (title, content, author, resdate, hits) VALUES (?, ?, ?, ?, ?)",
        board.Title, board.Content, board.Author, board.ResDate, board.Hits)
    return err
}

func (s *BoardService) UpdateBoard(board *models.Board) error {
    _, err := s.DB.Exec("UPDATE board SET title = ?, content = ?, author = ?, resdate = ?, hits = ? WHERE no = ?",
        board.Title, board.Content, board.Author, board.ResDate, board.Hits, board.No)
    return err
}

func (s *BoardService) DeleteBoard(no int) error {
    _, err := s.DB.Exec("DELETE FROM board WHERE no = ?", no)
    return err
}
```
<br><br>

### 2-2-3. board_controller.go

마지막으로, controllers 디렉터리에 board_controller.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// controllers/board_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "your-app/services"

    "github.com/gorilla/mux"
)

type BoardController struct {
    BoardService *services.BoardService
}

func (c *BoardController) GetBoardList(w http.ResponseWriter, r *http.Request) {
    boards, err := c.BoardService.GetBoardList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(boards)
}

func (c *BoardController) GetBoardDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    no, err := strconv.Atoi(vars["no"])
    if err != nil {
        http.Error(w, "Invalid board ID", http.StatusBadRequest)
        return
    }

    board, err := c.BoardService.GetBoardDetail(no)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(board)
}

func (c *BoardController) InsertBoard(w http.ResponseWriter, r *http.Request) {
    var board models.Board
    if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.BoardService.InsertBoard(&board); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *BoardController) UpdateBoard(w http.ResponseWriter, r *http.Request) {
    var board models.Board
    if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.BoardService.UpdateBoard(&board); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *BoardController) DeleteBoard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    no, err := strconv.Atoi(vars["no"])
    if err != nil {
        http.Error(w, "Invalid board ID", http.StatusBadRequest)
        return
    }

    if err := c.BoardService.DeleteBoard(no); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```
<br><br>

### 2-2-4. main.go

모든 것을 main.go 파일에서 연결하고 실행합니다:

```go
// main.go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "your-app/config"
    "your-app/controllers"
    "your-app/services"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", config.GetDBConnectionString())
    if err != nil {
        log.Fatal("데이터베이스 연결 실패:", err)
    }
    defer db.Close()

    // 서비스 초기화
    boardService := &services.BoardService{DB: db}

    // 컨트롤러 초기화
    boardController := &controllers.BoardController{BoardService: boardService}

    router := mux.NewRouter()

    // 라우트 정의
    router.HandleFunc("/company/boards/list", boardController.GetBoardList).Methods("GET")
    router.HandleFunc("/company/boards/detail/{no}", boardController.GetBoardDetail).Methods("GET")
    router.HandleFunc("/company/boards/insert", boardController.InsertBoard).Methods("POST")
    router.HandleFunc("/company/boards/update", boardController.UpdateBoard).Methods("POST")
    router.HandleFunc("/company/boards/delete/{no}", boardController.DeleteBoard).Methods("POST")

    // 서버 시작
    log.Fatal(http.ListenAndServe(":8087", router))
}
```

<br><br><br>

## 2-3. Qna 관련 기능 개발

### 2-3-1. qna.go

- models 디렉터리에 qna.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// models/qna.go

package models

import "time"

type Qna struct {
    Qno     int       `json:"qno"`
    Lev     int       `json:"lev"`
    Parno   int       `json:"parno"`
    Title   string    `json:"title"`
    Content string    `json:"content"`
    Author  string    `json:"author"`
    Resdate time.Time `json:"resdate"`
    Hits    int       `json:"hits"`
}
```

<br><br>

### 2-3-2. qna_service.go

- services 디렉터리에 qna_service.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// services/qna_service.go

package services

import (
    "database/sql"
    "your-app/models"
)

type QnaService struct {
    DB *sql.DB
}

func (s *QnaService) GetQnaList() ([]models.Qna, error) {
    var qnas []models.Qna
    rows, err := s.DB.Query("SELECT qno, lev, parno, title, content, author, resdate, hits FROM qna")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var qna models.Qna
        if err := rows.Scan(&qna.Qno, &qna.Lev, &qna.Parno, &qna.Title, &qna.Content, &qna.Author, &qna.Resdate, &qna.Hits); err != nil {
            return nil, err
        }
        qnas = append(qnas, qna)
    }
    return qnas, nil
}

func (s *QnaService) GetQnaDetail(qno int) (*models.Qna, error) {
    var qna models.Qna
    row := s.DB.QueryRow("SELECT qno, lev, parno, title, content, author, resdate, hits FROM qna WHERE qno = ?", qno)
    if err := row.Scan(&qna.Qno, &qna.Lev, &qna.Parno, &qna.Title, &qna.Content, &qna.Author, &qna.Resdate, &qna.Hits); err != nil {
        return nil, err
    }
    return &qna, nil
}

func (s *QnaService) InsertQna(qna *models.Qna) error {
    _, err := s.DB.Exec("INSERT INTO qna (lev, parno, title, content, author, resdate, hits) VALUES (?, ?, ?, ?, ?, ?, ?)",
        qna.Lev, qna.Parno, qna.Title, qna.Content, qna.Author, qna.Resdate, qna.Hits)
    return err
}

func (s *QnaService) UpdateQna(qna *models.Qna) error {
    _, err := s.DB.Exec("UPDATE qna SET lev = ?, parno = ?, title = ?, content = ?, author = ?, resdate = ?, hits = ? WHERE qno = ?",
        qna.Lev, qna.Parno, qna.Title, qna.Content, qna.Author, qna.Resdate, qna.Hits, qna.Qno)
    return err
}

func (s *QnaService) DeleteQna(qno int) error {
    _, err := s.DB.Exec("DELETE FROM qna WHERE qno = ?", qno)
    return err
}
```

<br><br>

### 2-3-3. qna_controller.go

- controllers 디렉터리에 qna_controller.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// controllers/qna_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "your-app/models"
    "your-app/services"

    "github.com/gorilla/mux"
)

type QnaController struct {
    QnaService *services.QnaService
}

func (c *QnaController) GetQnaList(w http.ResponseWriter, r *http.Request) {
    qnas, err := c.QnaService.GetQnaList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(qnas)
}

func (c *QnaController) GetQnaDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    qno, err := strconv.Atoi(vars["qno"])
    if err != nil {
        http.Error(w, "Invalid Qna ID", http.StatusBadRequest)
        return
    }

    qna, err := c.QnaService.GetQnaDetail(qno)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(qna)
}

func (c *QnaController) InsertQna(w http.ResponseWriter, r *http.Request) {
    var qna models.Qna
    if err := json.NewDecoder(r.Body).Decode(&qna); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.QnaService.InsertQna(&qna); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *QnaController) UpdateQna(w http.ResponseWriter, r *http.Request) {
    var qna models.Qna
    if err := json.NewDecoder(r.Body).Decode(&qna); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.QnaService.UpdateQna(&qna); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *QnaController) DeleteQna(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    qno, err := strconv.Atoi(vars["qno"])
    if err != nil {
        http.Error(w, "Invalid Qna ID", http.StatusBadRequest)
        return
    }

    if err := c.QnaService.DeleteQna(qno); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```

<br><br>

### 2-3-4. main.go

main.go 파일을 열고 Qna 관련 라우팅 내용을 추가하고 실행합니다:

```go
// main.go

package main

import (
    "database/sql"
    "log"
    "net/http"
    "your-app/config"
    "your-app/controllers"
    "your-app/services"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", config.GetDBConnectionString())
    if err != nil {
        log.Fatal("데이터베이스 연결 실패:", err)
    }
    defer db.Close()

    // 서비스 초기화
    boardService := &services.BoardService{DB: db}
    qnaService := &services.QnaService{DB: db}

    // 컨트롤러 초기화
    boardController := &controllers.BoardController{BoardService: boardService}
    qnaController := &controllers.QnaController{QnaService: qnaService}

    router := mux.NewRouter()

    // Board 라우트 정의
    router.HandleFunc("/company/boards/list", boardController.GetBoardList).Methods("GET")
    router.HandleFunc("/company/boards/detail/{no}", boardController.GetBoardDetail).Methods("GET")
    router.HandleFunc("/company/boards/insert", boardController.InsertBoard).Methods("POST")
    router.HandleFunc("/company/boards/update", boardController.UpdateBoard).Methods("POST")
    router.HandleFunc("/company/boards/delete/{no}", boardController.DeleteBoard).Methods("POST")

    // Qna 라우트 정의
    router.HandleFunc("/company/qna/list", qnaController.GetQnaList).Methods("GET")
    router.HandleFunc("/company/qna/detail/{qno}", qnaController.GetQnaDetail).Methods("GET")
    router.HandleFunc("/company/qna/insert", qnaController.InsertQna).Methods("POST")
    router.HandleFunc("/company/qna/update", qnaController.UpdateQna).Methods("POST")
    router.HandleFunc("/company/qna/delete/{qno}", qnaController.DeleteQna).Methods("POST")

    // 서버 시작
    log.Fatal(http.ListenAndServe(":8087", router))
}
```

<br><br><br>

## 2-4. Dataroom 관련 기능 개발

### 2-4-1. dataroom.go

- models 디렉터리에 dataroom.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// models/dataroom.go

package models

import "time"

type Dataroom struct {
    Dno      int       `json:"dno"`
    Title    string    `json:"title"`
    Content  string    `json:"content"`
    Author   string    `json:"author"`
    Datafile string    `json:"datafile"`
    ResDate  time.Time `json:"resdate"`
    Hits     int       `json:"hits"`
}
```

<br><br>

### 2-4-2. dataroom_service.go

- services 디렉터리에 dataroom_service.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// services/dataroom_service.go

package services

import (
    "database/sql"
    "your-app/models"
)

type DataroomService struct {
    DB *sql.DB
}

func (s *DataroomService) GetDataroomList() ([]models.Dataroom, error) {
    var datarooms []models.Dataroom
    rows, err := s.DB.Query("SELECT dno, title, content, author, datafile, resdate, hits FROM dataroom")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var dataroom models.Dataroom
        if err := rows.Scan(&dataroom.Dno, &dataroom.Title, &dataroom.Content, &dataroom.Author, &dataroom.Datafile, &dataroom.ResDate, &dataroom.Hits); err != nil {
            return nil, err
        }
        datarooms = append(datarooms, dataroom)
    }
    return datarooms, nil
}

func (s *DataroomService) GetDataroomDetail(dno int) (*models.Dataroom, error) {
    var dataroom models.Dataroom
    row := s.DB.QueryRow("SELECT dno, title, content, author, datafile, resdate, hits FROM dataroom WHERE dno = ?", dno)
    if err := row.Scan(&dataroom.Dno, &dataroom.Title, &dataroom.Content, &dataroom.Author, &dataroom.Datafile, &dataroom.ResDate, &dataroom.Hits); err != nil {
        return nil, err
    }
    return &dataroom, nil
}

func (s *DataroomService) InsertDataroom(dataroom *models.Dataroom) error {
    _, err := s.DB.Exec("INSERT INTO dataroom (title, content, author, datafile, resdate, hits) VALUES (?, ?, ?, ?, ?, ?)",
        dataroom.Title, dataroom.Content, dataroom.Author, dataroom.Datafile, dataroom.ResDate, dataroom.Hits)
    return err
}

func (s *DataroomService) UpdateDataroom(dataroom *models.Dataroom) error {
    _, err := s.DB.Exec("UPDATE dataroom SET title = ?, content = ?, author = ?, datafile = ?, resdate = ?, hits = ? WHERE dno = ?",
        dataroom.Title, dataroom.Content, dataroom.Author, dataroom.Datafile, dataroom.ResDate, dataroom.Hits, dataroom.Dno)
    return err
}

func (s *DataroomService) DeleteDataroom(dno int) error {
    _, err := s.DB.Exec("DELETE FROM dataroom WHERE dno = ?", dno)
    return err
}
```
<br><br>

### 2-4-3. dataroom_controller.go

- controllers 디렉터리에 dataroom_controller.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// controllers/dataroom_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "your-app/models"
    "your-app/services"

    "github.com/gorilla/mux"
)

type DataroomController struct {
    DataroomService *services.DataroomService
}

func (c *DataroomController) GetDataroomList(w http.ResponseWriter, r *http.Request) {
    datarooms, err := c.DataroomService.GetDataroomList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(datarooms)
}

func (c *DataroomController) GetDataroomDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    dno, err := strconv.Atoi(vars["dno"])
    if err != nil {
        http.Error(w, "Invalid Dataroom ID", http.StatusBadRequest)
        return
    }

    dataroom, err := c.DataroomService.GetDataroomDetail(dno)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(dataroom)
}

func (c *DataroomController) InsertDataroom(w http.ResponseWriter, r *http.Request) {
    var dataroom models.Dataroom
    if err := json.NewDecoder(r.Body).Decode(&dataroom); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.DataroomService.InsertDataroom(&dataroom); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *DataroomController) UpdateDataroom(w http.ResponseWriter, r *http.Request) {
    var dataroom models.Dataroom
    if err := json.NewDecoder(r.Body).Decode(&dataroom); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.DataroomService.UpdateDataroom(&dataroom); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *DataroomController) DeleteDataroom(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    dno, err := strconv.Atoi(vars["dno"])
    if err != nil {
        http.Error(w, "Invalid Dataroom ID", http.StatusBadRequest)
        return
    }

    if err := c.DataroomService.DeleteDataroom(dno); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```

<br><br>

### 2-4-4. main.go

- main.go 파일을 열고, Dataroom 기능 관련 내용을 추가한 후 연결하고 실행합니다:

```go
// main.go

package main

import (
    "database/sql"
    "log"
    "net/http"
    "your-app/config"
    "your-app/controllers"
    "your-app/services"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", config.GetDBConnectionString())
    if err != nil {
        log.Fatal("데이터베이스 연결 실패:", err)
    }
    defer db.Close()

    // 서비스 초기화
    boardService := &services.BoardService{DB: db}
    qnaService := &services.QnaService{DB: db}
    dataroomService := &services.DataroomService{DB: db}

    // 컨트롤러 초기화
    boardController := &controllers.BoardController{BoardService: boardService}
    qnaController := &controllers.QnaController{QnaService: qnaService}
    dataroomController := &controllers.DataroomController{DataroomService: dataroomService}

    router := mux.NewRouter()

    // Board 라우트 정의
    router.HandleFunc("/company/boards/list", boardController.GetBoardList).Methods("GET")
    router.HandleFunc("/company/boards/detail/{no}", boardController.GetBoardDetail).Methods("GET")
    router.HandleFunc("/company/boards/insert", boardController.InsertBoard).Methods("POST")
    router.HandleFunc("/company/boards/update", boardController.UpdateBoard).Methods("POST")
    router.HandleFunc("/company/boards/delete/{no}", boardController.DeleteBoard).Methods("POST")

    // Qna 라우트 정의
    router.HandleFunc("/company/qna/list", qnaController.GetQnaList).Methods("GET")
    router.HandleFunc("/company/qna/detail/{qno}", qnaController.GetQnaDetail).Methods("GET")
    router.HandleFunc("/company/qna/insert", qnaController.InsertQna).Methods("POST")
    router.HandleFunc("/company/qna/update", qnaController.UpdateQna).Methods("POST")
    router.HandleFunc("/company/qna/delete/{qno}", qnaController.DeleteQna).Methods("POST")

    // Dataroom 라우트 정의
    router.HandleFunc("/company/dataroom/list", dataroomController.GetDataroomList).Methods("GET")
    router.HandleFunc("/company/dataroom/detail/{dno}", dataroomController.GetDataroomDetail).Methods("GET")
    router.HandleFunc("/company/dataroom/insert", dataroomController.InsertDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/update", dataroomController.UpdateDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/delete/{dno}", dataroomController.DeleteDataroom).Methods("POST")

    // 서버 시작
    log.Fatal(http.ListenAndServe(":8087", router))
}
```

<br><br><br>

## 2-5. Product 관련 기능 개발

### 2-5-1. product.go

- models 디렉터리에 product.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// models/product.go

package models

import "time"

type Product struct {
    Pno      int       `json:"pno"`
    Cate     string    `json:"cate"`
    Pname    string    `json:"pname"`
    Pcontent string    `json:"pcontent"`
    Img1     string    `json:"img1"`
    Img2     string    `json:"img2"`
    Img3     string    `json:"img3"`
    ResDate  time.Time `json:"resdate"`
    Hits     int       `json:"hits"`
}
```

<br><br>

### 2-5-2. product_service.go

- services 디렉터리에 product_service.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// services/product_service.go

package services

import (
    "database/sql"
    "your-app/models"
)

type ProductService struct {
    DB *sql.DB
}

func (s *ProductService) GetProductList() ([]models.Product, error) {
    var products []models.Product
    rows, err := s.DB.Query("SELECT pno, cate, pname, pcontent, img1, img2, img3, resdate, hits FROM product")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.Pno, &product.Cate, &product.Pname, &product.Pcontent, &product.Img1, &product.Img2, &product.Img3, &product.ResDate, &product.Hits); err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (s *ProductService) GetProductDetail(pno int) (*models.Product, error) {
    var product models.Product
    row := s.DB.QueryRow("SELECT pno, cate, pname, pcontent, img1, img2, img3, resdate, hits FROM product WHERE pno = ?", pno)
    if err := row.Scan(&product.Pno, &product.Cate, &product.Pname, &product.Pcontent, &product.Img1, &product.Img2, &product.Img3, &product.ResDate, &product.Hits); err != nil {
        return nil, err
    }
    return &product, nil
}

func (s *ProductService) InsertProduct(product *models.Product) error {
    _, err := s.DB.Exec("INSERT INTO product (cate, pname, pcontent, img1, img2, img3, resdate, hits) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
        product.Cate, product.Pname, product.Pcontent, product.Img1, product.Img2, product.Img3, product.ResDate, product.Hits)
    return err
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
    _, err := s.DB.Exec("UPDATE product SET cate = ?, pname = ?, pcontent = ?, img1 = ?, img2 = ?, img3 = ?, resdate = ?, hits = ? WHERE pno = ?",
        product.Cate, product.Pname, product.Pcontent, product.Img1, product.Img2, product.Img3, product.ResDate, product.Hits, product.Pno)
    return err
}

func (s *ProductService) DeleteProduct(pno int) error {
    _, err := s.DB.Exec("DELETE FROM product WHERE pno = ?", pno)
    return err
}
```

<br><br>

### 2-5-3. product_controller.go

- controllers 디렉터리에 product_controller.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// controllers/product_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "your-app/models"
    "your-app/services"

    "github.com/gorilla/mux"
)

type ProductController struct {
    ProductService *services.ProductService
}

func (c *ProductController) GetProductList(w http.ResponseWriter, r *http.Request) {
    products, err := c.ProductService.GetProductList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProductDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pno, err := strconv.Atoi(vars["pno"])
    if err != nil {
        http.Error(w, "Invalid Product ID", http.StatusBadRequest)
        return
    }

    product, err := c.ProductService.GetProductDetail(pno)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func (c *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.ProductService.InsertProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.ProductService.UpdateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pno, err := strconv.Atoi(vars["pno"])
    if err != nil {
        http.Error(w, "Invalid Product ID", http.StatusBadRequest)
        return
    }

    if err := c.ProductService.DeleteProduct(pno); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```


<br><br>

### 2-5-4. main.go

- main.go 파일을 열고, Product 관련 라우팅 내용을 추가한 후 연결하고 실행합니다:

```go
// main.go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "your-app/config"
    "your-app/controllers"
    "your-app/services"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", config.GetDBConnectionString())
    if err != nil {
        log.Fatal("데이터베이스 연결 실패:", err)
    }
    defer db.Close()

    // 서비스 초기화
    boardService := &services.BoardService{DB: db}
    qnaService := &services.QnaService{DB: db}
    dataroomService := &services.DataroomService{DB: db}
    productService := &services.ProductService{DB: db}

    // 컨트롤러 초기화
    boardController := &controllers.BoardController{BoardService: boardService}
    qnaController := &controllers.QnaController{QnaService: qnaService}
    dataroomController := &controllers.DataroomController{DataroomService: dataroomService}
    productController := &controllers.ProductController{ProductService: productService}

    router := mux.NewRouter()

    // Board 라우트 정의
    router.HandleFunc("/company/boards/list", boardController.GetBoardList).Methods("GET")
    router.HandleFunc("/company/boards/detail/{no}", boardController.GetBoardDetail).Methods("GET")
    router.HandleFunc("/company/boards/insert", boardController.InsertBoard).Methods("POST")
    router.HandleFunc("/company/boards/update", boardController.UpdateBoard).Methods("POST")
    router.HandleFunc("/company/boards/delete/{no}", boardController.DeleteBoard).Methods("POST")

    // Qna 라우트 정의
    router.HandleFunc("/company/qna/list", qnaController.GetQnaList).Methods("GET")
    router.HandleFunc("/company/qna/detail/{qno}", qnaController.GetQnaDetail).Methods("GET")
    router.HandleFunc("/company/qna/insert", qnaController.InsertQna).Methods("POST")
    router.HandleFunc("/company/qna/update", qnaController.UpdateQna).Methods("POST")
    router.HandleFunc("/company/qna/delete/{qno}", qnaController.DeleteQna).Methods("POST")

    // Dataroom 라우트 정의
    router.HandleFunc("/company/dataroom/list", dataroomController.GetDataroomList).Methods("GET")
    router.HandleFunc("/company/dataroom/detail/{dno}", dataroomController.GetDataroomDetail).Methods("GET")
    router.HandleFunc("/company/dataroom/insert", dataroomController.InsertDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/update", dataroomController.UpdateDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/delete/{dno}", dataroomController.DeleteDataroom).Methods("POST")

    // Product 라우트 정의
    router.HandleFunc("/company/products/list", productController.GetProductList).Methods("GET")
    router.HandleFunc("/company/products/detail/{pno}", productController.GetProductDetail).Methods("GET")
    router.HandleFunc("/company/products/insert", productController.InsertProduct).Methods("POST")
    router.HandleFunc("/company/products/update", productController.UpdateProduct).Methods("POST")
    router.HandleFunc("/company/products/delete/{pno}", productController.DeleteProduct).Methods("POST")

    // 서버 시작
    log.Fatal(http.ListenAndServe(":8087", router))
}
```

<br><br><br>

## 2-6. Member 관련 기능 개발

### 2-6-1. member.go

- models 디렉터리에 member.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// models/member.go

package models

import "time"

type Member struct {
    ID       string    `json:"id"`
    PW       string    `json:"pw"`
    Name     string    `json:"name"`
    Birth    time.Time `json:"birth"`
    Email    string    `json:"email"`
    Tel      string    `json:"tel"`
    Addr1    string    `json:"addr1"`
    Addr2    string    `json:"addr2"`
    Postcode string    `json:"postcode"`
    RegDate  time.Time `json:"regdate"`
}
```

<br><br>

### 2-6-2. member_service.go

- services 디렉터리에 member_service.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// services/member_service.go

package services

import (
    "database/sql"
    "your-app/models"
)

type MemberService struct {
    DB *sql.DB
}

func (s *MemberService) GetMemberList() ([]models.Member, error) {
    var members []models.Member
    rows, err := s.DB.Query("SELECT id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate FROM member")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var member models.Member
        if err := rows.Scan(&member.ID, &member.PW, &member.Name, &member.Birth, &member.Email, &member.Tel, &member.Addr1, &member.Addr2, &member.Postcode, &member.RegDate); err != nil {
            return nil, err
        }
        members = append(members, member)
    }
    return members, nil
}

func (s *MemberService) GetMember(id string) (*models.Member, error) {
    var member models.Member
    row := s.DB.QueryRow("SELECT id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate FROM member WHERE id = ?", id)
    if err := row.Scan(&member.ID, &member.PW, &member.Name, &member.Birth, &member.Email, &member.Tel, &member.Addr1, &member.Addr2, &member.Postcode, &member.RegDate); err != nil {
        return nil, err
    }
    return &member, nil
}

func (s *MemberService) InsertMember(member *models.Member) error {
    _, err := s.DB.Exec("INSERT INTO member (id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
        member.ID, member.PW, member.Name, member.Birth, member.Email, member.Tel, member.Addr1, member.Addr2, member.Postcode, member.RegDate)
    return err
}

func (s *MemberService) UpdateMember(member *models.Member) error {
    _, err := s.DB.Exec("UPDATE member SET pw = ?, name = ?, birth = ?, email = ?, tel = ?, addr1 = ?, addr2 = ?, postcode = ? WHERE id = ?",
        member.PW, member.Name, member.Birth, member.Email, member.Tel, member.Addr1, member.Addr2, member.Postcode, member.ID)
    return err
}

func (s *MemberService) DeleteMember(id string) error {
    _, err := s.DB.Exec("DELETE FROM member WHERE id = ?", id)
    return err
}
```

<br><br>

### 2-6-3. member_controller.go

- controllers 디렉터리에 member_controller.go 파일을 생성하고 다음과 같이 작성합니다:

```go
// controllers/member_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "your-app/models"
    "your-app/services"

    "github.com/gorilla/mux"
)

type MemberController struct {
    MemberService *services.MemberService
}

func (c *MemberController) GetMemberList(w http.ResponseWriter, r *http.Request) {
    members, err := c.MemberService.GetMemberList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(members)
}

func (c *MemberController) GetMember(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    member, err := c.MemberService.GetMember(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(member)
}

func (c *MemberController) InsertMember(w http.ResponseWriter, r *http.Request) {
    var member models.Member
    if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.MemberService.InsertMember(&member); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *MemberController) UpdateMember(w http.ResponseWriter, r *http.Request) {
    var member models.Member
    if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.MemberService.UpdateMember(&member); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *MemberController) DeleteMember(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if err := c.MemberService.DeleteMember(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```

<br><br>

### 2-6-4. main.go

- main.go 파일을 열고, Member 관련 라우팅 내용을 추가한 후 연결하고 실행합니다:

```go
// main.go

package main

import (
    "database/sql"
    "log"
    "net/http"
    "your-app/config"
    "your-app/controllers"
    "your-app/services"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", config.GetDBConnectionString())
    if err != nil {
        log.Fatal("데이터베이스 연결 실패:", err)
    }
    defer db.Close()

    // 서비스 초기화
    boardService := &services.BoardService{DB: db}
    qnaService := &services.QnaService{DB: db}
    dataroomService := &services.DataroomService{DB: db}
    productService := &services.ProductService{DB: db}
    memberService := &services.MemberService{DB: db}

    // 컨트롤러 초기화
    boardController := &controllers.BoardController{BoardService: boardService}
    qnaController := &controllers.QnaController{QnaService: qnaService}
    dataroomController := &controllers.DataroomController{DataroomService: dataroomService}
    productController := &controllers.ProductController{ProductService: productService}
    memberController := &controllers.MemberController{MemberService: memberService}

    router := mux.NewRouter()

    // Board 라우트 정의
    router.HandleFunc("/company/boards/list", boardController.GetBoardList).Methods("GET")
    router.HandleFunc("/company/boards/detail/{no}", boardController.GetBoardDetail).Methods("GET")
    router.HandleFunc("/company/boards/insert", boardController.InsertBoard).Methods("POST")
    router.HandleFunc("/company/boards/update", boardController.UpdateBoard).Methods("POST")
    router.HandleFunc("/company/boards/delete/{no}", boardController.DeleteBoard).Methods("POST")

    // Qna 라우트 정의
    router.HandleFunc("/company/qna/list", qnaController.GetQnaList).Methods("GET")
    router.HandleFunc("/company/qna/detail/{qno}", qnaController.GetQnaDetail).Methods("GET")
    router.HandleFunc("/company/qna/insert", qnaController.InsertQna).Methods("POST")
    router.HandleFunc("/company/qna/update", qnaController.UpdateQna).Methods("POST")
    router.HandleFunc("/company/qna/delete/{qno}", qnaController.DeleteQna).Methods("POST")

    // Dataroom 라우트 정의
    router.HandleFunc("/company/dataroom/list", dataroomController.GetDataroomList).Methods("GET")
    router.HandleFunc("/company/dataroom/detail/{dno}", dataroomController.GetDataroomDetail).Methods("GET")
    router.HandleFunc("/company/dataroom/insert", dataroomController.InsertDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/update", dataroomController.UpdateDataroom).Methods("POST")
    router.HandleFunc("/company/dataroom/delete/{dno}", dataroomController.DeleteDataroom).Methods("POST")

    // Product 라우트 정의
    router.HandleFunc("/company/products/list", productController.GetProductList).Methods("GET")
    router.HandleFunc("/company/products/detail/{pno}", productController.GetProductDetail).Methods("GET")
    router.HandleFunc("/company/products/insert", productController.InsertProduct).Methods("POST")
    router.HandleFunc("/company/products/update", productController.UpdateProduct).Methods("POST")
    router.HandleFunc("/company/products/delete/{pno}", productController.DeleteProduct).Methods("POST")

    // Member 라우트 정의
    router.HandleFunc("/company/members/getMemberList", memberController.GetMemberList).Methods("GET")
    router.HandleFunc("/company/members/getMember/{id}", memberController.GetMember).Methods("GET")
    router.HandleFunc("/company/members/insert", memberController.InsertMember).Methods("POST")
    router.HandleFunc("/company/members/update", memberController.UpdateMember).Methods("POST")
    router.HandleFunc("/company/members/delete/{id}", memberController.DeleteMember).Methods("POST")

    // 서버 시작
    log.Fatal(http.ListenAndServe(":8087", router))
}
```

<br><br><br>

## 2-7. 프로젝트 애플리케이션 실행


```bash
go run main.go
```

GET /company/boards/list: 게시판 목록 조회
GET /company/boards/detail/{no}: 특정 게시판 항목 조회
POST /company/boards/insert: 새로운 게시판 항목 삽입
POST /company/boards/update: 게시판 항목 수정
POST /company/boards/delete/{no}: 게시판 항목 삭제


GET /company/qna/list: Qna 목록 조회
GET /company/qna/detail/{qno}: 특정 Qna 항목 조회
POST /company/qna/insert: 새로운 Qna 항목 삽입
POST /company/qna/update: Qna 항목 수정
POST /company/qna/delete/{qno}: Qna 항목 삭제


GET /company/dataroom/list: Dataroom 목록 조회
GET /company/dataroom/detail/{dno}: 특정 Dataroom 항목 조회
POST /company/dataroom/insert: 새로운 Dataroom 항목 삽입
POST /company/dataroom/update: Dataroom 항목 수정
POST /company/dataroom/delete/{dno}: Dataroom 항목 삭제
