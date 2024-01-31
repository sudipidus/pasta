#!bin/sh
process_name="$1"

# Use pkill to send a kill signal to the process
pkill -SIGUSR1 "$process_name"
