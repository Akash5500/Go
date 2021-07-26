Rest API With GOLANG using GIN, GORM and MySql

Install packages:
	Enter the following command to install it in the project

	go get github.com/go-sql-driver/mysql
	go get github.com/gin-gonic/gin
	go get github.com/jinzhu/gorm

The folder Structure is as follows:
	Config
		Database.go
	Controller
		user.go
	Models
		User.go
		UserModel.go
	Routes
		Routes.go
	go.mod
	go.sum
	main.go

Database Setup:
	connection to the Datbase is in folder Config and Create a new database in mysql, then we add the database configuration in Database.go.

Create Model :

	let’s create a model for Hotel that consists following fields:


	Hotel_Id    
	Name        
	Country     
	Address     
	Latitude    
	Longitude   
	Description 
	Room_count  
	Currency    
	Amenities   

Configure Routing:

	Similarly, we create the routing for our project. We create a group to adjust our routing related to Hotel and we have to take care of the group 
	while consuming the api. So, let’s create Routes.go file inside Routes folder. We will see the implementation of the controller functions below.

Create Controller:

	We handle http requests coming from front end in controller. We create different functions that handles our specific requests routed to controller by our 
	router. We have User.go inside Models to interact with the database. We respond to the user according to the data we get from our database. If we get no error, 
	we supply response as StatusOK and if we get the error, we supply error status. We create User.go file inside Controllers folder.

	inside the controllers.Go file we have implentation for :
  		1. GetAllHotels
		2. GetHotelById
		3. CreateHotel
		4. UpdateHotel
		5. DeleteHotel
		
		you can see the implentation in Controlers.go inside Controller folder
		
Handle Requests:

This is the crucial file which fetch data and interacts directly with our database. We create User.go file inside Models folder to handle the database requests.

Setting up Server:

This is the starter function of our project. We connect mysql, auto migrate our modal and setup router form here. We have to create main.go in root of the project.


RESULTS:

Endpoints
These are the endpoints we will use to create, update, read and delete the user data.

	1. GET user-api/user → Retrieves all the user data
	2. POST user-api/user → Add new user data
	3.GET user-api/user/{id} → Retrieve the single user data	
	4.PUT user-api/user/{id} → Update the user data
	5.DELETE user-api/user → Delete the user data
		


