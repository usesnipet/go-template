import { useDocumentTitle } from "@/hooks/use-document-title";

export function HomePage() {
  useDocumentTitle(() => "Home", []);
  return (
    <div>
      <h1>Home</h1>
    </div>
  )
}