This is a basic webserver written in GO.

Instructions on usage:
This program is to automatically manage some philips hue lightbulbs. The program will read in the state of the bulbs into an internal struct to make it easier to work with and can then send out PUT requests to change the state of the bulbs

Currently the program will read in the state of the bulbs and then act on that.
If the bulbs are on, then they will be changed to a colour temperature of 200, if they are off then they will be turned on and the colour temp will not be changed

Currently, the intention for this program is that it will read the state of the lights and, if they are on, then they will get gradually warmer white throughout the evening. If the lights are not on, then the program will keep track of what colour temperature the lights should be so that they can be corrected when they are turned on.
Due to the fact that Philips Hue light bulbs do not allow their colour temperature to be changed whilst the lights are off, this will involve the program storing a copy of the state of the lights and working on than until the lights are on.
Currently, the program will simply have to poll to determin when the lights are turned on and then correct them, as I do not know of a way to detect when the lights are turned on. Either that, or the user will have to turn the lights on via this program, however this functionality is currently not itended to be developed.
