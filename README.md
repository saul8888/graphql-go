### Graphql in go

**compile the code**

`go run server.go`

**getUser**

```

    query getUser{
        getUser(id:<id>) {
            password
    	    customer_name
  		    email
  		    phone_number
  		    password
  		    created_at
  		    update_at
        }
    }
```

**getTotalUser**

```

    query getTotalUser{
        getTotalUser(input:{limit:<int>,offset:<int>}) {
			customers{
                password
    		    customer_name
  			    email
  			    phone_number
  			    password
  			    created_at
  			    update_at
            }
  		    cant
        }
    }

```

**insertUser**

```

    mutation insertUser {
        insertUser(input:{customer_name:<string>,email:<string>,phone_number:<string>,password:<string>}) {
    	    password
    	    customer_name
  		    email
  		    phone_number
  		    password
  		    created_at
  		    update_at
        }
    }

```
**updateUser**

```

    mutation updateUser {
        updateUser(id:<id>,input:{email:<string>,phone_number:<string>,password:<string>}) {
    	    password
    	    customer_name
  		    email
  		    phone_number
  		    password
  		    created_at
  		    update_at
  }
}

```
**deleteUser**


```

    mutation deleteUser{
        deleteUser(id:<id>) 
    }

```


