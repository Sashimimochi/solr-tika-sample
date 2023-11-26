package main

import (
	"github.com/google/go-tika/tika"
	"context"
	"log"
	"os"
	"fmt"
	"strconv"
	"time"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Document struct {
	Title string
	Body string
	CreatedDate time.Time
	ModifiedDate time.Time
	Pages int
	CharLength int
}

func main() {
	err := createTable()
	if err != nil {
		log.Fatal(err)
	}
	filepath := "book.pdf"
	err = insertTable(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = selectTable()
	if err != nil {
		log.Fatal(err)
	}
}

func extract(filepath string) Document {
	err := tika.DownloadServer(context.Background(), tika.Version121, "tika-server.jar")
	if err != nil {
		log.Fatal(err)
	}
	s, err := tika.NewServer("tika-server.jar", "")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer s.Stop()

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, s.URL())
	meta, err := client.MetaRecursive(context.Background(), f)
	metadata := meta[0]
	pages, err := strconv.Atoi(metadata["xmpTPg:NPages"][0])
	if err != nil {
		log.Fatal(err)
	}
	createdDate, err := strpDate(metadata["created"][0])
	if err != nil {
		log.Fatal(err)
	}
	modifiedDate, err := strpDate(metadata["modified"][0])
	if err != nil {
		log.Fatal(err)
	}
	doc := Document{
		Title: metadata["title"][0],
		Body: metadata["X-TIKA:content"][0],
		CreatedDate: createdDate,
		ModifiedDate: modifiedDate,
		Pages: pages,
		CharLength: calcSum(metadata["pdf:charsPerPage"]),
	}
	return doc
}

func calcSum(strNums []string) int {
	sum := 0
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Errorf("数値に変換できません:%s", strNum)
		}
		sum += num
	}
	return sum
}

func strpDate(dateString string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func createTable() error {
	// MySQLデータベースに接続
	db, err := sql.Open("mysql", "mysql:mysql@tcp(mysql:3306)/book")
	if err != nil {
		return err
	}
	defer db.Close()

	// テーブルの作成クエリ
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS documents (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255),
			body TEXT,
			created_date DATETIME,
			modified_date DATETIME,
			pages INT,
			char_length INT
		)
	`
	// テーブル作成の実行
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("テーブルが作成されました")
	return nil
}

func insertTable(filepath string) error {
	// MySQLデータベースに接続
	db, err := sql.Open("mysql", "mysql:mysql@tcp(mysql:3306)/book")
	if err != nil {
		return err
	}
	defer db.Close()

	// 挿入するデータ
	document := extract(filepath)

	// データ挿入クエリ
	insertQuery := `
		INSERT INTO documents (title, body, created_date, modified_date, pages, char_length)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	// データ挿入の実行
	_, err = db.Exec(insertQuery, document.Title, document.Body, document.CreatedDate, document.ModifiedDate, document.Pages, document.CharLength)
	if err != nil {
		return err
	}
	fmt.Println("データが挿入されました")
	return nil
}

func selectTable() error {
	// MySQLデータベースに接続
	db, err := sql.Open("mysql", "mysql:mysql@tcp(mysql:3306)/book")
	if err != nil {
		return err
	}
	defer db.Close()

	// データを取得するクエリ
	selectQuery := `SELECT * FROM documents`

	// データ取得の実行
	rows, err := db.Query(selectQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("挿入されたデータ:")

	// 結果を取得して表示
	for rows.Next() {
		var (
			id            int
			title         string
			body          string
			createdDate   []uint8 // []uint8型に変更
			modifiedDate  []uint8 // []uint8型に変更
			pages         int
			charLength int
		)
		err := rows.Scan(&id, &title, &body, &createdDate, &modifiedDate, &pages, &charLength)
		if err != nil {
			return err
		}

		// []uint8型からtime.Time型への変換
		createdDateTime, err := time.Parse("2006-01-02 15:04:05", string(createdDate))
		if err != nil {
			return err
		}
		modifiedDateTime, err := time.Parse("2006-01-02 15:04:05", string(modifiedDate))
		if err != nil {
			return err
		}

		fmt.Printf("ID: %d, Title: %s, Body: %s, CreatedDate: %s, ModifiedDate: %s, Pages: %d, CharLength: %d\n", id, title, body, createdDateTime, modifiedDateTime, pages, charLength)
	}
	return nil
}
