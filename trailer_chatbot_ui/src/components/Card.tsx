export function Card({ children, className }: { children: React.ReactNode; className?: string }) {
  return <div className={`shadow-lg border rounded-lg p-4 bg-white ${className}`}>{children}</div>;
}

export function CardContent({ children, className }: { children: React.ReactNode; className?: string }) {
  return <div className={`p-4 ${className}`}>{children}</div>;
}
