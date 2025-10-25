#!/bin/bash

# Build the LSP server
echo "Building LSP server..."
cd ..
go build -o main main.go
cd vscode-extension

# Install VSCode extension dependencies
echo "Installing VSCode extension dependencies..."
npm install

# Compile TypeScript
echo "Compiling TypeScript..."
npm run compile

# Install vsce for packaging
echo "Installing vsce..."
npm install -g vsce

# Create VSIX package
echo "Creating VSIX package..."
npm run package

echo "VSCode extension setup complete!"
echo ""
echo "To install the extension:"
echo "1. Open VSCode"
echo "2. Press Ctrl+Shift+P (Cmd+Shift+P on Mac)"
echo "3. Type 'Extensions: Install from VSIX'"
echo "4. Select the .vsix file from this directory"
echo ""
echo "Or run: code --install-extension educationallsp-0.0.1.vsix"
