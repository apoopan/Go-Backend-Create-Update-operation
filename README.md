# Go-Backend-Create-Update-operation
Implements 3 functions:

    Create User
    Get all users
    Update User information

To run:

  1. Run three Mongod processes on ports 27011, 27012 and 27013 (on separate terminals) -> mongo --port 
  2. go build
  3. go run .
 

Creating a 3-Node Replica Set Cluster (on a single machine): Referred from here. In addition to this,execute rs.slaveOk() on the secondary nodes.

In order to view the collections secondary nodes, authenticate by -

  use admin
  db.auth(username,password)
   
