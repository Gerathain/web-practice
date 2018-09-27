This is a basic webserver written in GO.

### Current behaviour
Currently the program will read in the state of the bulbs and then act on that.
If the bulbs are on, then they will be changed to a colour temperature that is slightly warmer, if they are off then nothing will happen 

### Future behaviour
Currently the intention for this program is to add in some user controls, so that the user can control the rate at which the lights get warmer, as well as the time at which they start getting warmer.
The program will also be automated in some fashion, currently, I just intent to set it as a cron job, meaning that the code will not need to change. Obviously, this will not work on non unix systems, but currently, I don't mind.
