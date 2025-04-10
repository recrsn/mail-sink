// API types
export interface Email {
  id: string;
  from: string;
  to: string[];
  subject: string;
  textBody?: string;
  htmlBody?: string;
  attachments?: Attachment[];
  headers: Record<string, string[]>;
  receivedAt: string;
}

export interface Attachment {
  filename: string;
  contentType: string;
  size: number;
  content: string;
}

export interface ServerInfo {
  smtpPort: number;
  httpPort: number;
}

// API endpoints
const API_BASE = '/api';

// Fetch all emails
export async function fetchEmails(): Promise<Email[]> {
  const response = await fetch(`${API_BASE}/emails`);
  if (!response.ok) {
    throw new Error(`Failed to fetch emails: ${response.statusText}`);
  }
  return response.json();
}

// Fetch a single email by ID
export async function fetchEmailById(id: string): Promise<Email> {
  const response = await fetch(`${API_BASE}/emails/${id}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch email: ${response.statusText}`);
  }
  return response.json();
}

// Clear all emails
export async function clearEmails(): Promise<void> {
  const response = await fetch(`${API_BASE}/emails`, {
    method: 'DELETE',
  });
  if (!response.ok) {
    throw new Error(`Failed to clear emails: ${response.statusText}`);
  }
}

// Fetch server info
export async function fetchServerInfo(): Promise<ServerInfo> {
  const response = await fetch(`${API_BASE}/info`);
  if (!response.ok) {
    throw new Error(`Failed to fetch server info: ${response.statusText}`);
  }
  return response.json();
}