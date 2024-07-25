// main.go

package main

import (
	"company/config"
	"company/controllers"
	"company/services"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
