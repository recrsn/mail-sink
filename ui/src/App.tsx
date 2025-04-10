import { useEffect, useState } from 'react'
import { Email, ServerInfo, clearEmails, fetchEmails, fetchServerInfo } from './lib/api'
import { EmailList } from './components/EmailList'
import { EmailDetail } from './components/EmailDetail'
import { Separator } from './components/ui/separator'
import { Toaster } from './components/ui/sonner'
import { toast } from 'sonner'
import { ModeToggle } from './components/ui/mode-toggle'
import { Badge } from './components/ui/badge'
import { Server } from 'lucide-react'
import {
  ResizablePanelGroup,
  ResizablePanel,
  ResizableHandle
} from './components/ui/resizable'

function App() {
  const [emails, setEmails] = useState<Email[]>([])
  const [selectedEmail, setSelectedEmail] = useState<Email | null>(null)
  const [serverInfo, setServerInfo] = useState<ServerInfo | null>(null)
  const [loading, setLoading] = useState(true)

  // Fetch emails on load and periodically
  useEffect(() => {
    const loadEmails = async () => {
      try {
        setLoading(true)
        const data = await fetchEmails()
        setEmails(data)

        // If we have a selected email, update it with fresh data
        if (selectedEmail) {
          const updatedEmail = data.find(e => e.id === selectedEmail.id)
          if (updatedEmail) {
            setSelectedEmail(updatedEmail)
          }
        }
      } catch (err) {
        toast.error('Failed to load emails')
        console.error(err)
      } finally {
        setLoading(false)
      }
    }

    // Load emails immediately
    loadEmails()

    // Then refresh every 5 seconds
    const interval = setInterval(loadEmails, 5000)

    // Load server info
    fetchServerInfo().then(setServerInfo).catch(console.error)

    return () => clearInterval(interval)
  }, [selectedEmail])

  const handleRefresh = async () => {
    try {
      setLoading(true)
      const data = await fetchEmails()
      setEmails(data)
      toast.success('Emails refreshed')
    } catch (err) {
      toast.error('Failed to refresh emails')
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  const handleClearEmails = async () => {
    try {
      await clearEmails()
      setEmails([])
      setSelectedEmail(null)
      toast.success('All emails cleared')
    } catch (err) {
      toast.error('Failed to clear emails')
      console.error(err)
    }
  }

  return (
    <>
      <ModeToggle />
      <Toaster position="top-right" />
      <ResizablePanelGroup
        direction="horizontal"
        className="h-full border"
      >
        <ResizablePanel defaultSize={30} minSize={20}>
          <div className="p-4 h-full overflow-auto">
          <div className="flex flex-col gap-2">
            <div className="flex items-center gap-2">
              <h1 className="text-3xl font-bold">Mail Sink</h1>
            </div>
            {serverInfo && (
              <Badge variant="outline" className="flex items-center gap-2">
                <Server className="h-4 w-4" />
                <span>SMTP {serverInfo.smtpPort}</span>
                <div className="h-2 w-2 rounded-full bg-green-500"></div>
              </Badge>
            )}
          </div>
          <Separator className="my-4" />
            <EmailList
              emails={emails}
              loading={loading}
              selectedEmailId={selectedEmail?.id}
              onSelectEmail={(email) => setSelectedEmail(email)}
              onRefresh={handleRefresh}
              onClear={handleClearEmails}
            />
          </div>
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel defaultSize={70}>
          <div className="p-8 h-full">
            {selectedEmail ? (
              <EmailDetail email={selectedEmail} />
            ) : (
              <div className="flex items-center justify-center h-full">
                <p className="text-muted-foreground">Select an email to view its contents</p>
              </div>
            )}
          </div>
        </ResizablePanel>
      </ResizablePanelGroup>
    </>
  )
}

export default App
