#!/bin/bash
set -e

# Dev Machine Proxy - Install/Update Script

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INSTALL_PATH="$HOME/.local/bin/dev-machine-proxy"
SERVICE_NAME="dev-machine-proxy@${USER}.service"
PROJECTS_DIR="${PROJECTS_DIR:-$HOME/Projects}"

echo "=== Dev Machine Proxy Installer ==="
echo ""

# Check for Go
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

# Build the binary
echo "Building dev-machine-proxy..."
cd "$SCRIPT_DIR"
go build -o dev-machine-proxy .
echo "Build complete."

# Stop existing service if running
if systemctl --user is-active --quiet "$SERVICE_NAME" 2>/dev/null; then
    echo "Stopping existing service..."
    systemctl --user stop "$SERVICE_NAME"
fi

# Install binary
echo "Installing binary to $INSTALL_PATH..."
mkdir -p "$(dirname "$INSTALL_PATH")"
cp dev-machine-proxy "$INSTALL_PATH"
chmod 755 "$INSTALL_PATH"

# Create user systemd directory if it doesn't exist
mkdir -p ~/.config/systemd/user

# Generate service file with correct paths
echo "Installing systemd user service..."
cat > ~/.config/systemd/user/dev-machine-proxy.service << EOF
[Unit]
Description=Dev Machine Proxy - Service Discovery Dashboard
After=network.target

[Service]
Type=simple
ExecStart=$INSTALL_PATH -projects $PROJECTS_DIR
Environment=PATH=/usr/local/bin:/usr/bin:/bin:$HOME/.local/bin
Restart=on-failure
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=default.target
EOF

# Reload systemd
echo "Reloading systemd..."
systemctl --user daemon-reload

# Enable and start service
echo "Enabling and starting service..."
systemctl --user enable dev-machine-proxy.service
systemctl --user start dev-machine-proxy.service

# Enable lingering so service runs without login (optional, requires sudo)
if command -v loginctl &> /dev/null; then
    echo "Enabling user lingering (service will run at boot)..."
    loginctl enable-linger "$USER" 2>/dev/null || echo "Note: Could not enable lingering (may require sudo). Service will only run when logged in."
fi

echo ""
echo "=== Installation Complete ==="
echo ""
echo "Service status:"
systemctl --user status dev-machine-proxy.service --no-pager || true
echo ""
echo "Dashboard available at: http://localhost:9999"
echo ""
echo "Useful commands:"
echo "  systemctl --user status dev-machine-proxy    # Check status"
echo "  systemctl --user restart dev-machine-proxy   # Restart"
echo "  systemctl --user stop dev-machine-proxy      # Stop"
echo "  journalctl --user -u dev-machine-proxy -f    # View logs"
echo ""
echo "To change the projects directory, edit:"
echo "  ~/.config/systemd/user/dev-machine-proxy.service"
echo "Then run: systemctl --user daemon-reload && systemctl --user restart dev-machine-proxy"
