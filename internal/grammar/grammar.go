package grammar

import "net/http"

func HandleMarkdownGrammarCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, grammar endpoint!"))
}
