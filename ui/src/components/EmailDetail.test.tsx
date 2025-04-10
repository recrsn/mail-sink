import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { EmailDetail } from './EmailDetail';
import { type Email } from '@/lib/api';

// Mock email data
const mockEmail: Email = {
  id: 'email1',
  from: 'sender@example.com',
  to: ['recipient@example.com', 'recipient2@example.com'],
  subject: 'Test Subject',
  textBody: 'This is the text body of the email',
  htmlBody: '<p>This is the <strong>HTML</strong> body of the email</p>',
  attachments: [
    {
      filename: 'test.pdf',
      contentType: 'application/pdf',
      size: 1024,
      content: 'PDF content'
    }
  ],
  headers: {
    'Content-Type': ['multipart/mixed'],
    'From': ['sender@example.com'],
    'To': ['recipient@example.com', 'recipient2@example.com'],
    'Subject': ['Test Subject'],
    'Date': ['Mon, 01 Jan 2023 12:00:00 +0000']
  },
  receivedAt: new Date(2023, 0, 1, 12, 0, 0).toISOString()
};

describe('EmailDetail', () => {
  it('renders email details correctly', () => {
    render(<EmailDetail email={mockEmail} />);
    
    // Check basic email information
    expect(screen.getByText('Test Subject')).toBeInTheDocument();
    expect(screen.getByText(/sender@example.com/)).toBeInTheDocument();
    expect(screen.getByText(/recipient@example.com, recipient2@example.com/)).toBeInTheDocument();
    
    // Check tabs
    expect(screen.getByRole('tab', { name: 'Preview' })).toBeInTheDocument();
    expect(screen.getByRole('tab', { name: 'Plain Text' })).toBeInTheDocument();
    expect(screen.getByRole('tab', { name: 'Headers' })).toBeInTheDocument();
  });
});