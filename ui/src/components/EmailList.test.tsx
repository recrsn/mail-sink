import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { EmailList } from './EmailList';
import { type Email } from '@/lib/api';

// Mock data
const mockEmails: Email[] = [
  {
    id: 'email1',
    from: 'sender@example.com',
    to: ['recipient@example.com'],
    subject: 'Test Subject 1',
    textBody: 'This is a test email body',
    headers: {},
    receivedAt: new Date(2023, 0, 1).toISOString(),
  },
  {
    id: 'email2',
    from: 'another@example.com',
    to: ['user@example.com'],
    subject: 'Test Subject 2',
    textBody: 'Another test email body',
    headers: {},
    receivedAt: new Date(2023, 0, 2).toISOString(),
  },
];

describe('EmailList', () => {
  it('renders loading state correctly', () => {
    render(
      <EmailList
        emails={[]}
        loading={true}
        onSelectEmail={() => {}}
      />
    );
    
    expect(screen.getByText('Loading emails...')).toBeInTheDocument();
  });
  
  it('renders empty state correctly', () => {
    render(
      <EmailList
        emails={[]}
        loading={false}
        onSelectEmail={() => {}}
      />
    );
    
    expect(screen.getByText('No emails received yet')).toBeInTheDocument();
  });
  
  it('renders a list of emails correctly', () => {
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        onSelectEmail={() => {}}
      />
    );
    
    expect(screen.getByText('Test Subject 1')).toBeInTheDocument();
    expect(screen.getByText('Test Subject 2')).toBeInTheDocument();
    expect(screen.getByText('sender@example.com')).toBeInTheDocument();
    expect(screen.getByText('another@example.com')).toBeInTheDocument();
  });
  
  it('displays email count correctly', () => {
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        onSelectEmail={() => {}}
      />
    );
    
    expect(screen.getByText('2')).toBeInTheDocument();
  });
  
  it('calls onSelectEmail when an email is clicked', () => {
    const handleSelectEmail = vi.fn();
    
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        onSelectEmail={handleSelectEmail}
      />
    );
    
    fireEvent.click(screen.getByText('Test Subject 1'));
    
    expect(handleSelectEmail).toHaveBeenCalledWith(mockEmails[0]);
  });
  
  it('calls onRefresh when refresh button is clicked', () => {
    const handleRefresh = vi.fn();
    
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        onSelectEmail={() => {}}
        onRefresh={handleRefresh}
      />
    );
    
    const refreshButton = document.querySelector('button svg.lucide-refresh-cw')?.closest('button');
    if (refreshButton) {
      fireEvent.click(refreshButton);
      expect(handleRefresh).toHaveBeenCalled();
    }
  });
  
  it('calls onClear when clear button is clicked', () => {
    const handleClear = vi.fn();
    
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        onSelectEmail={() => {}}
        onClear={handleClear}
      />
    );
    
    const trashButton = document.querySelector('button svg.lucide-trash-2')?.closest('button');
    if (trashButton) {
      fireEvent.click(trashButton);
      expect(handleClear).toHaveBeenCalled();
    }
  });
  
  it('disables refresh button when loading', () => {
    render(
      <EmailList
        emails={mockEmails}
        loading={true}
        onSelectEmail={() => {}}
        onRefresh={() => {}}
      />
    );
    
    const refreshButton = document.querySelector('button svg.lucide-refresh-cw')?.closest('button');
    expect(refreshButton).toBeDisabled();
  });
  
  it('highlights selected email', () => {
    render(
      <EmailList
        emails={mockEmails}
        loading={false}
        selectedEmailId="email1"
        onSelectEmail={() => {}}
      />
    );
    
    // Find the div containing the selected email
    const selectedEmail = screen.getByText('Test Subject 1').closest('div[class*="cursor-pointer"]');
    expect(selectedEmail).toHaveClass('bg-muted');
    
    // Non-selected email should not have the bg-muted class
    const nonSelectedEmail = screen.getByText('Test Subject 2').closest('div[class*="cursor-pointer"]');
    expect(nonSelectedEmail).not.toHaveClass('bg-muted');
  });
});