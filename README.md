# Time Zone API
Get Current Time From Time Zone

```go
// get value of param (?zone=) from url
keys, _ := request.URL.Query()["zone"] 

// get current time of time zone (zone is getting from param in url)
location, _ := time.LoadLocation(string(keys[0])) 


fmt.Fprint(writer, "the time ", time.Now().In(location)) 
```

<img  src="https://raw.githubusercontent.com/kadir-ince/Time-Zone-API/main/medias/tokyo.png"/>
<img  src="https://raw.githubusercontent.com/kadir-ince/Time-Zone-API/main/medias/berlin.png"/>

[You can see time zones full list on this link](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
