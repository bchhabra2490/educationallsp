# Educational LSP Server From [this video](https://www.youtube.com/watch?v=YsdlcQoHqPY)

An simple Language Server Protocol (LSP) implementation in Go for educational purposes. This server demonstrates the basics of LSP communication and can be connected to Neovim for testing.

## Features

- Basic LSP message handling (initialize, shutdown, exit)
- JSON-RPC over stdio communication
- Neovim integration
- **AI-powered code explanations** using OpenAI (optional)
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
├── analysis/            # Analysis and explanation services
│   ├── state.go        # Document state management
│   └── explanation_service.go  # AI explanation service
├── openai/              # OpenAI API client
│   └── openai.go       # OpenAI completions client
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

## AI-Powered Code Explanations (Optional)

The server can provide AI-powered explanations of code lines when you hover over them in Neovim. To enable this feature:

1. **Get an OpenAI API key** from [OpenAI](https://platform.openai.com/api-keys)
2. **Set the environment variable**:
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```
3. **Restart the LSP server** - it will automatically detect the API key and enable AI explanations

**Note**: Without an API key, the server will still work but will only show the raw line content on hover instead of AI explanations.

## Connecting to VSCode

### Method 1: Install as VSCode Extension (Recommended)

1. **Navigate to the VSCode extension directory**:
   ```bash
   cd vscode-extension
   ```

2. **Run the installation script**:
   ```bash
   ./install.sh
   ```

3. **Install the extension in VSCode**:
   - Open VSCode
   - Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on Mac)
   - Type "Extensions: Install from VSIX"
   - Select the generated `.vsix` file

4. **Configure the extension** (optional):
   - Open VSCode Settings (`Ctrl+,`)
   - Search for "Educational LSP"
   - Set your OpenAI API key if you want AI explanations
   - Adjust the server path if needed

### Method 2: Manual VSCode Setup

1. **Build the LSP server**:
   ```bash
   go build -o main main.go
   ```

2. **Set environment variable** (optional, for AI features):
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```

3. **Configure VSCode settings.json**:
   ```json
   {
     "educationallsp.openaiApiKey": "your-api-key-here",
     "educationallsp.serverPath": "./main"
   }
   ```

### Testing in VSCode

1. **Open any file** in VSCode
2. **Hover over code** to see AI explanations (if API key is set)
3. **Use Ctrl+Space** for completion suggestions
4. **Check the Output panel** → "Educational LSP Server" for logs

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

1. **Open any file** in Neovim:
   ```bash
   nvim test.go
   # or
   nvim test.py
   # or
   nvim test.js
   # or any other file type
   ```

2. **Check for notifications** - you should see messages like:
   - "Educationallsp client started"
   - "Educationallsp LSP client initialized"
   - "Educationallsp attached to buffer X"

3. **Test hover functionality** - hover over any line to see AI explanations

4. **Test completion functionality**:
   - Enter insert mode (`i`)
   - Type some text
   - Press `Ctrl+Space` to trigger completion
   - Use `Tab`/`Shift+Tab` to navigate suggestions

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
- **Check file type**: The server is configured to attach to all file types
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
