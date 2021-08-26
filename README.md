# go-http-server-helper
Helper for http server by go - make your code clear

### ResponseWriterLazy
Same as http.ResponseWriter, except call methods WriteHeader and Write lazy.
To call lazy methods use method Done.

Example:
```
func SomeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		rw := &httpHelper.ResponseWriterLazy{
			Writer: response,
		}
		next.ServeHTTP(response, r)
		
		response.WriteHeader(200)

		if _, err := rw.Done(); err != nil {
			log.Println(err)
		}
	})
}
```
