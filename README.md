# timeline
Store data points or events in a timeline

There are two ways to store points with **timeline**

 - Send a *start* point followed by *stop* point. The response for the *start* point will include an id that must be provided to the *stop* point. Note that the *stop* point will not complain if you provide an invalid or non-existant start point id.
 - Send a *simple* data point, the difference is this has no followup *stop* point. If desired, you can always add a *duration* field to this.

The only required field for either of these is the *series* field. It will keep all of the points in the same series "grouped" together.

# Endpoints

## /api/point/simple

## /api/point/start

## /api/point/stop

# How to use

## Just use http get requests

Example:

    curl 'example.com/api/point/simple?series=temperatures&location=living-room&value=70'
