# HR GraphQL Server

This is the GraphQL server for the API endpoint for the horse racing game side project that I am working on. 
It's implemented in Golang and the GraphQL implementation is done through the library gqlgen.

Here is the [live demo](https://vntchang.dev/hrgraphql/).
an example query would be like this
```graphql
query {
  horse(horseName: "Deep Impact"){
    name
    sire{
      name
    }
    dam{
      name
    }
  }
}
```
the result for the example would be
```json
{
  "data": {
    "horse": {
      "name": "Deep Impact",
      "sire": {
        "name": "Sunday Silence"
      },
      "dam": {
        "name": "Wind in Her Hair"
      }
    }
  }
}
```
Currently, one can only query horses using the internal id and name of the horse (exact match). 
So I guess there really isn't much this graphql endpoint can do at this time.
If you know any racehorse's name you can try it out; see if my crawler had crawled the info for the horse.
