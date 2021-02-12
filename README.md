# logging-middleware
ログ出力するミドルウェアを自分で実装してみる


- ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux（公式ドキュメントより
）
