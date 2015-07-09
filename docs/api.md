# API


## measurements

Measurements allow you to store measurements on the timeline. Measurements consist of a name, a value, and a series. Optionally, you can add one or more tags which can be used to filter measurements when you retrieve them.

`/api/measurement`

Methods:

 - GET
 - POST

Parameters:

 - name
 - value
 - tags

## events

*Event*s are conceptually ways to store the occurence of an event on the timeline. An event has a name, but no value associated with it. This is the primary difference from a *measurement*. As an example, storing "cpu-load" with value of 1.2 would be a *measurement* on the timeline, where the occurrence of "machine-restarted" (no value associated) would be an *event*.

`/api/event`

Methods:

 - GET
 - POST

Parameters:

 - name
 - tags
