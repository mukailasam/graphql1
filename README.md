### GraphQL1
A Golang CRUD GraphQL API

### What is GraphQL
Wikipedia: GraphQL is an open-source data query and manipulation language for APIs and a query runtime engine.

GraphQL enables declarative data fetching where a client can specify exactly what data it needs from an API. Instead of multiple endpoints that return separate data, a GraphQL server exposes a single endpoint and responds with precisely the data a client asked for.

### Basic Concepts:
Schema: The foundation of a GraphQL server. It defines the types of data that can be queried and the relationships between them.

Types: GraphQL has scalar types (e.g., Int, String, Boolean) and complex types (e.g., Object types, Interface types, Union types).

Queries: Clients send queries to request specific data. A query describes the data structure the client expects in response.

Mutations: Used to modify data on the server. Similar to queries, but for operations that cause changes.


### Setup
+ Clone the repository\
`git clone https://github.com/mukailasam/graphql1`
+ Change directory into the program directory\
`cd graphql1`
+ Start the API server\
`go run .`

after starting the server start making use of the Queries

Note: There are 5 posts with the id(7, 8, 9, 10, 11)already created and they are stored in the sqlite database, but you can add yours by making use of create post query. besides creating post you can do all other stuff too, like reading a single post, all posts, updating a post and deleting a post. You can do all these stuff using the queries below

### Queries
Create a post\
`
http://127.0.0.1:8080/graphql?query=mutation+_{create(title:"First Post",body:"Hello, word"){title,body}}
`

Read all posts\
`
http://127.0.0.1:8080/graphql?query={posts{title,body}}
`

Read a Post\
`
http://127.0.0.1:8080/graphql?query={post(id:11){title,body}}
`

Update a post\
`
http://127.0.0.1:8080/graphql?query=mutation+_{update(id:11,title:"First Post",body:"Hey buddy"){title,body}}
`

Delete a post\
`
http://127.0.0.1:8080/graphql?query=mutation+_{delete(id:11){title,body}}
`
