# sensors
Sensors is event-driven microservice that will implement aggregate updating DB and emitting events. It is part of the project hosted in gowithdocker repo here. 

Sensors service will have API to update repository (CRUD)

Currently implementing CommandHandler: CUD commands from Command Queue(RabbitMQ) will update repository and emit sensor events

Todo:

Eventhandler will append events to EventStore(DynamoDB) also will provide endpoint to implement CQRS(Read requests)
