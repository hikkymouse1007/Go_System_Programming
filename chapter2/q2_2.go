package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// gzipのパターンを真似して作ってみる
	file, err := os.Create("test.csv")
	if err != nil {
		panic((err))
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"FirstName, LastNAme"}) //GoのSlice宣言
	writer.Flush()
	writer.Write([]string{"Hiroki, Mashimo"})
	writer.Flush()

	// これでまとめて書き込むこともできる模様
	// writer.Write([]string{"abc","efg", "zyx"})
	// writer.Write([]string{"1","2", "3"})
	// writer.Write([]string{"A","B", "C"})
	// writer.Flush() // ここで書き込みを実行する

	// godocのサンプル
	// records := [][]string{
	// 	{"first_name", "last_name", "username"},
	// 	{"Rob", "Pike", "rob"},
	// 	{"Ken", "Thompson", "ken"},
	// 	{"Robert", "Griesemer", "gri"},
	// }

	// w := csv.NewWriter(os.Stdout)

	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatalln("error writing record to csv:", err)
	// 	}
	// }

	// // Write any buffered data to the underlying writer (standard output).
	// w.Flush()

	// if err := w.Error(); err != nil {
	// 	log.Fatal(err)
	// }
}

// Output:
// first_name,last_name,username
// Rob,Pike,rob
// Ken,Thompson,ken
// Robert,Griesemer,gri
