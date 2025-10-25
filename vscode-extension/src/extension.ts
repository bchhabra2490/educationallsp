import * as vscode from 'vscode';
import * as path from 'path';
import { LanguageClient, LanguageClientOptions, ServerOptions } from 'vscode-languageclient/node';

let client: LanguageClient;

export function activate(context: vscode.ExtensionContext) {
    console.log('Educational LSP extension is now active!');

    // Get configuration
    const config = vscode.workspace.getConfiguration('educationallsp');
    const serverPath = config.get<string>('serverPath', './main');
    const openaiApiKey = config.get<string>('openaiApiKey', '');
    const supportedLanguages = config.get<string[]>('supportedLanguages', ['plaintext', 'markdown', 'text']);
    const priority = config.get<number>('priority', 0);

    // Resolve server path
    const serverExecutable = path.resolve(context.extensionPath, '..', '..', serverPath);
    
    // Set up environment variables
    const env = { ...process.env };
    if (openaiApiKey) {
        env.OPENAI_API_KEY = openaiApiKey;
    }

    // Create server options
    const serverOptions: ServerOptions = {
        command: serverExecutable,
        options: {
            env: env
        }
    };

    // Create document selector from configuration
    const documentSelector = supportedLanguages.map(lang => ({
        scheme: 'file',
        language: lang
    }));

    // Create client options
    const clientOptions: LanguageClientOptions = {
        // Register the server for configured file types to avoid conflicts
        documentSelector: documentSelector,
        synchronize: {
            fileEvents: vscode.workspace.createFileSystemWatcher('**/*')
        },
        // Add initialization options to avoid conflicts
        initializationOptions: {
            priority: priority
        }
    };

    // Create language client
    client = new LanguageClient(
        'educationallsp',
        'Educational LSP Server',
        serverOptions,
        clientOptions
    );

    // Start the client
    client.start();

    // Register commands
    const restartCommand = vscode.commands.registerCommand('educationallsp.restart', () => {
        client.restart();
        vscode.window.showInformationMessage('Educational LSP Server restarted');
    });

    const showLogsCommand = vscode.commands.registerCommand('educationallsp.showLogs', () => {
        const logPath = path.resolve(context.extensionPath, '..', '..', 'educationallsp.log');
        vscode.workspace.openTextDocument(logPath).then(doc => {
            vscode.window.showTextDocument(doc);
        });
    });

    context.subscriptions.push(restartCommand, showLogsCommand);
}

export function deactivate(): Thenable<void> | undefined {
    if (!client) {
        return undefined;
    }
    return client.stop();
}
