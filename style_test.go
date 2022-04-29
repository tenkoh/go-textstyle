package style_test

// ToDo
// []byteを受け取って文字単位で処理をする => utf8.DecodeRune([]byte) (rune, int)するだけ。テスト不要。
// ok a-zA-Z0-9のRegularの時だけ変換処理をする
// ok  uint8の特定の範囲の時にTrueを返す関数 -> internalに実装
// ok  上を使ってRegular書体の小文字、大文字、数字に合致するかを判定する関数 -> internalに実装
// ok 変換処理は複数パターン持たせる => interfaceを作成した
// 不要 []byte中に無効なbyte列が含まれている場合はErrShortSrcを返す
// 不要 変換した[]byteの末尾に無効なbyte列があれば内部にストックする
// 不要 ただしEOFの場合はストックせずにそのまま吐き出す
// 不要 次の変換では内部にストックしたbyte列を先頭にくっつけてから処理を開始する
// 不要 変換の先頭に無効なbyte列が4byteより多く溜まったらいったん吐き出す
// ok 書き込み先の[]byteの大きさが足りない場合は書き込まずにErrShortDstを返し、内部に保持しておく
