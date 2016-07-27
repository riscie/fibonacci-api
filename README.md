# Simple go fibonacci web api 
Very simple go (golang) web api returning fiobnacci numbers in JSON.

### Tech
  * go net/http
  * negroni (middleware)
  * httprouter (mux)


### Usage
  * `/` returns this help
  * `/fibonacci/5` returns: `{ [0,1,1,2,3,5] }` in JSON
