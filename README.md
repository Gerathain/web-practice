This is a basic webserver written in GO.

Instructions on usage:
This program is to automatically manage some philips hue lightbulbs. The program will read in the state of the bulbs into an internal struct to make it easier to work with and can then send out PUT requests to change the state of the bulbs

Currently the program will read in the state of the bulbs and then act on that.
If the bulbs are on, then they will be changed to a colour temperature of 200, if they are off then they will be turned on and the colour temp will not be changed
