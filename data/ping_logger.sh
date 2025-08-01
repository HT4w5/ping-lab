#!/bin/bash

# ==============================================================================
# SCRIPT: ping_logger.sh
# DESCRIPTION: Pings a specified host for one hour and logs the output.
# USAGE: ./ping_logger.sh <hostname>
# ==============================================================================

# Check if a hostname argument was provided.
if [ -z "$1" ]; then
    echo "Error: Missing hostname argument."
    echo "Usage: $0 <hostname>"
    exit 1
fi

# Assign the first argument to a variable for clarity.
HOSTNAME="$1"

# Create a timestamped filename for the log.
# The format is YYYY-MM-DD_HH-MM-SS.
TIMESTAMP=$(date +%Y-%m-%d_%H-%M-%S)
LOG_FILE="${HOSTNAME}-${TIMESTAMP}-ping.txt"

# Inform the user about the operation.
echo "Starting ping to ${HOSTNAME} for 1 hour."
echo "Output will be saved to ${LOG_FILE}"

# Run the ping command.
# -D: Print timestamps for each ping.
# -n: Numeric output only (no DNS lookup).
# -i 0.2: Interval of 0.2 seconds between pings.
# The 'timeout 3600' command will run the 'ping' command for exactly 3600 seconds (1 hour).
# The entire output is redirected to the log file.
timeout 3600 ping -D -n -i 0.2 "${HOSTNAME}" | tee "${LOG_FILE}"

# Check the exit status of the timeout command.
if [ $? -eq 124 ]; then
    echo "Ping completed successfully after 1 hour."
else
    echo "An error occurred during the ping command."
fi

echo "Script finished."
