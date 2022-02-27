# Echo Service Template

Backend echo framework template for microservices.
Use postgresql for database.

third party library dependency
- labstack echo v4
- joho godotenv
- lib pq

## Usage

change `test.env` to `.env`, or create your own .env file and add it to godotenv.Overload arguments
PG___ is for database connection.

### Database setup

you can set your database in `psql` to setup database manually.
I recommended you to create role and schema specific to this service.
Use createdb and createuser or login to `psql` and write sql create statement.
User need permission to login, and connect to that database. Also, privillage of 
usage and create table in schema that have the same name with that user.

#### Create tables

To create table, you can write sql directly in the cli postgresql terminal, or
use .sql with my setup.go program in `dbSetup` folder to automate your create tasks.
This help you create and drop your tables quickly for table test. Furthermore, 
you can execute more than one sql files in order easily.

Write utilities sql file to automate in dbSetup folder.
In setup.go simply add or remove your own script files you want to run to seting up 
your database in `scripts` variable in main function "in order(start from 0 index)".

run the setup.go from the root of the modules.
fo run ./dbSetup/setup.go

### Database connection.

Set "PG___" in .env varibles for database connection setting.
The database connection code live in models/dbConfig.go

### Server settings

Add or use your echo middleware in `serverMiddlewares` function.
file urlMapping.go is for map your api endpoint and use handlers from 
package `handlers`.

### Data Models

Put your database models code in the package `models`.