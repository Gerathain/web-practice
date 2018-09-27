This is a basic webserver written in GO.

Instructions on usage:
This program is to automatically manage some philips hue lightbulbs. The program will read in the state of the bulbs into an internal struct to make it easier to work with and can then send out PUT requests to change the state of the bulbs

Currently the program will read in the state of the bulbs and then act on that.
If the bulbs are on, then they will be changed to a colour temperature that is slightly warmer, if they are off then nothing will happen 

Currently the intention for this program is to add in some user controls, so that the user can control the rate at which the lights get warmer, as well as the time at which they start getting warmer.
The program will also be automated in some fashion, currently, I just intent to set it as a cron job, meaning that the code will not need to change. Obviously, this will not work on non unix systems, but currently, I don't mind.
