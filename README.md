# dng_dropframe_fix
Fills in missing frames in BRAW / ML DNG sequences. Crude, but it works!

I guess I wasn't using a fast enough SD card in my BMPCC OG to record BRAW, so I ended up missing a bunch of frames in the Cinema DNG sequence, meaning neither Resolve or Premiere would accept them as "image sequences" or register them as clips. Doing traditional batch renaming meant that the gaps between missing frames would close, resulting in out-of-sync audio. This Go script fills the gaps in DNG sequences, with simple copy and pasting of the previous frame, which allows them to be understood as clips. 

Better than nothing!

## Usage:

Install [Go](https://golang.org/doc/install), if you haven't already. 

Download the .go file, and copy it into the affected/broken DNG sequence folder. Run the script in your terminal of choice with "go run dng_dropframe_fix.go"

Run the script over and over again until you no longer recieve messages about new files being created. Should only take 3 or 4 runs. 

The script assumes your first file has this as a suffix - 000000.dng, and counts up from there.

I shouldn't have to say this, but please backup your footage before you run this script. 

Inspired by: JETHRO_THE_POTHOLE_FILLER / Jethro's Pothole Filler, a VB script that does the same thing, but I couldn't get it to run on my Windows 10 machine. 
