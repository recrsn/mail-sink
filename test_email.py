#!/usr/bin/env python3
"""
Simple script to send a test email to the SMTP sink.
"""
import smtplib
import sys
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.mime.application import MIMEApplication

def send_test_email(smtp_host="localhost", smtp_port=1025):
    # Create message
    msg = MIMEMultipart('alternative')
    msg['Subject'] = 'Test Email from SMTP Sink'
    msg['From'] = 'sender@example.com'
    msg['To'] = 'recipient@example.com'
    
    # Plain text version
    text = """
    Hello,
    
    This is a test email sent to the SMTP Sink.
    
    Regards,
    SMTP Sink Test
    """
    
    # HTML version
    html = """
    <html>
      <head></head>
      <body>
        <h1>Test Email</h1>
        <p>This is a <b>test email</b> sent to the SMTP Sink.</p>
        <p>It contains:</p>
        <ul>
          <li>Plain text part</li>
          <li>HTML part</li>
          <li>Attachment</li>
        </ul>
        <p>Regards,<br>SMTP Sink Test</p>
      </body>
    </html>
    """
    
    # Attach parts
    part1 = MIMEText(text, 'plain')
    part2 = MIMEText(html, 'html')
    msg.attach(part1)
    msg.attach(part2)
    
    # Create a simple attachment
    attachment = MIMEApplication("This is a test attachment content.")
    attachment.add_header('Content-Disposition', 'attachment', filename='test.txt')
    msg.attach(attachment)
    
    # Send the message
    try:
        smtp = smtplib.SMTP(smtp_host, smtp_port)
        smtp.sendmail(msg['From'], msg['To'], msg.as_string())
        smtp.quit()
        print(f"Test email sent successfully to {smtp_host}:{smtp_port}")
        return True
    except Exception as e:
        print(f"Failed to send email: {e}")
        return False

if __name__ == "__main__":
    host = "localhost"
    port = 1025
    
    # Parse command line arguments
    if len(sys.argv) > 1:
        port = int(sys.argv[1])
    if len(sys.argv) > 2:
        host = sys.argv[2]
    
    send_test_email(host, port)