package web

const themesCSS = `
/* Theme variables */
:root {
    /* Cyberpunk (default) */
    --bg-primary: #1a1a2e;
    --bg-secondary: #16213e;
    --bg-card: rgba(255, 255, 255, 0.05);
    --border-color: rgba(255, 255, 255, 0.1);
    --text-primary: #e4e4e4;
    --text-secondary: #aaa;
    --text-muted: #666;
    --accent-primary: #00d9ff;
    --accent-secondary: #00ff88;
    --accent-gradient: linear-gradient(180deg, #00d9ff, #00ff88);
    /* Default tag colors */
    --tag-docker-bg: rgba(13, 183, 237, 0.2);
    --tag-docker-text: #0db7ed;
    --tag-http-bg: rgba(0, 255, 136, 0.2);
    --tag-http-text: #00ff88;
    --tag-project-bg: rgba(255, 193, 7, 0.2);
    --tag-project-text: #ffc107;
    --tag-known-bg: rgba(156, 39, 176, 0.2);
    --tag-known-text: #ce93d8;
    /* Card and port badge defaults */
    --card-bg: var(--bg-card);
    --port-bg: rgba(0, 217, 255, 0.1);
    --port-text: #00d9ff;
}

[data-theme="catppuccin-mocha"] {
    --bg-primary: #1e1e2e;
    --bg-secondary: #181825;
    --bg-card: rgba(203, 166, 247, 0.05);
    --border-color: rgba(203, 166, 247, 0.15);
    --text-primary: #cdd6f4;
    --text-secondary: #bac2de;
    --text-muted: #6c7086;
    --accent-primary: #cba6f7;
    --accent-secondary: #f5c2e7;
    --accent-gradient: linear-gradient(180deg, #cba6f7, #f5c2e7);
    /* Catppuccin Mocha extended palette */
    --ctp-rosewater: #f5e0dc;
    --ctp-flamingo: #f2cdcd;
    --ctp-pink: #f5c2e7;
    --ctp-mauve: #cba6f7;
    --ctp-red: #f38ba8;
    --ctp-maroon: #eba0ac;
    --ctp-peach: #fab387;
    --ctp-yellow: #f9e2af;
    --ctp-green: #a6e3a1;
    --ctp-teal: #94e2d5;
    --ctp-sky: #89dceb;
    --ctp-sapphire: #74c7ec;
    --ctp-blue: #89b4fa;
    --ctp-lavender: #b4befe;
    --ctp-surface0: #313244;
    --ctp-surface1: #45475a;
    --ctp-surface2: #585b70;
    /* Tag colors */
    --tag-docker-bg: rgba(137, 220, 235, 0.2);
    --tag-docker-text: #89dceb;
    --tag-http-bg: rgba(166, 227, 161, 0.2);
    --tag-http-text: #a6e3a1;
    --tag-project-bg: rgba(249, 226, 175, 0.2);
    --tag-project-text: #f9e2af;
    --tag-known-bg: rgba(203, 166, 247, 0.2);
    --tag-known-text: #cba6f7;
    /* Card and port badge */
    --card-bg: #313244;
    --port-bg: rgba(137, 180, 250, 0.15);
    --port-text: #89b4fa;
}

[data-theme="catppuccin-macchiato"] {
    --bg-primary: #24273a;
    --bg-secondary: #1e2030;
    --bg-card: rgba(198, 160, 246, 0.05);
    --border-color: rgba(198, 160, 246, 0.15);
    --text-primary: #cad3f5;
    --text-secondary: #b8c0e0;
    --text-muted: #6e738d;
    --accent-primary: #c6a0f6;
    --accent-secondary: #f5bde6;
    --accent-gradient: linear-gradient(180deg, #c6a0f6, #f5bde6);
    /* Catppuccin Macchiato extended palette */
    --ctp-rosewater: #f4dbd6;
    --ctp-flamingo: #f0c6c6;
    --ctp-pink: #f5bde6;
    --ctp-mauve: #c6a0f6;
    --ctp-red: #ed8796;
    --ctp-maroon: #ee99a0;
    --ctp-peach: #f5a97f;
    --ctp-yellow: #eed49f;
    --ctp-green: #a6da95;
    --ctp-teal: #8bd5ca;
    --ctp-sky: #91d7e3;
    --ctp-sapphire: #7dc4e4;
    --ctp-blue: #8aadf4;
    --ctp-lavender: #b7bdf8;
    --ctp-surface0: #363a4f;
    --ctp-surface1: #494d64;
    --ctp-surface2: #5b6078;
    /* Tag colors */
    --tag-docker-bg: rgba(145, 215, 227, 0.2);
    --tag-docker-text: #91d7e3;
    --tag-http-bg: rgba(166, 218, 149, 0.2);
    --tag-http-text: #a6da95;
    --tag-project-bg: rgba(238, 212, 159, 0.2);
    --tag-project-text: #eed49f;
    --tag-known-bg: rgba(198, 160, 246, 0.2);
    --tag-known-text: #c6a0f6;
    /* Card and port badge */
    --card-bg: #363a4f;
    --port-bg: rgba(138, 173, 244, 0.15);
    --port-text: #8aadf4;
}

[data-theme="solarized-dark"] {
    --bg-primary: #002b36;
    --bg-secondary: #073642;
    --bg-card: rgba(147, 161, 161, 0.05);
    --border-color: rgba(147, 161, 161, 0.1);
    --text-primary: #839496;
    --text-secondary: #93a1a1;
    --text-muted: #586e75;
    --accent-primary: #268bd2;
    --accent-secondary: #2aa198;
    --accent-gradient: linear-gradient(180deg, #268bd2, #2aa198);
}

[data-theme="solarized-light"] {
    --bg-primary: #fdf6e3;
    --bg-secondary: #eee8d5;
    --bg-card: rgba(101, 123, 131, 0.08);
    --border-color: rgba(101, 123, 131, 0.15);
    --text-primary: #657b83;
    --text-secondary: #586e75;
    --text-muted: #93a1a1;
    --accent-primary: #268bd2;
    --accent-secondary: #2aa198;
    --accent-gradient: linear-gradient(180deg, #268bd2, #2aa198);
}

[data-theme="dracula"] {
    --bg-primary: #282a36;
    --bg-secondary: #1e1f29;
    --bg-card: rgba(189, 147, 249, 0.05);
    --border-color: rgba(189, 147, 249, 0.1);
    --text-primary: #f8f8f2;
    --text-secondary: #c0c0c0;
    --text-muted: #6272a4;
    --accent-primary: #bd93f9;
    --accent-secondary: #ff79c6;
    --accent-gradient: linear-gradient(180deg, #bd93f9, #ff79c6);
}

[data-theme="nord"] {
    --bg-primary: #2e3440;
    --bg-secondary: #3b4252;
    --bg-card: rgba(136, 192, 208, 0.05);
    --border-color: rgba(136, 192, 208, 0.1);
    --text-primary: #eceff4;
    --text-secondary: #d8dee9;
    --text-muted: #4c566a;
    --accent-primary: #88c0d0;
    --accent-secondary: #81a1c1;
    --accent-gradient: linear-gradient(180deg, #88c0d0, #81a1c1);
}

[data-theme="gruvbox-dark"] {
    --bg-primary: #282828;
    --bg-secondary: #1d2021;
    --bg-card: rgba(251, 184, 108, 0.05);
    --border-color: rgba(251, 184, 108, 0.1);
    --text-primary: #ebdbb2;
    --text-secondary: #d5c4a1;
    --text-muted: #665c54;
    --accent-primary: #fabd2f;
    --accent-secondary: #fe8019;
    --accent-gradient: linear-gradient(180deg, #fabd2f, #fe8019);
}

[data-theme="tokyo-night"] {
    --bg-primary: #1a1b26;
    --bg-secondary: #16161e;
    --bg-card: rgba(122, 162, 247, 0.05);
    --border-color: rgba(122, 162, 247, 0.1);
    --text-primary: #c0caf5;
    --text-secondary: #a9b1d6;
    --text-muted: #565f89;
    --accent-primary: #7aa2f7;
    --accent-secondary: #bb9af7;
    --accent-gradient: linear-gradient(180deg, #7aa2f7, #bb9af7);
}

[data-theme="one-dark"] {
    --bg-primary: #282c34;
    --bg-secondary: #21252b;
    --bg-card: rgba(97, 175, 239, 0.05);
    --border-color: rgba(97, 175, 239, 0.1);
    --text-primary: #abb2bf;
    --text-secondary: #9da5b4;
    --text-muted: #5c6370;
    --accent-primary: #61afef;
    --accent-secondary: #c678dd;
    --accent-gradient: linear-gradient(180deg, #61afef, #c678dd);
}

[data-theme="monokai"] {
    --bg-primary: #272822;
    --bg-secondary: #1e1f1c;
    --bg-card: rgba(166, 226, 46, 0.05);
    --border-color: rgba(166, 226, 46, 0.1);
    --text-primary: #f8f8f2;
    --text-secondary: #cfcfc2;
    --text-muted: #75715e;
    --accent-primary: #a6e22e;
    --accent-secondary: #f92672;
    --accent-gradient: linear-gradient(180deg, #a6e22e, #f92672);
}
`

const baseCSS = `
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
    background: linear-gradient(135deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
    min-height: 100vh;
    color: var(--text-primary);
    padding: 2rem;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
}

header {
    text-align: center;
    margin-bottom: 2rem;
    position: relative;
}

h1 {
    font-size: 2.5rem;
    font-weight: 300;
    color: var(--accent-primary);
    margin-bottom: 0.5rem;
}

.subtitle {
    color: var(--text-muted);
    font-size: 0.9rem;
}

.refresh-info {
    color: var(--text-muted);
    font-size: 0.8rem;
    margin-top: 0.5rem;
}

.config-link {
    position: absolute;
    top: 0;
    right: 0;
    color: var(--text-muted);
    text-decoration: none;
    padding: 0.5rem 1rem;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    transition: all 0.3s ease;
    font-size: 0.9rem;
}

.config-link:hover {
    color: var(--accent-primary);
    border-color: var(--accent-primary);
}

.services-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1.5rem;
}

.service-card {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 1.5rem;
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
}

.service-card:hover {
    transform: translateY(-4px);
    border-color: color-mix(in srgb, var(--accent-primary) 30%, transparent);
    box-shadow: 0 8px 32px color-mix(in srgb, var(--accent-primary) 10%, transparent);
}

.service-card.http {
    cursor: pointer;
}

.service-card.http::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 4px;
    height: 100%;
    background: var(--accent-gradient);
}

.service-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 1rem;
}

.service-name {
    font-size: 1.25rem;
    font-weight: 600;
    color: var(--text-primary);
}

.service-port {
    font-family: 'Courier New', monospace;
    font-size: 1.1rem;
    color: var(--port-text);
    background: var(--port-bg);
    padding: 0.25rem 0.75rem;
    border-radius: 6px;
}

.service-details {
    font-size: 0.85rem;
    color: var(--text-secondary);
    margin-bottom: 1rem;
}

.service-details p {
    margin: 0.25rem 0;
}

.service-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
}

.tag {
    font-size: 0.7rem;
    padding: 0.2rem 0.6rem;
    border-radius: 4px;
    text-transform: uppercase;
    font-weight: 600;
    letter-spacing: 0.5px;
}

.tag.docker {
    background: var(--tag-docker-bg);
    color: var(--tag-docker-text);
}

.tag.http {
    background: var(--tag-http-bg);
    color: var(--tag-http-text);
}

.tag.project {
    background: var(--tag-project-bg);
    color: var(--tag-project-text);
}

.tag.known-port {
    background: var(--tag-known-bg);
    color: var(--tag-known-text);
}

.empty-state {
    text-align: center;
    padding: 4rem 2rem;
    color: var(--text-muted);
}

.loading {
    text-align: center;
    padding: 4rem 2rem;
}

.spinner {
    width: 40px;
    height: 40px;
    border: 3px solid color-mix(in srgb, var(--accent-primary) 20%, transparent);
    border-top-color: var(--accent-primary);
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto 1rem;
}

@keyframes spin {
    to { transform: rotate(360deg); }
}

.source-badge {
    font-size: 0.65rem;
    color: var(--text-muted);
    text-transform: uppercase;
    letter-spacing: 1px;
}

.stats-container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.5rem;
}

.section-subtitle {
    font-size: 0.85rem;
    color: var(--text-muted);
    margin-left: 1.75rem;
    margin-top: -0.5rem;
    margin-bottom: 1rem;
}

.stat-card {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 1.25rem;
}

.stat-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 0.75rem;
}

.stat-title {
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.stat-value {
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--accent-primary);
}

.stat-chart {
    height: 80px;
    position: relative;
}

.stat-chart canvas {
    width: 100%;
    height: 100%;
}

.section {
    margin-bottom: 1rem;
}

.section-header {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    cursor: pointer;
    padding: 0.75rem 0;
    user-select: none;
}

.section-header:hover .section-title {
    color: var(--accent-primary);
}

.section-toggle {
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--text-muted);
    transition: transform 0.3s ease;
}

.section-toggle svg {
    width: 16px;
    height: 16px;
}

.section.collapsed .section-toggle {
    transform: rotate(-90deg);
}

.section-title {
    font-size: 1.1rem;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 1px;
    transition: color 0.3s ease;
}

.section-content {
    overflow: hidden;
    transition: max-height 0.3s ease, opacity 0.3s ease;
    max-height: 2000px;
    opacity: 1;
}

.section.collapsed .section-content {
    max-height: 0;
    opacity: 0;
}

.terminal-container {
    background: #1a1a2e;
    border-radius: 8px;
    padding: 0.5rem;
    min-height: 300px;
}

#terminal {
    height: 350px;
}
`

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dev Machine Services</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@xterm/xterm@5.5.0/css/xterm.css">
    {{CUSTOM_HEAD_HTML}}
    <style>
` + themesCSS + baseCSS + `
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1 id="page-title">Dev Machine Services</h1>
            <a href="/config" class="config-link">Settings</a>
        </header>

        <div class="section" id="performance-section">
            <div class="section-header" onclick="toggleSection('performance-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">Performance</span>
            </div>
            <p class="section-subtitle">Real-time CPU and memory usage (2 minute history)</p>
            <div class="section-content">
                <div class="stats-container">
                    <div class="stat-card">
                        <div class="stat-header">
                            <span class="stat-title">CPU Usage</span>
                            <span class="stat-value" id="cpu-value">--%</span>
                        </div>
                        <div class="stat-chart">
                            <canvas id="cpu-chart"></canvas>
                        </div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-header">
                            <span class="stat-title">Memory Usage</span>
                            <span class="stat-value" id="mem-value">--%</span>
                        </div>
                        <div class="stat-chart">
                            <canvas id="mem-chart"></canvas>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="section" id="services-section">
            <div class="section-header" onclick="toggleSection('services-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">Discovered Services</span>
            </div>
            <p class="section-subtitle">Auto-discovered services running on this machine <span class="refresh-info" id="refresh-info">(refreshes every 30 seconds)</span></p>
            <div class="section-content">
                <div id="services" class="services-grid">
                    <div class="loading">
                        <div class="spinner"></div>
                        <p>Discovering services...</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="section" id="terminal-section">
            <div class="section-header" onclick="toggleSection('terminal-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">Terminal</span>
            </div>
            <p class="section-subtitle">Remote shell access to this machine</p>
            <div class="section-content">
                <div class="terminal-container">
                    <div id="terminal"></div>
                </div>
            </div>
        </div>
    </div>

    <script src="https://cdn.jsdelivr.net/npm/@xterm/xterm@5.5.0/lib/xterm.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/@xterm/addon-fit@0.10.0/lib/addon-fit.min.js"></script>
    <script>
        let refreshInterval = 30000;

        let terminalFont = 'monospace';

        async function loadConfig() {
            try {
                const response = await fetch('/api/config');
                const config = await response.json();

                document.body.setAttribute('data-theme', config.theme);
                document.getElementById('page-title').textContent = config.title;
                document.title = config.title;

                refreshInterval = config.refreshInterval * 1000;
                document.getElementById('refresh-info').textContent =
                    '(refreshes every ' + config.refreshInterval + ' seconds)';

                // Set terminal font from config
                if (config.terminalFont) {
                    terminalFont = '"' + config.terminalFont + '", monospace';
                }

                return config;
            } catch (error) {
                console.error('Failed to load config:', error);
            }
        }

        async function loadServices() {
            try {
                const response = await fetch('/api/services');
                const services = await response.json();
                renderServices(services);
            } catch (error) {
                console.error('Failed to load services:', error);
                document.getElementById('services').innerHTML = ` + "`" + `
                    <div class="empty-state">
                        <p>Failed to load services</p>
                        <p style="font-size: 0.8rem; margin-top: 0.5rem;">${error.message}</p>
                    </div>
                ` + "`" + `;
            }
        }

        function renderServices(services) {
            const container = document.getElementById('services');

            if (!services || services.length === 0) {
                container.innerHTML = ` + "`" + `
                    <div class="empty-state">
                        <p>No services discovered</p>
                    </div>
                ` + "`" + `;
                return;
            }

            container.innerHTML = services.map(svc => ` + "`" + `
                <div class="service-card ${svc.isHttp ? 'http' : ''}"
                     ${svc.url ? ` + "`" + `onclick="window.open('${svc.url}', '_blank')"` + "`" + ` : ''}>
                    <div class="service-header">
                        <div>
                            <div class="service-name">${escapeHtml(svc.name)}</div>
                            <div class="source-badge">${svc.source}</div>
                        </div>
                        <div class="service-port">:${svc.port}</div>
                    </div>
                    <div class="service-details">
                        ${svc.container ? ` + "`" + `<p>Container: ${escapeHtml(svc.container)}</p>` + "`" + ` : ''}
                        ${svc.image ? ` + "`" + `<p>Image: ${escapeHtml(svc.image)}</p>` + "`" + ` : ''}
                        ${svc.process ? ` + "`" + `<p>Process: ${escapeHtml(svc.process)}</p>` + "`" + ` : ''}
                        ${svc.projectPath ? ` + "`" + `<p>Project: ${escapeHtml(svc.projectPath)}</p>` + "`" + ` : ''}
                        ${svc.description ? ` + "`" + `<p>${escapeHtml(svc.description)}</p>` + "`" + ` : ''}
                    </div>
                    <div class="service-tags">
                        ${(svc.tags || []).map(tag => ` + "`" + `<span class="tag ${tag}">${tag}</span>` + "`" + `).join('')}
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        function escapeHtml(text) {
            if (!text) return '';
            const div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }

        // Section collapse toggle
        function toggleSection(sectionId) {
            const section = document.getElementById(sectionId);
            if (section) {
                section.classList.toggle('collapsed');
                // Save state to localStorage
                const collapsed = section.classList.contains('collapsed');
                localStorage.setItem(sectionId + '-collapsed', collapsed);
            }
        }

        // Restore collapsed state from localStorage
        function restoreCollapsedState() {
            ['performance-section', 'services-section', 'terminal-section'].forEach(id => {
                const collapsed = localStorage.getItem(id + '-collapsed') === 'true';
                if (collapsed) {
                    document.getElementById(id)?.classList.add('collapsed');
                }
            });
        }

        // Chart drawing
        function drawChart(canvasId, data, color) {
            const canvas = document.getElementById(canvasId);
            if (!canvas) return;

            const ctx = canvas.getContext('2d');
            const dpr = window.devicePixelRatio || 1;
            const rect = canvas.getBoundingClientRect();

            canvas.width = rect.width * dpr;
            canvas.height = rect.height * dpr;
            ctx.scale(dpr, dpr);

            const width = rect.width;
            const height = rect.height;
            const padding = 2;

            ctx.clearRect(0, 0, width, height);

            if (!data || data.length < 2) return;

            // Draw filled area
            ctx.beginPath();
            ctx.moveTo(padding, height - padding);

            const stepX = (width - padding * 2) / (data.length - 1);
            data.forEach((val, i) => {
                const x = padding + i * stepX;
                const y = height - padding - (val / 100) * (height - padding * 2);
                if (i === 0) {
                    ctx.lineTo(x, y);
                } else {
                    ctx.lineTo(x, y);
                }
            });

            ctx.lineTo(width - padding, height - padding);
            ctx.closePath();

            // Fill with gradient
            const gradient = ctx.createLinearGradient(0, 0, 0, height);
            gradient.addColorStop(0, color + '40');
            gradient.addColorStop(1, color + '05');
            ctx.fillStyle = gradient;
            ctx.fill();

            // Draw line
            ctx.beginPath();
            data.forEach((val, i) => {
                const x = padding + i * stepX;
                const y = height - padding - (val / 100) * (height - padding * 2);
                if (i === 0) {
                    ctx.moveTo(x, y);
                } else {
                    ctx.lineTo(x, y);
                }
            });
            ctx.strokeStyle = color;
            ctx.lineWidth = 2;
            ctx.stroke();
        }

        async function loadStats() {
            try {
                const response = await fetch('/api/stats');
                const history = await response.json();

                if (history.stats && history.stats.length > 0) {
                    const cpuData = history.stats.map(s => s.cpuPercent);
                    const memData = history.stats.map(s => s.memoryPercent);

                    const latest = history.stats[history.stats.length - 1];
                    document.getElementById('cpu-value').textContent = latest.cpuPercent.toFixed(1) + '%';
                    document.getElementById('mem-value').textContent = latest.memoryPercent.toFixed(1) + '%';

                    // Get accent color from CSS
                    const style = getComputedStyle(document.body);
                    const accentPrimary = style.getPropertyValue('--accent-primary').trim() || '#00d9ff';
                    const accentSecondary = style.getPropertyValue('--accent-secondary').trim() || '#00ff88';

                    drawChart('cpu-chart', cpuData, accentPrimary);
                    drawChart('mem-chart', memData, accentSecondary);
                }
            } catch (error) {
                console.error('Failed to load stats:', error);
            }
        }

        // Terminal initialization
        let term = null;
        let termSocket = null;
        let fitAddon = null;

        function initTerminal() {
            if (term) return; // Already initialized

            const termContainer = document.getElementById('terminal');
            if (!termContainer) return;

            term = new Terminal({
                cursorBlink: true,
                fontSize: 14,
                fontFamily: terminalFont,
                theme: {
                    background: '#1a1a2e',
                    foreground: '#e4e4e4',
                    cursor: '#00d9ff',
                    cursorAccent: '#1a1a2e',
                    selectionBackground: 'rgba(0, 217, 255, 0.3)',
                }
            });

            fitAddon = new FitAddon.FitAddon();
            term.loadAddon(fitAddon);
            term.open(termContainer);
            fitAddon.fit();

            // Connect WebSocket
            const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
            const wsUrl = protocol + '//' + window.location.host + '/ws/terminal';
            termSocket = new WebSocket(wsUrl);
            termSocket.binaryType = 'arraybuffer';

            termSocket.onopen = () => {
                // Send initial size
                sendTerminalSize();
            };

            termSocket.onmessage = (event) => {
                if (event.data instanceof ArrayBuffer) {
                    term.write(new Uint8Array(event.data));
                } else {
                    term.write(event.data);
                }
            };

            termSocket.onclose = () => {
                term.write('\r\n\x1b[31m[Connection closed]\x1b[0m\r\n');
            };

            termSocket.onerror = (error) => {
                console.error('WebSocket error:', error);
                term.write('\r\n\x1b[31m[Connection error]\x1b[0m\r\n');
            };

            // Send terminal input to server
            term.onData((data) => {
                if (termSocket && termSocket.readyState === WebSocket.OPEN) {
                    termSocket.send(data);
                }
            });

            // Handle resize
            window.addEventListener('resize', () => {
                if (fitAddon) {
                    fitAddon.fit();
                    sendTerminalSize();
                }
            });

            // Also refit when section is expanded
            const termSection = document.getElementById('terminal-section');
            if (termSection) {
                const observer = new MutationObserver(() => {
                    if (!termSection.classList.contains('collapsed') && fitAddon) {
                        setTimeout(() => {
                            fitAddon.fit();
                            sendTerminalSize();
                        }, 100);
                    }
                });
                observer.observe(termSection, { attributes: true, attributeFilter: ['class'] });
            }
        }

        function sendTerminalSize() {
            if (termSocket && termSocket.readyState === WebSocket.OPEN && term) {
                const size = { cols: term.cols, rows: term.rows };
                termSocket.send(JSON.stringify(size));
            }
        }

        // Initial load
        restoreCollapsedState();
        loadConfig().then(() => {
            loadServices();
            loadStats();
            initTerminal();
            // Set up refresh with configured interval
            setInterval(loadServices, refreshInterval);
            // Stats refresh more frequently (every 2 seconds to match server collection)
            setInterval(loadStats, 2000);
        });
    </script>
</body>
</html>`

const configHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Settings - Dev Machine Proxy</title>
    <style>
` + themesCSS + baseCSS + `
        .config-form {
            max-width: 600px;
            margin: 0 auto;
        }

        .form-group {
            margin-bottom: 1.5rem;
        }

        .form-group label {
            display: block;
            font-size: 0.9rem;
            font-weight: 600;
            color: var(--text-primary);
            margin-bottom: 0.5rem;
        }

        .form-group .description {
            font-size: 0.8rem;
            color: var(--text-muted);
            margin-bottom: 0.5rem;
        }

        .form-group input[type="text"],
        .form-group input[type="number"],
        .form-group textarea {
            width: 100%;
            padding: 0.75rem 1rem;
            background: var(--bg-card);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            color: var(--text-primary);
            font-size: 1rem;
            transition: border-color 0.3s ease;
        }

        .form-group textarea {
            font-family: 'Courier New', monospace;
            font-size: 0.85rem;
            min-height: 120px;
            resize: vertical;
        }

        .form-group input:focus,
        .form-group textarea:focus {
            outline: none;
            border-color: var(--accent-primary);
        }

        .theme-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
            gap: 1rem;
        }

        .theme-option {
            background: var(--bg-card);
            border: 2px solid var(--border-color);
            border-radius: 10px;
            padding: 1rem;
            cursor: pointer;
            transition: all 0.3s ease;
        }

        .theme-option:hover {
            border-color: color-mix(in srgb, var(--accent-primary) 50%, transparent);
        }

        .theme-option.selected {
            border-color: var(--accent-primary);
            background: color-mix(in srgb, var(--accent-primary) 10%, transparent);
        }

        .theme-option .theme-name {
            font-weight: 600;
            color: var(--text-primary);
            margin-bottom: 0.25rem;
        }

        .theme-option .theme-desc {
            font-size: 0.75rem;
            color: var(--text-muted);
        }

        .theme-preview {
            height: 8px;
            border-radius: 4px;
            margin-bottom: 0.75rem;
        }

        .button-row {
            display: flex;
            align-items: center;
            gap: 1rem;
            margin-top: 2rem;
        }

        .btn {
            padding: 0.75rem 1.5rem;
            border-radius: 8px;
            font-size: 1rem;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s ease;
            border: none;
            text-decoration: none;
            display: inline-block;
        }

        .btn-primary {
            background: var(--accent-primary);
            color: var(--bg-primary);
        }

        .btn-primary:hover {
            opacity: 0.9;
            transform: translateY(-2px);
        }

        .btn-secondary {
            background: var(--bg-card);
            color: var(--text-primary);
            border: 1px solid var(--border-color);
        }

        .btn-secondary:hover {
            border-color: var(--accent-primary);
        }

        .back-link {
            display: inline-block;
            color: var(--text-muted);
            text-decoration: none;
            margin-bottom: 2rem;
            transition: color 0.3s ease;
        }

        .back-link:hover {
            color: var(--accent-primary);
        }

        .save-status {
            color: var(--accent-secondary);
            opacity: 0;
            transition: opacity 0.3s ease;
        }

        .save-status.show {
            opacity: 1;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-link">&larr; Back to Dashboard</a>

        <header>
            <h1>Settings</h1>
            <p class="subtitle">Customize your dashboard</p>
        </header>

        <div class="config-form">
            <div class="form-group">
                <label for="title">Dashboard Title</label>
                <p class="description">The title shown at the top of the dashboard</p>
                <input type="text" id="title" placeholder="Dev Machine Services">
            </div>

            <div class="form-group">
                <label for="refresh">Refresh Interval (seconds)</label>
                <p class="description">How often to refresh the service list</p>
                <input type="number" id="refresh" min="5" max="300" value="30">
            </div>

            <div class="form-group">
                <label>Theme</label>
                <p class="description">Choose your preferred color scheme</p>
                <div class="theme-grid" id="theme-grid">
                    <!-- Themes loaded dynamically -->
                </div>
            </div>

            <div class="form-group">
                <label for="terminal-font">Terminal Font Family</label>
                <p class="description">CSS font-family for the terminal (e.g., "MesloLGS NF", "Fira Code")</p>
                <input type="text" id="terminal-font" placeholder="MesloLGS NF">
            </div>

            <div class="form-group">
                <label for="custom-head">Custom Head HTML</label>
                <p class="description">Custom HTML to inject in &lt;head&gt; for loading web fonts, stylesheets, etc.</p>
                <textarea id="custom-head" placeholder="<link href='https://fonts.googleapis.com/...' rel='stylesheet'>"></textarea>
            </div>

            <div class="button-row">
                <button class="btn btn-primary" onclick="saveConfig()">Save Settings</button>
                <a href="/" class="btn btn-secondary">&larr; Back to Dashboard</a>
                <span class="save-status" id="save-status">Saved!</span>
            </div>
        </div>
    </div>

    <script>
        let currentTheme = 'cyberpunk';
        const themePreviewColors = {
            'cyberpunk': 'linear-gradient(90deg, #00d9ff, #00ff88)',
            'catppuccin-mocha': 'linear-gradient(90deg, #cba6f7, #f5c2e7)',
            'catppuccin-macchiato': 'linear-gradient(90deg, #c6a0f6, #f5bde6)',
            'solarized-dark': 'linear-gradient(90deg, #268bd2, #2aa198)',
            'solarized-light': 'linear-gradient(90deg, #268bd2, #2aa198)',
            'dracula': 'linear-gradient(90deg, #bd93f9, #ff79c6)',
            'nord': 'linear-gradient(90deg, #88c0d0, #81a1c1)',
            'gruvbox-dark': 'linear-gradient(90deg, #fabd2f, #fe8019)',
            'tokyo-night': 'linear-gradient(90deg, #7aa2f7, #bb9af7)',
            'one-dark': 'linear-gradient(90deg, #61afef, #c678dd)',
            'monokai': 'linear-gradient(90deg, #a6e22e, #f92672)'
        };

        async function loadConfig() {
            try {
                const response = await fetch('/api/config');
                const config = await response.json();

                document.getElementById('title').value = config.title;
                document.getElementById('refresh').value = config.refreshInterval;
                document.getElementById('terminal-font').value = config.terminalFont || '';
                document.getElementById('custom-head').value = config.customHeadHtml || '';
                currentTheme = config.theme;
                document.body.setAttribute('data-theme', config.theme);

                updateThemeSelection();
            } catch (error) {
                console.error('Failed to load config:', error);
            }
        }

        async function loadThemes() {
            try {
                const response = await fetch('/api/themes');
                const themes = await response.json();

                const grid = document.getElementById('theme-grid');
                grid.innerHTML = themes.map(theme => ` + "`" + `
                    <div class="theme-option" data-theme="${theme.id}" onclick="selectTheme('${theme.id}')">
                        <div class="theme-preview" style="background: ${themePreviewColors[theme.id] || '#666'}"></div>
                        <div class="theme-name">${theme.name}</div>
                        <div class="theme-desc">${theme.description}</div>
                    </div>
                ` + "`" + `).join('');

                updateThemeSelection();
            } catch (error) {
                console.error('Failed to load themes:', error);
            }
        }

        function selectTheme(themeId) {
            currentTheme = themeId;
            document.body.setAttribute('data-theme', themeId);
            updateThemeSelection();
        }

        function updateThemeSelection() {
            document.querySelectorAll('.theme-option').forEach(el => {
                el.classList.toggle('selected', el.dataset.theme === currentTheme);
            });
        }

        async function saveConfig() {
            const config = {
                title: document.getElementById('title').value,
                refreshInterval: parseInt(document.getElementById('refresh').value, 10),
                theme: currentTheme,
                terminalFont: document.getElementById('terminal-font').value,
                customHeadHtml: document.getElementById('custom-head').value
            };

            try {
                const response = await fetch('/api/config', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(config)
                });

                if (response.ok) {
                    const status = document.getElementById('save-status');
                    status.classList.add('show');
                    setTimeout(() => status.classList.remove('show'), 2000);
                }
            } catch (error) {
                console.error('Failed to save config:', error);
                alert('Failed to save settings');
            }
        }

        // Load everything
        loadThemes();
        loadConfig();
    </script>
</body>
</html>`
