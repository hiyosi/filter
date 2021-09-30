## SafetyCulture fork

Forking parent repo in order to remove usage of 'bson' and update to use go modules. Doing this allows us to install this module without usage of Bazaar and opening our GOVCS to bzr. Removing 'bson' removes support for REGEX filters but these are not support within our SCIM setup.

### Install

```
go get github.com/hiyosi/filter
```

### Usage

```.go
func main() {
	env := filter.Env{}
	s := new(filter.Scanner)
	s.Init("ham eq \"spam\"")
	statements := filter.Parse(s)
	query, err := filter.Evaluate(statements[0], env)
	if err != nil {
	        fmt.Printf("error")
	}

	fmt.Printf("%s\n", query)  // result is output as string.
 }
```

```
map[ham:spam]
```

As above, result data type is `map[string]interface{}`.

### Generate parser.go from parser.go.y

```
$ cd filter
$ make
```

### Implemented operators are follows

| status | Operator | Description                       |
| :----- | :------: | :-------------------------------- |
| ✓      |    eq    | equal                             |
| ✓      |    ne    | not equal                         |
| ✓      |    co    | contains                          |
| ✓      |    sw    | starts with                       |
| ✓      |    ew    | ends with                         |
| ✓      |    pr    | present (has value?               |
| ✓      |    gt    | greater than                      |
| ✓      |    ge    | greater than equal                |
| ✓      |    lt    | less than                         |
| ✓      |    le    | less than equal                   |
| ✓      |   and    | Logical and                       |
| ✓      |    or    | Logical or                        |
| ✓      |   not    | Not function                      |
| ✓      |    ()    | Precedence grouping               |
| ✓      |    []    | Complex attribute filter grouping |

- https://tools.ietf.org/html/draft-ietf-scim-api-14#section-3.2.2

### result example

- Input data

```
userName eq "bjensen"
name.familyName co "O'Malley"
userName sw "J"
title pr
(title pr)
meta.lastModified gt "2011-05-13T04:42:34Z"
meta.lastModified ge "2011-05-13T04:42:34Z"
meta.lastModified lt "2011-05-13T04:42:34Z"
meta.lastModified le "2011-05-13T04:42:34Z"
title pr and userType eq "Employee"
title pr or userType eq "Intern"
schemas eq "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
userType eq "Employee" and (emails co "example.com" or emails co "example.org")
userType eq "Employee" and ( emails co "example.com" or emails co "example.org")
userType eq "Employee" and (emails co "example.com" or emails co "example.org" )
userType eq "Employee" and ( emails co "example.com" or emails co "example.org" )
userType ne "Employee" and not ( emails co "example.com" or  emails co "example.org" )
userType eq "Employee" and ( emails.type eq "work" )
userType eq "Employee" and emails[type eq "work" and value co "@example.com"]
emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]
```

- result

```
map[userName:bjensen]
map[name.familyName:map[$regex:{O'Malley }]]
map[userName:map[$regex:{^J }]]
map[title:map[$exists:true]]
map[title:map[$exists:true]]
map[meta.lastModified:map[$gt:2011-05-13 04:42:34 +0000 UTC]]
map[meta.lastModified:map[$gte:2011-05-13 04:42:34 +0000 UTC]]
map[meta.lastModified:map[$lt:2011-05-13 04:42:34 +0000 UTC]]
map[meta.lastModified:map[$lte:2011-05-13 04:42:34 +0000 UTC]]
map[$and:[map[title:map[$exists:true]] map[userType:Employee]]]
map[$or:[map[title:map[$exists:true]] map[userType:Intern]]]
map[schemas:urn:ietf:params:scim:schemas:extension:enterprise:2.0:User]
map[$and:[map[userType:Employee] map[$or:[map[emails:map[$regex:{example\.com }]] map[emails:map[$regex:{example\.org }]]]]]]
map[$and:[map[userType:Employee] map[$or:[map[emails:map[$regex:{example\.com }]] map[emails:map[$regex:{example\.org }]]]]]]
map[$and:[map[userType:Employee] map[$or:[map[emails:map[$regex:{example\.com }]] map[emails:map[$regex:{example\.org }]]]]]]
map[$and:[map[userType:Employee] map[$or:[map[emails:map[$regex:{example\.com }]] map[emails:map[$regex:{example\.org }]]]]]]
map[$and:[map[userType:map[$ne:Employee]] map[$or:[map[emails:map[$not:map[$regex:{example\.com }]]] map[emails:map[$not:map[$regex:{example\.org }]]]]]]]
map[$and:[map[userType:Employee] map[emails.type:work]]]
map[$and:[map[userType:Employee] map[$and:[map[emails.type:work] map[emails.value:map[$regex:{@example\.com }]]]]]]
map[$or:[map[$and:[map[emails.type:work] map[emails.value:map[$regex:{@example\.com }]]]] map[$and:[map[ims.type:xmpp] map[ims.value:map[$regex:{@foo\.com }]]]]]]
```

### TODO

- Specified to attribute name with full path with schema URI.( To disambiguate duplicate names between schemas)
  - e.g. `filter=urn:ietf:params:scim:schemas:core:2.0:User:userName sw "J"`
- Evaluator do not use the Env{} structure.
- Do not use panic.
