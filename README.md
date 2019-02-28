# ngram-db
_A fast in-memory ngram database, written in Golang._

## N-gram sets

This is a key-value data store of **n-gram sets**. Here, an n-gram set is a collection of n-gram counts corresponding to a corpus of text.

Users can create n-gram sets, add text to them, and then query the database to get counts and other information about n-grams in specific sets.

The goal of this project is to provide a simpler way to persist, process,
and retrieve n-gram data. 

### Running Server
This project can be built with:
```
go build
```

And then the database can be started with:
```
./ngramdb --server --port=3000
```

### Using Client

To query the database, you need to open a TCP connection to it.

There is a client library for connecting to the database included in this repository.

```
import "github.com/ngramdb"

client := ngramdb.client.New("localhost:3000")
response := client.send("ADD SET english(3);") 
```

### Query Syntax

#### Creating Sets
To add a new set of 3-grams:
```
ADD SET setName(3)
```

An example response:
```
{ "success": true }
```

#### Add text to a set
The server automatically breaks received text up into n-grams of the appropriate size. 

```
ADD TEXT "abcdef" TO setName
```

An example response:
```
{ "success": true }
```

#### All n-grams of a given size in a set
Each set represents a collection of n-grams. All the n-grams for a given
set can be retrieved with this query.

```
GET NGRAMS(3) IN setName
```

An example response:
```
{ 
  "success": true,
  "total": 100
  "ngrams": {
    "aaa": 70,
    "aab": 20,
    "aac": 10
  }
}
```

#### Single N-gram counts
Getting the count of an ngram in the set:
```
GET COUNT OF "abc" IN setName
```

An example response:
```
{ "success": true, count: 10 }
```

#### Single n-gram frequency

Getting the frequency of an ngram in the set:
```
GET FREQ OF "abc" IN setName
```

An example response:
```
{ 
  "success": true,
  "frequency": .75,
  "count": 75,
  "total": 100
}
```

#### Probable Completions
Completions are a more interesting application of n-grams. 

The query specifies an n-gram with some characters missing
(replaced with a placeholder `.`).

The server will indicate which characters are most likely,
assuming the string follows the same distribution as the 
others in the given set. 

This could be used, for example, to implement auto-complete for a search box.

```
GET COMPLETIONS OF "aa." IN setName
```

An example response:
```
{ 
  "success": true,
  "completions": [
    { 
      "ngram": "aaa",
      "probability": 0.7
    },
    
    { 
      "ngram": "aab",
      "probability": 0.2
    },
        
    { 
      "ngram": "aac",
      "probability": 0.1
    }
  ]
}
```

#### Probable Sets
Another application of n-grams is that you have unknown text,
and you want to know which set it likely belongs in.

You may have a set for English text and a set for French text.
You want to guess the language of new text, based on which n-gram set it is most similar to.

To find out which sets are most likely to contain the n-gram "aaa":
```
GET PROBABLE SETS OF "aaa"
```

An example response:
```
{ 
  "success": true,
  "sets": [
    { 
      "name": "english",
      "probability": 0.7
    },
    
    { 
      "name": "french",
      "probability": 0.2
    },
        
    { 
      "name": "spanish",
      "probability": 0.1
    }
  ]
}
```

#### Errors
Certain queries can result in errors. Below is a list of possible errors.
```
{ 
  "success": false, 
  "error": "SET_ALREADY_EXISTS",
  "message": "A set with the name 'setName' already exists"
}

{ 
  "success": false, 
  "error": "SET_NOT_EXIST",
  "message": "There is no set with the name 'setName'"
}
```
