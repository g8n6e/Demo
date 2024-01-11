# Jurassic Park
Welcome to Jurassic Park, the safest and most fun place to store your dinosaurs. Feel free to use this API to track your dinosaurs.

## How to use this API

First make sure you have go installed, you can use this resource if you need assistance https://go.dev/doc/install
### Start the app
> cd $directoryLocationofAppCode  
> go run .

By default in debug it should listen on http://127.0.0.1:8080/ which you can use as the root url on all of the following RESTful API endpoints

# Dinosaur Commands

## GET /dinosaurs
### Get All Dinosaurs

## POST /dinosaur 
### Add a Dinosaur
example request body:
>{  
>    "name" : "testbronty",  
>    "speciesId" : 2,  
>    "cageId" : 1  
>}

## GET /dinosaur/:id
### Get Dinosaur by id

## PATCH /dinosaur/:id
This patch command will update one or many elements on a dinosaur

example request body:
>{  
>    "name" : "testbronty",  
>    "speciesId" : 2,  
>    "cageId" : 1  
>}

or limited attributes

>{  
>    "name" : "testbronty"  
>}

# Cage Commands

## GET /cages
### Get All Cages

## POST /cage
### Add a Cage
example request body:
>{  
>    "Capacity" : 2,  
>    "Active" : 1  
>}

## GET /cage/:id
### Get Cage by id

## PATCH /cage/:id
This patch command will update one or many elements on a cage

example request body:
>{  
>    "Capacity" : 2,  
>    "Active" : 1  
>}

or limited attributes

>{  
>    "Capacity" : 3  
>}

# Specie Commands

## GET /species
### Get All Species

## POST /specie
### Add a Specie
example request body:
>{  
>    "name":"raptor",  
>    "diet":1  
>}  

## GET /specie/:id
### Get Specie by id

## PATCH /specie/:id
This patch command will update one or many elements on a specie

example request body:
>{  
>    "name":"raptor",  
>    "diet":1  
>}  

or limited attributes

>{  
>    "diet" : 0  
>}

### Notes:
Diet is an Enumeration:
> Herbivore = 0  
> Carnivore = 1  

## Goals
- All requests should respond with the correct HTTP status codes and a response, if necessary, 
representing either the success or error conditions.
- Data should be persisted using some flavor of SQL.
- Each dinosaur must have a name.
- Each dinosaur is considered an herbivore or a carnivore, depending on its species.
- Carnivores can only be in a cage with other dinosaurs of the same species.
- Each dinosaur must have a species (See enumerated list below, feel free to add others).
- Herbivores cannot be in the same cage as carnivores.
- Use Carnivore dinosaurs like Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus.
- Use Herbivores like Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops.

## Future Objectives

### If I had more time

- Figure out how to get GORM to honor foreign keys correctly to be able to preload associations on objects and streamline validation of associated traits

- Clean up test cases and scaffold out some utilities to reduce duplication of code throughout testing

### Options for Scaling

 - create an api gateway to load balance requests across multiple service instances

 - replace sqlite with an always on sql cluster such as amazon aurora postgresql or mssql fronted with either a pub/sub or streaming (kafka) writer interface utizing CQRS where multiple service instances can asyncronously queue their requests to the db worker service

  - replace sqlite with a distributed db such as cockroach so multiple service instances can directly connect to db without creating connection/request congestion
