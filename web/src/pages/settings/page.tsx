import { useDocumentTitle } from "@/hooks/use-document-title";

export function SettingsPage() {
  useDocumentTitle(() => "Settings", []);
  return (
    <div>
      <h1>Settings</h1>
    </div>
  )
}