package main

import "fmt"

func main() {
	books := []Book{
		{
			Title:      "Go言語入門",
			Author:     "山田太郎",
			Price:      2500,
			Categories: []string{"プログラミング", "技術書"},
		},
		{
			Title:      "データ構造とアルゴリズム",
			Author:     "佐藤次郎",
			Price:      3200,
			Categories: []string{"計算機科学", "教育"},
		},
		{
			Title:      "効率的Go",
			Author:     "鈴木花子",
			Price:      2800,
			Categories: []string{"プログラミング", "中級者向け"},
		},
	}

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	authorName := "山田太郎"
	fmt.Println("著者: ", authorName, ", タイトル: ", authorBooksMap[authorName].Title)
}

// 構造体 Book を定義 ---

// フィールド
// Title ... 文字列、本のタイトル
// Author ... 文字列、著者
// Price ... 整数、価格
// Categories ... 文字列のスライス、カテゴリ

// --------------------

type Book struct {
	Title string
	Author string
	Price int
	Categories []string
}

// Bookの配列を引数とし、本の合計価格を戻り値とする関数 `Total` を実装する
func Total(books []Book) (int){
	sum := 0
	for _, book := range books {
		sum += book.Price
	}
	return sum
}

// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap(books []Book) (map[string]Book){
	mp := make(map[string]Book)
	for _, book := range books {
		mp[book.Author] = book
	}
	return mp
}
