# dng_dropframe_fix
Fills in missing frames in BRAW / ML DNG sequences. Crude, but it works!

I guess I wasn't using a fast enough SD card in my BMPCC OG to record BRAW, so I ended up missing a bunch of frames in the Cinema DNG sequence, meaning neither Resolve or Premiere would accept them as "image sequences" or register them as clips. Doing traditional batch renaming meant that the video frames would close the gaps, resulting in out-of-sync audio. This Go script fills the gaps in DNG sequences, with simple copy and pasting of the previous frame, which allows them to be understood as clips. Better than nothing!

## Usage:

Drop the go script into the DNG sequence folder. Run the script in your terminal of choice with "go run dng_dropframe_fix.go"

Run the script over and over again until you no longer recieve messages about new files being created. Should only take 3 or 4 runs. 

Oh yea, you need to have Go installed.
