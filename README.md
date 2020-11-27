Layer the code, instead of grouping it.

We try to keep packages on a given layer as independent as possible.
Imports can go down, but never up.
App imports business and foundation, business import foundation.
There can be also horizontal imports - inside the same layer. 

## First thing to focus

## Logging
Before we start any project, we have to understand what the purpose of our logs are.
You cannot just log as an insurance policy - too much noise. A nightmare.

- Do the logs represent data or is it to debug the app?
This can be the difference between structure log or just human readable log.
Better is to use it to debug the app.

- Logging isn't free. It causes allocation on the heap, which causes GC, which cause latency.
Think about it before logging it.

- Log everything to standard out. Not out job to manage disk or whatever.
That's a OPs job - to figure it out where to put it - CLI tools.

- Never use a singleton or global variable for our logger.
If anything else (apart from foundation) needs to log, we gonna use the right level of
precision for it: function parameters, closure, receivers, etc.

#### Advice:
Don't play with logging level - log everything you need, either in production, staging or development.
Have the discussion to realize what needs to be logged.

## Configuration

- Avoid configuration files - it creates a whole lot of work to manage them.

- Try to make default values work across dev, staging and production environments. 

### Global or package level variables
We should use with these policies in mind:
1 - Doesn't matter when the variable is initialized as long as it happens before main
2 - The initialisation does not require the configuration system
3 - the only code that can use the variable is the one where it is declared in.

### Policies

- Only the maingo file can import and hit the configuration system, whatever it is.
No other app, business or foundational package is allowed to do it.
All business and foundation must ask for their configuration with precision - parameters, config struct.

- Configuration is read once and then passed

- Configuration default has to be overwritten.
This has to work on command line with command line flags. The other way we are going to do is environment variables.
 

## App Layers

### App
App layer is at the top
Holds all the code that represents the application level concern
Presentation, CLI, UI, UX.

#### Policies
* Cannot log

### Business
The code that holds the business logic - the purpose - the problem we are trying to solve.
DB, external systems, data modeling.
Probably not reusable across multiple projects.
If it is being used, we ask on a macro level, why does this second project exists.
Why is it not part of this one.

#### Policies
* Can log

### Foundation
Code that is not business oriented.
Code that can be reusable.
Can be moved out to another project or repo - the kit repo.
They are not aloud to LOG - you should not set a pattern on how this code 
should be used.

#### Policies
* Cannot log


### Vendor

Third part libraries. We own in our project - we do not need to worry about people removing their code, etc.

## Handler Functions

Will do 3 things, and only 3 things - request, business, response

### Take request and validate
### Know what to do regarding business logic - the purpose of the handler
### Formulate a response - without error handling

We need also two different and very specific handlers.
The aliveness and the readness handlers.

### Readness handler
It tells our CI that we are ready to receive request

### Aliveness
In the case of this project, tells kubernets that the application is alive

### Loging on the handler function

We use the receiver approach.
Creating a struct that has a field of the types we need, we can create a value of that type
and use the handler function of the type at the same time.

