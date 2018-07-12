package scripts

import()

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
}