package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Write() メソッドで書き込まれた内容を淡々とためておいて
	// あとでまとめて結果を受け取れる bytes.Buffer があります
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// この例では、今までの例と同じように、バイト列に変換した文字列をWrite() メソッドに渡しています。しかし bytes.Buffer には、特別に文字列を受け取れる WriteString() というメソッドもあります。そのため、Write() メソッドを呼んで いる行は次のように書き換えることもできます。
// buffer.WriteString("bytes.Buffer example\n")

// しかし、WriteString() は io.Writer のメソッドではないため、他の構造体では 使えません。代わりに、次の io.WriteString() 関数を使えばキャストは不要になり ます。
// io.WriteString(&buffer, "bytes.Buffer example\n")
