package main

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func main() {
	// 引数の検証
	if len(os.Args) != 4 {
		fmt.Println("Usage: file-converter markdown <inputfile> <outputfile>")
		os.Exit(1)
	}

	command := os.Args[1]
	inputFilePath := os.Args[2]
	outputFilePath := os.Args[3]

	// コマンドの検証
	if command != "markdown" {
		fmt.Println("Invalid command. Supported command: markdown")
		os.Exit(1)
	}

	// ファイルを読み込む
	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		os.Exit(1)
	}

	// MarkdownをHTMLに変換
	convertedData := markdownToHTML(string(inputData))

	// HTMLをファイルに書き込む
	err = os.WriteFile(outputFilePath, []byte(convertedData), 0644)
	if err != nil {
		fmt.Printf("Failed to write output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Markdown successfully converted to HTML!")
}

// MarkdownをHTMLに変換する関数
func markdownToHTML(markdownText string) string {
	// HTMLのレンダリングオプションを設定
	renderer := html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	})

	// MarkdownをHTMLに変換
	return string(markdown.ToHTML([]byte(markdownText), nil, renderer))
}
