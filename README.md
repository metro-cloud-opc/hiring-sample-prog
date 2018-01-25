# hiring-sample-prog

1. L53-57: No look-up in the map
    - L54: `key` and `map[key].tinyName` are the same -- unnecessary access
    - L55-56: Not breaking from the loop
2. L39: If multiple URLs come at the same second, all but the last will be lost
3. L38: missing check: same URL, if sent multiple times, will create multiple entries in the map
4. L25: Error check wrong (in `ListenAndServe`)
5. Program-wide issue: hashmap is not thread-safe
6. Program-wide issue: Variable namings/style -- should be criticised because inconsistent, bonus for criticizing L39-41 name: `__`

## If they are done quickly:

ask them how they would add a feature of "desired-tiny-name" for a given URL -- and which part(s) would they modify. Questions to be asked not by interviewer but by candidate might include:
- where would the name come from
- what validations

Once identified, we suggest that they write some pseudocode in the locations they envisioned.

Somewhat exemplary implementation:

```go
	if reqTinyName := req.PostFormValue("tinyName"); reqTinyName != "" {
		if bookmarks[reqTinyName] != nil {
			http.Error(w, "Already exists", 400)
			return
		}
		__ = reqTinyName
	}
```
