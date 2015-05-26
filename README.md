#MyCommander server application
#[MyCommander.info](http://mycommander.info/)

#Description
Golang server application

#Usage
After you download binary for your distribution, start the server

    ./my_commander_arch -config /path/to/config -voice /path/for/voice/config

On Android side, configure the IP of the running server. You can find your IP running
command on your computer

    ifconfig
or similar
. The IP must be entered in value like

    192.168.1.2

For using voice, you need to download libary

    sudo apt-get install libttspico0 libttspico-utils libttspico-data
If you dont want to use voice, dont start the program with voice configuration

    ./my_commander_arch -config /path/to/config
For more information, see the configuration section

#Configuration
###Standard configuration
    
    conf/cmd.cfg
    
    chrome=/usr/bin/google-chrome
    firefox=/usr/bin/firefox
Left side is a word that will match your said text on Android. If match, exec will happen on right sideSo, when you say 'chrome', golang will execute /usr/bin/google-chrome</p>
###Voice configuration

    conf/voice.cfg

    name=Eva
    open=opening
    close=closing
    
Same as standard configuration, but we have new config option that is 'name'
.Name is used when you say 'Who are you?'. The commander will have your Android username and phone model so you can interact with diffrent devices

#Compiling from source
There is makefile for so you can compile for target distribution.
For Linux 64bit (amd64)

    make amd64
For Linux 32bit (386)
    
    make 386
For Linux arm (arm)

    make arm

#Need help?
Drop the email at

    marin.basic02@gmail.com
Or open a issue at github
.You are free to commit new options, update documentation, etc..

#License
This library is under the MIT License
#Author
Marin Basic 
