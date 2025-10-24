-- Check if the binary exists
local binary_path = "/Users/b-eq/Desktop/projects/educationallsp/main"
if vim.fn.executable(binary_path) == 0 then
    vim.notify("educationallsp binary not found at " .. binary_path, vim.log.levels.ERROR)
    return
end

local client = vim.lsp.start_client {
    name = "educationallsp",    
    cmd = {binary_path},
    root_dir = vim.fn.getcwd(),
    capabilities = vim.lsp.protocol.make_client_capabilities(),
    on_init = function(client, result)
        vim.notify("Educationallsp LSP client initialized", vim.log.levels.INFO)
    end,
    on_attach = function(client, bufnr)
        vim.notify("Educationallsp attached to buffer " .. bufnr, vim.log.levels.INFO)
    end,
}

if not client then
    vim.notify("Failed to start educationallsp client", vim.log.levels.ERROR)
    return
end

vim.notify("Educationallsp client started", vim.log.levels.INFO)

-- Attach to markdown files
vim.api.nvim_create_autocmd("FileType", {
    pattern = "markdown",
    callback = function()
        vim.lsp.buf_attach_client(0, client)
    end,
})