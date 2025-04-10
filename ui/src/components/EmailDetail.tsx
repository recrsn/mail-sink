import { useState } from "react"
import { Attachment, Email } from "@/lib/api"
import { Badge } from "@/components/ui/badge"
import { Separator } from "@/components/ui/separator"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog"

interface EmailDetailProps {
  email: Email
}

export function EmailDetail({ email }: EmailDetailProps) {
  const [activeTab, setActiveTab] = useState<string>("preview")
  const [openAttachment, setOpenAttachment] = useState<Attachment | null>(null)

  return (
    <div className="h-full flex flex-col">
      <div className="mb-4">
        <h2 className="text-xl font-bold mb-2">{email.subject || "(No Subject)"}</h2>
        <div className="grid grid-cols-2 gap-4 text-sm">
          <div>
            <span className="font-semibold">From:</span> {email.from}
          </div>
          <div>
            <span className="font-semibold">To:</span> {email.to.join(", ")}
          </div>
          <div>
            <span className="font-semibold">Date:</span> {new Date(email.receivedAt).toLocaleString()}
          </div>
          {email.attachments && email.attachments.length > 0 && (
            <div>
              <span className="font-semibold">Attachments:</span>{" "}
              {email.attachments.map((attachment, index) => (
                <Badge 
                  key={index} 
                  variant="outline" 
                  className="ml-1 cursor-pointer hover:bg-accent"
                  onClick={() => setOpenAttachment(attachment)}
                >
                  {attachment.filename}
                </Badge>
              ))}
            </div>
          )}
        </div>
      </div>
      <Separator className="mb-4" />
      <div className="flex-grow overflow-auto">
        <Tabs value={activeTab} onValueChange={setActiveTab} className="h-full flex flex-col">
          <TabsList className="mb-4">
            {email.htmlBody && <TabsTrigger value="preview">Preview</TabsTrigger>}
            {email.textBody && <TabsTrigger value="text">Plain Text</TabsTrigger>}
            <TabsTrigger value="headers">Headers</TabsTrigger>
          </TabsList>
          
          {email.htmlBody && (
            <TabsContent value="preview" className="mt-0 flex-grow">
              <div 
                className="border p-4 rounded-md overflow-auto max-h-[60vh]"
                dangerouslySetInnerHTML={{ __html: email.htmlBody }}
              />
            </TabsContent>
          )}
          
          {email.textBody && (
            <TabsContent value="text" className="mt-0 flex-grow">
              <pre className="border p-4 rounded-md overflow-auto max-h-[60vh] whitespace-pre-wrap">
                {email.textBody}
              </pre>
            </TabsContent>
          )}
          
          <TabsContent value="headers" className="mt-0 flex-grow">
            <div className="border p-4 rounded-md overflow-auto max-h-[60vh]">
              <table className="w-full">
                <tbody>
                  {Object.entries(email.headers).map(([key, values]) => (
                    <tr key={key} className="border-b last:border-b-0">
                      <td className="py-2 pr-4 font-semibold align-top">{key}</td>
                      <td className="py-2 break-words">{values.join("\n")}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </TabsContent>
        </Tabs>
      </div>

      {/* Attachment Dialog */}
      {openAttachment && (
        <Dialog open={!!openAttachment} onOpenChange={(open) => !open && setOpenAttachment(null)}>
          <DialogContent className="max-w-3xl">
            <DialogHeader>
              <DialogTitle>{openAttachment.filename}</DialogTitle>
            </DialogHeader>
            <div className="text-sm mb-2">
              <span className="font-semibold">Type:</span> {openAttachment.contentType} | 
              <span className="font-semibold ml-2">Size:</span> {(openAttachment.size / 1024).toFixed(2)} KB
            </div>
            {openAttachment.contentType.startsWith("image/") ? (
              <div className="border rounded-md p-2 bg-gray-50 text-center">
                <img 
                  src={`data:${openAttachment.contentType};base64,${openAttachment.content}`} 
                  alt={openAttachment.filename}
                  className="max-h-[60vh] max-w-full mx-auto"
                />
              </div>
            ) : openAttachment.contentType.startsWith("text/") ? (
              <pre className="border rounded-md p-4 overflow-auto max-h-[60vh] whitespace-pre-wrap bg-gray-50">
                {atob(openAttachment.content)}
              </pre>
            ) : (
              <div className="border rounded-md p-4 text-center">
                <p className="mb-4">Download attachment:</p>
                <a 
                  href={`data:${openAttachment.contentType};base64,${openAttachment.content}`} 
                  download={openAttachment.filename}
                  className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 inline-block"
                >
                  Download {openAttachment.filename}
                </a>
              </div>
            )}
          </DialogContent>
        </Dialog>
      )}
    </div>
  )
}