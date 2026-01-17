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

.title-row {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;
}

.logo {
    width: 48px;
    height: 48px;
}

h1 {
    font-size: 2.5rem;
    font-weight: 300;
    color: var(--accent-primary);
    margin-bottom: 0;
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

.project-status-icons {
    display: flex;
    gap: 0.4rem;
    align-items: center;
}

.status-icon {
    display: inline-flex;
    align-items: center;
    gap: 0.15rem;
    font-family: 'Courier New', monospace;
    font-size: 0.9rem;
    color: var(--port-text);
}

.status-icon svg {
    width: 14px;
    height: 14px;
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

.usage-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
    gap: 1.5rem;
}

.usage-card {
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 1.25rem;
}

.usage-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1rem;
}

.usage-title {
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--text-primary);
}

.usage-subtitle {
    font-size: 0.75rem;
    color: var(--text-muted);
}

.usage-pill {
    padding: 0.25rem 0.6rem;
    border-radius: 999px;
    font-size: 0.7rem;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border: 1px solid var(--border-color);
    color: var(--text-secondary);
}

.usage-pill.ok {
    border-color: color-mix(in srgb, var(--accent-secondary) 60%, transparent);
    color: var(--accent-secondary);
}

.usage-pill.warn {
    border-color: color-mix(in srgb, var(--accent-primary) 60%, transparent);
    color: var(--accent-primary);
}

.usage-windows {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1rem;
}

.usage-window {
    background: color-mix(in srgb, var(--bg-secondary) 60%, transparent);
    border: 1px solid var(--border-color);
    border-radius: 10px;
    padding: 0.85rem;
}

.usage-window-header {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    margin-bottom: 0.35rem;
}

.usage-window-title {
    font-size: 0.75rem;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    color: var(--text-secondary);
}

.usage-window-value {
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--accent-primary);
}

.usage-window-meta {
    font-size: 0.75rem;
    color: var(--text-muted);
    margin-bottom: 0.5rem;
}

.usage-chart {
    height: 70px;
    margin-bottom: 0.5rem;
}

.usage-chart canvas {
    width: 100%;
    height: 100%;
}

.usage-forecast {
    font-size: 0.75rem;
    color: var(--text-secondary);
    line-height: 1.3;
}

.usage-forecast.warn {
    color: var(--accent-primary);
}

.top-processes {
    margin-top: 0.75rem;
    padding-top: 0.75rem;
    border-top: 1px solid var(--border-color);
    font-size: 0.8rem;
}

.process-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.25rem 0;
    color: var(--text-muted);
}

.process-row:hover {
    color: var(--text-primary);
}

.process-name {
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin-right: 0.5rem;
}

.process-value {
    font-family: 'Courier New', monospace;
    color: var(--port-text);
    min-width: 50px;
    text-align: right;
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

.section-summary {
    display: flex;
    gap: 0.5rem;
    margin-left: auto;
    align-items: center;
}

.summary-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.25rem;
    font-size: 0.75rem;
    font-weight: 600;
    padding: 0.2rem 0.5rem;
    border-radius: 4px;
    background: color-mix(in srgb, var(--accent-primary) 15%, transparent);
    color: var(--accent-primary);
    font-family: 'Courier New', monospace;
}

.summary-badge.secondary {
    background: color-mix(in srgb, var(--accent-secondary) 15%, transparent);
    color: var(--accent-secondary);
}

.summary-badge.warning {
    background: var(--tag-project-bg);
    color: var(--tag-project-text);
}

.summary-badge.muted {
    background: color-mix(in srgb, var(--text-muted) 15%, transparent);
    color: var(--text-muted);
}

.section-content {
    overflow: hidden;
    transition: max-height 0.3s ease, opacity 0.3s ease;
    max-height: 5000px;
    opacity: 1;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
}

.section.collapsed .section-content {
    max-height: 0;
    opacity: 0;
}

.daily-tasks-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.daily-tasks-input {
    display: flex;
    gap: 0.5rem;
}

.daily-tasks-input input {
    flex: 1;
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 0.75rem 1rem;
    color: var(--text-primary);
    font-size: 0.9rem;
}

.daily-tasks-input input:focus {
    outline: none;
    border-color: var(--accent-primary);
}

.daily-tasks-input button {
    background: var(--accent-primary);
    border: none;
    border-radius: 8px;
    padding: 0.75rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
}

.daily-tasks-input button svg {
    width: 20px;
    height: 20px;
    stroke: var(--bg-primary);
}

.daily-tasks-input button:hover {
    opacity: 0.9;
}

.daily-tasks-list {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.daily-task-item {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1rem;
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    transition: all 0.2s ease;
}

.daily-task-item:hover {
    border-color: var(--accent-primary);
}

.daily-task-item.completed {
    opacity: 0.7;
}

.daily-task-item.completed .task-name {
    text-decoration: line-through;
    color: var(--text-muted);
}

.task-checkbox {
    width: 22px;
    height: 22px;
    border: 2px solid var(--accent-primary);
    border-radius: 4px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    transition: all 0.2s ease;
}

.task-checkbox:hover {
    background: color-mix(in srgb, var(--accent-primary) 20%, transparent);
}

.task-checkbox.checked {
    background: var(--accent-primary);
}

.task-checkbox svg {
    width: 14px;
    height: 14px;
    stroke: var(--bg-primary);
    stroke-width: 3;
    opacity: 0;
}

.task-checkbox.checked svg {
    opacity: 1;
}

.task-name {
    flex: 1;
    font-size: 0.95rem;
    color: var(--text-primary);
    cursor: text;
    padding: 0.2rem 0.4rem;
    margin: -0.2rem -0.4rem;
    border-radius: 4px;
    transition: background 0.2s ease;
}

.task-name:hover {
    background: color-mix(in srgb, var(--accent-primary) 10%, transparent);
}

.task-name-input {
    flex: 1;
    font-size: 0.95rem;
    color: var(--text-primary);
    background: var(--bg-primary);
    border: 1px solid var(--accent-primary);
    border-radius: 4px;
    padding: 0.2rem 0.4rem;
    outline: none;
}

.task-streaks {
    display: flex;
    gap: 0.75rem;
    align-items: center;
    font-size: 0.8rem;
    color: var(--text-muted);
}

.streak-badge {
    display: flex;
    align-items: center;
    gap: 0.25rem;
    padding: 0.2rem 0.5rem;
    border-radius: 4px;
    background: color-mix(in srgb, var(--accent-secondary) 15%, transparent);
    color: var(--accent-secondary);
}

.streak-badge.current {
    background: color-mix(in srgb, var(--accent-primary) 15%, transparent);
    color: var(--accent-primary);
}

.streak-badge svg {
    width: 14px;
    height: 14px;
}

.task-delete {
    opacity: 0;
    cursor: pointer;
    padding: 0.25rem;
    border-radius: 4px;
    transition: opacity 0.2s ease;
}

.daily-task-item:hover .task-delete {
    opacity: 0.5;
}

.task-delete:hover {
    opacity: 1 !important;
    background: rgba(255, 0, 0, 0.1);
}

.task-delete svg {
    width: 16px;
    height: 16px;
    stroke: var(--text-muted);
}

.daily-tasks-empty {
    text-align: center;
    padding: 2rem;
    color: var(--text-muted);
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
            <div class="title-row">
                <svg class="logo" viewBox="0 0 64 64">
                    <rect width="64" height="64" rx="8" fill="var(--bg-secondary)"/>
                    <path d="M16 20 L28 32 L16 44" stroke="var(--accent-primary)" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" fill="none"/>
                    <line x1="32" y1="44" x2="48" y2="44" stroke="var(--accent-secondary)" stroke-width="4" stroke-linecap="round"/>
                </svg>
                <h1 id="page-title">Dev Machine Services</h1>
            </div>
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
                <div class="section-summary">
                    <span class="summary-badge" id="summary-cpu">CPU --%</span>
                    <span class="summary-badge secondary" id="summary-mem">MEM --%</span>
                </div>
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
                        <div class="top-processes" id="top-cpu-processes"></div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-header">
                            <span class="stat-title">Memory Usage</span>
                            <span class="stat-value" id="mem-value">--%</span>
                        </div>
                        <div class="stat-chart">
                            <canvas id="mem-chart"></canvas>
                        </div>
                        <div class="top-processes" id="top-mem-processes"></div>
                    </div>
                </div>
            </div>
        </div>

        <div class="section" id="ai-usage-section">
            <div class="section-header" onclick="toggleSection('ai-usage-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">AI Usage</span>
                <div class="section-summary">
                    <span class="summary-badge" id="summary-claude">Claude --%</span>
                    <span class="summary-badge secondary" id="summary-codex">Codex --%</span>
                </div>
            </div>
            <p class="section-subtitle">Usage from claude-usage and codex-usage</p>
            <div class="section-content">
                <div class="usage-grid">
                    <div class="usage-card">
                        <div class="usage-header">
                            <div>
                                <div class="usage-title">Claude</div>
                                <div class="usage-subtitle">5-hour and 7-day windows</div>
                            </div>
                            <div class="usage-pill" id="claude-status">Loading</div>
                        </div>
                        <div class="usage-windows">
                            <div class="usage-window">
                                <div class="usage-window-header">
                                    <span class="usage-window-title">5-hour</span>
                                    <span class="usage-window-value" id="claude-five-value">--%</span>
                                </div>
                                <div class="usage-window-meta" id="claude-five-reset">Reset: --</div>
                                <div class="usage-chart">
                                    <canvas id="claude-five-chart"></canvas>
                                </div>
                                <div class="usage-forecast" id="claude-five-forecast">--</div>
                            </div>
                            <div class="usage-window">
                                <div class="usage-window-header">
                                    <span class="usage-window-title">7-day</span>
                                    <span class="usage-window-value" id="claude-seven-value">--%</span>
                                </div>
                                <div class="usage-window-meta" id="claude-seven-reset">Reset: --</div>
                                <div class="usage-chart">
                                    <canvas id="claude-seven-chart"></canvas>
                                </div>
                                <div class="usage-forecast" id="claude-seven-forecast">--</div>
                            </div>
                        </div>
                    </div>
                    <div class="usage-card">
                        <div class="usage-header">
                            <div>
                                <div class="usage-title">Codex</div>
                                <div class="usage-subtitle">Primary and secondary windows</div>
                            </div>
                            <div class="usage-pill" id="codex-status">Loading</div>
                        </div>
                        <div class="usage-windows">
                            <div class="usage-window">
                                <div class="usage-window-header">
                                    <span class="usage-window-title">Primary</span>
                                    <span class="usage-window-value" id="codex-primary-value">--%</span>
                                </div>
                                <div class="usage-window-meta" id="codex-primary-reset">Reset: --</div>
                                <div class="usage-chart">
                                    <canvas id="codex-primary-chart"></canvas>
                                </div>
                                <div class="usage-forecast" id="codex-primary-forecast">--</div>
                            </div>
                            <div class="usage-window">
                                <div class="usage-window-header">
                                    <span class="usage-window-title">Secondary</span>
                                    <span class="usage-window-value" id="codex-secondary-value">--%</span>
                                </div>
                                <div class="usage-window-meta" id="codex-secondary-reset">Reset: --</div>
                                <div class="usage-chart">
                                    <canvas id="codex-secondary-chart"></canvas>
                                </div>
                                <div class="usage-forecast" id="codex-secondary-forecast">--</div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="empty-state" id="usage-error" style="display:none; margin-top: 1rem;">
                    <p>Usage data unavailable</p>
                    <p style="font-size: 0.8rem; margin-top: 0.5rem;" id="usage-error-details"></p>
                </div>
            </div>
        </div>

        <div class="section" id="projects-section">
            <div class="section-header" onclick="toggleSection('projects-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">Projects</span>
                <div class="section-summary">
                    <span class="summary-badge warning" id="summary-projects-dirty" style="display:none">0 changed</span>
                    <span class="summary-badge muted" id="summary-projects-total">0 total</span>
                </div>
            </div>
            <p class="section-subtitle">Project folders with git status</p>
            <div class="section-content">
                <div id="projects" class="services-grid">
                    <div class="loading">
                        <div class="spinner"></div>
                        <p>Scanning projects...</p>
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
                <div class="section-summary">
                    <span class="summary-badge" id="summary-services">0 services</span>
                </div>
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

        <div class="section" id="daily-tasks-section">
            <div class="section-header" onclick="toggleSection('daily-tasks-section')">
                <span class="section-toggle">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                </span>
                <span class="section-title">Daily Tasks</span>
                <div class="section-summary">
                    <span class="summary-badge warning" id="summary-tasks-remaining" style="display:none">0 remaining</span>
                    <span class="summary-badge secondary" id="summary-tasks-done" style="display:none">0 done</span>
                </div>
            </div>
            <p class="section-subtitle">Recurring habits and tasks - resets at midnight</p>
            <div class="section-content">
                <div class="daily-tasks-container">
                    <div class="daily-tasks-input">
                        <input type="text" id="new-task-input" placeholder="Add a new daily task..." onkeypress="if(event.key==='Enter')addDailyTask()">
                        <button onclick="addDailyTask()">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="12" y1="5" x2="12" y2="19"></line>
                                <line x1="5" y1="12" x2="19" y2="12"></line>
                            </svg>
                        </button>
                    </div>
                    <div id="daily-tasks" class="daily-tasks-list"></div>
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

                // Apply section visibility
                if (config.sections) {
                    toggleSectionVisibility('performance-section', config.sections.performance);
                    toggleSectionVisibility('ai-usage-section', config.sections.aiUsage);
                    toggleSectionVisibility('projects-section', config.sections.projects);
                    toggleSectionVisibility('services-section', config.sections.services);
                    toggleSectionVisibility('daily-tasks-section', config.sections.dailyTasks);
                    toggleSectionVisibility('terminal-section', config.sections.terminal);
                }

                // Apply section order
                if (config.sectionOrder && config.sectionOrder.length > 0) {
                    reorderSections(config.sectionOrder);
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

            // Update summary badge
            const count = services ? services.length : 0;
            document.getElementById('summary-services').textContent = count + ' service' + (count !== 1 ? 's' : '');

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

        async function loadProjects() {
            try {
                const response = await fetch('/api/projects');
                const projects = await response.json();
                renderProjects(projects);
            } catch (error) {
                console.error('Failed to load projects:', error);
                document.getElementById('projects').innerHTML = ` + "`" + `
                    <div class="empty-state">
                        <p>Failed to load projects</p>
                    </div>
                ` + "`" + `;
            }
        }

        function renderProjects(projects) {
            const container = document.getElementById('projects');

            // Update summary badges
            const total = projects ? projects.length : 0;
            const dirty = projects ? projects.filter(p => p.changedFiles > 0 || p.unpushed > 0).length : 0;

            const dirtyBadge = document.getElementById('summary-projects-dirty');
            const totalBadge = document.getElementById('summary-projects-total');

            if (dirty > 0) {
                dirtyBadge.textContent = dirty + ' changed';
                dirtyBadge.style.display = '';
            } else {
                dirtyBadge.style.display = 'none';
            }
            totalBadge.textContent = total + ' total';

            if (!projects || projects.length === 0) {
                container.innerHTML = ` + "`" + `
                    <div class="empty-state">
                        <p>No projects found</p>
                        <p style="font-size: 0.8rem; margin-top: 0.5rem;">Set the -projects flag to scan a directory</p>
                    </div>
                ` + "`" + `;
                return;
            }

            container.innerHTML = projects.map(proj => {
                let statusIcons = [];

                // Changed files icon (pencil)
                if (proj.changedFiles > 0) {
                    statusIcons.push(` + "`" + `
                        <span class="status-icon" title="${proj.changedFiles} changed files">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M12 20h9"></path>
                                <path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path>
                            </svg>
                            ${proj.changedFiles}
                        </span>
                    ` + "`" + `);
                }

                // Unpushed commits icon (upload/arrow up)
                if (proj.unpushed > 0) {
                    statusIcons.push(` + "`" + `
                        <span class="status-icon" title="${proj.unpushed} unpushed commits">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <line x1="12" y1="19" x2="12" y2="5"></line>
                                <polyline points="5 12 12 5 19 12"></polyline>
                            </svg>
                            ${proj.unpushed}
                        </span>
                    ` + "`" + `);
                }

                // Ahead icon (arrow up from line)
                if (proj.ahead > 0) {
                    statusIcons.push(` + "`" + `
                        <span class="status-icon" title="${proj.ahead} commits ahead of remote">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="18 15 12 9 6 15"></polyline>
                                <line x1="12" y1="9" x2="12" y2="21"></line>
                            </svg>
                            ${proj.ahead}
                        </span>
                    ` + "`" + `);
                }

                // Behind icon (arrow down)
                if (proj.behind > 0) {
                    statusIcons.push(` + "`" + `
                        <span class="status-icon" title="${proj.behind} commits behind remote">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="6 9 12 15 18 9"></polyline>
                                <line x1="12" y1="15" x2="12" y2="3"></line>
                            </svg>
                            ${proj.behind}
                        </span>
                    ` + "`" + `);
                }

                // Clean status icon (checkmark)
                if (proj.isGit && statusIcons.length === 0) {
                    statusIcons.push(` + "`" + `
                        <span class="status-icon" title="Clean - no changes">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="20 6 9 17 4 12"></polyline>
                            </svg>
                        </span>
                    ` + "`" + `);
                }

                const lastMod = proj.lastModified ? new Date(proj.lastModified).toLocaleDateString() : '';

                return ` + "`" + `
                <div class="service-card ${proj.changedFiles > 0 || proj.unpushed > 0 ? 'http' : ''}">
                    <div class="service-header">
                        <div>
                            <div class="service-name">${escapeHtml(proj.name)}</div>
                            <div class="source-badge">${proj.branch || 'no branch'}</div>
                        </div>
                        <div class="project-status-icons">
                            ${statusIcons.join('')}
                        </div>
                    </div>
                    <div class="service-details">
                        ${lastMod ? ` + "`" + `<p>Last activity: ${lastMod}</p>` + "`" + ` : ''}
                    </div>
                    <div class="service-tags">
                        ${(proj.tags || []).map(tag => ` + "`" + `<span class="tag ${tag}">${tag}</span>` + "`" + `).join('')}
                    </div>
                </div>
            ` + "`" + `}).join('');
        }

        // Daily Tasks Functions
        async function loadDailyTasks() {
            try {
                const response = await fetch('/api/daily-tasks');
                const tasks = await response.json();
                renderDailyTasks(tasks);
            } catch (error) {
                console.error('Failed to load daily tasks:', error);
            }
        }

        function renderDailyTasks(tasks) {
            const container = document.getElementById('daily-tasks');

            // Update summary badges
            const total = tasks ? tasks.length : 0;
            const done = tasks ? tasks.filter(t => t.completedToday).length : 0;
            const remaining = total - done;

            const remainingBadge = document.getElementById('summary-tasks-remaining');
            const doneBadge = document.getElementById('summary-tasks-done');

            if (remaining > 0) {
                remainingBadge.textContent = remaining + ' remaining';
                remainingBadge.style.display = '';
            } else {
                remainingBadge.style.display = 'none';
            }

            if (done > 0) {
                doneBadge.textContent = done + ' done';
                doneBadge.style.display = '';
            } else {
                doneBadge.style.display = 'none';
            }

            if (!tasks || tasks.length === 0) {
                container.innerHTML = ` + "`" + `
                    <div class="daily-tasks-empty">
                        <p>No daily tasks yet</p>
                        <p style="font-size: 0.8rem; margin-top: 0.5rem;">Add a task above to start tracking your habits</p>
                    </div>
                ` + "`" + `;
                return;
            }

            container.innerHTML = tasks.map(task => ` + "`" + `
                <div class="daily-task-item ${task.completedToday ? 'completed' : ''}" data-id="${task.id}">
                    <div class="task-checkbox ${task.completedToday ? 'checked' : ''}" onclick="toggleDailyTask('${task.id}')">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor">
                            <polyline points="20 6 9 17 4 12"></polyline>
                        </svg>
                    </div>
                    <span class="task-name" onclick="editDailyTask('${task.id}', this)" title="Click to edit">${escapeHtml(task.name)}</span>
                    <div class="task-streaks">
                        ${task.currentStreak > 0 ? ` + "`" + `
                            <span class="streak-badge current" title="Current streak">
                                <svg viewBox="0 0 24 24" fill="currentColor" stroke="none">
                                    <path d="M12 2C6.5 8 4 12 4 15c0 4.4 3.6 8 8 8s8-3.6 8-8c0-3-2.5-7-8-13z"/>
                                </svg>
                                ${task.currentStreak}
                            </span>
                        ` + "`" + ` : ''}
                        ${task.longestStreak > 0 ? ` + "`" + `
                            <span class="streak-badge" title="Longest streak">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M6 9l6 6 6-6"/>
                                </svg>
                                ${task.longestStreak}
                            </span>
                        ` + "`" + ` : ''}
                    </div>
                    <div class="task-delete" onclick="deleteDailyTask('${task.id}')" title="Delete task">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="18" y1="6" x2="6" y2="18"></line>
                            <line x1="6" y1="6" x2="18" y2="18"></line>
                        </svg>
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        async function addDailyTask() {
            const input = document.getElementById('new-task-input');
            const name = input.value.trim();
            if (!name) return;

            try {
                await fetch('/api/daily-tasks', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ name })
                });
                input.value = '';
                loadDailyTasks();
            } catch (error) {
                console.error('Failed to add task:', error);
            }
        }

        async function toggleDailyTask(id) {
            try {
                await fetch('/api/daily-tasks/toggle', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ id })
                });
                loadDailyTasks();
            } catch (error) {
                console.error('Failed to toggle task:', error);
            }
        }

        async function deleteDailyTask(id) {
            if (!confirm('Delete this daily task? This will also delete its streak history.')) return;

            try {
                await fetch('/api/daily-tasks', {
                    method: 'DELETE',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ id })
                });
                loadDailyTasks();
            } catch (error) {
                console.error('Failed to delete task:', error);
            }
        }

        function editDailyTask(id, element) {
            const currentName = element.textContent;
            const input = document.createElement('input');
            input.type = 'text';
            input.className = 'task-name-input';
            input.value = currentName;

            const saveEdit = async () => {
                const newName = input.value.trim();
                if (newName && newName !== currentName) {
                    try {
                        await fetch('/api/daily-tasks', {
                            method: 'PUT',
                            headers: { 'Content-Type': 'application/json' },
                            body: JSON.stringify({ id, name: newName })
                        });
                        loadDailyTasks();
                    } catch (error) {
                        console.error('Failed to update task:', error);
                        loadDailyTasks();
                    }
                } else {
                    loadDailyTasks();
                }
            };

            input.addEventListener('blur', saveEdit);
            input.addEventListener('keypress', (e) => {
                if (e.key === 'Enter') {
                    e.preventDefault();
                    input.blur();
                }
            });
            input.addEventListener('keydown', (e) => {
                if (e.key === 'Escape') {
                    loadDailyTasks();
                }
            });

            element.replaceWith(input);
            input.focus();
            input.select();
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

        // Section visibility toggle
        function toggleSectionVisibility(sectionId, visible) {
            const section = document.getElementById(sectionId);
            if (section) {
                section.style.display = visible ? '' : 'none';
            }
        }

        // Section order mapping (config key -> section id)
        const sectionIdMap = {
            'performance': 'performance-section',
            'aiUsage': 'ai-usage-section',
            'projects': 'projects-section',
            'services': 'services-section',
            'dailyTasks': 'daily-tasks-section',
            'terminal': 'terminal-section'
        };

        // Reorder sections based on config order
        function reorderSections(order) {
            const container = document.querySelector('.container');
            if (!container) return;

            // Get all section elements in order specified
            const sections = order.map(key => document.getElementById(sectionIdMap[key])).filter(Boolean);

            // Append sections in order (this moves them)
            sections.forEach(section => container.appendChild(section));
        }

        // Restore collapsed state from localStorage
        function restoreCollapsedState() {
            ['performance-section', 'ai-usage-section', 'projects-section', 'services-section', 'daily-tasks-section', 'terminal-section'].forEach(id => {
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

                    // Update summary badges
                    document.getElementById('summary-cpu').textContent = 'CPU ' + latest.cpuPercent.toFixed(0) + '%';
                    document.getElementById('summary-mem').textContent = 'MEM ' + latest.memoryPercent.toFixed(0) + '%';

                    // Get accent color from CSS
                    const style = getComputedStyle(document.body);
                    const accentPrimary = style.getPropertyValue('--accent-primary').trim() || '#00d9ff';
                    const accentSecondary = style.getPropertyValue('--accent-secondary').trim() || '#00ff88';

                    drawChart('cpu-chart', cpuData, accentPrimary);
                    drawChart('mem-chart', memData, accentSecondary);

                    // Render top processes
                    renderTopProcesses('top-cpu-processes', latest.topCPU, 'cpu');
                    renderTopProcesses('top-mem-processes', latest.topMemory, 'mem');
                }
            } catch (error) {
                console.error('Failed to load stats:', error);
            }
        }

        async function loadUsage() {
            try {
                const response = await fetch('/api/usage');
                const payload = await response.json();
                renderUsage(payload);
            } catch (error) {
                console.error('Failed to load usage:', error);
                showUsageError('Failed to load usage data.');
            }
        }

        function renderUsage(payload) {
            const errorBox = document.getElementById('usage-error');
            const errorDetails = document.getElementById('usage-error-details');

            if (!payload || !payload.latest) {
                showUsageError('No usage data available.');
                return;
            }

            if (payload.latest.errors && payload.latest.errors.length > 0) {
                showUsageError(payload.latest.errors.join(' | '));
            } else {
                errorBox.style.display = 'none';
                errorDetails.textContent = '';
            }

            const style = getComputedStyle(document.body);
            const accentPrimary = style.getPropertyValue('--accent-primary').trim() || '#00d9ff';
            const accentSecondary = style.getPropertyValue('--accent-secondary').trim() || '#00ff88';

            if (payload.latest.claude) {
                const claude = payload.latest.claude;
                updateWindowDisplay({
                    prefix: 'claude-five',
                    windowData: claude.fiveHour,
                    history: payload.history ? payload.history.claudeFiveHour : [],
                    color: accentPrimary
                });
                updateWindowDisplay({
                    prefix: 'claude-seven',
                    windowData: claude.sevenDay,
                    history: payload.history ? payload.history.claudeSevenDay : [],
                    color: accentPrimary
                });

                document.getElementById('summary-claude').textContent =
                    'Claude ' + formatPercent(claude.fiveHour.current.usedPercent);
                setUsagePill('claude-status', claude.fiveHour.forecast);
            }

            if (payload.latest.codex) {
                const codex = payload.latest.codex;
                updateWindowDisplay({
                    prefix: 'codex-primary',
                    windowData: codex.primary,
                    history: payload.history ? payload.history.codexPrimary : [],
                    color: accentSecondary
                });

                if (codex.secondary) {
                    updateWindowDisplay({
                        prefix: 'codex-secondary',
                        windowData: codex.secondary,
                        history: payload.history ? payload.history.codexSecondary : [],
                        color: accentSecondary
                    });
                } else {
                    setWindowUnavailable('codex-secondary');
                }

                document.getElementById('summary-codex').textContent =
                    'Codex ' + formatPercent(codex.primary.current.usedPercent);
                setUsagePill('codex-status', codex.primary.forecast);
            }
        }

        function updateWindowDisplay({ prefix, windowData, history, color }) {
            if (!windowData || !windowData.current) {
                setWindowUnavailable(prefix);
                return;
            }

            document.getElementById(prefix + '-value').textContent =
                formatPercent(windowData.current.usedPercent);
            document.getElementById(prefix + '-reset').textContent =
                'Reset: ' + formatReset(windowData.current.resetAt, windowData.current.resetInSeconds);
            document.getElementById(prefix + '-forecast').textContent =
                formatForecast(windowData.forecast, windowData.current.resetAt);
            document.getElementById(prefix + '-forecast').classList.toggle('warn', windowData.forecast && windowData.forecast.willExhaust);

            const series = (history || []).map(point => point.value);
            drawChart(prefix + '-chart', series, color);
        }

        function setWindowUnavailable(prefix) {
            document.getElementById(prefix + '-value').textContent = 'N/A';
            document.getElementById(prefix + '-reset').textContent = 'Reset: --';
            document.getElementById(prefix + '-forecast').textContent = 'Forecast unavailable';
        }

        function setUsagePill(elementId, forecast) {
            const pill = document.getElementById(elementId);
            if (!pill) return;
            if (forecast && forecast.willExhaust) {
                pill.textContent = 'At risk';
                pill.classList.add('warn');
                pill.classList.remove('ok');
            } else {
                pill.textContent = 'On track';
                pill.classList.add('ok');
                pill.classList.remove('warn');
            }
        }

        function formatPercent(value) {
            if (typeof value !== 'number') return '--%';
            return value.toFixed(0) + '%';
        }

        function formatReset(resetAt, resetInSeconds) {
            if (resetInSeconds && resetInSeconds > 0) {
                return 'in ' + formatDuration(resetInSeconds);
            }
            if (resetAt) {
                const date = new Date(resetAt);
                if (!isNaN(date)) {
                    return date.toLocaleString();
                }
            }
            return '--';
        }

        function formatForecast(forecast, resetAt) {
            if (!forecast) return 'Forecast unavailable';

            const rate = forecast.ratePerHour;
            if (!rate || rate <= 0) {
                return 'Stable usage';
            }

            if (forecast.willExhaust && forecast.exhaustAt) {
                const exhaustAt = new Date(forecast.exhaustAt);
                return 'Exhaust in ' + formatDuration(forecast.hoursToExhaust * 3600) +
                    ' (' + exhaustAt.toLocaleTimeString() + ')';
            }

            if (resetAt) {
                const resetDate = new Date(resetAt);
                if (!isNaN(resetDate)) {
                    return 'Projected ' + forecast.projectedAtReset.toFixed(0) + '% at reset';
                }
            }

            return 'Projected ' + forecast.projectedAtReset.toFixed(0) + '% at reset';
        }

        function formatDuration(seconds) {
            if (!seconds || seconds < 0) return '--';
            const rounded = Math.max(0, Math.floor(seconds));
            const hours = Math.floor(rounded / 3600);
            const minutes = Math.floor((rounded % 3600) / 60);
            if (hours > 0 && minutes > 0) return hours + 'h ' + minutes + 'm';
            if (hours > 0) return hours + 'h';
            return minutes + 'm';
        }

        function showUsageError(message) {
            const errorBox = document.getElementById('usage-error');
            const errorDetails = document.getElementById('usage-error-details');
            if (errorBox && errorDetails) {
                errorDetails.textContent = message;
                errorBox.style.display = '';
            }
        }

        function renderTopProcesses(containerId, processes, type) {
            const container = document.getElementById(containerId);
            if (!container || !processes || processes.length === 0) {
                if (container) container.innerHTML = '';
                return;
            }

            container.innerHTML = processes.map(proc => {
                const value = type === 'cpu'
                    ? proc.cpuPercent.toFixed(1) + '%'
                    : proc.memoryMB.toFixed(0) + ' MB';
                return ` + "`" + `
                    <div class="process-row" title="PID: ${proc.pid}">
                        <span class="process-name">${escapeHtml(proc.name)}</span>
                        <span class="process-value">${value}</span>
                    </div>
                ` + "`" + `;
            }).join('');
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
            loadProjects();
            loadServices();
            loadStats();
            loadDailyTasks();
            loadUsage();
            initTerminal();
            // Set up refresh with configured interval
            setInterval(loadProjects, refreshInterval);
            setInterval(loadServices, refreshInterval);
            // Stats refresh more frequently (every 2 seconds to match server collection)
            setInterval(loadStats, 2000);
            setInterval(loadUsage, 60000);
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

        .checkbox-group {
            display: flex;
            flex-wrap: wrap;
            gap: 1rem;
        }

        .checkbox-item {
            display: flex;
            align-items: center;
            gap: 0.5rem;
            cursor: pointer;
        }

        .checkbox-item input[type="checkbox"] {
            width: 18px;
            height: 18px;
            accent-color: var(--accent-primary);
            cursor: pointer;
        }

        .checkbox-item label {
            cursor: pointer;
            margin-bottom: 0;
        }

        .section-order-list {
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }

        .section-order-item {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            padding: 0.75rem 1rem;
            background: var(--bg-secondary);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            cursor: grab;
            transition: all 0.2s ease;
        }

        .section-order-item:hover {
            border-color: var(--accent-primary);
        }

        .section-order-item.dragging {
            opacity: 0.5;
            border-color: var(--accent-primary);
            background: color-mix(in srgb, var(--accent-primary) 10%, var(--bg-secondary));
        }

        .section-order-item.drag-over {
            border-color: var(--accent-primary);
            border-style: dashed;
        }

        .drag-handle {
            display: flex;
            align-items: center;
            color: var(--text-muted);
            cursor: grab;
        }

        .drag-handle svg {
            width: 20px;
            height: 20px;
        }

        .section-order-item input[type="checkbox"] {
            width: 18px;
            height: 18px;
            accent-color: var(--accent-primary);
            cursor: pointer;
        }

        .section-order-item label {
            flex: 1;
            cursor: pointer;
            margin-bottom: 0;
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

            <div class="form-group">
                <label>Section Order & Visibility</label>
                <p class="description">Drag to reorder, toggle visibility with checkboxes</p>
                <div id="section-order-list" class="section-order-list">
                    <div class="section-order-item" data-section="performance" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-performance" checked>
                        <label for="section-performance">Performance</label>
                    </div>
                    <div class="section-order-item" data-section="aiUsage" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-ai-usage" checked>
                        <label for="section-ai-usage">AI Usage</label>
                    </div>
                    <div class="section-order-item" data-section="projects" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-projects" checked>
                        <label for="section-projects">Projects</label>
                    </div>
                    <div class="section-order-item" data-section="services" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-services" checked>
                        <label for="section-services">Services</label>
                    </div>
                    <div class="section-order-item" data-section="dailyTasks" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-daily-tasks" checked>
                        <label for="section-daily-tasks">Daily Tasks</label>
                    </div>
                    <div class="section-order-item" data-section="terminal" draggable="true">
                        <span class="drag-handle">
                            <svg viewBox="0 0 24 24" fill="currentColor"><path d="M9 5h2v2H9zm4 0h2v2h-2zM9 9h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/></svg>
                        </span>
                        <input type="checkbox" id="section-terminal" checked>
                        <label for="section-terminal">Terminal</label>
                    </div>
                </div>
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

                // Load section visibility
                if (config.sections) {
                    document.getElementById('section-performance').checked = config.sections.performance !== false;
                    document.getElementById('section-ai-usage').checked = config.sections.aiUsage !== false;
                    document.getElementById('section-projects').checked = config.sections.projects !== false;
                    document.getElementById('section-services').checked = config.sections.services !== false;
                    document.getElementById('section-daily-tasks').checked = config.sections.dailyTasks !== false;
                    document.getElementById('section-terminal').checked = config.sections.terminal !== false;
                }

                // Load section order
                if (config.sectionOrder && config.sectionOrder.length > 0) {
                    reorderSectionItems(config.sectionOrder);
                }

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
                customHeadHtml: document.getElementById('custom-head').value,
                sections: {
                    performance: document.getElementById('section-performance').checked,
                    aiUsage: document.getElementById('section-ai-usage').checked,
                    projects: document.getElementById('section-projects').checked,
                    services: document.getElementById('section-services').checked,
                    dailyTasks: document.getElementById('section-daily-tasks').checked,
                    terminal: document.getElementById('section-terminal').checked
                },
                sectionOrder: getSectionOrder()
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

        // Get current section order from DOM
        function getSectionOrder() {
            const list = document.getElementById('section-order-list');
            return Array.from(list.querySelectorAll('.section-order-item')).map(item => item.dataset.section);
        }

        // Reorder section items in settings based on config
        function reorderSectionItems(order) {
            const list = document.getElementById('section-order-list');
            order.forEach(sectionKey => {
                const item = list.querySelector(` + "`" + `[data-section="${sectionKey}"]` + "`" + `);
                if (item) list.appendChild(item);
            });
        }

        // Drag and drop functionality
        let draggedItem = null;

        function initDragAndDrop() {
            const list = document.getElementById('section-order-list');
            const items = list.querySelectorAll('.section-order-item');

            items.forEach(item => {
                item.addEventListener('dragstart', handleDragStart);
                item.addEventListener('dragend', handleDragEnd);
                item.addEventListener('dragover', handleDragOver);
                item.addEventListener('dragenter', handleDragEnter);
                item.addEventListener('dragleave', handleDragLeave);
                item.addEventListener('drop', handleDrop);
            });
        }

        function handleDragStart(e) {
            draggedItem = this;
            this.classList.add('dragging');
            e.dataTransfer.effectAllowed = 'move';
        }

        function handleDragEnd(e) {
            this.classList.remove('dragging');
            document.querySelectorAll('.section-order-item').forEach(item => {
                item.classList.remove('drag-over');
            });
            draggedItem = null;
        }

        function handleDragOver(e) {
            e.preventDefault();
            e.dataTransfer.dropEffect = 'move';
        }

        function handleDragEnter(e) {
            e.preventDefault();
            if (this !== draggedItem) {
                this.classList.add('drag-over');
            }
        }

        function handleDragLeave(e) {
            this.classList.remove('drag-over');
        }

        function handleDrop(e) {
            e.preventDefault();
            this.classList.remove('drag-over');

            if (draggedItem && this !== draggedItem) {
                const list = document.getElementById('section-order-list');
                const items = Array.from(list.querySelectorAll('.section-order-item'));
                const draggedIndex = items.indexOf(draggedItem);
                const targetIndex = items.indexOf(this);

                if (draggedIndex < targetIndex) {
                    this.parentNode.insertBefore(draggedItem, this.nextSibling);
                } else {
                    this.parentNode.insertBefore(draggedItem, this);
                }
            }
        }

        // Load everything
        loadThemes();
        loadConfig();
        initDragAndDrop();
    </script>
</body>
</html>`
