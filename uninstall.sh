#!/bin/bash
echo "Uninstalling ProcPipe..."
if [ -f "/usr/local/bin/procpipe" ]; then
    sudo rm /usr/local/bin/procpipe
    echo "✅ Uninstalled from /usr/local/bin/procpipe"
else
    echo "⚠️  ProcPipe not found in /usr/local/bin"
fi
