# go-http-server-helper
Helper for http server by go - make your code clear

### ResponseWriterLazy
Same as http.ResponseWriter, except call methods WriteHeader and Write lazy.
To call lazy methods use method Done.