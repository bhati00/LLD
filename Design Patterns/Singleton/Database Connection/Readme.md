### 🧩 Problem Description: Database Connection Pool Manager   
 
    You are tasked with designing a Connection Pool Manager — a system that manages a fixed set of reusable database connections. Instead of creating a new connection every time the database is accessed (which is expensive), the system should allow connections to be reused efficiently, ensuring better performance, resource utilization, and thread safety.  

###  Functional Requirements  
- Singleton Instance

    There should be only one global instance of the connection pool in the application.

- Initialization with Fixed Pool Size

    On startup, the system should initialize N number of database connections and maintain them.

- Connection Acquisition

    - A method getConnection() should:

        - Return an available connection if one exists. 

        - Handle the situation gracefully if no connection is available (block, queue, or error).

- Connection Release

    - A method releaseConnection(conn) should:

    - Return the connection back to the pool after use.

- Thread Safety

    - Multiple threads should be able to request/release connections safely without race conditions.

## Solution 
let's try build up the solution by following some steps :

- ### *Nature of the problem* :   
    -  we'll consider this problem in the category of **resource pooling + concurrency control.**. Due to allocation of resources in multithreaded environment
    - Design pattern : if observe carefully, the problem requires us to have one golbally shared pool throughout the application and all clients must access the same instance. which is the perfect use case of  **singleton pattern**
    - life cycle management : We need to manage the lifecycle of connections. since they create, distroy or lost. 
    - **concurrency** : it requires us to handle thready safety   

- ### *Understanding requirements in depth* :    

    we always divide the requirements into the two parts functional and non functional    
    1. functional   
          
        - pool size : we need to ensure that pool size should be created on demand or eager loading 
        - get connection : we need to think about the behavior when all connections are in use
        - release connection : how connectinos should be released (manually or auto release)





