import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { fetchEmails, fetchEmailById, clearEmails, fetchServerInfo } from './api';

// Setup mock for global fetch
const mockFetch = vi.fn();
global.fetch = mockFetch;

describe('API', () => {
  beforeEach(() => {
    vi.resetAllMocks();
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe('fetchEmails', () => {
    it('fetches emails successfully', async () => {
      const mockEmails = [
        { id: 'email1', from: 'sender1@example.com' },
        { id: 'email2', from: 'sender2@example.com' }
      ];

      mockFetch.mockResolvedValueOnce({
        ok: true,
        json: async () => mockEmails,
      });

      const result = await fetchEmails();

      expect(mockFetch).toHaveBeenCalledWith('/api/emails');
      expect(result).toEqual(mockEmails);
    });

    it('throws an error when fetch fails', async () => {
      mockFetch.mockResolvedValueOnce({
        ok: false,
        statusText: 'Not Found',
      });

      await expect(fetchEmails()).rejects.toThrow('Failed to fetch emails: Not Found');
    });
  });

  describe('fetchEmailById', () => {
    it('fetches a single email by ID successfully', async () => {
      const mockEmail = { id: 'email1', from: 'sender@example.com' };

      mockFetch.mockResolvedValueOnce({
        ok: true,
        json: async () => mockEmail,
      });

      const result = await fetchEmailById('email1');

      expect(mockFetch).toHaveBeenCalledWith('/api/emails/email1');
      expect(result).toEqual(mockEmail);
    });

    it('throws an error when fetch fails', async () => {
      mockFetch.mockResolvedValueOnce({
        ok: false,
        statusText: 'Not Found',
      });

      await expect(fetchEmailById('nonexistent')).rejects.toThrow('Failed to fetch email: Not Found');
    });
  });

  describe('clearEmails', () => {
    it('clears emails successfully', async () => {
      mockFetch.mockResolvedValueOnce({
        ok: true,
      });

      await clearEmails();

      expect(mockFetch).toHaveBeenCalledWith('/api/emails', {
        method: 'DELETE',
      });
    });

    it('throws an error when clear fails', async () => {
      mockFetch.mockResolvedValueOnce({
        ok: false,
        statusText: 'Internal Server Error',
      });

      await expect(clearEmails()).rejects.toThrow('Failed to clear emails: Internal Server Error');
    });
  });

  describe('fetchServerInfo', () => {
    it('fetches server info successfully', async () => {
      const mockServerInfo = { smtpPort: 1025, httpPort: 8025 };

      mockFetch.mockResolvedValueOnce({
        ok: true,
        json: async () => mockServerInfo,
      });

      const result = await fetchServerInfo();

      expect(mockFetch).toHaveBeenCalledWith('/api/info');
      expect(result).toEqual(mockServerInfo);
    });

    it('throws an error when fetch fails', async () => {
      mockFetch.mockResolvedValueOnce({
        ok: false,
        statusText: 'Internal Server Error',
      });

      await expect(fetchServerInfo()).rejects.toThrow('Failed to fetch server info: Internal Server Error');
    });
  });
});