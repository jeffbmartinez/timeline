# timeline
Store timeline events

There are two ways to store events with **timeline**

 - Send a *start* event followed by *stop* event. The response for the *single* event will include an id that must be provided to the *stop* event. Note that the *stop* event will not complain if you provide an invalid or non-existant event id.
 - Send a *single* event, which includes a duration.

# Endpoints

## /api/event/single

## /api/event/start

## /api/event/stop
