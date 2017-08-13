GOMSX: an MSX Emulator
======================

Why?
----

There are many emulators you can use to emulate a MSX computer. I'm not trying to code the
best emulator. This is a learning experience for me, to understand how the MSX computer works.

What's the best way of learning how something works? by building it, of course!!!

How to RUN it
-------------

If you want to try the emulator for youlself, you can download binaries here (only Linux AMD64):
https://drive.google.com/open?id=0B8__MZ9xDS79SkRpQzFneDVoOTA

Make sure you have the following libraries installed:
    - libsdl2
    - libsdl2-image
    - libsdl2-ttf

Bundled with the emulator is the C-BIOS rom file. It provides a free implementation of the
BIOS routines of the MSX. No BASIC.

If you want to run BASIC, you can find a MSX1.ROM system file elsewhere.

The file "softwaredb.xml" is useful in aiding the emulator to apply the correct memory mapper
for the MSX cartridge games. It's not required, but usually you can't play games without it.

To run it:

    $ chmod +x gomsx
    $ ./gomsx -cart game.rom

Help:

    $ ./gomsx -h

Compilation
-----------

So, you want to check the source, eh? First, you'll need a golang installation. In Ubuntu:

    $ sudo apt-get install golang

Go needs to have its workspace set up. Look for the documentation in https://golang.org/doc/

Next, clone the source:

    $ git clone https://github.com/pnegre/gomsx.git

Get the dependencies:

    $ go get github.com/pnegre/gogame
    $ git clone https://github.com/pnegre/z80

And, you just build the program:

    $ go build

Happy hacking!!!
