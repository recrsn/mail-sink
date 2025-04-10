import { Moon, Sun } from "lucide-react"
import { useTheme } from "../ui/theme-provider"
import { Button } from "@/components/ui/button"

export function ModeToggle({ className, ...props }: React.ComponentProps<typeof Button>) {
  const { setTheme, theme } = useTheme()

  return (
    <Button
      variant="outline"
      size="icon"
      onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
      className={`fixed bottom-6 left-6 rounded-full h-12 w-12 shadow-md z-50 transition-all duration-500 hover:scale-110 ${className}`}
      {...props}
    >
      <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all duration-500 dark:-rotate-90 dark:scale-0" />
      <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all duration-500 dark:rotate-0 dark:scale-100" />
      <span className="sr-only">Toggle theme</span>
    </Button>
  )
}
