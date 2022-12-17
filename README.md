## Design consideration of microservices
- Data management: microservices should are designed to be stateless, meaning they should not rely on data stored in memory or as sessions objects. Instead, they should retrieve data directly from a database (e.g. mySQL) as needed. This helps ensure microservices can be scaled horizontally and can recover from failure without losing data. As such, the login feature of a typical ride-sharing platform will be omitted in the simulation, with direct access to the functionalities in the Console Application instead.

- Resilience: microservices should be designed to be resilient and self-healing, meaning they should be able to handle failure and continue operating without breaking the flow of the entire system. This was achieved through appropriate error handling and data validation.

- Passenger & Driver struct: For both of these two structs, a new variable "Username" was added, which is uneditable after creation. The purpose of this new addition is to act as a unique identifier for each Passenger and Driver Account and to estabilish a primary key in the database, since all other information in the account can be changed.

## Architecture diagram

## Instructions for setting up and running the microservices
1. Go to the 2 Github Links below and download the source code:
   1. https://github.com/s10205426/ETI_Assignment1_admin
   2. https://github.com/s10205426/ETI_Assignment1_microservices

2. Retrieve the SQL script from the first Github link to setup the database (Hostname: 127.0.0.1 & Port: 3307)

3. 
