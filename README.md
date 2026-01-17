# Dev Machine Proxy

A lightweight service discovery dashboard for development machines. Automatically discovers services running on your machine and presents them in a clean web UI.

Perfect for use with Netbird or other VPN solutions - access all your dev services from a single landing page.

## Features

- **Port scanning** - Detects all listening TCP ports via `/proc/net/tcp`
- **Docker integration** - Identifies containers and extracts names from images/labels
- **Project folder scanning** - Searches config files (docker-compose, .env, package.json, etc.) for port references
- **HTTP probing** - Connects to ports to detect HTTP services and extract titles/server info
- **Known port database** - Maps common ports (3000, 5432, 8080, etc.) to service names
- **System monitoring** - Real-time CPU and memory usage charts with top processes by CPU/memory
- **AI usage tracking** - Monitor Claude and Codex rate limit usage with forecasting (requires optional CLI tools)
- **Daily tasks** - Track recurring habits with streak counting and daily reset
- **Web terminal** - Built-in shell access via xterm.js
- **Themeable** - 11 color themes including Catppuccin, Dracula, Nord, and more
- **Customizable layout** - Drag-and-drop section ordering, show/hide sections
- **Auto-refresh** - Dashboard updates every 30 seconds (configurable)

## Installation

```bash
# Clone or download the repo
cd dev-machine-proxy

# Run the install script
./install.sh

# Or with a custom projects directory
PROJECTS_DIR=/path/to/projects ./install.sh
```

The install script will:
- Build the binary
- Install to `/usr/local/bin/`
- Set up a systemd user service
- Enable the service to start at boot

## Optional Dependencies

### AI Usage Tracking

The AI Usage section displays rate limit information for Claude and Codex. This feature requires external CLI tools to be installed and available in your PATH:

- **`claude-usage`** - For Claude Code usage tracking. Must output JSON with `five_hour` and `seven_day` utilization data.
- **`codex-usage`** - For Codex CLI usage tracking. Must output JSON with `rate_limit.primary_window` data.

If these tools are not installed, the AI Usage section will show an error message but the rest of the dashboard will function normally. You can hide the section entirely via the Settings page.

The monitor collects usage data every 5 minutes and maintains 7 days of history for forecasting. Data is stored in `~/.config/dev-machine-proxy/usage-history.json`.

## Usage

After installation, the dashboard is available at: **http://localhost:9999**

### Command Line Options

```bash
dev-machine-proxy [options]

Options:
  -port int
        Port to serve the dashboard on (default 9999)
  -projects string
        Directory containing project folders to scan for port references
  -refresh duration
        How often to refresh service discovery (default 30s)
```

### Examples

```bash
# Basic usage
dev-machine-proxy

# Custom port and projects directory
dev-machine-proxy -port 8888 -projects ~/code

# Faster refresh interval
dev-machine-proxy -refresh 10s -projects ~/Projects
```

## Service Management

```bash
# Check status
systemctl --user status dev-machine-proxy

# Restart
systemctl --user restart dev-machine-proxy

# Stop
systemctl --user stop dev-machine-proxy

# View logs
journalctl --user -u dev-machine-proxy -f

# Disable autostart
systemctl --user disable dev-machine-proxy
```

## How Discovery Works

Services are identified using multiple sources, in priority order:

1. **Docker containers** - If a port belongs to a Docker container:
   - Name comes from `com.docker.compose.service` label, `org.opencontainers.image.title` label, or the image name
   - Project path comes from `com.docker.compose.project.working_dir` label

2. **Project folder scanning** - For non-Docker services, scans the projects directory for config files containing port references:
   - `docker-compose.yml`, `compose.yaml`
   - `.env`, `.env.local`, `.env.development`
   - `package.json`, `Makefile`, `Dockerfile`
   - `config.json`, `config.yaml`, `appsettings.json`

3. **HTTP probing** - Connects to the port and checks:
   - HTML `<title>` tag for app names (Grafana, Prometheus, etc.)
   - `Server` header for framework detection (Express, Flask, etc.)

4. **Known ports** - Falls back to common port conventions (5432=PostgreSQL, etc.)

5. **Process name** - Uses the process name from `/proc` as a last resort

## Configuration

### Service Configuration

To change command-line options after installation, edit the service file:

```bash
# Edit the service
nano ~/.config/systemd/user/dev-machine-proxy.service

# Reload and restart
systemctl --user daemon-reload
systemctl --user restart dev-machine-proxy
```

### Dashboard Settings

Access the Settings page via the gear icon in the dashboard header. You can configure:

- **Theme** - Choose from 11 color themes
- **Terminal font** - Custom font-family for the terminal
- **Section visibility** - Show/hide individual dashboard sections
- **Section order** - Drag and drop to reorder sections

Settings are stored in `~/.config/dev-machine-proxy/config.json`.

### Data Storage

The application stores data in `~/.config/dev-machine-proxy/`:

- `config.json` - Dashboard settings (theme, sections, etc.)
- `daily-tasks.json` - Daily tasks and completion history
- `usage-history.json` - AI usage metrics history (7 days)

## Updating

To update to a new version:

```bash
cd dev-machine-proxy
git pull  # if using git
./install.sh
```

The install script handles stopping the service, rebuilding, and restarting.

## Uninstalling

```bash
# Stop and disable the service
systemctl --user stop dev-machine-proxy
systemctl --user disable dev-machine-proxy

# Remove files
rm ~/.config/systemd/user/dev-machine-proxy.service
sudo rm /usr/local/bin/dev-machine-proxy

# Reload systemd
systemctl --user daemon-reload
```

## Building from Source

```bash
# Requires Go 1.21+
go build -o dev-machine-proxy .
```

## FAQ

### How do I run this on a remote or public server?

**You don't.**

This application is designed exclusively for local development machines accessed via a VPN like Netbird, Tailscale, or similar. It is **not** intended for public internet exposure and lacks the security features required for such deployment.

**Why this is dangerous on a public server:**

1. **No authentication** - Anyone with network access can view all your running services
2. **Remote terminal access** - The built-in terminal provides full shell access to the machine with the same privileges as the running process. On a public server, this is equivalent to leaving an SSH port open with no password.
3. **Service enumeration** - Exposes detailed information about your infrastructure that could be used for reconnaissance
4. **No TLS** - All traffic is unencrypted HTTP

**Safe usage patterns:**

- Run on your development laptop/desktop
- Access via VPN (Netbird, Tailscale, WireGuard, etc.)
- Bind to localhost and proxy through an authenticated reverse proxy if needed
- Use firewall rules to restrict access to trusted networks only

**If you need something for production:**

Consider purpose-built tools like:
- **Service mesh dashboards** (Istio, Linkerd)
- **Container orchestration UIs** (Portainer, Rancher)
- **Monitoring stacks** (Grafana + Prometheus)

All of which have proper authentication, authorization, and security features.

### What's the Terminal feature?

The Terminal section provides a web-based shell using xterm.js. It connects via WebSocket to a PTY on the host machine, giving you the same access as an SSH session. This is convenient for quick commands without switching windows, but reinforces why this tool should never be exposed publicly.

### Can I disable the terminal?

Yes - go to Settings (gear icon) and uncheck "Terminal" in the section visibility options. You can also hide any other section you don't need.

### Why does service X show as "unknown"?

The discovery system works in priority order: Docker labels > project scanning > HTTP probing > known ports > process name. If a service shows as unknown, it likely:
- Is not running in Docker
- Has no config files with port references in your projects directory
- Doesn't respond to HTTP probes
- Uses a non-standard port

You can improve detection by ensuring Docker containers have proper labels or by organizing your projects in the scanned directory.

### What are the AI Usage forecasts?

The AI Usage section tracks your Claude and Codex rate limit consumption over time. Based on your usage pattern over the last few hours, it forecasts:

- **Projected usage at reset** - What percentage you'll likely be at when the window resets
- **Time to exhaustion** - If you're on pace to hit 100%, when that will happen
- **Status pill** - "On track" (green) if you'll be fine, "At risk" (orange) if you might hit the limit

This helps you pace your AI tool usage throughout the day.

### How do Daily Tasks work?

Daily Tasks is a simple habit tracker built into the dashboard:

- Add tasks you want to complete every day
- Check them off as you complete them
- Tasks reset automatically at midnight
- Streaks track consecutive days of completion
- "Current streak" shows your active streak, "Longest streak" shows your record

Task data persists in `~/.config/dev-machine-proxy/daily-tasks.json`.

## License

MIT
