# Educational LSP Server From [this video](https://www.youtube.com/watch?v=YsdlcQoHqPY)

An simple Language Server Protocol (LSP) implementation in Go for educational purposes. This server demonstrates the basics of LSP communication and can be connected to Neovim for testing.

## Features

- Basic LSP message handling (initialize, shutdown, exit)
- JSON-RPC over stdio communication
- Neovim integration
- Logging for debugging
## Project Structure

```
educationallsp/
├── main.go              # Main LSP server entry point
├── rpc/                 # RPC message handling
│   ├── rpc.go          # Message encoding/decoding
│   └── rpc_test.go     # Tests for RPC functions
├── lsp/                 # LSP protocol definitions
│   ├── lsp.go          # Basic LSP message types
│   └── initialize.go   # Initialize request/response types
├── go.mod              # Go module definition
└── educationallsp.log  # Server log file (created at runtime)
```

## Building the Server

1. Make sure you have Go installed (version 1.22.1 or later)
2. Build the server:
   ```bash
   go build -o main main.go
   ```
3. The binary will be created as `main` in the project directory

## Connecting to Neovim

### Method 1: Automatic Configuration (Recommended)

1. **Create Neovim config directory** (if it doesn't exist):
   ```bash
   mkdir -p ~/.config/nvim
   ```

2. **Create the LSP configuration file**:
   ```bash
   # Copy the provided configuration
   cp /path/to/your/educationallsp/load_test_lsp.lua ~/.config/nvim/lua/educationallsp.lua
   ```

3. **Create or update your Neovim init file** (`~/.config/nvim/init.lua`):
   ```lua
   -- Load the LSP configuration
   local educationallsp = require('educationallsp')
   educationallsp.setup()
   
   -- Optional: Add other Neovim settings
   vim.opt.number = true
   vim.opt.relativenumber = true
   ```

4. **Restart Neovim** - the LSP server will now load automatically!

### Method 2: Manual Loading

If you prefer to load the LSP manually each time:

1. Open Neovim
2. Source the configuration file:
   ```vim
   :source /path/to/your/educationallsp/load_test_lsp.lua
   ```

### Method 3: Using vim.cmd in init.lua

Add this line to your `~/.config/nvim/init.lua`:
```lua
vim.cmd('source /path/to/your/educationallsp/load_test_lsp.lua')
```

## Testing the Connection

1. **Open a markdown file** in Neovim:
   ```bash
   nvim test.md
   ```

2. **Check for notifications** - you should see messages like:
   - "Educationallsp client started"
   - "Educationallsp LSP client initialized"
   - "Educationallsp attached to buffer X"

3. **Check the log file** (`educationallsp.log`) to see what messages the server is receiving:
   ```bash
   tail -f educationallsp.log
   ```

4. **Verify LSP status** in Neovim:
   ```vim
   :LspInfo
   ```

## Troubleshooting

### LSP Server Not Starting
- **Check binary path**: Make sure the path in your Lua config matches where you built the binary
- **Check permissions**: Ensure the binary is executable (`chmod +x main`)
- **Check logs**: Look at `educationallsp.log` for error messages

### Neovim Not Loading Configuration
- **Check config location**: Ensure files are in `~/.config/nvim/`
- **Check syntax**: Validate your Lua syntax
- **Check Neovim version**: Ensure you're using Neovim 0.5+ (LSP support)

### No LSP Features Working
- **Check file type**: The server is configured to attach to markdown files
- **Check LSP status**: Use `:LspInfo` to see if the client is attached
- **Check logs**: Look for error messages in the log file

## Development

### Adding New LSP Features

1. **Define message types** in the `lsp/` directory
2. **Handle new methods** in `main.go`'s `handleMessage` function
3. **Add tests** for new functionality
4. **Rebuild** the server: `go build -o main main.go`

### Debugging

- **Server logs**: Check `educationallsp.log` for server-side debugging
- **Neovim logs**: Use `:messages` to see Neovim notifications
- **LSP status**: Use `:LspInfo` and `:LspLog` for LSP-specific debugging

## Next Steps

This is a basic LSP implementation. To make it more functional, consider adding:

- Text document synchronization
- Diagnostics (error/warning reporting)
- Code completion
- Hover information
- Go to definition
- Symbol search

## Resources

- [LSP Specification](https://microsoft.github.io/language-server-protocol/)
- [Neovim LSP Documentation](https://neovim.io/doc/user/lsp.html)
- [Go JSON-RPC](https://pkg.go.dev/encoding/json)
