import { Email } from "@/lib/api";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { cn, formatDistanceToNow } from "@/lib/utils";
import { RefreshCw, Trash2 } from "lucide-react";

interface EmailListProps {
  emails: Email[];
  loading: boolean;
  selectedEmailId?: string;
  onSelectEmail: (email: Email) => void;
  onRefresh?: () => void;
  onClear?: () => void;
}

export function EmailList({
  emails,
  loading,
  selectedEmailId,
  onSelectEmail,
  onRefresh,
  onClear,
}: EmailListProps) {
  return (
    <div className="flex flex-col">
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <span className="font-medium">Emails</span>
          <Badge variant="outline">{emails.length}</Badge>
        </div>
        <div className="flex items-center gap-2">
          <Button
            variant="outline"
            size="icon"
            onClick={onRefresh}
            disabled={loading}
            className="h-8 w-8"
          >
            <RefreshCw
              className={`h-4 w-4 ${loading ? "animate-spin" : ""}`}
            />
          </Button>
          <Button
            variant="destructive"
            size="icon"
            onClick={onClear}
            className="h-8 w-8"
          >
            <Trash2 className="h-4 w-4" />
          </Button>
        </div>
      </div>

      {loading && emails.length === 0 ? (
        <div className="text-center py-8 flex-grow">
          <p className="text-muted-foreground">Loading emails...</p>
        </div>
      ) : emails.length === 0 ? (
        <div className="text-center py-8 flex-grow">
          <p className="text-muted-foreground">No emails received yet</p>
        </div>
      ) : (
        <div className="overflow-auto flex-grow">
          <div className="flex flex-col gap-2 pt-0">
            {emails.map((email) => (
              <div
                key={email.id}
                className={cn(
                  "cursor-pointer",
                  "flex flex-col items-start gap-2 rounded-lg border p-3 text-left text-sm transition-all hover:bg-accent",
                  email.id === selectedEmailId && "bg-muted"
                )}
                onClick={() => onSelectEmail(email)}
              >
                <div className="flex w-full flex-col gap-1">
                  <div className="flex items-center">
                    <div className="flex items-center gap-2">
                      <div className="font-semibold">{email.from}</div>
                    </div>
                    <div
                      className={cn(
                        "ml-auto text-xs",
                        email.id === selectedEmailId
                          ? "text-foreground"
                          : "text-muted-foreground"
                      )}
                    >
                      {formatDistanceToNow(new Date(email.receivedAt))}
                    </div>
                  </div>
                  <div className="text-xs font-medium">{email.subject}</div>
                </div>
                <div className="line-clamp-2 text-xs text-muted-foreground">
                  {(
                    email.textBody ||
                    email.htmlBody?.replace(/<[^>]*>/g, "") ||
                    ""
                  ).substring(0, 300)}
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
